package vagrant

import (
  "github.com/gruffwizard/kabnet/cluster"
  //"github.com/gruffwizard/kabnet/util"
)
/*
apiVersion: kubeadm.k8s.io/v1beta1
kind: JoinConfiguration
discovery:
  bootstrapToken:
    token: qspy68.jy731iep1jb59uni
    apiServerEndpoint: "haproxy-0.test.dev.kab:6443"
    unsafeSkipCAVerification: true
nodeRegistration:
  name: master-1.test.dev.kab
controlPlane:
  localAPIEndpoint:
    advertiseAddress:


*/


/*

apiVersion: kubeadm.k8s.io/v1beta2
caCertPath: /etc/kubernetes/pki/ca.crt
discovery:
  bootstrapToken:
    apiServerEndpoint: kube-apiserver:6443
    token: abcdef.0123456789abcdef
    unsafeSkipCAVerification: true
  timeout: 5m0s
  tlsBootstrapToken: abcdef.0123456789abcdef
kind: JoinConfiguration
nodeRegistration:
  criSocket: /var/run/dockershim.sock
  name: master-1.test.dev.kab
  taints: null

*/

func genMaster1Deploy(c *cluster.Cluster,m *cluster.Machine) *cluster.Deploy {

  d := c.NewDeploy("first master")

  m.Cluster().FindMachine("haproxy-0")




  d.AddCommand("kubeadm init --apiserver-advertise-address=%s --apiserver-cert-extra-sans=%s  --node-name %s --pod-network-cidr=172.16.0.0/16",
                m.IP,m.IP,m.Fqdn())
  d.AddCommand("export KUBECONFIG=/etc/kubernetes/admin.conf")
  d.AddCommand("kubectl apply -f https://raw.githubusercontent.com/ecomm-integration-ballerina/kubernetes-cluster/master/calico/rbac-kdd.yaml")
  d.AddCommand("kubectl apply -f https://raw.githubusercontent.com/ecomm-integration-ballerina/kubernetes-cluster/master/calico/calico.yaml")
  d.AddCommand("kubeadm token create --print-join-command > /vagrant/join.sh")
  d.AddCommand("kubeadm token create --print-join-command > /vagrant/mjoin.sh")
  d.AddCommand("echo  '--experimental-control-plane' >> /vagrant/mjoin.sh")
  d.AddCommand("chmod +x /vagrant/join.sh")
  d.AddCommand("chmod +x /vagrant/mjoin.sh")
  d.AddCommand("cp /etc/kubernetes/admin.conf /vagrant/config.kube")
  d.AddCommand("tar zcvf /vagrant/certs.tar.gz /etc/kubernetes/admin.conf /etc/kubernetes/pki/ca.crt /etc/kubernetes/pki/ca.key /etc/kubernetes/pki/sa.key /etc/kubernetes/pki/sa.pub /etc/kubernetes/pki/front-proxy-ca.crt /etc/kubernetes/pki/front-proxy-ca.key /etc/kubernetes/pki/etcd/ca.crt /etc/kubernetes/pki/etcd/ca.key")
  return d
}

package vagrant

import (
  "github.com/gruffwizard/kabnet/cluster"
)

// kubeadm config images pull

func genWorkerDeployment(c *cluster.Cluster) *cluster.Deploy {

d := c.NewDeploy("worker")
d.AddCommand("sudo --user=vagrant mkdir -p /home/vagrant/.kube")
d.AddCommand("cp /vagrant/config.kube  /home/vagrant/.kube/config")
d.AddCommand("chown $(id -u vagrant):$(id -g vagrant) /home/vagrant/.kube/config")
d.AddCommand("/vagrant/join.sh")
return d

}


/*

kubeadm init --apiserver-advertise-address=192.168.50.10
 --apiserver-cert-extra-sans=192.168.50.10
 --node-name=master-0.test.dev.kab
  --pod-network-cidr=172.16.0.0/16

kubeadm join 192.168.50.10:6443 --token 9mo6qb.siitxxziy767n3l0 \
    --discovery-token-ca-cert-hash sha256:72e9d26293db0058a4813497a43751d26f1e2a3b1ddba74cd33dcf2cd67f9429
root@master-0:/home/vagrant#
*/

package vagrant

import (
  "github.com/gruffwizard/kabnet/cluster"
  "github.com/gruffwizard/kabnet/util"
)

// kubeadm config images pull
func genEnableAPTCachingDeploy(c *cluster.Cluster,m *cluster.Machine) *cluster.Deploy {
  d := c.NewDeploy("apt cache - use host filesystem as cache for apt")

  d.AddCommand("mkdir -p /vagrant/machines/%s/cache",m.Name)
  apt:= util.CreateText()
  apt.Add("Dir::Cache{Archives /vagrant/machines/%s/cache}",m.Name)
  am:=c.FindMachine("apt-cache")
  if am!=nil {
    apt.Add("Acquire::http::Proxy \"http://%s:3142/apt-cacher/\";",am.IP)
  }
  d.AddFile("apt.conf","/etc/apt/apt.conf",apt.AsString())
  d.AddCommand("apt-get update")

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

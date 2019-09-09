
package vagrant


import (
  "github.com/gruffwizard/kabnet/cluster"
)

func genKubeDeploy(c *cluster.Cluster) *cluster.Deploy {



d := c.NewDeploy("kubernetes")

// /etc/apt/apt.conf.d
// Acquire::http::Proxy "http://user:password@proxy.server:port/";

d.AddCommand("apt-get update")
d.AddCommand("apt-get install -y apt-transport-https curl docker.io")
d.AddCommand("curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -")
d.AddCommand("echo 'deb https://apt.kubernetes.io/ kubernetes-xenial main' > /etc/apt/sources.list.d/kubernetes.list")
d.AddCommand("apt-get update")
d.AddCommand("apt-get install -y kubelet kubeadm kubectl")
d.AddCommand("apt-mark hold kubelet kubeadm kubectl")
d.AddCommand("swapoff -a")
return d
}

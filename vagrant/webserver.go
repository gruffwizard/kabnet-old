package vagrant


import (

  "github.com/gruffwizard/kabnet/cluster"


)

// kubeadm config images pull
func genWebServerDeploy(c *cluster.Cluster) *cluster.Deploy {
  d := c.NewDeploy("web server")
  d.AddCommand("apt-get update")
  d.AddCommand("apt-get install -y apache2")
  return d

}

package vagrant


import (

  "github.com/gruffwizard/kabnet/cluster"
   "github.com/gruffwizard/kabnet/util"

)

// k

func genAPTCacheServer(c *cluster.Cluster) *cluster.Deploy {

    d := c.NewDeploy("apt cache")

    d.AddCommand("apt-get update")
    d.AddCommand("DEBIAN_FRONTEND=noninteractive apt-get install -yqq apache2 apt-cacher")
    d.AddCommand("mkdir -p /vagrant/cache/apt")

    cache:= util.CreateText()
    cache.Add("AUTOSTART=1")
    cache.Add("cache_dir=/vagrant/cache/apt")
    cache.Add("distinct_namespaces=1")
    cache.Add("generate_reports=1")
    cache.Add("allowed_hosts = *")

    d.AddFile("apt-cacher","/etc/default/apt-cacher",cache.AsString())
    d.AddCommand("service apache2 restart")
    d.AddCommand("service apt-cacher restart")


    return d

}

package vagrant

import (
  "log"
  "fmt"
  "github.com/gruffwizard/kabnet/cluster"
)


func HA_Generator(output string,masters int,workers int) {

    validate(&masters,&workers)
    log.Print("Vagrant: HA generator")
    log.Printf("Vagrant: masters =%d   workers = %d",masters,workers)

    c := cluster.NewCluster()

    c.Network.Domain="dev.kab"
    c.Network.Cluster="test"
    c.Network.NetworkPrefix="192.168.50"
    c.Network.SubnetMask="255.255.255.0/24"

    c.AddMachine("gateway")
    c.AddMachine("apt-cache")
    c.AddMachine("nameserver")
    c.AddMachine("proxy")

    for m := 0; m < masters; m++ {
      master := c.AddMachine("master")
      master.Name = fmt.Sprintf("master-%d",m)
    }

    for w := 0; w < workers; w++ {
      worker := c.AddMachine("worker")
      worker.Name = fmt.Sprintf("worker-%d",w)
    }

    // apply roles

    for _,m:= range c.Machines {

        m.SetHW(2,2)

        // everyone caches locally

        switch m.Type {
          case "proxy"      :  m.AddDeploy(genHaProxyDeployment(c))
                               m.AddDeploy(genSetHostName(c,m))
                               m.AddDeploy(genEnableAPTCachingDeploy(c,m))

          case "gateway"    :  m.AddDeploy(genGatewayDeploy(c))
                               m.AddDeploy(genDHCPServerDeploy(c))


          case "nameserver" :  m.AddDeploy(genEnableAPTCachingDeploy(c,m))
                               m.AddDeploy(genNameServerDeployment(c))

          case "apt-cache"  :  m.AddDeploy(genAPTCacheServer(c))
                               m.AddDeploy(genSetHostName(c,m))

          case "master"     :  m.AddDeploy(genEnableAPTCachingDeploy(c,m))
                               m.AddDeploy(genKubeDeploy(c))
                               if m.Name=="master-0" { m.AddDeploy(genMaster1Deploy(c,m))}
                               m.AddDeploy(genSetHostName(c,m))

          case "worker"     :
                               m.AddDeploy(genEnableAPTCachingDeploy(c,m))
                               m.AddDeploy(genKubeDeploy(c))
                               m.AddDeploy(genSetHostName(c,m))

          default: log.Panicf("unknown type %s",m.Type)

         }

    }




    c.Print()

    generate(output,c)

}





func generate(path string,c *cluster.Cluster) {

  writeVagrantFile(path,c)
  c.WriteDeployment(path)

}

func validate(masters *int,workers *int) {


    if *masters < 3  {
      *masters = 3
      log.Print(".Adjusted masters count to 3 ")
    }

    if *workers < 3  {
      *workers = 3
      log.Print(".Adjusted workers count to 3 ")

    }

}

package cluster

import (
  "fmt"
)

type Network struct {
  Domain string  `json:"domain"`
  Cluster string `json:"subdomain"`
//  Subnet string  `json:"subnet"`
  Gateway string `json:"gateway"`
  NetworkPrefix string `json:"netprefix"`
  SubnetMask string
//  IPP string
}


type Cluster struct {
    Network Network
    Machines  []*Machine
}

type HWDef struct {
    CPUs int
    RAM  int
}

type Machine struct {
  ID int
  cluster *Cluster
  Name string
  Type string
  IP   string
  MacAddr  string
  SSHPort int
  HWDef HWDef
  deploy *Deploy

}

type Deploy struct {
    Name string
    steps [] interface{}
}


type File struct {
  LocalName string
  Path string
  Content string
  Replace bool
  Local bool
}

func NewCluster() *Cluster {

    c := new(Cluster)
    return c

}

func (c *Cluster) NewDeploy(name string) *Deploy {

  var d Deploy
  d.Name=name
  return &d
}

func (c *Cluster) AddMachine(mtype string) *Machine {

    m := new(Machine)
    m.cluster = c
    m.Type=mtype
    m.Name=mtype
    m.ID=len(c.Machines)+1
    m.IP=fmt.Sprintf("%s.%d",c.Network.NetworkPrefix,m.ID)
    m.SSHPort=1200+m.ID
    m.MacAddr=fmt.Sprintf("10002781%04d",m.ID)
    m.deploy=c.NewDeploy("1st")
    c.Machines=append(c.Machines,m)
    return m
}

package cluster

import (
  "fmt"

)

func  (c *Cluster) FindMachine(n string) *Machine {
  for _,m := range c.Machines {
    if m.Name==n  { return m}
  }
  return nil
}

func (c *Cluster) Nameserver() *Machine {
  for _,m := range c.Machines {
    if m.Type=="nameserver"  { return m}
  }
  return nil
}


func (c *Cluster) Gateway() *Machine {
  for _,m := range c.Machines {
    if m.Type=="gateway"  { return m}
  }
  return nil
}


func (c *Cluster) Print() {

    for i,m := range c.Machines {
      hw:=fmt.Sprintf("CPUs %d, RAM %dGb",m.HWDef.CPUs,m.HWDef.RAM)
      fmt.Printf("%2d %-15s %s (%s)\n",i,m.Name,m.IP,hw)
    }
}


func (c *Cluster)  WriteDeployment(localdir string) {

    for _,m :=range c.Machines {
        m.WriteDeployment(localdir)
    }
}

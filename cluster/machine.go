package cluster


import (
  "log"
  "os"
    "github.com/gruffwizard/kabnet/util"
)


func (m *Machine) Fqdn() string {

  return m.Cqdn()+"."+m.Cluster().Network.Domain

}


func (m *Machine) Cluster() *Cluster {

  return m.cluster

}


func (m *Machine) Cqdn() string {

  return m.Name+"."+m.Cluster().Network.Cluster

}

func (m *Machine) AddDeploy(d *Deploy) {

    if m.deploy==nil {
      m.deploy=d
    } else {
    m.deploy.Add(d)
    }

}

func (m *Machine)  Deploy() *Deploy {
  return  m.deploy
}

func (m *Machine)  HasHostActions() bool {
  if  m.deploy == nil { return false}
  return len(m.deploy.steps) > 0
}

func (m Machine) MacAddrColon() string  {
  return util.ToForm(m.MacAddr,":")
}

func (m* Machine) SetHW(cpus int, ramgb int ) {
    m.HWDef.CPUs=2
    m.HWDef.RAM =2
}

func (m* Machine) WriteDeployment(localdir string) {

    d := m.Deploy()

    path := localdir+"/machines/"+m.Name
    os.MkdirAll(path, os.ModePerm)

    cmdscript, err := os.Create(path+"/kabnet.sh")
    if err!=nil { log.Fatal(err)}
    defer cmdscript.Close()

    d.Write(cmdscript,m,path)


}

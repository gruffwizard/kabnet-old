package cluster

import (
  "fmt"
  "log"
    "github.com/gruffwizard/kabnet/util"
  "os"
)



func (d *Deploy) Add(n *Deploy) {

    if n.Name==""  {log.Panic("missing deploy name")}

    log.Printf("add deploy %s",n.Name)
    d.addStep(n)

}



func (d *Deploy) AddFile(l string, p string,c string) {
  var f File
  f.LocalName=l
  f.Path=p
  f.Content=c
  d.addStep(f)
}



func (f *Deploy) AddCommand(p string,args ...interface {}) {
  c := fmt.Sprintf(p,args...)
  f.addStep(c)

}


func (f *Deploy) addStep(c interface{}) {

  f.steps=append(f.steps,c)
}


func (f *Deploy) Print() {

  log.Printf("Deploy stage:%s",f.Name)

  for _,step := range f.steps {

        s, ok := step.(File)
        if ok { log.Print("Deploy file:"+s.LocalName+" -> "+s.Path) }
        t, ok := step.(string)
        if ok { log.Print("Deploy cmd :",t) }
        d, ok := step.(*Deploy)
        if ok { d.Print() }

  }

}


func (d *Deploy) Write(cmdscript *os.File,m *Machine,path string) {

  cmdscript.WriteString("# step "+d.Name+"\n")

  for _,step := range d.steps {

    s, ok := step.(File)
    if ok {
      cmdscript.WriteString("cp /vagrant/machines/"+m.Name+"/"+s.LocalName+" "+s.Path+"\n")
      util.WriteFile(path,s.LocalName,s.Content)

    }
    q, ok := step.(string)
    if ok {
      cmdscript.WriteString(q+"\n")

     }
    dp, ok := step.(*Deploy)
    if ok {
      dp.Write(cmdscript,m,path)

    }
    

  }
}

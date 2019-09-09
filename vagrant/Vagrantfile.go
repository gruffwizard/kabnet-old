package vagrant

import (
  "github.com/gruffwizard/kabnet/util"
  "github.com/gruffwizard/kabnet/cluster"
)

func writeVagrantFile(outdir string, c *cluster.Cluster) {

  //cluster *model.Cluster
  f := util.CreateFile(outdir,"Vagrantfile")
  defer f.Close()

  util.Emit(f,"# -*- mode: ruby -*-")
  util.Emit(f,"# vi: set ft=ruby :")
  util.Emit(f,"Vagrant.configure('2') do |config|")

  util.Emit(f,"   config.vm.box=\"ubuntu/xenial64\"")
  util.Emit(f,"   config.vm.box_version =\"20190903.0\"")

  util.Emit(f,"   config.ssh.insert_key = true")



    for  _,m := range  c.Machines    {

      util.Emit(f," #")
      util.Emit(f," # Machine config for %s",m.Name)
      util.Emit(f," #")


      util.Emit(f,"   config.vm.provider \"virtualbox\" do |v|")
      util.Emit(f,"     v.customize [\"modifyvm\", :id, \"--memory\", %d]",m.HWDef.RAM*1024)
      util.Emit(f,"     v.customize [\"modifyvm\", :id, \"--cpus\", %d]",m.HWDef.CPUs)
      util.Emit(f,"   end ")


      util.Emit(f,"  config.vm.define  '%s' do |m|",m.Name)
      util.Emit(f,"   m.vm.hostname    = '%s'",m.Name)

      if m.Type=="gateway"  {
            util.Emit(f,"   m.vm.network 'private_network',  ip: '172.30.1.5'")
            util.Emit(f,"   m.vm.network 'private_network',  ip: '%s' , virtualbox__intnet: 'intnet'",m.IP)

            for _,pm := range c.Machines {
              util.Emit(f,"   m.vm.network 'forwarded_port', guest: %d, host: %d",pm.SSHPort,pm.SSHPort)
            }

      } else {
            util.Emit(f,"   m.ssh.host = '172.30.1.5'")
            util.Emit(f,"   m.ssh.guest_port = %d",m.SSHPort)
            util.Emit(f,"   m.vm.network 'private_network',  :adapter=>1, mac: '%s', virtualbox__intnet: 'intnet' , auto_config: false",m.MacAddr)

          }



      util.Emit(f,"   m.ssh.insert_key = false")
      if m.HasHostActions() {
        util.Emit(f,"   m.vm.provision 'shell', path: '%s/kabnet.sh'",outdir+"/machines/"+m.Name)
      }

      util.Emit(f,"  end")
    }

    util.Emit(f," end")

}

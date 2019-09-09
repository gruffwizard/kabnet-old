package vagrant

import (

  "github.com/gruffwizard/kabnet/cluster"
  "github.com/gruffwizard/kabnet/util"
  "strings"
)


func genNameServerDeployment(c *cluster.Cluster) *cluster.Deploy {

d := c.NewDeploy("dns")

d.AddCommand("apt-get update")
d.AddCommand("apt-get install -y bind9 resolvconf")
d.AddFile("bind9","/etc/default/bind9",
  `RESOLVCONF=no
  OPTIONS="-u bind -4"`)
d.AddCommand("systemctl restart bind9")

//d.AddCommand("echo 'nameserver "+c.Nameserver().IP+"' > /etc/resolvconf/resolv.conf.d/head")
//d.AddCommand("service resolvconf restart")

  options:= util.CreateText()
  options.Add("options {")
  options.Add("        directory \"/var/cache/bind\";")
  options.Add("        auth-nxdomain no;")
  options.Add("        listen-on port 53 { localhost; %s.0/24; };",c.Network.NetworkPrefix)
  options.Add("        allow-query { localhost; %s.0/24; };",c.Network.NetworkPrefix)
  options.Add("        forwarders { 8.8.8.8; };")
  options.Add("        recursion yes;")
  options.Add("        };")



  d.AddFile("named.conf.options","/etc/bind/named.conf.options",options.AsString())

  zone:= util.CreateText()
  zone.Add("zone \"%s\" {",c.Network.Domain)
  zone.Add("     type master;")
  zone.Add("     file \"/etc/bind/db.%s\";",c.Network.Domain)
  zone.Add("     allow-update { none; };")

  zone.Add("};")


  d.AddFile("named.conf.local","/etc/bind/named.conf.local",zone.AsString())

  text:= util.CreateText()
  text.Add("$TTL 1W")
  text.Add("@   IN  SOA     ns1.%s. root (",c.Network.Domain)
  text.Add("2019070700  ; serial")
  text.Add("3H          ; refresh (3 hours)")
  text.Add("30M         ; retry (30 minutes)")
  text.Add("2W          ; expiry (2 weeks)")
  text.Add("1W )        ; minimum (1 week)")
  text.Add(" IN     NS      ns1.%s." ,c.Network.Domain)
  text.Add(" IN     MX 10   smtp.%s.",c.Network.Domain)

  text.Add("%-25s   IN  A  %s","ns1",c.Nameserver().IP)
  text.Add("%-25s   IN  A  %s","smtp",c.Nameserver().IP)


  for _,m := range c.Machines {
      name:=m.Cqdn()
      text.Add("%-25s   IN  A  %s",name,m.IP)
  }

  //text.Add("*.apps.%s            IN      A       %s",c.Network.Cluster,c.Router().IP)
  addRevEntries(text,c)


  d.AddFile("bind.db","/etc/bind/db."+c.Network.Domain,text.AsString())
  d.AddCommand("systemctl restart bind9")
  d.AddCommand("named-checkconf")
  d.AddCommand("named-checkzone "+c.Network.Domain+" /etc/bind/db."+c.Network.Domain)



return d

}

func genSetHostName(c *cluster.Cluster, m *cluster.Machine) *cluster.Deploy {

  md := c.NewDeploy("set host name")
  md.AddCommand("apt-get update")
  md.AddCommand("apt-get install -y resolvconf")
  md.AddCommand("echo 'nameserver "+c.Nameserver().IP+"' > /etc/resolvconf/resolv.conf.d/head")
  md.AddCommand("service resolvconf restart")
  md.AddCommand("hostnamectl set-hostname "+m.Fqdn()+" --static")
  md.AddCommand("hostname "+m.Fqdn())

  return md
}

func addRevEntries(text *util.Text, c *cluster.Cluster) {
  machines := c.Machines

  for _,mx := range machines {
    lastOctetA := strings.Split(mx.IP,".")
    lastOctet  :=  lastOctetA[len(lastOctetA)-1]
    text.Add("%3s IN PTR  %s",lastOctet,mx.Fqdn()) //id,c.Network.Domain)

  }
}

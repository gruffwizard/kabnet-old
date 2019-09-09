package vagrant

import (
  "strconv"
  "github.com/gruffwizard/kabnet/cluster"
  "github.com/gruffwizard/kabnet/util"
)

func genGatewayDeploy(c *cluster.Cluster) *cluster.Deploy {

  d := c.NewDeploy("gateway - ip tables")

  d.AddCommand("apt-get update")
  d.AddCommand("apt-get install -yqq isc-dhcp-server")
  d.AddCommand("export GWLAN=`ip route | grep "+c.Gateway().IP+" | cut -d' ' -f3`")
  d.AddCommand("export GWWAN=`ip route | grep default  | cut -d' ' -f5`")

  d.AddCommand("echo 'net.ipv4.ip_forward=1' >> /etc/sysctl.conf")
  d.AddCommand("echo 1 > /proc/sys/net/ipv4/ip_forward")
  d.AddCommand("iptables --table nat --append POSTROUTING --out-interface $GWWAN -j MASQUERADE")
  d.AddCommand("iptables --append FORWARD --in-interface $GWLAN  -j ACCEPT")

    for _,pm := range c.Machines {
     d.AddCommand("iptables -t nat -A PREROUTING -p tcp -d "+"172.30.1.5" +
                  " --dport "+strconv.Itoa(pm.ID+1200)+" -j DNAT --to-destination "+
                  pm.IP+":22")
    }
  return d
}

func genDHCPServerDeploy(c *cluster.Cluster) *cluster.Deploy {


      d := c.NewDeploy("dhcp server")

      text:= util.CreateText()

      text.Add("#configuration file for ISC dhcpd")
      text.Add("ddns-update-style none;")
      text.Add("option domain-name \"%s\";",c.Network.Domain)
      text.Add("option bootfile-name \"pxelinux.0\";")
      text.Add("option domain-name-servers 8.8.8.8,%s;",c.Nameserver().IP)
      text.Add("default-lease-time 3600;")
      text.Add("max-lease-time 7200;")
      text.Add("authoritative;")
      text.Add("subnet %s.0 netmask 255.255.255.0 {",c.Network.NetworkPrefix)
      text.Add("    option routers %s.1;",c.Network.NetworkPrefix)
      text.Add("    option subnet-mask 255.255.255.0;")
      text.Add("    range %s.100 %s.200;",c.Network.NetworkPrefix,c.Network.NetworkPrefix)
      text.Add("}")

    for _,mac := range c.Machines {
      text.Add("host %s {",mac.Name)
      text.Add("  hardware ethernet %s;",mac.MacAddrColon())
      text.Add("  fixed-address %s;",mac.IP)
      text.Add("}")
    }

        d.AddFile("dhcp.conf","/etc/dhcp/dhcpd.conf",text.AsString())

        d.AddCommand("dhcpd -t")
        d.AddCommand("export DHLAN=`ip route | grep "+c.Gateway().IP+" | cut -d' ' -f3`")
        d.AddCommand("echo \"INTERFACES=$DHLAN\" > /etc/default/isc-dhcp-server")
        d.AddCommand("sudo systemctl restart isc-dhcp-server.service")

    return d

}

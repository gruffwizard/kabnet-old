package vagrant


import (

    "github.com/gruffwizard/kabnet/cluster"
    "github.com/gruffwizard/kabnet/util"

)

func addMainServerEntries(text *util.Text,c *cluster.Cluster) {

  for _,m := range c.Machines {
      if m.Type=="master" {
        text.Add("    server %s %s:6443 check",m.Name,m.IP)
      }
  }

}

func addWorkerEntries(text *util.Text,c *cluster.Cluster) {

  for _,m := range c.Machines {
      if m.Type=="worker"  {
        text.Add("    server %s %s:6443 check",m.Name,m.IP)
      }
  }

}


func genHaProxyDeployment(c *cluster.Cluster) *cluster.Deploy {

text:= util.CreateText()

text.Add(`global
        log /dev/log    local0
        log /dev/log    local1 notice
        chroot /var/lib/haproxy
        stats socket /run/haproxy/admin.sock mode 660 level admin
        stats timeout 30s
        user haproxy
        group haproxy
        daemon

        # Default SSL material locations
        ca-base /etc/ssl/certs
        crt-base /etc/ssl/private

        # Default ciphers to use on SSL-enabled listening sockets.
        # For more information, see ciphers(1SSL). This list is from:
        #  https://hynek.me/articles/hardening-your-web-servers-ssl-ciphers/
        ssl-default-bind-ciphers ECDH+AESGCM:DH+AESGCM:ECDH+AES256:DH+AES256:ECDH+AES128:DH+AES:ECDH+3DES:DH+3DES:RSA+AESGCM:RSA+AES:RSA+3DES:!aNULL:!MD5:!DSS
        ssl-default-bind-options no-sslv3

`)

text.Add(`defaults
        log     global
        mode    http
        option  httplog
        option  dontlognull
        timeout connect 5000
        timeout client  50000
        timeout server  50000
        errorfile 400 /etc/haproxy/errors/400.http
        errorfile 403 /etc/haproxy/errors/403.http
        errorfile 408 /etc/haproxy/errors/408.http
        errorfile 500 /etc/haproxy/errors/500.http
        errorfile 502 /etc/haproxy/errors/502.http
        errorfile 503 /etc/haproxy/errors/503.http
        errorfile 504 /etc/haproxy/errors/504.http`)


text.Add("listen stats")
text.Add("    bind :9000")
text.Add("    mode http")
text.Add("    stats enable")
text.Add("    stats uri /")
text.Add("    monitor-uri /healthz")

text.Add("frontend openshift-api-server")
text.Add("    bind *:6443")
text.Add("    default_backend openshift-api-server")
text.Add("    mode tcp")
text.Add("    option tcplog")

text.Add("backend openshift-api-server")
text.Add("    balance source")
text.Add("    mode tcp")

addMainServerEntries(text,c)

text.Add("frontend machine-config-server")
text.Add("    bind *:22623")
text.Add("    default_backend machine-config-server")
text.Add("    mode tcp")
text.Add("    option tcplog")

text.Add("backend machine-config-server")
text.Add("    balance source")
text.Add("    mode tcp")

addMainServerEntries(text,c)

text.Add("frontend ingress-http")
text.Add("    bind *:80")
text.Add("    default_backend ingress-http")
text.Add("    mode tcp")
text.Add("    option tcplog")

text.Add("backend ingress-http")
text.Add("    balance source")
text.Add("    mode tcp")

addWorkerEntries(text,c)

text.Add("frontend ingress-https")
text.Add("    bind *:443")
text.Add("    default_backend ingress-https")
text.Add("    mode tcp")
text.Add("    option tcplog")

text.Add("backend ingress-https")
text.Add("    balance source")
text.Add("    mode tcp")

addWorkerEntries(text,c)

d := c.NewDeploy("proxy")

d.AddCommand("apt-get update")
d.AddCommand("apt-get install haproxy -y")
d.AddFile("haproxy.conf","/etc/haproxy/haproxy.cfg",text.AsString())

return d
}

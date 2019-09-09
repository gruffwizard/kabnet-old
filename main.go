
package main

import (

  "github.com/gruffwizard/kabnet/cmd"

)


var (
    VERSION = "latest"
)

func main() {
  cmd.Version=VERSION
  cmd.Execute()
}

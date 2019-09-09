
package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"

)

var Version string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "kabnet",
  Short: "Kubenetes cluster generator using kubeadm",
  Long: `Generates deployment config and related scripts to
         deploy kubernetes cluster based on kubeadm`,

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  rootCmd.AddCommand(versionCmd)
}


var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "version",
  Long:  "version",

 Run: func(cmd *cobra.Command, args []string) {
   fmt.Printf("kabnet version %s\n",Version)
}, }

func initConfig() {

}

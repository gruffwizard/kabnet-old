package cmd

import (

  "github.com/spf13/cobra"
  "github.com/gruffwizard/kabnet/vagrant"
  "fmt"
)

var PMasters string
var PWorkers string
var POutputDir string

var vagrantCmd = &cobra.Command{
  Use:   "vagrant",
  Short: "vagrant",
  Long: "",
}

var generateCmd = &cobra.Command{
  Use:   "gen",
  Short: "generage",
  Long: "generate vagrant based deployment",
}

var haCmd = &cobra.Command{
  Use:   "ha",
  Short: "high-availabilty",
  Long: "generate high availability deployment",
 Run: func(cmd *cobra.Command, args []string) {

    workers :=  validateAsNumber(PWorkers,"invalid value for -w option")
    masters :=  validateAsNumber(PMasters,"invalid value for -m option")
    validateAsExistingDirectory(POutputDir,"invalid or missing value for -o option")

    vagrant.HA_Generator(POutputDir,masters,workers)
  },
}

var simpleCmd = &cobra.Command{
  Use:   "min",
  Short: "minimal",
  Long: "generate minimal deployment",
 Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("minimal vagrant generation")
  },
}


var deployCmd = &cobra.Command{
  Use:   "deploy",
  Short: "deploy",
  Long: "deploy kubernetes using vagrant deployment",
 Run: func(cmd *cobra.Command, args []string) {

  },
}

func init() {
        rootCmd.AddCommand(vagrantCmd)

        vagrantCmd.AddCommand(generateCmd)
        vagrantCmd.AddCommand(deployCmd)

        generateCmd.AddCommand(haCmd)
        haCmd.Flags().StringVarP(&PMasters, "masters", "m", "", "number of masters in cluster")
        haCmd.Flags().StringVarP(&PWorkers, "workers", "w", "", "number of workers in cluster")
        haCmd.Flags().StringVarP(&POutputDir, "output", "o", "", "output directory")

        generateCmd.AddCommand(simpleCmd)
        simpleCmd.Flags().StringVarP(&PMasters, "masters", "m", "", "number of masters in cluster")
        simpleCmd.Flags().StringVarP(&PWorkers, "workers", "w", "", "number of workers in cluster")
        simpleCmd.Flags().StringVarP(&POutputDir, "output", "o", "", "output directory")


}

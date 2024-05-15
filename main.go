package main

import "github.com/kaliwin/Needle/network/dns"

func main() {
	//var rootCmd = &cobra.Command{
	//	Use:     "myapp",
	//	Short:   "My sample Cobra app",
	//	Long:    "This is a sample app to demonstrate Cobra.",
	//	Version: "1.0.0",
	//}
	//
	////var bi = &cobra.Command{
	////	Use:   "version",
	////	Short: "Print the version number of Hugo",
	////	Long:  `All software has versions. This is Hugo's`,
	////	Run: func(cmd *cobra.Command, args []string) {
	////		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	////	},
	////}
	//
	////rootCmd.AddCommand(bi)
	////rootCmd.RemoveCommand(versionCmd) // 移除命令
	//rootCmd.CompletionOptions.DisableDefaultCmd = true
	//
	//rootCmd.Execute()

	dns.ServeDNS(":53", ".", "192.168.3.108")

}

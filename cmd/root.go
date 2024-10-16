// main.go
package letter

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config represents the application's configuration
// type Config struct {
// 	// Define your config fields here
// 	FNAME  string // full path filename in use
// 	Server string
// 	Caz    string
// }

var Version = "v0.1.1"
var cfgFile string
var collectionName string
var collectionFilter string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "letter",
	Short: "Newsletter publishing from the cli",
	Long:  "Publish newsletters using a simple server and the command line",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Letter> what do you want to do?")
		fmt.Println("VERSION>", Version)
		fmt.Println("CONFIG>", cfgFile)
		fmt.Println("CONFIG>", viper.GetString("pocketbase.address"))
		if collectionName != "" {
			fmt.Println("COLLECTION>", collectionName)
			fmt.Println("FILTER>", collectionFilter)	
		} else {
			fmt.Println("TEST>", viper.GetString("test.email"))
		}

		// fmt.Println("COLLECTION (viper)>", viper.GetString("pocketbase.collection"))
		// fmt.Println("FILTER (viper)>", viper.GetString("pocketbase.filter"))
		// fmt.Println("CONFIG>", config.FNAME)
		// fmt.Println("CONFIG>", config.Server)
		// fmt.Println("CONFIG>", config.Caz)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&config.FNAME, "config", "", "config file (default is ./letter.toml)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./letter.toml)")
	rootCmd.PersistentFlags().StringVarP(&collectionName, "collection", "c", "",  "collection name")
	rootCmd.PersistentFlags().StringVarP(&collectionFilter, "filter", "f", "", "collection filter")
	// sendCmd.Flags().StringVarP(&collectionFlag, "collection", "c", "", "email addresses collection name")
	// Add other persistent flags here

	// Bind the persistent flags to the config struct fields
	// viper.BindPFlag("pocketbase.collection", rootCmd.PersistentFlags().Lookup("collection"))
	viper.BindPFlag("pocketbase.filter", rootCmd.PersistentFlags().Lookup("filter"))
	viper.BindPFlag("pocketbase.collection", rootCmd.PersistentFlags().Lookup("collection"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		// viper.SetConfigFile(config.FNAME)
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in home directory with name ".yourapp" (without extension).
		viper.AddConfigPath("./")
		viper.SetConfigName("letter")
	}

	// viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		cfgFile=viper.ConfigFileUsed()
		// Optionally unmarshal the config into a struct
		// err := viper.Unmarshal(&config)
		// if err != nil {
		// 	fmt.Println("ERR>", err)
		// 	return
		// }
		// collectionName   = viper.GetString("pocketbase.collection")
		collectionFilter = viper.GetString("pocketbase.filter")
	}
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}

func GetFilter() string {
	return collectionFilter
}

func GetCollection() string {
	return collectionName
}

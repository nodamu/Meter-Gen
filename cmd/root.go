/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var startdate int64

var meterid string

var mqaddress string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "meter-gen",
	Short: "A smart meter value gen app",
	Long:  "A smart meter value gen app",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		meterid, _ := cmd.Flags().GetString("meterid")
		startdate, _ := cmd.Flags().GetInt64("startdate")
		mqaddress, _ := cmd.Flags().GetString("mqaddress")

		for {

			// genMeterData(startdate, meterid)

			Fire(startdate, meterid, mqaddress)

			//Sleep for 5 minutes
			time.Sleep(time.Second * 5)

			startdate = startdate + 5
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().Int64Var(&startdate, "startdate", 0, "unix timestamp for startdate for meter timestamp eg. 1624309765")
	rootCmd.PersistentFlags().StringVar(&meterid, "meterid", "", "meter id")
	rootCmd.PersistentFlags().StringVar(&mqaddress, "mqaddress", "", "Rabbit connection string eg. amqp://guest:guest@localhost:5672/ ")

	// // Cobra also supports local flags, which will only run
	// // when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rand.Seed(time.Now().UnixNano())
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".meter-gen" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".meter-gen")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

type MeterPayload struct {
	Energy    float32   `json:"energy_kwh"`
	TimeStamp time.Time `json:"timestamp"`
	MeterId   string    `json:"smart_meter_id"`
}

func genMeterData(startTime int64, meter_id string) MeterPayload {
	// test_uuid, err := uuid.NewRandom()
	// layout := "2006-01-02 15:04"
	start_time := time.Unix(startTime, 0)

	payload := MeterPayload{
		Energy:    randomFloat32(),
		TimeStamp: start_time,
		MeterId:   meter_id,
	}

	// // Marshall json into bytes
	// meter_payload, err := json.Marshal(payload)
	// if err != nil {
	// 	fmt.Errorf("Could not marshall struct into json")
	// }
	// fmt.Print(string(meter_payload))

	return payload
}

func randomFloat32() float32 {
	return 0.0 + rand.Float32()*(1-0.0)
}

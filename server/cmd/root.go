// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/valasek/timesheet/server/logger"
	"github.com/valasek/timesheet/server/version"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "timesheet",
	Version: version.Version,
	Short:   "Web based timesheet application with DB persistence",
	Long: `Web based timesheet application with DB persistence.
	
Application reads DB and server configuration from timesheet.toml, loads default data if DB is empty and launch web GUI.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetVersionTemplate(rootCmd.Version)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./timesheet.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("version", "v", false, "Prints application versions")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// get current folder
		curDir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in current directory "./" (without extension).
		viper.AddConfigPath(curDir)
		viper.SetConfigName("timesheet")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("config file:      ", viper.ConfigFileUsed())
	}

	initLogger()
	logger.Log.Info("config file ", viper.ConfigFileUsed())
}

func initLogger() {
	formatter := nested.Formatter{
		HideKeys:       true,
		NoColors:       true,
		NoFieldsColors: true,
		FieldsOrder:    []string{"component", "category"},
	}

	logger.Log.SetFormatter(&formatter)
	logger.Log.SetLevel(logrus.InfoLevel)
	logFolder := viper.GetString("logFolder")

	hook, err := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename: filepath.Join(".", logFolder, "general.log"),
			MaxSize:  100,
		},
		logrus.InfoLevel,
		&formatter,
		&lumberjackrus.LogFileOpts{
			logrus.InfoLevel: &lumberjackrus.LogFile{
				Filename: filepath.Join(".", logFolder, "info.log"),
				MaxSize:  100,
			},
			logrus.WarnLevel: &lumberjackrus.LogFile{
				Filename: filepath.Join(logFolder, "error.log"),
				MaxSize:  100,
			},
			logrus.ErrorLevel: &lumberjackrus.LogFile{
				Filename: filepath.Join(logFolder, "error.log"),
				MaxSize:  100,
			},
		},
	)

	if err != nil {
		panic(err)
	}
	logger.Log.AddHook(hook)
	logger.Log.Info("log folder ", logFolder)
}

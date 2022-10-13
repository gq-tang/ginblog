package cmd

import (
	"bytes"
	"io/ioutil"

	"ginblog/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var version string

var rootCmd = &cobra.Command{
	Use:   "blog",
	Short: "blog is simple blog site",
	Long:  "blog is simple blog,it use gin frame work.",
	RunE:  run,
}

func Execute(v string) {
	version = v
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "path to configuration file (optional)")
	rootCmd.PersistentFlags().Int("log_level", 4, "debug=5, info=4, error=2, fatal=1, panic=0")

	viper.BindPFlag("general.log_level", rootCmd.PersistentFlags().Lookup("log_level"))
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(configCmd)
}

func initConfig() {
	if cfgFile != "" {
		b, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("err loading config file")
		}
		viper.SetConfigType("toml")
		if err := viper.ReadConfig(bytes.NewBuffer(b)); err != nil {
			log.WithError(err).WithField("config", cfgFile).Fatal("err loading config file")
		}
	} else {
		viper.SetConfigName("gin-blog")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.config/gin-blog")
		viper.AddConfigPath("/etc/gin-blog")
		if err := viper.ReadInConfig(); err != nil {
			switch err.(type) {
			case viper.ConfigFileNotFoundError:
				log.Warning("Deprecation warning! no configuration file found,falling back on environment variables.")
			default:
				log.WithError(err).Fatal("read configuration file  error")
			}
		}
	}

	if err := viper.Unmarshal(&config.C); err != nil {
		log.WithError(err).Fatal("unmarshal config error")
	}
	if config.C.General.UploadPath == "" {
		log.Warn("the upload file path is not set,will set by default[../uploadfile/]")
		config.C.General.UploadPath = "../uploadfile/"
	}
}

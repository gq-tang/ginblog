package cmd

import (
	"github.com/gq-tang/ginblog/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

const configTemplate = `[general]
# Log level
#
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level={{ .General.LogLevel }} 
# server port
port={{ .General.Port }}
# server mode
# options:debug, release
mode={{ .General.Mode }}
[mysql]
# data source name
dsn={{ .MySQL.DSN }}
#
auto_migrate={{ .MySQL.AutoMigrate }} `

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print the blog server configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, config.C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}

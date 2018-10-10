package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gq-tang/ginblog/config"
	"github.com/gq-tang/ginblog/migrations"
	"github.com/gq-tang/ginblog/routers"
	"github.com/gq-tang/ginblog/storage"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func run(cmd *cobra.Command, args []string) error {
	tasks := []func() error{
		setLogLevel,
		setServerMode,
		setMySQLConnection,
		runDatabaseMigrations,
	}

	for _, f := range tasks {
		if err := f(); err != nil {
			log.Fatal(err)
		}
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.C.General.Port),
		Handler: routers.Engine(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}

func setMySQLConnection() error {
	log.Info("connecting mysql.")
	db, err := storage.OpenDB("mysql", config.C.MySQL.DSN)
	if err != nil {
		return err
	}
	config.C.MySQL.DB = db
	return nil
}

func runDatabaseMigrations() error {
	if config.C.MySQL.AutoMigrate {
		log.Info("applying database migrations")
		m := &migrate.AssetMigrationSource{
			Asset:    migrations.Asset,
			AssetDir: migrations.AssetDir,
			Dir:      "",
		}
		n, err := migrate.Exec(config.C.MySQL.DB.DB.DB, "mysql", m, migrate.Up)
		if err != nil {
			return errors.Wrap(err, "applying migrations error")
		}
		log.WithField("count", n).Info("migrations applied")
	}
	return nil
}

func setServerMode() error {
	if strings.ToLower(config.C.General.Mode) == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	return nil
}

func setLogLevel() error {
	if runtime.GOOS == "windows" {
		log.SetFormatter(&log.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2006/01/02 15:04:05",
		})
	}
	log.SetLevel(log.Level(config.C.General.LogLevel))
	return nil
}

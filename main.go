package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "embed"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/kkgo-software-engineering/workshop/router"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

//go:embed db/01-init-pocket.sql
var Sql_01_init_pocket string

func initCoundPocketTable(db *sql.DB) {
	if _, err := db.Exec(Sql_01_init_pocket); err != nil {
		log.Fatal("can't create table ", err)
	}
}

func main() {
	cfg := config.New().All()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		logger.Fatal("unable to configure database", zap.Error(err))
	}

	initCoundPocketTable(sql)

	e := router.RegRoute(cfg, logger, sql)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Hostname, cfg.Server.Port)

	go func() {
		err := e.Start(addr)
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("unexpected shutdown the server", zap.Error(err))
		}
		logger.Info("gracefully shutdown the server")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	gCtx := context.Background()
	ctx, cancel := context.WithTimeout(gCtx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal("unexpected shutdown the server", zap.Error(err))
	}
}

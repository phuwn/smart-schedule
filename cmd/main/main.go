package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/phuwn/smart-schedule/pkg/handler"
	"github.com/phuwn/tools/log"
	"github.com/phuwn/tools/util/db"
)

func main() {
	db.Start()
	defer db.Close()

	addr := ":996"
	r := handler.Router()
	errs := make(chan error)

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		c := make(chan os.Signal)
		defer close(c)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		cancel()
	}()

	go func() {
		log.Status("listening on port%s", addr)
		errs <- http.ListenAndServe(addr, r)
	}()

	for err := range errs {
		if err != nil {
			log.Status("got terminated %s", err.Error())
			close(errs)
			return
		}
	}
}

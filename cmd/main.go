package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kozhamseitova/balance-service/internal/app"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to run initialize app: %s", err.Error())
	}

	errGroup, groupCtx := errgroup.WithContext(ctx)
	errGroup.Go(func() error {
		return app.Run(ctx)
	})
	errGroup.Go(func() error {
		<-groupCtx.Done()
		log.Println("stopping application...")
		return app.Stop(ctx)
	})

	if err := errGroup.Wait(); err != nil {
		fmt.Printf("exit reason: %s\n", err)
	}
}

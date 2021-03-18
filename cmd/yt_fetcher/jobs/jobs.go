package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hi20160616/yt_fetcher/internal/job"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		defer cancel()
		return start(ctx)
	})
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			defer cancel()
			os.Exit(1)
		case <-ctx.Done():
			defer cancel()
		}
		return ctx.Err()
	})

	if err := g.Wait(); err != nil {
		log.Printf("JOB: Update channels error: %v", err)
	}
}

func start(ctx context.Context) error {
	// Funs: updateChannels every determined time
	for {
		// Act on the hour
		if time.Now().Minute() == 0 {
			log.Println("Channels update Start ...")
			if err := job.UpdateChannels(false); err != nil {
				return err
			}
		}
	}
}

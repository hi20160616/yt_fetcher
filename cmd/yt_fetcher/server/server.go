package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	"github.com/hi20160616/yt_fetcher/internal/data"
	"github.com/hi20160616/yt_fetcher/internal/service"
	"golang.org/x/sync/errgroup"
)

func InitFetcherCase() *biz.FetcherCase {
	fetcherRepo := data.NewFetcherRepo()
	fetcherCase := biz.NewFetcherCase(fetcherRepo)
	return fetcherCase
}

func main() {
	opts := service.Options{Address: ":10000"}

	fc := InitFetcherCase()
	fservice := service.NewFetcherServer(fc)

	s := service.NewServer(opts)
	pb.RegisterYoutubeFetcherServer(s, fservice)

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		defer cancel()
		return s.Start(ctx)
	})
	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			defer cancel()
			s.Stop(ctx)
		case <-ctx.Done():
			defer cancel()
			s.Stop(ctx)
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("yt_fetcher server main error: %v", err)
	}
}

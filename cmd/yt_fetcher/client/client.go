package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		// log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewYoutubeFetcherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// get channel: pass test
	cid := "UCMUnInmOkrWN4gof9KlhNmQ"
	channel := &pb.Channel{Cid: cid}
	channel, err = c.GetSetCname(ctx, channel)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(channel.Cname)

	// pass test
	// channel, err = c.GetChannel(ctx, channel)
	// if err != nil {
	//         log.Println(err)
	// }
	// fmt.Println(channel)

	// get video: pass test
	// v, err := c.GetVideo(ctx, &pb.Video{Vid: "-2u6RirE7aI"})
	// if err != nil {
	//         log.Printf("c.GetVideo err: %+v", err)
	// }
	// fmt.Println(v)

	// get videos pass test
	// responses, err := c.GetVideos(ctx, &pb.Channel{Cid: "UCMUnInmOkrWN4gof9KlhNmQ"})
	// if err != nil {
	//         log.Printf("c. GetVideos err: %+v", err)
	// }
	//
	// videos := responses.GetVideos()
	// fmt.Println(videos)

	// r, err := c.GetVideos(ctx, &pb.Channel{Cid: "UCMUnInmOkrWN4gof9KlhNmQ"})
	// if err != nil {
	//         log.Fatalf("could not get videos: %v", err)
	// }
	// fmt.Println(r.Videos)
	// for _, e := range r.Videos {
	//         fmt.Println(e.Title)
	// }
}

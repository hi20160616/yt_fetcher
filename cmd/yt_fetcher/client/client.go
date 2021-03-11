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
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewYoutubeFetcherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// search videos: pass test
	vs := &pb.Videos{Keywords: []string{"老高"}}
	vs, err = c.SearchVideos(ctx, vs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(vs.Videos))
	fmt.Println(vs.Keywords)

	// get channel name: pass test
	// cid := "UCMUnInmOkrWN4gof9KlhNmQ"
	// channel := &pb.Channel{Id: cid}
	// channel, err = c.GetSetCname(ctx, channel)
	// if err != nil {
	//         log.Println(err)
	// }
	// fmt.Println(channel.Name)

	// get channel: pass test
	// cid := "UCMUnInmOkrWN4gof9KlhNmQ"
	// channel := &pb.Channel{Id: cid}
	// channel, err = c.GetChannel(ctx, channel)
	// if err != nil {
	//         log.Println(err)
	// }
	// fmt.Println(channel)

	// get video: pass test
	// v, err := c.GetVideo(ctx, &pb.Video{Id: "-2u6RirE7aI"})
	// if err != nil {
	//         log.Printf("c.GetVideo err: %+v", err)
	// }
	// fmt.Println(v)

	// get videos pass test
	// res, err := c.GetVideos(ctx, &pb.Channel{Id: "UCMUnInmOkrWN4gof9KlhNmQ"})
	// if err != nil {
	//         log.Printf("c. GetVideos err: %+v", err)
	// }
	//
	// videos := res.GetVideos()
	// fmt.Println(videos)

	// get channels
	// cs := &pb.Channels{}
	// res, err := c.GetChannels(ctx, cs)
	// if err != nil {
	//         log.Println(err)
	// }
	// for _, c := range res.Channels {
	//         fmt.Println(c)
	// }
}

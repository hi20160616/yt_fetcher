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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// get video
	v, err := c.GetVideo(ctx, &pb.Video{Id: "nyfAij5B9fM"})
	fmt.Println(v)

	// get videos
	// r, err := c.GetVideos(ctx, &pb.Channel{Url: "https://www.youtube.com/channel/UCCtTgzGzQSWVzCG0xR7U-MQ/videos", Name: "亮生活 / Bright Side"})
	// if err != nil {
	//         log.Fatalf("could not get videos: %v", err)
	// }
	// fmt.Println(r.Chan)
	// for _, e := range r.Videos {
	//         fmt.Println(e.Title)
	// }
}

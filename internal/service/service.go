package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/url"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedYoutubeFetcherServer
	*grpc.Server
	address string
	fc      *biz.FetcherCase
}

type Options struct {
	Address string
}

func NewServer(opts Options) *Server {
	return &Server{Server: grpc.NewServer(), address: opts.Address}
}

func NewFetcherServer(fc *biz.FetcherCase) pb.YoutubeFetcherServer {
	return &Server{fc: fc}
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	log.Printf("\ngrpc server start at: %s", s.address)
	return s.Serve(lis)
}

func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	log.Printf("grpc server gracefully stopped.")
	return nil
}

func (s *Server) GetVideo(ctx context.Context, in *pb.Video) (*pb.Video, error) {
	log.Println("Get info from Video and set it to Video")
	// in.Title = "Title get from url"
	// in.Id = "url from `in` pre-set"
	// in.Description = "get and set after fetch from url"
	// in.Author = "this is channel name that pre-set"
	// in.LastUpdated = in.GetLastUpdated() // "this timestamp is get from video url pre-setted"
	// return &pb.Video{
	//         Title:       in.Title,
	//         LastUpdated: in.LastUpdated,
	// }, nil
	v, err := s.fc.GetVideo(in)
	if err != nil {
		fmt.Printf("%+v", err)
		return nil, errors.WithMessagef(err, "service.Server.GetVideo err")
	}
	fmt.Println(v.Title)
	return v, nil
}

func (s *Server) GetVideos(ctx context.Context, in *pb.Channel) (*pb.Videos, error) {
	log.Println("Get videos")
	log.Println("fetch video links from channel")
	// dto -> do
	u, err := url.Parse(in.Url)
	if err != nil {
		return nil, err
	}
	f := &biz.Fetcher{Entrance: u}

	// call biz
	videos, err := s.fc.GetVideos(f)
	if err != nil {
		return nil, err
	}
	return &pb.Videos{Videos: videos}, nil
}

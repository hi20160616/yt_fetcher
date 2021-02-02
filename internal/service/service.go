package service

import (
	"context"
	"log"
	"net"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedYoutubeFetcherServer
	*grpc.Server
	address string
}

type Options struct {
	Address string
}

func NewServer(opts Options) *Server {
	return &Server{Server: grpc.NewServer(), address: opts.Address}
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

// TODO: test serve and client, implement methods for fetch video
func (s *Server) GetVideo(ctx context.Context, in *pb.Video) (*pb.Video, error) {
	log.Println("Get info from Video and set it to Video")
	in.Title = "Title get from url"
	in.Id = "url from `in` pre-set"
	in.Description = "get and set after fetch from url"
	in.Author = "this is channel name that pre-set"
	in.LastUpdated = in.GetLastUpdated() // "this timestamp is get from video url pre-setted"
	return &pb.Video{
		Title:       in.Title,
		LastUpdated: in.LastUpdated,
	}, nil
}

// TODO: implement methods for fetch videos
func (s *Server) GetVideos(ctx context.Context, in *pb.Channel) (*pb.Videos, error) {
	log.Println("Get videos")
	log.Println("fetch video links from channel")
	return &pb.Videos{Chan: &pb.Channel{Url: in.Url, Name: in.Name}}, nil
}

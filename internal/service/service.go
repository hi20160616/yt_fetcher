package service

import (
	"context"
	"log"
	"net"

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

// Get info from Video and set it to Video
func (s *Server) GetVideo(ctx context.Context, in *pb.Video) (*pb.Video, error) {
	log.Println("Get one video")
	v, err := s.fc.GetVideo(in)
	if err != nil {
		log.Printf("GetVideo err: %+v", err)
		return nil, errors.WithMessage(err, "service.Server.GetVideo err")
	}
	return v, nil
}

func (s *Server) GetVideoIds(ctx context.Context, in *pb.Channel) (*pb.Channel, error) {
	log.Println("Get video ids")
	s.fc.GetVideoIds(in)
	return in, nil
}

// Get videos' info
func (s *Server) GetVideos(ctx context.Context, in *pb.Channel) (*pb.Videos, error) {
	log.Println("Get videos")
	// call biz
	videos, err := s.fc.GetVideos(in)
	if err != nil {
		log.Printf("GetVideos parse url err: %+v", err)
		return nil, errors.WithMessage(err, "service.Server.GetVideos err")
	}
	return &pb.Videos{Videos: videos}, nil
}

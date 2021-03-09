package service

import (
	"context"
	"log"
	"net"

	pb "github.com/hi20160616/yt_fetcher/api/yt_fetcher/api"
	"github.com/hi20160616/yt_fetcher/internal/biz"
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

// GetVideo implements api.GetVideo, it get video's info and set it back to `in`
func (s *Server) GetVideo(ctx context.Context, in *pb.Video) (*pb.Video, error) {
	return s.fc.GetVideo(in)
}

// GetVideoIds implements api.GetVideoIds, it get videoIds by `in.Url` and set it back to `in`
func (s *Server) GetVideoIds(ctx context.Context, in *pb.Channel) (*pb.Channel, error) {
	return s.fc.GetVideoIds(in)
}

// GetVideos implements api.GetVideos, it get videos by `in.VideoIds` and set it back to `in`
func (s *Server) GetVideos(ctx context.Context, in *pb.Channel) (*pb.Videos, error) {
	return s.fc.GetVideos(in)
}

func (s *Server) GetSetCname(ctx context.Context, in *pb.Channel) (*pb.Channel, error) {
	return s.fc.GetChannelName(in)
}

func (s *Server) GetChannel(ctx context.Context, in *pb.Channel) (*pb.Channel, error) {
	return s.fc.GetChannel(in)
}

func (s *Server) GetChannels(ctx context.Context, in *pb.Channels) (*pb.Channels, error) {
	return s.fc.GetChannels(in)
}

func (s *Server) GetVideosFromTo(ctx context.Context, in *pb.Videos) (*pb.Videos, error) {
	return s.fc.GetVideosFromTo(in)
}

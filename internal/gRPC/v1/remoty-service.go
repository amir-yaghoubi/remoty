package service

import (
	"context"

	pb "github.com/amir-yaghoubi/remoty/pkg/api/v1"
	remoteSrv "github.com/amir-yaghoubi/remoty/pkg/service"
	log "github.com/sirupsen/logrus"
)

const apiVersion = "v1"

func New(logger *log.Logger, srv *remoteSrv.Service) pb.RemotyServiceServer {
	return &grpcService{logger, srv}
}

type grpcService struct {
	logger *log.Logger
	srv    *remoteSrv.Service
}

func (s *grpcService) AddToQueue(ctx context.Context, req *pb.Link) (*pb.Void, error) {
	url := req.GetUrl()
	log := s.logger.WithField("link", url)

	err := s.srv.AddToQueue(ctx, url)
	if err != nil {
		log.Error(err)
	}
	return &pb.Void{}, err
}
func (s *grpcService) Download(ctx context.Context, req *pb.Link) (*pb.Void, error) {
	url := req.GetUrl()
	log := s.logger.WithField("link", url)

	err := s.srv.Download(ctx, url)
	if err != nil {
		log.Error(err)
	}
	return &pb.Void{}, err
}
func (s *grpcService) StartDownload(ctx context.Context, req *pb.Void) (*pb.Void, error) {
	err := s.srv.StartQueue(ctx)

	if err != nil {
		s.logger.Error(err)
	}

	return &pb.Void{}, err
}

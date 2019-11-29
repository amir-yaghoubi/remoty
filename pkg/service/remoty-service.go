package service

import (
	"context"
	"errors"
	"net/url"

	"github.com/amir-yaghoubi/remoty/pkg/idm"
	log "github.com/sirupsen/logrus"
)

//ErrInvalidLink validation error for link
var ErrInvalidLink = errors.New("invalid link")

// New returns a Service instance
func New(logger *log.Logger, idmCtrl *idm.Controller) *Service {
	return &Service{logger, idmCtrl}
}

// Service remotly service
type Service struct {
	logger  *log.Logger
	idmCtrl *idm.Controller
}

func (s *Service) validateLink(link string) error {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return ErrInvalidLink
	}

	return nil
}

//AddToQueue add link to idm queue list
func (s *Service) AddToQueue(ctx context.Context, link string) error {
	err := s.validateLink(link)
	if err != nil {
		return err
	}

	return s.idmCtrl.AddToQueue(ctx, link)
}

//Download start downloading given link via idm
func (s *Service) Download(ctx context.Context, link string) error {
	err := s.validateLink(link)
	if err != nil {
		return err
	}

	return s.idmCtrl.Download(ctx, link)
}

//StartQueue start downloading idm queue list
func (s *Service) StartQueue(ctx context.Context) error {
	return s.idmCtrl.StartQueue(ctx)
}

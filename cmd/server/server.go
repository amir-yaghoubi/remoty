package main

import (
	"errors"
	"net"
	"os"

	"github.com/amir-yaghoubi/remoty/pkg/service"

	grpcSrv "github.com/amir-yaghoubi/remoty/internal/gRPC/v1"
	"github.com/amir-yaghoubi/remoty/pkg/api/v1"
	"github.com/amir-yaghoubi/remoty/pkg/idm"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	port := getPort()

	idmPath, err := getIdmPath()
	if err != nil {
		logger.Fatal(err)
	}

	idmController := idm.New(idmPath)
	remotyService := service.New(logger, idmController)

	srv := grpc.NewServer()
	grpc := grpcSrv.New(logger, remotyService)

	api.RegisterRemotyServiceServer(srv, grpc)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.WithField("port", port).Info("grpc server started")
	err = srv.Serve(listener)
	if err != nil {
		logger.Fatalln(err)
	}
}

func getPort() string {
	p := os.Getenv("RT_PORT")
	if len(p) == 0 {
		p = "1995"
	}
	return ":" + p
}

func getIdmPath() (string, error) {
	idm := os.Getenv("RT_IDM")
	if len(idm) == 0 {
		return "", errors.New("you have to set RT_IDM")
	}

	f, err := os.Stat(idm)
	if err != nil {
		return "", err
	}

	if f.IsDir() {
		return "", errors.New("please point RT_IDM to IDMan.exe")
	}

	return idm, nil
}

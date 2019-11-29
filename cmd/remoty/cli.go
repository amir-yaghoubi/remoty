package main

import (
	"net/url"
	"os"

	"github.com/amir-yaghoubi/remoty/cmd/remoty/idm"
	pb "github.com/amir-yaghoubi/remoty/pkg/api/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func main() {
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	addr, err := getRPCAddr()
	if err != nil {
		logger.Fatal(err)
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewRemotyServiceClient(conn)

	rootCmd := cobra.Command{
		Use:     "remoty",
		Short:   "remoty is a cli tool to communicate with remoty server",
		Version: "1.0.0",
	}

	idmCommandControl := idm.New(client, logger)

	rootCmd.AddCommand(idmCommandControl.Command())

	err = rootCmd.Execute()
	if err != nil {
		logger.Fatal(err)
	}

}

func getRPCAddr() (string, error) {
	addr := os.Getenv("RT_ADDR")
	if len(addr) == 0 {
		return "localhost:1995", nil
	}

	_, err := url.ParseRequestURI(addr)
	if err != nil {
		return "", err
	}

	return addr, nil
}

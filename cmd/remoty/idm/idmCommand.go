package idm

import (
	"context"
	"time"

	pb "github.com/amir-yaghoubi/remoty/pkg/api/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type cobraHandler = func(*cobra.Command, []string)

//New returns a new IdmCommandControl instance
func New(client pb.RemotyServiceClient, logger *log.Logger) *IdmCommandControl {
	return &IdmCommandControl{client, logger}
}

//IdmCommandControl is the internet download manager controller
type IdmCommandControl struct {
	client pb.RemotyServiceClient
	logger *log.Logger
}

func (ctrl *IdmCommandControl) addHandler() cobraHandler {
	return func(cmd *cobra.Command, args []string) {
		link := pb.Link{Url: args[0]}

		ctx := context.Background()
		ctx, cancelF := context.WithTimeout(ctx, time.Second*10)
		defer cancelF()

		_, err := ctrl.client.AddToQueue(ctx, &link)
		if err != nil {
			ctrl.logger.Error(err)
			return
		}
		ctrl.logger.Info("✅   Link added to the queue.")
	}
}

func (ctrl *IdmCommandControl) downloadHandler() cobraHandler {
	return func(cmd *cobra.Command, args []string) {
		link := pb.Link{Url: args[0]}

		ctx := context.Background()
		ctx, cancelF := context.WithTimeout(ctx, time.Second*10)
		defer cancelF()

		_, err := ctrl.client.Download(ctx, &link)
		if err != nil {
			ctrl.logger.Error(err)
			return
		}
		ctrl.logger.Info("✅   Start downloading...")
	}
}

func (ctrl *IdmCommandControl) startQueueHandler() cobraHandler {
	return func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		ctx, cancelF := context.WithTimeout(ctx, time.Second*10)
		defer cancelF()

		_, err := ctrl.client.StartDownload(ctx, &pb.Void{})
		if err != nil {
			ctrl.logger.Error(err)
			return
		}
		ctrl.logger.Info("✅   Start downloading...")
	}
}

//Command return cobra command for idm
func (ctrl *IdmCommandControl) Command() *cobra.Command {
	cmd := cobra.Command{
		Use:        "idm",
		Short:      "Internet download manager commands",
		SuggestFor: []string{"id", "im"},
	}

	addCmd := &cobra.Command{
		Use:  "add",
		Args: cobra.MinimumNArgs(1),
		Run:  ctrl.addHandler(),
	}

	downloadCmd := &cobra.Command{
		Use:  "download",
		Args: cobra.MinimumNArgs(1),
		Run:  ctrl.downloadHandler(),
	}

	startCmd := &cobra.Command{
		Use: "start",
		Run: ctrl.startQueueHandler(),
	}

	cmd.AddCommand(addCmd, downloadCmd, startCmd)
	return &cmd
}

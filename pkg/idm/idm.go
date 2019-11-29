/*
	Full documentation on IDM command line parameters available on:
	https://www.internetdownloadmanager.com/support/command_line.html

*/

package idm

import (
	"context"
	"os/exec"
)

// New returns an IdmController instance
func New(idmPath string) *Controller {
	return &Controller{idmPath: idmPath}
}

// Controller controls host IDM
type Controller struct {
	idmPath string
}

// AddToQueue Add given link to IDM main queue
func (c *Controller) AddToQueue(ctx context.Context, link string) error {
	cmd := exec.CommandContext(ctx, c.idmPath, "/a", "/d", link)
	return cmd.Run()
}

// Download Starts downloading given link
func (c *Controller) Download(ctx context.Context, link string) error {
	cmd := exec.CommandContext(ctx, c.idmPath, "/n", "/d", link)
	return cmd.Run()
}

// StartQueue Starts downloading main queue
func (c *Controller) StartQueue(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, c.idmPath, "/s")
	return cmd.Run()
}

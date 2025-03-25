package modules

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/mewmewlab/eki/pkg/utils"
	"github.com/pocketbase/pocketbase/core"
)

// Unit is the docker image used in eki task
type Unit struct {
	ID      string
	Image   string
	Version string

	ContainerConfig *container.Config
	HostConfig      *container.HostConfig
}

func NewUnit(image, version string) *Unit {
	if version == "" {
		version = "latest"
	}
	return &Unit{
		Image:   image,
		Version: version,
	}
}

func (u *Unit) String() string {
	return fmt.Sprintf("%s:%s", u.Image, u.Version)
}

func (u *Unit) Pull(app core.App) error {
	reader, err := utils.DockerClient.ImagePull(context.Background(), u.String(), image.PullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		return err
	}
	return nil
}

func (u *Unit) ImageInfo(ctx context.Context, app core.App) error {
	utils.DockerClient.ImageList(ctx, image.ListOptions{
		Filters: filters.NewArgs(filters.Arg("reference", u.String())),
	})
	return nil
}

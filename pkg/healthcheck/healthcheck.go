// Package healthcheck checks for missing binaries that the user needs to
// install.
package healthcheck

import (
	"fmt"

	"github.com/garethgeorge/freegoreleaser/internal/pipe/chocolatey"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/docker"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/nix"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sbom"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sign"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/snapcraft"
	"github.com/garethgeorge/freegoreleaser/pkg/context"
)

// Healthchecker should be implemented by pipes that want checks.
type Healthchecker interface {
	fmt.Stringer

	// Dependencies return the binaries of the dependencies needed.
	Dependencies(ctx *context.Context) []string
}

// Healthcheckers is the list of healthchekers.
//
//nolint:gochecknoglobals
var Healthcheckers = []Healthchecker{
	system{},
	snapcraft.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	chocolatey.Pipe{},
	nix.NewPublish(),
}

type system struct{}

func (system) String() string                           { return "system" }
func (system) Dependencies(_ *context.Context) []string { return []string{"git", "go"} }

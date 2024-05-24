// Package custompublishers provides a Pipe that executes a custom publisher
package custompublishers

import (
	"github.com/garethgeorge/freegoreleaser/internal/exec"
	"github.com/garethgeorge/freegoreleaser/pkg/context"
)

// Pipe for custom publisher.
type Pipe struct{}

func (Pipe) String() string                     { return "custom publisher" }
func (Pipe) Skip(ctx *context.Context) bool     { return len(ctx.Config.Publishers) == 0 }
func (Pipe) Publish(ctx *context.Context) error { return exec.Execute(ctx, ctx.Config.Publishers) }

package testlib

import (
	"testing"

	"github.com/garethgeorge/freegoreleaser/internal/pipe"
)

func TestAssertSkipped(t *testing.T) {
	AssertSkipped(t, pipe.Skip("skip"))
}

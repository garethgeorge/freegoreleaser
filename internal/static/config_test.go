package static

import (
	"bytes"
	"testing"

	"github.com/garethgeorge/freegoreleaser/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestExampleConfig(t *testing.T) {
	_, err := config.LoadReader(bytes.NewReader(ExampleConfig))
	require.NoError(t, err)
	require.NotEmpty(t, ExampleConfig)
}

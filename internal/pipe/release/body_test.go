package release

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/garethgeorge/freegoreleaser/internal/artifact"
	"github.com/garethgeorge/freegoreleaser/internal/golden"
	"github.com/garethgeorge/freegoreleaser/internal/testctx"
	"github.com/garethgeorge/freegoreleaser/internal/testlib"
	"github.com/garethgeorge/freegoreleaser/pkg/config"
	"github.com/garethgeorge/freegoreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

func TestDescribeBody(t *testing.T) {
	changelog := "feature1: description\nfeature2: other description"
	ctx := testctx.New()
	ctx.ReleaseNotes = changelog
	out, err := describeBody(ctx)
	require.NoError(t, err)

	golden.RequireEqual(t, out.Bytes())
}

func TestDontEscapeHTML(t *testing.T) {
	changelog := "<h1>test</h1>"
	ctx := testctx.New()
	ctx.ReleaseNotes = changelog

	out, err := describeBody(ctx)
	require.NoError(t, err)
	require.Contains(t, out.String(), changelog)
}

func TestDescribeBodyMultipleChecksums(t *testing.T) {
	ctx := testctx.NewWithCfg(
		config.Project{
			Release: config.Release{
				Header: "## Yada yada yada\nsomething\n",
				Footer: `
---

## Checksums

` + "```\n{{ range $key, $value := .Checksums }}{{ $value }} {{ $key }}\n{{ end }}```\n",
			},
		},
		testctx.WithCurrentTag("v1.0"),
		func(ctx *context.Context) { ctx.ReleaseNotes = "nothing" },
	)

	checksums := map[string]string{
		"bar.zip": "f674623cf1edd0f753e620688cedee4e7c0e837ac1e53c0cbbce132ffe35fd52",
		"foo.zip": "271a74b75a12f6c3affc88df101f9ef29af79717b1b2f4bdd5964aacf65bcf40",
	}
	for name, check := range checksums {
		checksumPath := filepath.Join(t.TempDir(), name+".sha256")
		ctx.Artifacts.Add(&artifact.Artifact{
			Name: name + ".sha256",
			Path: checksumPath,
			Type: artifact.Checksum,
			Extra: map[string]interface{}{
				artifact.ExtraChecksumOf: name,
				artifact.ExtraRefresh: func() error {
					return os.WriteFile(checksumPath, []byte(check), 0o644)
				},
			},
		})
	}

	require.NoError(t, ctx.Artifacts.Refresh())

	out, err := describeBody(ctx)
	require.NoError(t, err)

	golden.RequireEqual(t, out.Bytes())
}

func TestDescribeBodyWithHeaderAndFooter(t *testing.T) {
	changelog := "feature1: description\nfeature2: other description"
	ctx := testctx.NewWithCfg(
		config.Project{
			Release: config.Release{
				Header: "## Yada yada yada\nsomething\n",
				Footer: `
---

Get images at docker.io/foo/bar:{{.Tag}}

---

Get GoReleaser Pro at https://goreleaser.com/pro

---

## Checksums

` + "```\n{{ .Checksums }}\n```" + `
				`,
			},
		},
		testctx.WithCurrentTag("v1.0"),
		func(ctx *context.Context) { ctx.ReleaseNotes = changelog },
	)

	checksumPath := filepath.Join(t.TempDir(), "checksums.txt")
	checksumContent := "f674623cf1edd0f753e620688cedee4e7c0e837ac1e53c0cbbce132ffe35fd52  foo.zip"
	ctx.Artifacts.Add(&artifact.Artifact{
		Name: "checksums.txt",
		Path: checksumPath,
		Type: artifact.Checksum,
		Extra: map[string]interface{}{
			artifact.ExtraRefresh: func() error {
				return os.WriteFile(checksumPath, []byte(checksumContent), 0o644)
			},
		},
	})

	require.NoError(t, ctx.Artifacts.Refresh())

	out, err := describeBody(ctx)
	require.NoError(t, err)

	golden.RequireEqual(t, out.Bytes())
}

func TestDescribeBodyWithInvalidHeaderTemplate(t *testing.T) {
	ctx := testctx.NewWithCfg(config.Project{
		Release: config.Release{
			Header: "## {{ .Nop }\n",
		},
	})
	_, err := describeBody(ctx)
	testlib.RequireTemplateError(t, err)
}

func TestDescribeBodyWithInvalidFooterTemplate(t *testing.T) {
	ctx := testctx.NewWithCfg(config.Project{
		Release: config.Release{
			Footer: "{{ .Nops }",
		},
	})
	_, err := describeBody(ctx)
	testlib.RequireTemplateError(t, err)
}

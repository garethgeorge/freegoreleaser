// Package pipeline provides generic errors for pipes to use.
package pipeline

import (
	"fmt"

	"github.com/garethgeorge/freegoreleaser/internal/pipe/announce"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/archive"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/aur"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/before"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/brew"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/build"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/changelog"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/checksums"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/chocolatey"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/defaults"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/dist"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/docker"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/effectiveconfig"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/env"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/git"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/gomod"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/krew"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/metadata"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/nfpm"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/nix"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/notary"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/partial"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/prebuild"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/publish"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/reportsizes"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sbom"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/scoop"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/semver"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sign"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/snapcraft"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/snapshot"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sourcearchive"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/universalbinary"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/upx"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/winget"
	"github.com/garethgeorge/freegoreleaser/pkg/context"
)

// Piper defines a pipe, which can be part of a pipeline (a series of pipes).
type Piper interface {
	fmt.Stringer

	// Run the pipe
	Run(ctx *context.Context) error
}

// BuildPipeline contains all build-related pipe implementations in order.
//
//nolint:gochecknoglobals
var BuildPipeline = []Piper{
	// load and validate environment variables
	env.Pipe{},
	// get and validate git repo state
	git.Pipe{},
	// parse current tag to a semver
	semver.Pipe{},
	// load default configs
	defaults.Pipe{},
	// setup things for partial builds/releases
	partial.Pipe{},
	// snapshot version handling
	snapshot.Pipe{},
	// run global hooks before build
	before.Pipe{},
	// ensure ./dist is clean
	dist.Pipe{},
	// setup metadata options
	metadata.Pipe{},
	// creates a metadta.json files in the dist directory
	metadata.MetaPipe{},
	// setup gomod-related stuff
	gomod.Pipe{},
	// run prebuild stuff
	prebuild.Pipe{},
	// proxy gomod if needed
	gomod.CheckGoModPipe{},
	// proxy gomod if needed
	gomod.ProxyPipe{},
	// writes the actual config (with defaults et al set) to dist
	effectiveconfig.Pipe{},
	// build
	build.Pipe{},
	// universal binary handling
	universalbinary.Pipe{},
	// notarize macos apps
	notary.MacOS{},
	// upx
	upx.Pipe{},
}

// BuildCmdPipeline is the pipeline run by goreleaser build.
//
//nolint:gochecknoglobals
var BuildCmdPipeline = append(
	BuildPipeline,
	reportsizes.Pipe{},
	metadata.ArtifactsPipe{},
)

// Pipeline contains all pipe implementations in order.
//
//nolint:gochecknoglobals
var Pipeline = append(
	BuildPipeline,
	// builds the release changelog
	changelog.Pipe{},
	// archive in tar.gz, zip or binary (which does no archiving at all)
	archive.Pipe{},
	// archive the source code using git-archive
	sourcearchive.Pipe{},
	// archive via fpm (deb, rpm) using "native" go impl
	nfpm.Pipe{},
	// archive via snapcraft (snap)
	snapcraft.Pipe{},
	// create SBOMs of artifacts
	sbom.Pipe{},
	// checksums of the files
	checksums.Pipe{},
	// sign artifacts
	sign.Pipe{},
	// create arch linux aur pkgbuild
	aur.Pipe{},
	// create nixpkgs
	nix.NewBuild(),
	// winget installers
	winget.Pipe{},
	// create brew tap
	brew.Pipe{},
	// krew plugins
	krew.Pipe{},
	// create scoop buckets
	scoop.Pipe{},
	// create chocolatey pkg and publish
	chocolatey.Pipe{},
	// reports artifacts sizes to the log and to artifacts.json
	reportsizes.Pipe{},
	// create and push docker images
	docker.Pipe{},
	// publishes artifacts
	publish.New(),
	// creates a artifacts.json files in the dist directory
	metadata.ArtifactsPipe{},
	// announce releases
	announce.Pipe{},
)

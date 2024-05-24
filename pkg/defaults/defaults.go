// Package defaults make the list of Defaulter implementations available
// so projects extending GoReleaser are able to use it, namely, GoDownloader.
package defaults

import (
	"fmt"

	"github.com/garethgeorge/freegoreleaser/internal/pipe/archive"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/artifactory"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/aur"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/blob"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/bluesky"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/brew"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/build"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/changelog"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/checksums"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/chocolatey"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/discord"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/docker"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/gomod"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/ko"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/krew"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/linkedin"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/mastodon"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/mattermost"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/milestone"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/nfpm"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/nix"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/notary"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/opencollective"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/project"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/reddit"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/release"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sbom"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/scoop"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sign"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/slack"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/smtp"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/snapcraft"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/snapshot"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/sourcearchive"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/teams"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/telegram"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/twitter"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/universalbinary"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/upload"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/upx"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/webhook"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/winget"
	"github.com/garethgeorge/freegoreleaser/pkg/context"
)

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

// Defaulters is the list of defaulters.
//
//nolint:gochecknoglobals
var Defaulters = []Defaulter{
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	changelog.Pipe{},
	gomod.Pipe{},
	build.Pipe{},
	universalbinary.Pipe{},
	notary.MacOS{},
	upx.Pipe{},
	sourcearchive.Pipe{},
	archive.Pipe{},
	nfpm.Pipe{},
	snapcraft.Pipe{},
	checksums.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	artifactory.Pipe{},
	blob.Pipe{},
	upload.Pipe{},
	aur.Pipe{},
	nix.Pipe{},
	winget.Pipe{},
	brew.Pipe{},
	krew.Pipe{},
	ko.Pipe{},
	scoop.Pipe{},
	discord.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	teams.Pipe{},
	twitter.Pipe{},
	smtp.Pipe{},
	mastodon.Pipe{},
	mattermost.Pipe{},
	milestone.Pipe{},
	linkedin.Pipe{},
	telegram.Pipe{},
	webhook.Pipe{},
	chocolatey.Pipe{},
	opencollective.Pipe{},
	bluesky.Pipe{},
}

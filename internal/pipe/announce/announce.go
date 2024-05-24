// Package announce contains the announcing pipe.
package announce

import (
	"fmt"

	"github.com/garethgeorge/freegoreleaser/internal/middleware/errhandler"
	"github.com/garethgeorge/freegoreleaser/internal/middleware/logging"
	"github.com/garethgeorge/freegoreleaser/internal/middleware/skip"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/bluesky"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/discord"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/linkedin"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/mastodon"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/mattermost"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/opencollective"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/reddit"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/slack"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/smtp"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/teams"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/telegram"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/twitter"
	"github.com/garethgeorge/freegoreleaser/internal/pipe/webhook"
	"github.com/garethgeorge/freegoreleaser/internal/skips"
	"github.com/garethgeorge/freegoreleaser/internal/tmpl"
	"github.com/garethgeorge/freegoreleaser/pkg/context"
)

// Announcer should be implemented by pipes that want to announce releases.
type Announcer interface {
	fmt.Stringer
	Announce(ctx *context.Context) error
}

//nolint:gochecknoglobals
var announcers = []Announcer{
	// XXX: keep asc sorting
	bluesky.Pipe{},
	discord.Pipe{},
	linkedin.Pipe{},
	mastodon.Pipe{},
	mattermost.Pipe{},
	opencollective.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	smtp.Pipe{},
	teams.Pipe{},
	telegram.Pipe{},
	twitter.Pipe{},
	webhook.Pipe{},
}

// Pipe that announces releases.
type Pipe struct{}

func (Pipe) String() string { return "announcing" }

func (Pipe) Skip(ctx *context.Context) (bool, error) {
	if skips.Any(ctx, skips.Announce) {
		return true, nil
	}
	return tmpl.New(ctx).Bool(ctx.Config.Announce.Skip)
}

// Run the pipe.
func (Pipe) Run(ctx *context.Context) error {
	memo := errhandler.Memo{}
	for _, announcer := range announcers {
		_ = skip.Maybe(
			announcer,
			logging.PadLog(announcer.String(), memo.Wrap(announcer.Announce)),
		)(ctx)
	}
	if memo.Error() != nil {
		return fmt.Errorf("failed to announce release: %w", memo.Error())
	}
	return nil
}

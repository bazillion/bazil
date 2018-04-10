package ping // import "bazil.org/bazil/cli/server/ping"

import (
	clibazil "bazil.org/bazil/cli"
	"bazil.org/bazil/cliutil/subcommands"
	"bazil.org/bazil/server/control/wire"
	"golang.org/x/net/context"
)

type pingCommand struct {
	subcommands.Description
}

func (cmd *pingCommand) Run() error {
	ctx := context.Background()
	client, err := clibazil.Bazil.Control()
	if err != nil {
		return err
	}
	if _, err := client.Ping(ctx, &wire.PingRequest{}); err != nil {
		return err
	}
	return nil
}

var ping = pingCommand{
	Description: "ping bazil server",
}

func init() {
	subcommands.Register(&ping)
}

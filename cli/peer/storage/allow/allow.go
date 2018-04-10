package allow // import "bazil.org/bazil/cli/peer/storage/allow"

import (
	clibazil "bazil.org/bazil/cli"
	"bazil.org/bazil/cliutil/subcommands"
	"bazil.org/bazil/peer"
	"bazil.org/bazil/server/control/wire"
	"golang.org/x/net/context"
)

type allowCommand struct {
	subcommands.Description
	Arguments struct {
		PubKey  peer.PublicKey
		Storage string
	}
}

func (cmd *allowCommand) Run() error {
	req := &wire.PeerStorageAllowRequest{
		Pub:     cmd.Arguments.PubKey[:],
		Backend: cmd.Arguments.Storage,
	}
	ctx := context.Background()
	client, err := clibazil.Bazil.Control()
	if err != nil {
		return err
	}
	if _, err := client.PeerStorageAllow(ctx, req); err != nil {
		// TODO unwrap error
		return err
	}
	return nil
}

var allow = allowCommand{
	Description: "allow a peer",
}

func init() {
	subcommands.Register(&allow)
}

package peer

import (
	"bazil.org/bazil/db"
	"bazil.org/bazil/peer"
	"bazil.org/bazil/peer/wire"
	"bazil.org/bazil/server"
	"bazil.org/bazil/util/grpcedtls"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcpeer "google.golang.org/grpc/peer"
)

func (p *peers) auth(ctx context.Context) (*peer.PublicKey, error) {
	grpcPeer, ok := grpcpeer.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	auth, ok := grpcPeer.AuthInfo.(*grpcedtls.Auth)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "unauthenticated")
	}
	pub := (*peer.PublicKey)(auth.PeerPub)
	getPeer := func(tx *db.Tx) error {
		_, err := tx.Peers().Get(pub)
		return err
	}
	if err := p.app.DB.View(getPeer); err != nil {
		if err == db.ErrPeerNotFound {
			return nil, grpc.Errorf(codes.PermissionDenied, "permission denied")
		}
		return nil, err
	}
	return pub, nil
}

type peers struct {
	app *server.App
}

func New(app *server.App) *grpc.Server {
	auth := &grpcedtls.Authenticator{
		Config: app.GetTLSConfig,
		// TODO Lookup:
	}
	srv := grpc.NewServer(
		grpc.Creds(auth),
	)
	rpc := &peers{app: app}
	wire.RegisterPeerServer(srv, rpc)
	return srv
}

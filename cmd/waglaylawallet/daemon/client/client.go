package client

import (
	"context"
	"time"

	"github.com/waglayla/waglaylad/cmd/waglaylawallet/daemon/server"

	"github.com/pkg/errors"

	"github.com/waglayla/waglaylad/cmd/waglaylawallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the waglaylawalletd server, and returns the client instance
func Connect(address string) (pb.PyrinwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("waglaylawallet daemon is not running, start it with `waglaylawallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewPyrinwalletdClient(conn), func() {
		conn.Close()
	}, nil
}

package server

import (
	"context"
	"github.com/waglayla/waglaylad/cmd/waglaylawallet/daemon/pb"
	"github.com/waglayla/waglaylad/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}

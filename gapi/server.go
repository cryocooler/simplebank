package gapi

import (
	"fmt"

	db "github.com/cryocooler/simplebank/db/sqlc"
	"github.com/cryocooler/simplebank/pb"
	"github.com/cryocooler/simplebank/token"
	"github.com/cryocooler/simplebank/util"
	"github.com/cryocooler/simplebank/worker"
)

// servers gRPC service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// create new Server instance.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}

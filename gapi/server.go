package gapi

import (
	"fmt"

	db "github.com/cryocooler/simplebank/db/sqlc"
	"github.com/cryocooler/simplebank/pb"
	"github.com/cryocooler/simplebank/token"
	"github.com/cryocooler/simplebank/util"
)

// servers gRPC service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// create new Server instance.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

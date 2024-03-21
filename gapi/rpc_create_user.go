package gapi

import (
	"context"

	db "github.com/cryocooler/simplebank/db/sqlc"
	"github.com/cryocooler/simplebank/pb"
	"github.com/cryocooler/simplebank/util"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullname(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "user already exists %s", err)

		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)

	}

	rsp := &pb.CreateUserResponse{

		User: convertUser(user),
	}
	return rsp, nil
}

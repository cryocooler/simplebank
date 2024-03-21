package gapi

import (
	"context"
	"database/sql"

	db "github.com/cryocooler/simplebank/db/sqlc"
	"github.com/cryocooler/simplebank/pb"
	"github.com/cryocooler/simplebank/util"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found %s", err)
		}
		return nil, status.Errorf(codes.Internal, "could not retrieve user %s", err)
	}
	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect password %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		req.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create a token %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not refresh token %s", err)

	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create session %s", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return rsp, nil

}

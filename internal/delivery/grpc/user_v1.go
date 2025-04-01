package grpc

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	desc "github.com/vrnxx/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math/rand"
)

type UserV1GRPCServer struct {
	desc.UnimplementedUserV1Server
}

func NewV1GRPCServer() *UserV1GRPCServer {
	return &UserV1GRPCServer{}
}

func (s *UserV1GRPCServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	return &desc.CreateResponse{Id: rand.Int63n(99999)}, nil
}

func (s *UserV1GRPCServer) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	roles := [2]desc.Role{desc.Role_admin, desc.Role_user}
	return &desc.GetResponse{
		Id:        rand.Int63n(99999),
		Name:      gofakeit.FirstName(),
		Email:     gofakeit.Email(),
		Role:      roles[rand.Intn(len(roles))],
		CreatedAt: timestamppb.Now(),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *UserV1GRPCServer) Update(ctx context.Context, req *desc.UpdateRequest) (*desc.UpdateResponse, error) {
	return &desc.UpdateResponse{}, nil
}

func (s *UserV1GRPCServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	return &desc.DeleteResponse{}, nil
}

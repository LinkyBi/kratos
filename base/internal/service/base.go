package service

import (
	"base/internal/biz"
	"context"

	pb "base/api/base/v1"
)

type BaseService struct {
	baseUsecase *biz.BaseUsecase
	pb.UnimplementedBaseServer
}

func NewBaseService(baseUsecase *biz.BaseUsecase) *BaseService {
	return &BaseService{baseUsecase: baseUsecase}
}

func (s *BaseService) CreateBase(ctx context.Context, req *pb.CreateBaseRequest) (*pb.CreateBaseReply, error) {
	return &pb.CreateBaseReply{}, nil
}
func (s *BaseService) UpdateBase(ctx context.Context, req *pb.UpdateBaseRequest) (*pb.UpdateBaseReply, error) {
	return &pb.UpdateBaseReply{}, nil
}
func (s *BaseService) DeleteBase(ctx context.Context, req *pb.DeleteBaseRequest) (*pb.DeleteBaseReply, error) {
	return &pb.DeleteBaseReply{}, nil
}
func (s *BaseService) GetBase(ctx context.Context, req *pb.GetBaseRequest) (*pb.GetBaseReply, error) {
	return &pb.GetBaseReply{}, nil
}
func (s *BaseService) ListBase(ctx context.Context, req *pb.ListBaseRequest) (*pb.ListBaseReply, error) {
	return &pb.ListBaseReply{}, nil
}

package service

import (
	"context"

	pb "kvdbproxy/api/kvdb/v1"
	"kvdbproxy/internal/biz"
)

type KvdbService struct {
	pb.UnimplementedKvdbServer

	uc *biz.KvdbUsecase
}

func NewKvdbService(uc *biz.KvdbUsecase) *KvdbService {
	return &KvdbService{
		uc: uc,
	}
}

func (s *KvdbService) ListDB(ctx context.Context, req *pb.ListDBRequest) (*pb.ListDBReply, error) {
	return &pb.ListDBReply{}, nil
}
func (s *KvdbService) SearchPrefix(ctx context.Context, req *pb.SearchPrefixRequest) (*pb.SearchPrefixReply, error) {
	return &pb.SearchPrefixReply{}, nil
}
func (s *KvdbService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchReply, error) {
	return &pb.SearchReply{}, nil
}

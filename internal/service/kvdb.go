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
	listDB, err := s.uc.ListDB(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListDBReply{
		Name: listDB.Name,
	}, nil
}
func (s *KvdbService) SearchPrefix(ctx context.Context, req *pb.SearchPrefixRequest) (*pb.SearchPrefixReply, error) {
	lists, err := s.uc.SearchPrefix(ctx, req.KeyPrefix)
	if err != nil {
		return nil, err
	}

	values := make([]*pb.KvdbValue, 0)

	for _, v := range lists {
		values = append(values, &pb.KvdbValue{Key: v.Key, Value: v.Value})
	}
	return &pb.SearchPrefixReply{Kv: values}, nil
}
func (s *KvdbService) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchReply, error) {
	return &pb.SearchReply{}, nil
}

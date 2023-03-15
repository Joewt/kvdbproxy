package biz

import (
	"context"

	v1 "kvdbproxy/api/kvdb/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_KEY_NOT_FOUND.String(), "key not found")
)

// Kvdb is a Kvdb model.
type Kvdb struct {
	Hello string
}

// KvdbListDB model
type KvdbListDB struct {
	Name []string
}

type KvdbSearchPrefix struct {
	Key   string
	Value string
}

// KvdbRepo is a Greater repo.
type KvdbRepo interface {
	ListDB(ctx context.Context) (*KvdbListDB, error)
	SearchPrefix(ctx context.Context, prefix string) ([]*KvdbSearchPrefix, error)
}

// KvdbUsecase is a Kvdb usecase.
type KvdbUsecase struct {
	repo KvdbRepo
	log  *log.Helper
}

// NewKvdbUsecase new a Kvdb usecase.
func NewKvdbUsecase(repo KvdbRepo, logger log.Logger) *KvdbUsecase {
	return &KvdbUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *KvdbUsecase) ListDB(ctx context.Context) (*KvdbListDB, error) {
	return uc.repo.ListDB(ctx)
}

func (uc *KvdbUsecase) SearchPrefix(ctx context.Context, prefix string) ([]*KvdbSearchPrefix, error) {
	return uc.repo.SearchPrefix(ctx, prefix)
}

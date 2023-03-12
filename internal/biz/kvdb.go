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

// KvdbRepo is a Greater repo.
type KvdbRepo interface {
	Save(context.Context, *Kvdb) (*Kvdb, error)
	Update(context.Context, *Kvdb) (*Kvdb, error)
	FindByID(context.Context, int64) (*Kvdb, error)
	ListByHello(context.Context, string) ([]*Kvdb, error)
	ListAll(context.Context) ([]*Kvdb, error)
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

// CreateGreeter creates a Kvdb, and returns the new Kvdb.
func (uc *KvdbUsecase) CreateKvdb(ctx context.Context, g *Kvdb) (*Kvdb, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

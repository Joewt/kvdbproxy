package data

import (
	"context"

	"kvdbproxy/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type kvdbRepo struct {
	data *Data
	log  *log.Helper
}

// NewKvdbRepo .
func NewKvdbRepo(data *Data, logger log.Logger) biz.KvdbRepo {
	return &kvdbRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *kvdbRepo) Save(ctx context.Context, g *biz.Kvdb) (*biz.Kvdb, error) {
	return g, nil
}

func (r *kvdbRepo) Update(ctx context.Context, g *biz.Kvdb) (*biz.Kvdb, error) {
	return g, nil
}

func (r *kvdbRepo) FindByID(context.Context, int64) (*biz.Kvdb, error) {
	return nil, nil
}

func (r *kvdbRepo) ListByHello(context.Context, string) ([]*biz.Kvdb, error) {
	return nil, nil
}

func (r *kvdbRepo) ListAll(context.Context) ([]*biz.Kvdb, error) {
	return nil, nil
}

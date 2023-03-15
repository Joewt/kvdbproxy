package data

import (
	"context"
	"encoding/json"

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

func (r *kvdbRepo) ListDB(ctx context.Context) (*biz.KvdbListDB, error) {
	var resp biz.KvdbListDB
	return &resp, nil
}

func (r *kvdbRepo) SearchPrefix(ctx context.Context, prefix string) ([]*biz.KvdbSearchPrefix, error) {
	result := make([]*biz.KvdbSearchPrefix, 0)
	_, err := Paginate(r.data.db, &PageRequest{}, func(key, value []byte) error {
		keyValue, _ := json.Marshal(string(key))
		jsonValue, _ := json.Marshal(string(value))
		result = append(result, &biz.KvdbSearchPrefix{Key: string(keyValue), Value: string(jsonValue)})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	dbm "github.com/tendermint/tm-db"
	"kvdbproxy/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewKvdbRepo)

// Data .
type Data struct {
	db dbm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// TODO: 这里默认使用GoLevelDB调试，后续根据请求header进行动态调用
	db, err := dbm.NewDB(c.Database.Name, dbm.GoLevelDBBackend, c.Database.Dir)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
		db: db,
	}, cleanup, nil
}

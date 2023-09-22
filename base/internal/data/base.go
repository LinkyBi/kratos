package data

import (
	"context"

	"base/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// 连接数据客户端
type baseRepo struct {
	data *Data
	log  *log.Helper
}

func NewBaseRepo(data *Data, logger log.Logger) biz.BaseRepo {
	return &baseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// 实现操作的代码
func (r *baseRepo) Save(ctx context.Context, g *biz.Base) (*biz.Base, error) {
	return g, nil
}

func (r *baseRepo) Update(ctx context.Context, g *biz.Base) (*biz.Base, error) {
	return g, nil
}

func (r *baseRepo) FindByID(context.Context, int64) (*biz.Base, error) {
	return nil, nil
}

func (r *baseRepo) ListByHello(context.Context, string) ([]*biz.Base, error) {
	return nil, nil
}

func (r *baseRepo) ListAll(context.Context) ([]*biz.Base, error) {
	return nil, nil
}

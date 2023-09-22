package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

/*var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)*/

//errorreason这个我不知道为什么都要删掉，其他正常就是把greeter的定义修改成自己想要的

// 定义base的数据
type Base struct {
	account  string
	password string
	id       string
}

// 定义操作接口BaseRepo
type BaseRepo interface {
	Save(context.Context, *Base) (*Base, error)
	Update(context.Context, *Base) (*Base, error)
	FindByID(context.Context, int64) (*Base, error)
	ListByHello(context.Context, string) ([]*Base, error)
	ListAll(context.Context) ([]*Base, error)
}

// 定义Base日志
type BaseUsecase struct {
	repo BaseRepo
	log  *log.Helper
}

// 初始化BaseUsecase
func NewBaseUsecase(repo BaseRepo, logger log.Logger) *BaseUsecase {
	return &BaseUsecase{repo: repo, log: log.NewHelper(logger)}
}

// 具体操作用的都是type BaseRepo定义的操作，也就是我们api中定义的操作
func (uc *BaseUsecase) CreateBase(ctx context.Context, g *Base) (*Base, error) {
	uc.log.WithContext(ctx).Infof("CreateBase: %v", g.id)
	return uc.repo.Save(ctx, g)
}

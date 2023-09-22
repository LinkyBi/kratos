package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
// var ProviderSet = wire.NewSet(NewGreeterUsecase)原本

var ProviderSet = wire.NewSet(NewBaseUsecase)

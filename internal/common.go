package internal

import (
	"go.uber.org/zap"
)

type Common struct {
	*Config
	L *zap.Logger
}

func NewCommon(cfg Config) *Common {
	l, _ := zap.NewProduction()

	c := Common{
		Config: &cfg,
		L:      l,
	}

	return &c
}

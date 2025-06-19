package middlewares

import "lolymarsh/pkg/configs"

type Middleware struct {
	conf *configs.Config
}

func NewMiddleware(conf *configs.Config) *Middleware {
	return &Middleware{conf: conf}
}

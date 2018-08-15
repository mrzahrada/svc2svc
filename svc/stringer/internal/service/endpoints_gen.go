package service

import "context"

type CountEndpoint func(context.Context, string) (int, error)

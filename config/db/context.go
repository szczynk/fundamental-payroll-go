package db

import (
	"context"
	"time"
)

func NewContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

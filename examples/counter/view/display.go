package view

import (
	"context"

	"github.com/go-po/po/streams"
)

type Display struct{}

func (display Display) Handle(ctx context.Context, msg streams.Message) error {
	return nil
}

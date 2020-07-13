package command

import (
	"context"

	"github.com/go-po/po"
	"github.com/go-po/po/streams"
)

// A CommandHandler with a non-pointer receiver
type Stateless struct {
}

func (less Stateless) Handle(ctx context.Context, msg streams.Message) error {
	return nil
}

func (less Stateless) Execute(appender po.TransactionAppender) error {
	return nil
}

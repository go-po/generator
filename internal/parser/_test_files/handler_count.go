package _test_files

import (
	"context"
	"fmt"

	"github.com/go-po/po"
	"github.com/go-po/po/streams"
)

type Counter struct {
	Count int
	Texts []string
}

func (cmd *Counter) Handle(ctx context.Context, msg streams.Message) error {
	switch event := msg.Data.(type) {
	case Increment:
		cmd.Count = cmd.Count + event.Value
	default:
	}
}

func (cmd *Counter) Execute(appender po.TransactionAppender) error {
	msg := Message{
		Text: fmt.Sprintf("Count is: %d", cmd.Count),
	}
	appender.Append(msg)
	return nil
}

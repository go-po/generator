package command

import (
	"context"
	"fmt"

	event2 "github.com/go-po/generator/examples/counter/event"
	"github.com/go-po/po"
	"github.com/go-po/po/streams"
)

type Counter struct {
	Count int
	Texts []string
}

func (cmd *Counter) Handle(ctx context.Context, msg streams.Message) error {
	switch event := msg.Data.(type) {
	case event2.Increment:
		cmd.Count = cmd.Count + event.Value
	default:
	}
	return nil
}

// Executes the command
func (cmd *Counter) Execute(appender po.TransactionAppender) error {
	msg := event2.Message{
		Text: fmt.Sprintf("Count is: %d", cmd.Count),
	}
	appender.Append(msg)
	return nil
}

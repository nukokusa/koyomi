package koyomi

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

type DeleteOption struct {
	CalendarID string `required:"" help:"Calendar identifier"`
	EventID    string `required:"" help:"Identifier of the event"`
}

func (k *Koyomi) Delete(ctx context.Context, opt *DeleteOption) error {
	if err := k.cs.Delete(ctx, opt.CalendarID, opt.EventID); err != nil {
		return errors.Wrap(err, "error Delete")
	}

	log.Printf("[DEBUG] deleted event: CalendarID=%s, EventID=%s", opt.CalendarID, opt.EventID)

	return nil
}

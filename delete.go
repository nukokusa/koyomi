package koyomi

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

type DeleteOption struct {
	CalendarID string `help:"calendar ID" name:"calendar-id"`
	EventID    string `help:"event ID" name:"event-id"`
}

func (k *Koyomi) Delete(ctx context.Context, opt *DeleteOption) error {
	ev, err := k.cs.Events.Get(opt.CalendarID, opt.EventID).Do()
	if err != nil {
		return errors.Wrap(err, "error Events.Get")
	}

	if err := k.cs.Events.Delete(opt.CalendarID, ev.Id).Do(); err != nil {
		return errors.Wrap(err, "error Events.Delete")
	}

	log.Printf("[DEBUG] deleted event: calendar_id=%s, event_id=%s", opt.CalendarID, ev.Id)

	return nil
}

package koyomi

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
)

type UpdateOption struct {
	CalendarID  string `help:"calendar ID" name:"calendar-id"`
	EventID     string `help:"event ID" name:"event-id"`
	Summary     string `help:"event summary" name:"summary"`
	Description string `help:"event descriptuon" name:"description"`
	StartAt     string `help:"event start at" name:"start-at"`
	EndAt       string `help:"event end at" name:"end-at"`
}

func (k *Koyomi) Update(ctx context.Context, opt *UpdateOption) error {
	ev := &calendar.Event{}
	if opt.Summary != "" {
		ev.Summary = opt.Summary
	}
	if opt.Description != "" {
		ev.Description = opt.Description
	}
	if opt.StartAt != "" {
		startAt, err := time.ParseInLocation(layout, opt.StartAt, loc)
		if err != nil {
			return errors.Wrap(err, "error ParseInLocation")
		}
		ev.Start = &calendar.EventDateTime{DateTime: startAt.Format(time.RFC3339), TimeZone: locationName}
	}
	if opt.EndAt != "" {
		endAt, err := time.ParseInLocation(layout, opt.EndAt, loc)
		if err != nil {
			return errors.Wrap(err, "error ParseInLocation")
		}
		ev.End = &calendar.EventDateTime{DateTime: endAt.Format(time.RFC3339), TimeZone: locationName}
	}
	if _, err := k.cs.Events.Patch(opt.CalendarID, opt.EventID, ev).Do(); err != nil {
		return errors.Wrap(err, "error Events.Patch")
	}

	log.Printf("[DEBUG] updated event: calendar_id=%s, event_id=%s", opt.CalendarID, ev.Id)

	return nil
}

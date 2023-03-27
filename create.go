package koyomi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
)

type CreateOption struct {
	CalendarID  string `help:"calendar ID" name:"calendar-id"`
	Summary     string `help:"event summary" name:"summary"`
	Description string `help:"event descriptuon" name:"description"`
	StartAt     string `help:"event start at" name:"start-at"`
	EndAt       string `help:"event end at" name:"end-at"`
}

func (k *Koyomi) Create(ctx context.Context, opt *CreateOption) error {
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
		ev.Start = &calendar.EventDateTime{DateTime: startAt.Format(time.RFC3339)}
	}
	if opt.EndAt != "" {
		endAt, err := time.ParseInLocation(layout, opt.EndAt, loc)
		if err != nil {
			return errors.Wrap(err, "error ParseInLocation")
		}
		ev.End = &calendar.EventDateTime{DateTime: endAt.Format(time.RFC3339)}
	}

	var err error
	ev, err = k.cs.Events.Insert(opt.CalendarID, ev).Do()
	if err != nil {
		return errors.Wrap(err, "error Events.Insert")
	}

	log.Printf("[DEBUG] inserted event: calendar_id=%s, event_id=%s", opt.CalendarID, ev.Id)
	fmt.Println(ev.Id)

	return nil
}

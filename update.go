package koyomi

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/tkuchiki/parsetime"
	"google.golang.org/api/calendar/v3"
)

type UpdateOption struct {
	CalendarID  string `required:"" help:"Calendar identifier"`
	EventID     string `required:"" help:"Identifier of the event"`
	Summary     string `help:"Title of the event" short:"s"`
	Description string `help:"Description of the event" short:"d"`
	StartTime   string `help:"The start time of the event"`
	EndTime     string `help:"The end time of the event"`
}

func (k *Koyomi) Update(ctx context.Context, opt *UpdateOption) error {
	event := &calendar.Event{
		Id:          opt.EventID,
		Summary:     opt.Summary,
		Description: opt.Description,
	}
	if opt.StartTime != "" {
		p, err := parsetime.NewParseTime()
		if err != nil {
			return errors.Wrap(err, "error parsetime.NewParseTime")
		}
		t, err := p.Parse(opt.StartTime)
		if err != nil {
			return errors.Wrap(err, "error Parse")
		}
		event.Start = &calendar.EventDateTime{DateTime: t.Format(time.RFC3339)}
	}
	if opt.EndTime != "" {
		p, err := parsetime.NewParseTime()
		if err != nil {
			return errors.Wrap(err, "error parsetime.NewParseTime")
		}
		t, err := p.Parse(opt.EndTime)
		if err != nil {
			return errors.Wrap(err, "error Parse")
		}
		event.End = &calendar.EventDateTime{DateTime: t.Format(time.RFC3339)}
	}

	var err error
	event, err = k.cs.Patch(ctx, opt.CalendarID, event)
	if err != nil {
		return errors.Wrap(err, "error Patch")
	}

	log.Printf("[DEBUG] updated event: CalendarID=%s, EventID=%s", opt.CalendarID, event.Id)

	return nil
}

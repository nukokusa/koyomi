package koyomi

import (
	"context"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
	"github.com/tkuchiki/parsetime"
)

type UpdateOption struct {
	CalendarID  string `required:"" help:"Calendar identifier"`
	ID          string `required:"" help:"Identifier of the event"`
	Summary     string `help:"Title of the event"`
	Description string `help:"Description of the event"`
	StartTime   string `help:"The start time of the event" short:"s"`
	EndTime     string `help:"The end time of the event" short:"e"`
}

func (k *Koyomi) Update(ctx context.Context, opt *UpdateOption) error {
	event := &Event{
		ID:          opt.ID,
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
		event.StartTime = t
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
		event.EndTime = t
	}

	var err error
	event, err = k.cs.Patch(ctx, opt.CalendarID, event)
	if err != nil {
		return errors.Wrap(err, "error Patch")
	}

	log.Printf("[DEBUG] updated event: calendar_id=%s, id=%s", opt.CalendarID, event.ID)

	if err := json.NewEncoder(k.stdout).Encode(event); err != nil {
		return errors.Wrap(err, "error Encode")
	}
	return nil
}

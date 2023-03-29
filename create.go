package koyomi

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Songmu/flextime"
	"github.com/pkg/errors"
	"github.com/tkuchiki/parsetime"
)

type CreateOption struct {
	CalendarID  string `required:"" help:"Calendar identifier"`
	Summary     string `help:"Title of the event"`
	Description string `help:"Descriptuon of the event"`
	StartTime   string `help:"The start time of the event" short:"s"`
	EndTime     string `help:"The end time of the event" short:"e"`
}

func (k *Koyomi) Create(ctx context.Context, opt *CreateOption) error {
	p, err := parsetime.NewParseTime()
	if err != nil {
		return errors.Wrap(err, "error parsetime.NewParseTime")
	}
	startTime := flextime.Now()
	if opt.StartTime != "" {
		startTime, err = p.Parse(opt.StartTime)
		if err != nil {
			return errors.Wrap(err, "error Parse")
		}
	}
	endTime := flextime.Now()
	if opt.EndTime != "" {
		endTime, err = p.Parse(opt.EndTime)
		if err != nil {
			return errors.Wrap(err, "error Parse")
		}
	}

	event := &Event{
		Summary:     opt.Summary,
		Description: opt.Description,
		StartTime:   startTime,
		EndTime:     endTime,
	}
	event, err = k.cs.Insert(ctx, opt.CalendarID, event)
	if err != nil {
		return errors.Wrap(err, "error Insert")
	}

	log.Printf("[DEBUG] inserted event: CalendarID=%s, ID=%s", opt.CalendarID, event.ID)

	if err := json.NewEncoder(k.stdout).Encode(event); err != nil {
		return errors.Wrap(err, "error Encode")
	}
	return nil
}

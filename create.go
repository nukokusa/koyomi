package koyomi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/tkuchiki/parsetime"
	"google.golang.org/api/calendar/v3"
)

type CreateOption struct {
	CalendarID  string `required:"" help:"Calendar identifier"`
	Summary     string `help:"Title of the event" short:"s"`
	Description string `help:"Descriptuon of the event" short:"d"`
	StartTime   string `required:"" help:"The start time of the event"`
	EndTime     string `required:"" help:"The end time of the event"`
}

func (k *Koyomi) Create(ctx context.Context, opt *CreateOption) error {
	event := &calendar.Event{
		Summary:     opt.Summary,
		Description: opt.Description,
	}

	p, err := parsetime.NewParseTime()
	if err != nil {
		return errors.Wrap(err, "error parsetime.NewParseTime")
	}
	startTime, err := p.Parse(opt.StartTime)
	if err != nil {
		return errors.Wrap(err, "error Parse")
	}
	event.Start = &calendar.EventDateTime{DateTime: startTime.Format(time.RFC3339)}
	endTime, err := p.Parse(opt.EndTime)
	if err != nil {
		return errors.Wrap(err, "error Parse")
	}
	event.End = &calendar.EventDateTime{DateTime: endTime.Format(time.RFC3339)}

	event, err = k.cs.Insert(ctx, opt.CalendarID, event)
	if err != nil {
		return errors.Wrap(err, "error Insert")
	}

	log.Printf("[DEBUG] inserted event: CalendarID=%s, EventID=%s", opt.CalendarID, event.Id)
	fmt.Println(event.Id)

	return nil
}

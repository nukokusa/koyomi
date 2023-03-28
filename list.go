package koyomi

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/tkuchiki/parsetime"
)

type ListOption struct {
	CalendarID string `required:"" help:"Calendar identifier"`
	StartTime  string `required:"" help:"The start time of the event" short:"s"`
	EndTime    string `required:"" help:"The end time of the event" short:"e"`
}

func (k *Koyomi) List(ctx context.Context, opt *ListOption) error {
	p, err := parsetime.NewParseTime()
	if err != nil {
		return errors.Wrap(err, "error parsetime.NewParseTime")
	}
	startTime, err := p.Parse(opt.StartTime)
	if err != nil {
		return errors.Wrap(err, "error Parse")
	}
	endTime, err := p.Parse(opt.EndTime)
	if err != nil {
		return errors.Wrap(err, "error Parse")
	}

	events, err := k.cs.List(ctx, opt.CalendarID, startTime, endTime)
	if err != nil {
		return errors.Wrap(err, "error List")
	}

	if err := json.NewEncoder(k.stdout).Encode(events); err != nil {
		return errors.Wrap(err, "error Encode")
	}
	return nil
}

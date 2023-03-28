package koyomi

import (
	"context"
	"os"

	"github.com/olekukonko/tablewriter"
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

	rows := make([][]string, 0, len(events))
	for _, event := range events {
		rows = append(rows, []string{
			event.Id,
			event.Summary,
			event.Start.DateTime,
			event.End.DateTime,
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "summary", "start_time", "end_time"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(rows)
	table.Render()

	return nil
}

package koyomi

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

type DeleteOption struct {
	CalendarID string `required:"" help:"Calendar identifier"`
	ID         string `required:"" help:"Identifier of the event"`
}

func (k *Koyomi) Delete(ctx context.Context, opt *DeleteOption) error {
	if err := k.cs.Delete(ctx, opt.CalendarID, opt.ID); err != nil {
		return errors.Wrap(err, "error Delete")
	}

	log.Printf("[DEBUG] deleted event: CalendarID=%s, ID=%s", opt.CalendarID, opt.ID)

	return nil
}

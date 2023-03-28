package koyomi

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
)

var (
	Version string
	loc     *time.Location
)

const (
	locationName string = "Asia/Tokyo"
	layout       string = "2006-01-02 15:04:05"
)

func init() {
	var err error
	loc, err = time.LoadLocation(locationName)
	if err != nil {
		log.Panicln(err)
	}
	time.Local = loc
}

type Koyomi struct {
	cs CalendarService
}

func New(ctx context.Context, credentialPath string) (*Koyomi, error) {
	cs, err := newCalendarService(ctx, credentialPath)
	if err != nil {
		return nil, errors.Wrap(err, "error newCalendarService")
	}
	return &Koyomi{
		cs: cs,
	}, nil
}

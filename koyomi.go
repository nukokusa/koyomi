package koyomi

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
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
	cs *calendar.Service
}

func New(ctx context.Context, credentialPath string) (*Koyomi, error) {
	cs, err := calendar.NewService(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		return nil, errors.Wrap(err, "error calendar.NewService")
	}
	return &Koyomi{
		cs: cs,
	}, nil
}

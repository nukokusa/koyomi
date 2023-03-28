package koyomi

import (
	"context"
	"sort"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type CalendarService interface {
	List(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*calendar.Event, error)
	Insert(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error)
	Patch(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error)
	Delete(ctx context.Context, calendarID, eventID string) error
}

type calendarService struct {
	cs *calendar.Service
}

func newCalendarService(ctx context.Context, credentialPath string) (CalendarService, error) {
	cs, err := calendar.NewService(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		return nil, errors.Wrap(err, "error calendar.NewService")
	}
	return &calendarService{
		cs: cs,
	}, nil
}

func (s *calendarService) List(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*calendar.Event, error) {
	result := []*calendar.Event{}
	pageToken := ""
	for {
		evs, err := s.cs.Events.List(calendarID).
			TimeMin(startTime.Format(time.RFC3339)).
			TimeMax(endTime.Format(time.RFC3339)).
			SingleEvents(true).
			PageToken(pageToken).
			Context(ctx).Do()
		if err != nil {
			return nil, errors.Wrap(err, "error Events.List")
		}
		result = append(result, evs.Items...)
		if evs.NextPageToken == "" {
			break
		}
		pageToken = evs.NextPageToken
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Start.DateTime < result[j].Start.DateTime
	})
	return result, nil
}

func (s *calendarService) Insert(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error) {
	ev, err := s.cs.Events.Insert(calendarID, event).Context(ctx).Do()
	if err != nil {
		return nil, errors.Wrap(err, "error Events.Insert")
	}
	return ev, nil
}

func (s *calendarService) Patch(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error) {
	ev, err := s.cs.Events.Patch(calendarID, event.Id, event).Context(ctx).Do()
	if err != nil {
		return nil, errors.Wrap(err, "error Events.Patch")
	}
	return ev, nil
}

func (s *calendarService) Delete(ctx context.Context, calendarID, eventID string) error {
	if err := s.cs.Events.Delete(calendarID, eventID).Context(ctx).Do(); err != nil {
		return errors.Wrap(err, "error Events.Delete")
	}
	return nil
}

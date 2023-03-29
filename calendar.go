package koyomi

import (
	"context"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/pkg/errors"
	"github.com/shogo82148/go-retry"
	"github.com/tkuchiki/parsetime"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type Event struct {
	ID          string    `json:"id"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
}

func NewEvent(ev *calendar.Event) (*Event, error) {
	var startTime, endTime time.Time
	p, err := parsetime.NewParseTime()
	if err != nil {
		return nil, errors.Wrap(err, "error parsetime.NewParseTime")
	}
	if ev.Start != nil {
		v := ev.Start.DateTime
		if v == "" {
			v = ev.Start.Date
		}
		startTime, err = p.Parse(v)
		if err != nil {
			return nil, errors.Wrap(err, "error Parse")
		}
	}
	if ev.End != nil {
		v := ev.End.DateTime
		if v == "" {
			v = ev.End.Date
		}
		endTime, err = p.Parse(v)
		if err != nil {
			return nil, errors.Wrap(err, "error Parse")
		}
	}
	return &Event{
		ID:          ev.Id,
		Summary:     ev.Summary,
		Description: ev.Description,
		StartTime:   startTime,
		EndTime:     endTime,
	}, nil
}

type CalendarService interface {
	List(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*Event, error)
	Insert(ctx context.Context, calendarID string, event *Event) (*Event, error)
	Patch(ctx context.Context, calendarID string, event *Event) (*Event, error)
	Delete(ctx context.Context, calendarID, eventID string) error
}

type calendarService struct {
	cs     *calendar.Service
	policy retry.Policy
}

func newCalendarService(ctx context.Context, credentialPath string) (CalendarService, error) {
	cs, err := calendar.NewService(ctx, option.WithCredentialsFile(credentialPath))
	if err != nil {
		return nil, errors.Wrap(err, "error calendar.NewService")
	}
	policy := retry.Policy{
		MinDelay: time.Second,
		MaxDelay: 100 * time.Second,
		MaxCount: 10,
	}
	return &calendarService{
		cs:     cs,
		policy: policy,
	}, nil
}

func (s *calendarService) List(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*Event, error) {
	evs := []*calendar.Event{}
	req := s.cs.Events.List(calendarID).
		TimeMin(startTime.Format(time.RFC3339)).
		TimeMax(endTime.Format(time.RFC3339)).
		SingleEvents(true)

	pageToken := ""
	for {
		retrier := s.policy.Start(ctx)
		var resp *calendar.Events
		for retrier.Continue() {
			var err error
			resp, err = req.PageToken(pageToken).Context(ctx).Do()
			if err == nil {
				break
			}
			if apiError, ok := err.(*googleapi.Error); ok {
				if apiError.Code == http.StatusTooManyRequests {
					log.Printf("[WARN] reached to Too Many Requests: calendar_id=%s", calendarID)
					continue
				}
			}
			return nil, errors.Wrap(err, "error Events.List")
		}
		evs = append(evs, resp.Items...)
		if resp.NextPageToken == "" {
			break
		}
		pageToken = resp.NextPageToken
	}

	result := make([]*Event, 0, len(evs))
	for _, ev := range evs {
		event, err := NewEvent(ev)
		if err != nil {
			return nil, errors.Wrap(err, "error NewEvent")
		}
		result = append(result, event)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].StartTime.Before(result[j].StartTime)
	})
	return result, nil
}

func (s *calendarService) Insert(ctx context.Context, calendarID string, event *Event) (*Event, error) {
	ev := &calendar.Event{
		Id:          event.ID,
		Summary:     event.Summary,
		Description: event.Description,
		Start: &calendar.EventDateTime{
			DateTime: event.StartTime.Format(time.RFC3339),
		},
		End: &calendar.EventDateTime{
			DateTime: event.EndTime.Format(time.RFC3339),
		},
	}
	req := s.cs.Events.Insert(calendarID, ev)

	retrier := s.policy.Start(ctx)
	var resp *calendar.Event
	for retrier.Continue() {
		var err error
		resp, err = req.Context(ctx).Do()
		if err == nil {
			break
		}
		if apiError, ok := err.(*googleapi.Error); ok {
			if apiError.Code == http.StatusTooManyRequests {
				log.Printf("[WARN] reached to Too Many Requests: calendar_id=%s", calendarID)
				continue
			}
		}
		return nil, errors.Wrap(err, "error Events.Insert")
	}

	result, err := NewEvent(resp)
	if err != nil {
		return nil, errors.Wrap(err, "error NewEvent")
	}
	return result, nil
}

func (s *calendarService) Patch(ctx context.Context, calendarID string, event *Event) (*Event, error) {
	ev := &calendar.Event{
		Id:          event.ID,
		Summary:     event.Summary,
		Description: event.Description,
	}
	if !event.StartTime.IsZero() {
		ev.Start.DateTime = event.StartTime.Format(time.RFC3339)
	}
	if !event.EndTime.IsZero() {
		ev.End.DateTime = event.EndTime.Format(time.RFC3339)
	}
	req := s.cs.Events.Patch(calendarID, ev.Id, ev)

	retrier := s.policy.Start(ctx)
	var resp *calendar.Event
	for retrier.Continue() {
		var err error
		resp, err = req.Context(ctx).Do()
		if err == nil {
			break
		}
		if apiError, ok := err.(*googleapi.Error); ok {
			if apiError.Code == http.StatusTooManyRequests {
				log.Printf("[WARN] reached to Too Many Requests: calendar_id=%s", calendarID)
				continue
			}
		}
		return nil, errors.Wrap(err, "error Events.Insert")
	}

	result, err := NewEvent(resp)
	if err != nil {
		return nil, errors.Wrap(err, "error NewEvent")
	}
	return result, nil
}

func (s *calendarService) Delete(ctx context.Context, calendarID, eventID string) error {
	req := s.cs.Events.Delete(calendarID, eventID)
	retrier := s.policy.Start(ctx)
	for retrier.Continue() {
		err := req.Context(ctx).Do()
		if err == nil {
			break
		}
		if apiError, ok := err.(*googleapi.Error); ok {
			if apiError.Code == http.StatusTooManyRequests {
				log.Printf("[WARN] reached to Too Many Requests: calendar_id=%s", calendarID)
				continue
			}
		}
		return errors.Wrap(err, "error Events.Delete")
	}

	return nil
}

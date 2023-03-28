package koyomi

import (
	"context"
	"errors"
	"time"

	"google.golang.org/api/calendar/v3"
)

type calendarServiceMock struct {
	ListMock   func(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*calendar.Event, error)
	InsertMock func(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error)
	PatchMock  func(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error)
	DeleteMock func(ctx context.Context, calendarID, eventID string) error
}

func newCalendarServiceMock() *calendarServiceMock {
	return &calendarServiceMock{
		ListMock: func(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*calendar.Event, error) {
			return nil, errors.New("not implememted")
		},
		InsertMock: func(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error) {
			return nil, errors.New("not implememted")
		},
		PatchMock: func(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error) {
			return nil, errors.New("not implememted")
		},
		DeleteMock: func(ctx context.Context, calendarID, eventID string) error {
			return errors.New("not implememted")
		},
	}
}

func (s *calendarServiceMock) List(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*calendar.Event, error) {
	return s.ListMock(ctx, calendarID, startTime, endTime)
}
func (s *calendarServiceMock) Insert(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error) {
	return s.InsertMock(ctx, calendarID, event)
}
func (s *calendarServiceMock) Patch(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error) {
	return s.PatchMock(ctx, calendarID, event)
}
func (s *calendarServiceMock) Delete(ctx context.Context, calendarID, eventID string) error {
	return s.DeleteMock(ctx, calendarID, eventID)
}

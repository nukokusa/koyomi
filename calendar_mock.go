package koyomi

import (
	"context"
	"errors"
	"time"
)

type calendarServiceMock struct {
	ListMock   func(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*Event, error)
	InsertMock func(ctx context.Context, calendarID string, event *Event) (*Event, error)
	UpdateMock func(ctx context.Context, calendarID string, event *Event) (*Event, error)
	DeleteMock func(ctx context.Context, calendarID, eventID string) error
}

func newCalendarServiceMock() *calendarServiceMock {
	return &calendarServiceMock{
		ListMock: func(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*Event, error) {
			return nil, errors.New("not implememted")
		},
		InsertMock: func(ctx context.Context, calendarID string, event *Event) (*Event, error) {
			return nil, errors.New("not implememted")
		},
		UpdateMock: func(ctx context.Context, calendarID string, event *Event) (*Event, error) {
			return nil, errors.New("not implememted")
		},
		DeleteMock: func(ctx context.Context, calendarID, eventID string) error {
			return errors.New("not implememted")
		},
	}
}

func (s *calendarServiceMock) List(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*Event, error) {
	return s.ListMock(ctx, calendarID, startTime, endTime)
}
func (s *calendarServiceMock) Insert(ctx context.Context, calendarID string, event *Event) (*Event, error) {
	return s.InsertMock(ctx, calendarID, event)
}
func (s *calendarServiceMock) Update(ctx context.Context, calendarID string, event *Event) (*Event, error) {
	return s.UpdateMock(ctx, calendarID, event)
}
func (s *calendarServiceMock) Delete(ctx context.Context, calendarID, eventID string) error {
	return s.DeleteMock(ctx, calendarID, eventID)
}

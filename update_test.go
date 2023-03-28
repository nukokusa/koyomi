package koyomi_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/nukokusa/koyomi"
	"google.golang.org/api/calendar/v3"
)

func TestKoyomi_Update(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	var result *calendar.Event
	cs := koyomi.NewCalendarServiceMock()
	cs.PatchMock = func(ctx context.Context, calendarID string, event *calendar.Event) (*calendar.Event, error) {
		result = event
		return event, nil
	}

	k := &koyomi.Koyomi{}
	k.SetCalendarService(cs)

	opt := &koyomi.UpdateOption{
		CalendarID:  "dummy-calendar-id",
		EventID:     "dummy-id",
		Summary:     "dummy-summary",
		Description: "dummy-description",
		StartTime:   "2023-03-01 12:00:00",
		EndTime:     "2023-03-02 13:00:00",
	}

	if err := k.Update(ctx, opt); err != nil {
		t.Fatal("error Update", err)
	}

	expected := &calendar.Event{
		Id:          "dummy-id",
		Summary:     "dummy-summary",
		Description: "dummy-description",
		Start: &calendar.EventDateTime{
			DateTime: "2023-03-01T12:00:00+09:00",
		},
		End: &calendar.EventDateTime{
			DateTime: "2023-03-02T13:00:00+09:00",
		},
	}

	opts := []cmp.Option{
		cmpopts.IgnoreUnexported(calendar.Event{}, calendar.EventDateTime{}),
	}
	if diff := cmp.Diff(expected, result, opts...); diff != "" {
		t.Errorf("error mismatch (-want +got):\n%s", diff)
	}
}

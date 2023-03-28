package koyomi_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nukokusa/koyomi"
)

func TestKoyomi_Create(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cs := koyomi.NewCalendarServiceMock()
	cs.InsertMock = func(ctx context.Context, calendarID string, event *koyomi.Event) (*koyomi.Event, error) {
		event.ID = "dummy-id"
		return event, nil
	}

	k := &koyomi.Koyomi{}
	k.SetCalendarService(cs)
	var b bytes.Buffer
	k.SetStdout(&b)

	opt := &koyomi.CreateOption{
		CalendarID:  "dummy-calendar-id",
		Summary:     "dummy-summary",
		Description: "dummy-description",
		StartTime:   "2023-03-01 12:00:00",
		EndTime:     "2023-03-02 13:00:00",
	}

	if err := k.Create(ctx, opt); err != nil {
		t.Fatal("error Create", err)
	}

	expected := `{"id":"dummy-id","summary":"dummy-summary","description":"dummy-description","start_time":"2023-03-01T12:00:00+09:00","end_time":"2023-03-02T13:00:00+09:00"}
`
	if diff := cmp.Diff(expected, b.String()); diff != "" {
		t.Errorf("error mismatch (-want +got):\n%s", diff)
	}
}

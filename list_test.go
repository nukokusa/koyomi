package koyomi_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/nukokusa/koyomi"
)

func TestKoyomi_List(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	parse := func(str string) time.Time {
		tt, err := time.Parse(time.RFC3339, str)
		if err != nil {
			t.Fatal(err)
		}
		return tt
	}

	cs := koyomi.NewCalendarServiceMock()
	cs.ListMock = func(ctx context.Context, calendarID string, startTime, endTime time.Time) ([]*koyomi.Event, error) {
		return []*koyomi.Event{
			{
				ID:          "dummy-id-a",
				Summary:     "dummy-summary-a",
				Description: "dummy-description-a",
				StartTime:   parse("2023-03-01T12:00:00+09:00"),
				EndTime:     parse("2023-03-01T13:00:00+09:00"),
			},
			{
				ID:          "dummy-id-b",
				Summary:     "dummy-summary-b",
				Description: "dummy-description-b",
				StartTime:   parse("2023-03-01T15:00:00+09:00"),
				EndTime:     parse("2023-03-01T16:00:00+09:00"),
			},
		}, nil
	}

	k := &koyomi.Koyomi{}
	k.SetCalendarService(cs)
	var b bytes.Buffer
	k.SetStdout(&b)

	opt := &koyomi.ListOption{
		CalendarID: "dummy-calendar-id",
		StartTime:  "2023-03-01 00:00:00",
		EndTime:    "2023-03-02 00:00:00",
	}

	if err := k.List(ctx, opt); err != nil {
		t.Fatal("error List", err)
	}

	expected := `[{"id":"dummy-id-a","summary":"dummy-summary-a","description":"dummy-description-a","start_time":"2023-03-01T12:00:00+09:00","end_time":"2023-03-01T13:00:00+09:00"},{"id":"dummy-id-b","summary":"dummy-summary-b","description":"dummy-description-b","start_time":"2023-03-01T15:00:00+09:00","end_time":"2023-03-01T16:00:00+09:00"}]
`
	if diff := cmp.Diff(expected, b.String()); diff != "" {
		t.Errorf("error mismatch (-want +got):\n%s", diff)
	}
}

package koyomi_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/google/go-cmp/cmp"
	"github.com/nukokusa/koyomi"
)

func TestKoyomi_Create(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	now, err := time.Parse(time.RFC3339, "2023-04-01T12:00:00+09:00")
	if err != nil {
		t.Fatal(err)
	}
	restore := flextime.Fix(now)
	defer restore()

	cs := koyomi.NewCalendarServiceMock()
	cs.InsertMock = func(ctx context.Context, calendarID string, event *koyomi.Event) (*koyomi.Event, error) {
		event.ID = "dummy-id"
		return event, nil
	}

	k := &koyomi.Koyomi{}
	k.SetCalendarService(cs)
	var b bytes.Buffer
	k.SetStdout(&b)

	tests := []struct {
		opt      *koyomi.CreateOption
		expected string
	}{
		{
			opt: &koyomi.CreateOption{
				CalendarID:  "dummy-calendar-id",
				Summary:     "dummy-summary",
				Description: "dummy-description",
				StartTime:   "2023-03-01 12:00:00",
				EndTime:     "2023-03-02 13:00:00",
			},
			expected: `{"id":"dummy-id","summary":"dummy-summary","description":"dummy-description","start_time":"2023-03-01T12:00:00+09:00","end_time":"2023-03-02T13:00:00+09:00"}
`,
		},
		{
			opt: &koyomi.CreateOption{
				CalendarID:  "dummy-calendar-id",
				Summary:     "dummy-summary",
				Description: "dummy-description",
			},
			expected: `{"id":"dummy-id","summary":"dummy-summary","description":"dummy-description","start_time":"2023-04-01T12:00:00+09:00","end_time":"2023-04-01T12:00:00+09:00"}
`,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			defer b.Reset()

			if err := k.Create(ctx, tt.opt); err != nil {
				t.Fatal("error Create", err)
			}

			if diff := cmp.Diff(tt.expected, b.String()); diff != "" {
				t.Errorf("error mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

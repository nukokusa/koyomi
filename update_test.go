package koyomi_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nukokusa/koyomi"
)

func TestKoyomi_Update(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	cs := koyomi.NewCalendarServiceMock()
	cs.PatchMock = func(ctx context.Context, calendarID string, event *koyomi.Event) (*koyomi.Event, error) {
		return event, nil
	}

	k := &koyomi.Koyomi{}
	k.SetCalendarService(cs)
	var b bytes.Buffer
	k.SetStdout(&b)

	opt := &koyomi.UpdateOption{
		CalendarID:  "dummy-calendar-id",
		ID:          "dummy-id",
		Summary:     "dummy-summary",
		Description: "dummy-description",
		StartTime:   "2023-03-01 12:00:00",
		EndTime:     "2023-03-02 13:00:00",
	}

	if err := k.Update(ctx, opt); err != nil {
		t.Fatal("error Update", err)
	}

	expected := `{"id":"dummy-id","summary":"dummy-summary","description":"dummy-description","start_time":"2023-03-01T12:00:00+09:00","end_time":"2023-03-02T13:00:00+09:00"}
`
	if diff := cmp.Diff(expected, b.String()); diff != "" {
		t.Errorf("error mismatch (-want +got):\n%s", diff)
	}
}

package koyomi

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestGet(t *testing.T) {
	t.Parallel()

	cl := newCalendarMock()
	cl.GetMock = func(key, uid string) (*Event, error) {
		uidByKey := map[string]string{"schedule": "101"}
		u, ok := uidByKey[key]
		if !ok || u != uid {
			return nil, ErrEventNotFound
		}

		return &Event{
			Key:         key,
			UID:         uid,
			Summary:     "summary",
			Description: "description",
			StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
			EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
		}, nil
	}

	type Input struct {
		Key string
		UID string
	}
	type Test struct {
		Input  Input
		Expect *Event
		Error  error
	}
	tests := []Test{
		{
			Input: Input{
				Key: "schedule",
				UID: "101",
			},
			Expect: &Event{
				Key:         "schedule",
				UID:         "101",
				Summary:     "summary",
				Description: "description",
				StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
			},
			Error: nil,
		},
		{
			Input: Input{
				Key: "event",
				UID: "101",
			},
			Expect: nil,
			Error:  ErrEventNotFound,
		},
		{
			Input: Input{
				Key: "schedule",
				UID: "102",
			},
			Expect: nil,
			Error:  ErrEventNotFound,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			event, err := cl.Get(test.Input.Key, test.Input.UID)
			if test.Error != nil {
				if err == nil {
					t.Fatal("error must not be nil")
				}
				if errors.Cause(err) != test.Error {
					t.Errorf("get error %s, want %s", errors.Cause(err).Error(), test.Error.Error())
				}
				return
			}
			if err != nil {
				t.Fatal("error event create", err.Error())
			}
			if !reflect.DeepEqual(event, test.Expect) {
				t.Errorf("get result %v, want %v", event, test.Expect)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	t.Parallel()

	errAlreadyRegistered := errors.New("error already registered")

	cl := newCalendarMock()
	cl.CreateMock = func(event *Event) error {
		uidByKey := map[string]string{"schedule": "101"}
		u, ok := uidByKey[event.Key]
		if ok && u == event.UID {
			return errAlreadyRegistered
		}

		return nil
	}

	type Test struct {
		Input *Event
		Error error
	}
	tests := []Test{
		{
			Input: &Event{
				Key:         "schedule",
				UID:         "102",
				Summary:     "summary",
				Description: "description",
				StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
			},
			Error: nil,
		},
		{
			Input: &Event{
				Key:         "event",
				UID:         "101",
				Summary:     "summary",
				Description: "description",
				StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
			},
			Error: nil,
		},
		{
			Input: &Event{
				Key:         "schedule",
				UID:         "101",
				Summary:     "summary",
				Description: "description",
				StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
			},
			Error: errAlreadyRegistered,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			err := cl.Create(test.Input)
			if test.Error != nil {
				if err == nil {
					t.Fatal("error must not be nil")
				}
				if errors.Cause(err) != test.Error {
					t.Errorf("get error %s, want %s", errors.Cause(err).Error(), test.Error.Error())
				}
				return
			}
			if err != nil {
				t.Fatal("error event create", err.Error())
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	cl := newCalendarMock()
	cl.UpdateMock = func(event *Event) error {
		uidByKey := map[string]string{"schedule": "100"}
		u, ok := uidByKey[event.Key]
		if !ok || u != event.UID {
			return ErrEventNotFound
		}

		return nil
	}

	type Test struct {
		Input *Event
		Error error
	}
	tests := []Test{
		{
			Input: &Event{
				Key:         "schedule",
				UID:         "101",
				Summary:     "summary",
				Description: "description",
				StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
			},
			Error: ErrEventNotFound,
		},
		{
			Input: &Event{
				Key:         "schedule",
				UID:         "102",
				Summary:     "summary",
				Description: "description",
				StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
			},
			Error: ErrEventNotFound,
		},
		{
			Input: &Event{
				Key:         "event",
				UID:         "101",
				Summary:     "summary",
				Description: "description",
				StartAt:     time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
				EndAt:       time.Date(2019, 8, 1, 23, 59, 59, 0, time.UTC),
			},
			Error: ErrEventNotFound,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			err := cl.Update(test.Input)
			if test.Error != nil {
				if err == nil {
					t.Fatal("error must not be nil")
				}
				if errors.Cause(err) != test.Error {
					t.Errorf("get error %s, want %s", errors.Cause(err).Error(), test.Error.Error())
				}
				return
			}
			if err != nil {
				t.Fatal("error event update", err.Error())
			}
		})
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	cl := newCalendarMock()
	cl.DeleteMock = func(key, uid string) error {
		uidByKey := map[string]string{"schedule": "101"}
		u, ok := uidByKey[key]
		if !ok || u != uid {
			return ErrEventNotFound
		}

		return nil
	}

	type Input struct {
		Key string
		UID string
	}
	type Test struct {
		Input Input
		Error error
	}
	tests := []Test{
		{
			Input: Input{
				Key: "schedule",
				UID: "101",
			},
			Error: nil,
		},
		{
			Input: Input{
				Key: "event",
				UID: "101",
			},
			Error: ErrEventNotFound,
		},
		{
			Input: Input{
				Key: "schedule",
				UID: "102",
			},
			Error: ErrEventNotFound,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case:%d", i), func(t *testing.T) {
			err := cl.Delete(test.Input.Key, test.Input.UID)
			if test.Error != nil {
				if err == nil {
					t.Fatal("error must not be nil")
				}
				if errors.Cause(err) != test.Error {
					t.Errorf("get error %s, want %s", errors.Cause(err).Error(), test.Error.Error())
				}
				return
			}
			if err != nil {
				t.Fatal("error event create", err.Error())
			}
		})
	}
}

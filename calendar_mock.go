package koyomi

import "github.com/pkg/errors"

type calendarMock struct {
	calendarImpl
	GetMock    func(key, uid string) (*Event, error)
	CreateMock func(event *Event) error
	UpdateMock func(event *Event) error
	DeleteMock func(key, uid string) error
}

func newCalendarMock() *calendarMock {
	return &calendarMock{
		calendarImpl: calendarImpl{},
		GetMock: func(key, uid string) (*Event, error) {
			return nil, errors.New("not implemented")
		},
		CreateMock: func(event *Event) error {
			return errors.New("not implmented")
		},
		UpdateMock: func(event *Event) error {
			return errors.New("not implmented")
		},
		DeleteMock: func(key, uid string) error {
			return errors.New("not implemented")
		},
	}
}

func (mock *calendarMock) Get(key, uid string) (*Event, error) {
	return mock.GetMock(key, uid)
}

func (mock *calendarMock) Create(event *Event) error {
	return mock.CreateMock(event)
}

func (mock *calendarMock) Update(event *Event) error {
	return mock.UpdateMock(event)
}

func (mock *calendarMock) Delete(key, uid string) error {
	return mock.DeleteMock(key, uid)
}

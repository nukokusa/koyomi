package koyomi

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

// Config is a koyomi config
type Config struct {
	CalendarID   string
	Email        string
	PrivateKey   []byte
	PrivateKeyID string
}

// Service is koyomi service
type Service struct {
	Config Config
}

// New create a koyomi Service from given Config
func New(conf Config) *Service {
	return &Service{
		Config: conf,
	}
}

// Event is XXX
type Event struct {
	Key         string
	UID         string
	Summary     string
	Description string
	StartAt     time.Time
	EndAt       time.Time
}

// Get is XXX
func (srv *Service) Get(key, uid string) (*Event, error) {
	cl, err := srv.newCalendar()
	if err != nil {
		return nil, errors.Wrap(err, "error new calendar")
	}
	event, err := cl.Get(key, uid)
	if err != nil {
		return nil, errors.Wrap(err, "error get")
	}

	log.Printf("key:%s, uid:%s, summary:%s, description:%s, startAt:%s, endAt:%s",
		event.Key,
		event.UID,
		event.Summary,
		event.Description,
		event.StartAt.Format(time.RFC3339),
		event.EndAt.Format(time.RFC3339),
	)

	return event, nil
}

// Create is XXX
func (srv *Service) Create(event *Event) error {
	cl, err := srv.newCalendar()
	if err != nil {
		return errors.Wrap(err, "error new calendar")
	}
	if err := validate(event); err != nil {
		return errors.Wrap(err, "error validate event")
	}
	if err := cl.Create(event); err != nil {
		return errors.Wrap(err, "error create")
	}

	return nil
}

// Update is XXX
func (srv *Service) Update(event *Event) error {
	cl, err := srv.newCalendar()
	if err != nil {
		return errors.Wrap(err, "error new calendar")
	}
	if err := validate(event); err != nil {
		return errors.Wrap(err, "error validate event")
	}
	if err := cl.Update(event); err != nil {
		return errors.Wrap(err, "error update")
	}

	return nil
}

// Delete is XXX
func (srv *Service) Delete(key, uid string) error {
	cl, err := srv.newCalendar()
	if err != nil {
		return errors.Wrap(err, "error new calendar")
	}
	if key == "" {
		return errors.New("error key is empty")
	}
	if uid == "" {
		return errors.New("error uid is empty")
	}

	if err := cl.Delete(key, uid); err != nil {
		return errors.Wrap(err, "error delete")
	}

	return nil
}

func (srv *Service) newCalendar() (Calendar, error) {
	conf := CalendarConfig{
		CalendarID:   srv.Config.CalendarID,
		Email:        srv.Config.Email,
		PrivateKey:   srv.Config.PrivateKey,
		PrivateKeyID: srv.Config.PrivateKeyID,
	}

	cl, err := NewCalendar(conf)
	if err != nil {
		return nil, errors.Wrap(err, "error new calendar")
	}

	return cl, nil
}

func validate(event *Event) error {
	if event.Key == "" {
		return errors.New("error key is empty")
	}
	if event.UID == "" {
		return errors.New("error uid is empty")
	}
	if event.StartAt.IsZero() {
		return errors.New("error start_at is zero")
	}
	if event.EndAt.IsZero() {
		return errors.New("error end_at is zero")
	}

	return nil
}

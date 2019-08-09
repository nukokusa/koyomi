package koyomi

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	gcal "google.golang.org/api/calendar/v3"
)

var (
	// ErrEventNotFound is XXX
	ErrEventNotFound = errors.New("error event not found")
)

// Calendar is XXX
type Calendar interface {
	Get(key, uid string) (*Event, error)
	Create(event *Event) error
	Update(event *Event) error
	Delete(key, uid string) error
}

// CalendarConfig is a Calendar Config
type CalendarConfig struct {
	CalendarID   string `yaml:"calendar_id"`
	Email        string `yaml:"email"`
	PrivateKey   []byte `yaml:"private_key"`
	PrivateKeyID string `yaml:"private_key_id"`
}

// NewCalendar create a calendar from given Config
func NewCalendar(conf CalendarConfig) (Calendar, error) {
	if conf.CalendarID == "" || conf.Email == "" || string(conf.PrivateKey) == "" || conf.PrivateKeyID == "" {
		return nil, errors.New("error invalid calendar config")
	}

	jwtConf := &jwt.Config{
		Email:        conf.Email,
		PrivateKeyID: conf.PrivateKeyID,
		PrivateKey:   conf.PrivateKey,
		TokenURL:     google.JWTTokenURL,
		Scopes: []string{
			gcal.CalendarScope,
			gcal.CalendarEventsScope,
		},
	}
	httpClient := jwtConf.Client(context.Background())

	srv, err := gcal.New(httpClient)
	if err != nil {
		return nil, errors.Wrap(err, "error create google calendar client")
	}

	return &calendarImpl{
		config:  conf,
		service: srv,
	}, nil
}

type calendarImpl struct {
	config  CalendarConfig
	service *gcal.Service
}

func (c *calendarImpl) Get(key, uid string) (*Event, error) {
	ev, err := c.get(key, uid)
	if err != nil {
		return nil, errors.Wrap(err, "error get event")
	}

	startAt, err := time.Parse(time.RFC3339, ev.Start.DateTime)
	if err != nil {
		return nil, errors.Wrap(err, "error parse start")
	}

	endAt, err := time.Parse(time.RFC3339, ev.End.DateTime)
	if err != nil {
		return nil, errors.Wrap(err, "error parse end")
	}

	return &Event{
		Key:         key,
		UID:         uid,
		Summary:     ev.Summary,
		Description: ev.Description,
		StartAt:     startAt,
		EndAt:       endAt,
	}, nil
}

func (c *calendarImpl) get(key, uid string) (*gcal.Event, error) {
	property := fmt.Sprintf("%s=%s", key, uid)
	evs, err := c.service.Events.List(c.config.CalendarID).PrivateExtendedProperty(property).Do()
	if err != nil {
		return nil, errors.Wrap(err, "error list event")
	}
	if len(evs.Items) == 0 {
		return nil, ErrEventNotFound
	}

	return evs.Items[0], nil
}

func (c *calendarImpl) Create(event *Event) error {
	registered, err := c.get(event.Key, event.UID)
	if err != nil && errors.Cause(err) != ErrEventNotFound {
		return errors.Wrap(err, "error get event")
	}
	if registered != nil {
		return errors.Errorf("error already registered. key:%s, uid:%s", event.Key, event.UID)
	}

	ev := &gcal.Event{
		Summary:     event.Summary,
		Description: event.Description,
		Start:       &gcal.EventDateTime{DateTime: event.StartAt.Format(time.RFC3339)},
		End:         &gcal.EventDateTime{DateTime: event.EndAt.Format(time.RFC3339)},
		ExtendedProperties: &gcal.EventExtendedProperties{
			Private: map[string]string{
				event.Key: event.UID,
			},
		},
	}

	if _, err := c.service.Events.Insert(c.config.CalendarID, ev).Do(); err != nil {
		return errors.Wrap(err, "error insert event")
	}

	return nil
}

func (c *calendarImpl) Update(event *Event) error {
	ev, err := c.get(event.Key, event.UID)
	if err != nil {
		return errors.Wrap(err, "error get event")
	}

	ev.Summary = event.Summary
	ev.Description = event.Description
	ev.Start = &gcal.EventDateTime{DateTime: event.StartAt.Format(time.RFC3339)}
	ev.End = &gcal.EventDateTime{DateTime: event.EndAt.Format(time.RFC3339)}

	if _, err = c.service.Events.Update(c.config.CalendarID, ev.Id, ev).Do(); err != nil {
		return errors.Wrap(err, "error update event")
	}

	return nil
}

func (c *calendarImpl) Delete(key, uid string) error {
	ev, err := c.get(key, uid)
	if err != nil {
		return errors.Wrap(err, "error get event")
	}

	if err := c.service.Events.Delete(c.config.CalendarID, ev.Id).Do(); err != nil {
		return errors.Wrap(err, "error delete event")
	}

	return nil
}

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/hashicorp/logutils"
	"github.com/nukokusa/koyomi"
	"gopkg.in/yaml.v2"
)

var (
	config      string
	key         string
	uid         string
	summary     string
	description string
	startAtStr  string
	endAtStr    string
	logLevel    string
	startAt     time.Time
	endAt       time.Time
)

// Config is a koyomi config
type Config struct {
	CalendarID   string `yaml:"calendar_id"`
	Email        string `yaml:"email"`
	PrivateKey   string `yaml:"private_key"`
	PrivateKeyID string `yaml:"private_key_id"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("[ERROR] required sub command")
	}

	f := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	f.StringVar(&config, "config", "config.yml", "config yaml file")
	f.StringVar(&key, "key", "", "event key")
	f.StringVar(&uid, "uid", "", "event uid")
	f.StringVar(&summary, "summary", "", "event summary")
	f.StringVar(&description, "description", "", "event description")
	f.StringVar(&startAtStr, "startat", "", "event start at")
	f.StringVar(&endAtStr, "endat", "", "event end at")
	f.StringVar(&logLevel, "loglevel", "INFO", "logging level: INFO, WARN, ERROR")
	f.Parse(os.Args[2:])

	logFilter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(logLevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(logFilter)

	data, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal("[ERROR] open config file", err.Error())
	}

	var conf Config
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatal("[ERROR] unmarshal config file", err.Error())
	}

	if startAtStr != "" {
		startAt, err = time.Parse(time.RFC3339, startAtStr)
		if err != nil {
			log.Fatal("[ERROR] failed to parse start at", err.Error())
		}
	}
	if endAtStr != "" {
		endAt, err = time.Parse(time.RFC3339, endAtStr)
		if err != nil {
			log.Fatal("[ERROR] failed to parse end at", err.Error())
		}
	}

	switch os.Args[1] {
	case "create":
		create(&conf)
	case "update":
		update(&conf)
	case "delete":
		delete(&conf)
	default:
		log.Fatalf("[ERROR] invalid sub command: %s", os.Args[1])
	}
}

func service(conf *Config) *koyomi.Service {
	koyomiConf := koyomi.Config{
		CalendarID:   conf.CalendarID,
		Email:        conf.Email,
		PrivateKey:   []byte(conf.PrivateKey),
		PrivateKeyID: conf.PrivateKeyID,
	}

	return koyomi.New(koyomiConf)
}

func create(conf *Config) {
	srv := service(conf)

	event := &koyomi.Event{
		Key:         key,
		UID:         uid,
		Summary:     summary,
		Description: description,
		StartAt:     startAt,
		EndAt:       endAt,
	}

	if err := srv.Create(event); err != nil {
		log.Fatalf("[ERROR] koyomi.Create: failed to create: %s", err)
	}
	return
}

func update(conf *Config) {
	srv := service(conf)

	event := &koyomi.Event{
		Key:         key,
		UID:         uid,
		Summary:     summary,
		Description: description,
		StartAt:     startAt,
		EndAt:       endAt,
	}
	if err := srv.Update(event); err != nil {
		log.Fatalf("[ERROR] koyomi.Update: failed to update: %s", err)
	}
}

func delete(conf *Config) {
	srv := service(conf)

	if err := srv.Delete(key, uid); err != nil {
		log.Fatalf("[ERROR] koyomi.Delete: failed to delete: %s", err)
	}
}

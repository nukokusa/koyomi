package koyomi

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/hashicorp/logutils"
)

type CLIOptions struct {
	Create         *CreateOption    `cmd:"" help:"create event"`
	Update         *UpdateOption    `cmd:"" help:"update event"`
	Delete         *DeleteOption    `cmd:"" help:"delete event"`
	CredentialPath string           `help:"JSON credential file for access to calendar" name:"credential" default:"credential.json"`
	LogLevel       string           `help:"logging level: DEBUG, INFO, WARN, ERROR" enum:"DEBUG,INFO,WARN,ERROR" name:"loglevel" default:"INFO"`
	Version        kong.VersionFlag `help:"show Version" name:"version" short:"v"`
}

func Run(ctx context.Context, args []string) error {
	var opts CLIOptions
	parser, err := kong.New(&opts, kong.Vars{"version": "app " + Version})
	if err != nil {
		return err
	}
	kctx, err := parser.Parse(args)
	if err != nil {
		return err
	}

	logFilter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(opts.LogLevel),
		Writer:   os.Stderr,
	}
	log.SetOutput(logFilter)
	log.Println("[DEBUG] parsed args")

	s, err := New(ctx, opts.CredentialPath)
	if err != nil {
		return err
	}
	cmd := strings.Fields(kctx.Command())[0]
	return s.Dispatch(ctx, cmd, &opts)
}

func (k *Koyomi) Dispatch(ctx context.Context, command string, opts *CLIOptions) error {
	switch command {
	case "create":
		return k.Create(ctx, opts.Create)
	case "update":
		return k.Update(ctx, opts.Update)
	case "delete":
		return k.Delete(ctx, opts.Delete)
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

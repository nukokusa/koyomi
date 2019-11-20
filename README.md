# Koyomi

[![Test Status](https://github.com/nukokusa/koyomi/workflows/test/badge.svg?branch=master)][actions]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[actions]: https://github.com/nukokusa/koyomi/actions?workflow=test
[license]: https://github.com/nukokusa/koyomi/blob/master/LICENSE

Koyomi is a simple schedule client for Google Calendar.

## Example

```golang
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nukokusa/koyomi"
)

func main() {
	conf := koyomi.Config{
		CalendarID:   "xxx",
		Email:        "xxx",
		PrivateKey:   []byte("xxx"),
		PrivateKeyID: "xxx",
	}

	srv := koyomi.New(conf)

	event := &koyomi.Event{
		Key:         "key",
		UID:         "1001",
		Summary:     "summary",
		Description: "description",
		StartAt:     time.Now(),
		EndAt:       time.Now(),
	}
	if err := srv.Create(event); err != nil {
		fmt.Printf("failed to create event: %s", err)
		os.Exit(1)
	}

	event.Summary = "updated summary"
	if err := srv.Update(event); err != nil {
		fmt.Printf("failed to update event: %s", err)
		os.Exit(1)
	}

	if err := srv.Delete("key", "1001"); err != nil {
		fmt.Printf("failed to delete event: %s", err)
		os.Exit(1)
	}
}
```

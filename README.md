# koyomi

[![CircleCI](https://circleci.com/gh/nukokusa/koyomi/tree/master.svg?style=svg)](https://circleci.com/gh/nukokusa/koyomi/tree/master)

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

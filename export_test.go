package koyomi

import "io"

var (
	NewCalendarServiceMock = newCalendarServiceMock
)

type ExportKoyomi = Koyomi

func (k *ExportKoyomi) SetCalendarService(cs CalendarService) {
	k.cs = cs
}

func (k *ExportKoyomi) SetStdout(w io.Writer) {
	k.stdout = w
}

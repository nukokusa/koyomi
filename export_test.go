package koyomi

var (
	NewCalendarServiceMock = newCalendarServiceMock
)

type ExportKoyomi = Koyomi

func (k *ExportKoyomi) SetCalendarService(cs CalendarService) {
	k.cs = cs
}

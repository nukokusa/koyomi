# koyomi

[![CI Status](https://github.com/nukokusa/koyomi/actions/workflows/ci.yml/badge.svg)][actions]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[actions]: https://github.com/nukokusa/koyomi/actions?workflow=ci
[license]: https://github.com/nukokusa/koyomi/blob/master/LICENSE

Koyomi is a simple schedule client for Google Calendar.

## Usage

```
Usage: koyomi <command>

Flags:
  -h, --help                            Show context-sensitive help.
      --credential="credential.json"    JSON credential file for access to calendar
      --loglevel="INFO"                 Logging level: DEBUG, INFO, WARN, ERROR
  -v, --version                         Show Version

Commands:
  list --calendar-id=STRING --start-time=STRING --end-time=STRING
    List events

  create --calendar-id=STRING --start-time=STRING --end-time=STRING
    Creates an event

  update --calendar-id=STRING --event-id=STRING
    Updates an event

  delete --calendar-id=STRING --event-id=STRING
    Deletes an event
```

### List Events

```
Usage: koyomi list --calendar-id=STRING --start-time=STRING --end-time=STRING

List events

Flags:
      --calendar-id=STRING              Calendar identifier
  -s, --start-time=STRING               The start time of the event
  -e, --end-time=STRING                 The end time of the event
```

### Creates an Event

```
Usage: koyomi create --calendar-id=STRING --start-time=STRING --end-time=STRING

Creates an event

Flags:
      --calendar-id=STRING              Calendar identifier
      --summary=STRING                  Title of the event
      --description=STRING              Descriptuon of the event
  -s, --start-time=STRING               The start time of the event
  -e, --end-time=STRING                 The end time of the event
```

### Updates an Event

```
Usage: koyomi update --calendar-id=STRING --event-id=STRING

Updates an event

Flags:
      --calendar-id=STRING              Calendar identifier
      --event-id=STRING                 Identifier of the event
      --summary=STRING                  Title of the event
      --description=STRING              Description of the event
  -s, --start-time=STRING               The start time of the event
  -e, --end-time=STRING                 The end time of the event
```

### Deletes an Event

```
Usage: koyomi delete --calendar-id=STRING --event-id=STRING

Deletes an event

Flags:
      --calendar-id=STRING              Calendar identifier
      --event-id=STRING                 Identifier of the event
```

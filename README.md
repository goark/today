# [today] -- Show information for today.

[![ci status](https://github.com/goark/today/workflows/ci/badge.svg)](https://github.com/goark/today/actions)
[![build status](https://github.com/goark/today/workflows/build/badge.svg)](https://github.com/goark/today/actions)
[![CodeQL status](https://github.com/goark/today/workflows/CodeQL/badge.svg)](https://github.com/goark/today/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/today/main/LICENSE)
[![GitHub release](http://img.shields.io/github/release/goark/today.svg)](https://github.com/goark/today/releases/latest)

This package requires Go 1.27 or later.

## Overview

today prints date information for a specified day (or today by default).

The output includes:

- Gregorian date and weekday
- Japanese era and year in era
- Zodiac information
- Optional holiday information
- Optional solar-term information
- Optional custom events from a JSON file

## Usage

```
$ today --help
Usage of today

Usage:
  today [flags] [yyyy-mm-dd]

Flags:
  -a, --all-events          show all events information
      --debug               enable debug mode
      --event-file string   path to the event file (default "~/.config/today/event.json")
  -h, --holiday             show holiday information
  -j, --json                output information in JSON format
  -e, --other-events        show other events information
  -s, --solar-term          show solar term information
      --temp-dir string     path to the temporary directory (default "~/.cache/today")
  -v, --version             show version information
```

## Date Argument

The date argument is optional.

- No argument: use the current date
- With argument: use one of the supported formats

Supported formats:

- `yyyy-mm-dd` (example: `2026-09-09`)
- `yyyymmdd` (example: `20260909`)
- RFC3339 (example: `2026-09-09T00:00:00+09:00`)

## Output Modes

### Text output (default)

```sh
today 2026-09-25
```

Example (shape):

```
2026-09-25 Friday
令和8年 (丙午) 9月25日 金曜日 (壬寅)
```

### JSON output

Use `--json` or `-j` to print JSON.

```sh
today --json 2026-09-25
```

Example (shape):

```json
{
  "year": 2026,
  "month": 9,
  "day": 25,
  "weekday": "Friday",
  "japanese_era": "令和",
  "year_in_era": 8,
  "weekday_jp": "金曜日",
  "zodiac_year": "丙午",
  "zodiac_day": "壬寅",
  "events": []
}
```

## Event Options

The following flags control optional event sources:

- `--holiday` / `-h`: import holiday events
- `--solar-term` / `-s`: import solar-term, moon-phase, and eclipse events
- `--other-events` / `-e`: import custom events from a JSON file
- `--all-events` / `-a`: enable all event sources above

### Custom event file

Set custom event file path with `--event-file`.

```sh
today --other-events --event-file ./event.json 2026-09-25
```

Event file format:

```json
[
  {
    "start": "2026-09-25",
    "end": "2026-09-25",
    "title": "中秋の名月"
  }
]
```

Example (plain text):

```
2026-09-25 Friday
令和8年 (丙午) 9月25日 金曜日 (壬寅)
中秋の名月
```

Example (JSON format):

```json
{
  "year": 2026,
  "month": 9,
  "day": 25,
  "weekday": "Friday",
  "japanese_era": "令和",
  "year_in_era": 8,
  "weekday_jp": "金曜日",
  "zodiac_year": "丙午",
  "zodiac_day": "壬寅",
  "events": [
    "中秋の名月"
  ]
}
```

Rules:

- `start` and `end` accept the same date formats as CLI input
- event is included when target date is in the inclusive range [`start`, `end`]

## Temp Directory

`--temp-dir` controls the cache/temp directory used when fetching external calendar data.

```sh
today --holiday --temp-dir /tmp/today-cache 2026-09-09
```

Default value:

- Linux: `~/.cache/today`

## Debug

Use `--debug` to print structured error details.

```sh
today --debug --event-file ./not-found.json --other-events
```

[today]: https://github.com/goark/today "goark/today: Show information for today"

package today

import (
	"fmt"
	"io"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/koyomi"
	"github.com/goark/koyomi/value"
	"github.com/goark/koyomi/zodiac"
	"github.com/goark/today/internal/ecode"
	"github.com/goark/today/internal/events"
)

var (
	holidayEventTitlesFetcher = fetchHolidayEventTitles
	solarTermEventTitlesFetcher = fetchSolarTermEventTitles
)

// InfoDate represents detailed information about a specific date.
type InfoDate struct {
	date        value.DateJp
	Year        int      `json:"year"`
	Month       int      `json:"month"`
	Day         int      `json:"day"`
	Weekday     string   `json:"weekday"`
	JapaneseEra string   `json:"japanese_era"`
	YearInEra   int      `json:"year_in_era"`
	WeekdayJp   string   `json:"weekday_jp"`
	ZodiacYear  string   `json:"zodiac_year"`
	ZodiacDay   string   `json:"zodiac_day"`
	Events      []string `json:"events"`
}

// NewInfoDate creates a new InfoDate instance for the given date.
func NewInfoDate(dt value.DateJp) *InfoDate {
	era, yr := dt.YearEra()
	ky, sy := zodiac.ZodiacYearNumber(dt.Year())
	kd, sd := zodiac.ZodiacDayNumber(dt)
	return &InfoDate{
		date:        dt,
		Year:        dt.Year(),
		Month:       int(dt.Month()),
		Day:         dt.Day(),
		Weekday:     dt.Weekday().String(),
		JapaneseEra: era.String(),
		YearInEra:   yr,
		WeekdayJp:   dt.WeekdayJp().StringJp(),
		ZodiacYear:  fmt.Sprintf("%v%v", ky, sy),
		ZodiacDay:   fmt.Sprintf("%v%v", kd, sd),
		Events:      []string{},
	}
}

// Strings returns a slice of strings representing the detailed information of the date.
func (i *InfoDate) Strings() []string {
	if i == nil {
		return []string{}
	}
	ss := []string{
		fmt.Sprintf(
			"%04d-%02d-%02d %s",
			i.Year,
			i.Month,
			i.Day,
			i.Weekday,
		),
		fmt.Sprintf(
			"%v%v年 (%v) %d月%v日 %s (%v)",
			i.JapaneseEra,
			i.YearInEra,
			i.ZodiacYear,
			i.Month,
			i.Day,
			i.WeekdayJp,
			i.ZodiacDay,
		),
	}
	if len(i.Events) > 0 {
		ss = append(ss, i.Events...)
	}
	return ss
}

// String returns a string representation of the detailed information of the date.
func (i *InfoDate) String() string {
	return strings.Join(i.Strings(), "\n")
}

func (i *InfoDate) ImportHoliday() error {
	if i == nil {
		return errs.Wrap(ecode.ErrNullPointer)
	}
	titles, err := holidayEventTitlesFetcher(i.date)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("date", i.date.String()))
	}
	i.Events = append(i.Events, titles...)
	return nil
}

func (i *InfoDate) ImportSolarTerm() error {
	if i == nil {
		return errs.Wrap(ecode.ErrNullPointer)
	}
	titles, err := solarTermEventTitlesFetcher(i.date)
	if err != nil {
		return errs.Wrap(err, errs.WithContext("date", i.date.String()))
	}
	i.Events = append(i.Events, titles...)
	return nil
}

func fetchHolidayEventTitles(dt value.DateJp) ([]string, error) {
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.Holiday),
		koyomi.WithStartDate(dt),
		koyomi.WithEndDate(dt),
	).Get()
	if err != nil {
		return nil, err
	}
	titles := make([]string, 0, len(k.Events()))
	for _, e := range k.Events() {
		titles = append(titles, e.Title)
	}
	return titles, nil
}

func fetchSolarTermEventTitles(dt value.DateJp) ([]string, error) {
	k, err := koyomi.NewSource(
		koyomi.WithCalendarID(koyomi.SolarTerm, koyomi.MoonPhase, koyomi.Eclipse),
		koyomi.WithStartDate(dt),
		koyomi.WithEndDate(dt),
	).Get()
	if err != nil {
		return nil, err
	}
	titles := make([]string, 0, len(k.Events()))
	for _, e := range k.Events() {
		titles = append(titles, e.Title)
	}
	return titles, nil
}

// ImportEvents imports events from the provided reader and adds the titles of events containing the date to the InfoDate instance.
func (i *InfoDate) ImportEvents(r io.Reader) error {
	evts, err := events.ImportEvents(r)
	if err != nil {
		return errs.Wrap(err)
	}
	ss, err := events.GetEventsContaining(i.date, evts)
	if err != nil {
		return errs.Wrap(err)
	}
	i.Events = append(i.Events, ss...)
	return nil
}

/* Copyright 2026 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

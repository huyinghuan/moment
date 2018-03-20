package moment

import (
	"time"
)

type Mymonent struct {
	Now        time.Time
	StartDay   time.Time
	EndDay     time.Time
	StartWeek  time.Time
	EndWeek    time.Time
	StartMonth time.Time
	EndMonth   time.Time
	StartYear  time.Time
	EndYear    time.Time
}

func New(date time.Time) *Mymonent {
	now := date
	location := now.Location()
	weekday := now.Weekday()
	year, month, day := now.Date()

	start := time.Date(year, month, day, 0, 0, 0, 0, location)
	end := start.AddDate(0, 0, 1).Add(time.Nanosecond * -1)

	startMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
	endMonth := startMonth.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	startYear := time.Date(year, 1, 1, 0, 0, 0, 0, location)
	endYear := startYear.AddDate(1, 0, 0).Add(time.Nanosecond * -1)

	startWeek := start.AddDate(0, 0, int(weekday)*-1)
	endWeek := startWeek.AddDate(0, 0, 7).Add(time.Nanosecond * -1)

	return &Mymonent{
		Now:        now,
		StartDay:   start,
		EndDay:     end,
		StartMonth: startMonth,
		EndMonth:   endMonth,
		EndYear:    endYear,
		StartWeek:  startWeek,
		EndWeek:    endWeek,
	}
}

func NewUinx(unix int64) *Mymonent {
	return New(time.Unix(unix, 0))
}

//New
func Now() *Mymonent {
	return New(time.Now())
}

//Today get today begin and end time
func (moment *Mymonent) Today() (time.Time, time.Time) {
	return moment.StartDay, moment.EndDay
}

//TodayUnix get today begin and end unix time
func (moment *Mymonent) TodayUnix() (int64, int64) {
	start, end := moment.Today()
	return start.Unix(), end.Unix()
}

//LastPeriodDay 最近的一段时间
func (moment *Mymonent) LastPeriodDay(day int) time.Time {
	return moment.Now.AddDate(0, 0, day*-1)
}

//LastPeriodDay 最近的一段时间unix
func (moment *Mymonent) LastPeriodDayUinx(day int) int64 {
	return moment.LastPeriodDay(day).Unix()
}

//Format  format date to string
/*
	YYYY	2014	4 or 2 digit year
	M MM	1..12	Month number
	D DD	1..31	Day of month
	Do	1st..31st	Day of month with ordinal
	DDD DDDD	1..365	Day of year
	X	1410715640.579	Unix timestamp
	x	1410715640579	Unix ms timestamp
*/
func (moment *Mymonent) Format(format string) string {
	//Mon Jan 2 15:04:05 MST 2006
	return ""
}

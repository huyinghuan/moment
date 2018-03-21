package moment

import (
	"log"
	"strings"
	"time"
)

type Mymonent struct {
	Unix int64
	Now  time.Time
}

//three contructor

//New Get mymoment ref
func New(date time.Time) *Mymonent {
	now := date
	// location := now.Location()
	// weekday := now.Weekday()
	// year, month, day := now.Date()

	// start := time.Date(year, month, day, 0, 0, 0, 0, location)
	// end := start.AddDate(0, 0, 1).Add(time.Nanosecond * -1)

	// startMonth := time.Date(year, month, 1, 0, 0, 0, 0, location)
	// endMonth := startMonth.AddDate(0, 1, 0).Add(time.Nanosecond * -1)

	// startYear := time.Date(year, 1, 1, 0, 0, 0, 0, location)
	// endYear := startYear.AddDate(1, 0, 0).Add(time.Nanosecond * -1)

	// startWeek := start.AddDate(0, 0, int(weekday)*-1)
	// endWeek := startWeek.AddDate(0, 0, 7).Add(time.Nanosecond * -1)

	return &Mymonent{
		Unix: now.Unix(),
		Now:  now,
		// StartDay:   start,
		// EndDay:     end,
		// StartMonth: startMonth,
		// EndMonth:   endMonth,
		// EndYear:    endYear,
		// StartWeek:  startWeek,
		// EndWeek:    endWeek,
	}
}

// NewFromUinx use unix time create a  mymoment ref
func NewFromUinx(unix int64) *Mymonent {
	return New(time.Unix(unix, 0))
}

//Now Get now mymonent struct ref
func Now() *Mymonent {
	return New(time.Now())
}

// ====================

//Today get today begin and end time
func (moment *Mymonent) Today() (*Mymonent, *Mymonent) {
	year, month, day := moment.Now.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, moment.Now.Location())
	end := start.AddDate(0, 0, 1).Add(time.Nanosecond * -1)
	return New(start), New(end)
}

//ThisMonth get today begin and end time
func (moment *Mymonent) ThisMonth() (*Mymonent, *Mymonent) {
	now := moment.Now
	year, month, _ := now.Date()
	startMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	endMonth := startMonth.AddDate(0, 1, 0).Add(time.Nanosecond * -1)
	return New(startMonth), New(endMonth)
}

//ThisYear get today begin and end time
func (moment *Mymonent) ThisYear() (*Mymonent, *Mymonent) {
	now := moment.Now
	year, _, _ := now.Date()
	startYear := time.Date(year, 1, 1, 0, 0, 0, 0, now.Location())
	endYear := startYear.AddDate(1, 0, 0).Add(time.Nanosecond * -1)
	return New(startYear), New(endYear)
}

//LastDay 几天以前
func (moment *Mymonent) LastDay(day int) *Mymonent {
	return New(moment.Now.AddDate(0, 0, day*-1))
}

//Uinx 获取时间戳
func (moment *Mymonent) Uinx() int64 {
	return moment.Now.Unix()
}

//Format  format date to string
/*
	YYYY	2014	4 or 2 digit year
	MM	1..12	Month number
	DD 1..31	Day of month
	HH 0..23	Hours (24 hour time)
	mm 0..59	Minutes
	ss 0..59	Seconds
	2006-01-02T15:04:05
*/
func (moment *Mymonent) Format(format string) string {
	//Mon Jan 2 15:04:05 MST 2006
	//2006-01-02T15:04:05
	format = strings.Replace(format, "YYYY", "2006", -1)
	format = strings.Replace(format, "MM", "01", -1)
	format = strings.Replace(format, "DD", "02", -1)
	format = strings.Replace(format, "HH", "15", -1)
	format = strings.Replace(format, "mm", "04", -1)
	format = strings.Replace(format, "ss", "05", -1)
	return moment.Now.Format(format)
}

func (moment *Mymonent) AddDate(years int, month int, day int) *Mymonent {
	return New(moment.Now.AddDate(years, month, day))
}

func (moment *Mymonent) SubDate(years int, month int, day int) *Mymonent {
	return New(moment.Now.AddDate(years*-1, month*-1, day*-1))
}

// such as "300ms", "-1.5h" or "2h45m1s"
func (moment *Mymonent) Add(duration string) *Mymonent {
	complex, err := time.ParseDuration(duration)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return New(moment.Now.Add(complex))
}

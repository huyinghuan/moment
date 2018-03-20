package moment

import (
	"log"
	"testing"
)

func Test_Today(t *testing.T) {
	now := New()

	log.Println("start day", now.StartDay)
	log.Println("end day", now.EndDay)
	log.Println("start week", now.StartWeek)
	log.Println("end week", now.EndWeek)
	log.Println("start month", now.StartMonth)
	log.Println("end month", now.EndMonth)
	log.Println("start year", now.StartYear)
	log.Println("end year", now.EndYear)
}

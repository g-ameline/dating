package dating

import (
	"log"
	"math/rand"
	"time"
)

const DateTime = "2006-01-02 15:04:05"
const DateOnly = "2006-01-02"

func Parse_date_from_database(record string) time.Time {
	layout := DateTime
	date, err := time.Parse(layout, record)
	if err != nil {
		log.Fatalln(err, "could not parse the time from db")
	}
	return date
}

func Just_the_date(date time.Time) string {
	date_only := date.Format(DateOnly)
	return date_only
}

func Rdm_date() string {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2099, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	date := time.Unix(sec, 0)
	date_text := date.Format(DateTime)
	return date_text
}

func That_moment() string {
	return time.Now().Format(DateTime)
}

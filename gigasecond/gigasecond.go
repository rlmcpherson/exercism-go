package gigasecond

import (
	"log"
	"time"
)

const TestVersion = 2

var Birthday time.Time

func init() {
	var err error
	Birthday, err = time.Parse("2006-01-02", "1900-10-10")
	if err != nil {
		log.Fatal(err)
	}

}

func AddGigasecond(t time.Time) time.Time {
	return t.Add(1E9 * time.Second)
}

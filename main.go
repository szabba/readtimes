package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

var (
	format string
	sched  Schedule

	errReadUnsupported = errors.New("unsupported platform")
)

func main() {

	flag.StringVar(&format, "fmt", "15 04", "Time format to use. Specified as a Go time package layout.")

	flag.Var(&sched, "sched", "A cron-like schedule. Minutes and hours are mandatory. Optionally, seconds can be specified.")
	sched.MustSet("0 *")

	flag.Parse()

	for {
		msg := fmt.Sprintf("It is now %s.", time.Now().Format(format))

		err := read(msg)
		if errors.Is(err, errReadUnsupported) {
			log.Fatal(err)
		}
		if err != nil {
			log.Print(err)
		}

		next := sched.Next()
		left := time.Until(next)

		<-time.After(left)
	}
}

type Schedule struct {
	raw    string
	parsed cron.Schedule
}

func (s *Schedule) Next() time.Time {
	now := time.Now()
	return s.parsed.Next(now)
}

func (s *Schedule) MustSet(raw string) {
	err := s.Set(raw)
	if err != nil {
		panic(err)
	}
}

var _ flag.Value = new(Schedule)

func (s *Schedule) Set(raw string) error {
	p := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour)
	parsed, err := p.Parse(raw)

	if err != nil {
		return err
	}

	s.raw, s.parsed = raw, parsed
	return nil
}

func (s *Schedule) String() string { return s.raw }

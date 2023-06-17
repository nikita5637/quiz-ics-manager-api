package generator

import (
	"bytes"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/google/uuid"
)

const (
	tz = "Europe/Moscow"
)

// Generate ...
func (g *Generator) Generate(summary, address, description, url string, gameDateTime time.Time) ([]byte, error) {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}

	now := time.Now().In(loc)
	gameDateTime = gameDateTime.In(loc)

	cal := ics.NewCalendar()
	cal.SetVersion("2.0")
	cal.SetCalscale("GREGORIAN")
	cal.SetTzid(tz)

	event := cal.AddEvent(uuid.NewString())
	cal.AddVEvent(event)

	event.SetCreatedTime(now)
	event.SetDescription(description)
	event.SetEndAt(gameDateTime.Add(2 * time.Hour))
	event.SetDtStampTime(now)
	event.SetStartAt(gameDateTime)
	event.SetLocation(string(bytes.Replace([]byte(address), []byte{0xd}, []byte{}, -1)))
	event.SetSequence(0)
	event.SetSummary(summary)
	event.SetTimeTransparency(ics.TransparencyOpaque)
	event.SetURL(url, ics.WithValue("URI"))

	alarm := event.AddAlarm()
	alarm.SetAction(ics.ActionDisplay)
	alarm.SetTrigger("-PT2H")
	alarm.SetProperty(ics.ComponentPropertyDescription, "Reminder")

	return []byte(cal.Serialize()), nil

}

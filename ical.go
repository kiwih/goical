package goical

import (
	"io"
	"text/template"
	"time"
)

type IcalEventable interface {
	GetUId() string
	GetStartTime() time.Time
	GetEndTime() *time.Time //allows for nil, meaning that it won't be specified
	GetOrganizerName() string
	GetOrganizerEmail() string
	GetLocation() string
	GetSummary() string
}

type Ical struct {
	Events []IcalEventable
}

const (
	IcalTimeFormat = "20060102T150405Z0700"
)

func NewIcal() *Ical {
	return new(Ical)
}

func (i *Ical) AddEvent(e IcalEventable) {
	i.Events = append(i.Events, e)
}

func (i *Ical) Write(w io.Writer) error {
	return icalTemplate.Execute(w, i)
}

func FormatIcalTime(t time.Time) string {
	return t.In(time.UTC).Format(IcalTimeFormat)
}

var funcMap = template.FuncMap{
	"FormatIcalTime": FormatIcalTime,
}

var icalTemplate = template.Must(template.
	New("icalTemplate").
	Funcs(funcMap).
	Parse(`BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN{{range $event := .Events}}
BEGIN:VEVENT
UID:{{$event.GetUId}}
ORGANIZER;CN={{$event.GetOrganizerName}}:MAILTO={{$event.GetOrganizerEmail}}
DTSTART:{{FormatIcalTime $event.GetStartTime}}{{if $event.GetEndTime}}
DTEND:{{FormatIcalTime $event.GetEndTime}}{{end}}
SUMMARY:{{$event.GetSummary}}
LOCATION:{{$event.GetLocation}}
END:VEVENT{{end}}
END:VCALENDAR`))

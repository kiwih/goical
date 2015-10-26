package ical

import (
	"bytes"
	"testing"
	"time"
)

//A mock event for us to use
type icalTestEvent struct {
	UId            string
	StartTime      time.Time
	EndTime        time.Time
	OrganizerName  string
	OrganizerEmail string
	Location       string
	Summary        string
}

func (e *icalTestEvent) GetUId() string {
	return e.UId
}

func (e *icalTestEvent) GetStartTime() time.Time {
	return e.StartTime
}

func (e *icalTestEvent) GetEndTime() *time.Time {
	return &e.EndTime
}

func (e *icalTestEvent) GetOrganizerName() string {
	return e.OrganizerName
}

func (e *icalTestEvent) GetOrganizerEmail() string {
	return e.OrganizerEmail
}

func (e *icalTestEvent) GetLocation() string {
	return e.Location
}

func (e *icalTestEvent) GetSummary() string {
	return e.Summary
}

var testEvent = icalTestEvent{
	UId:            "abcd-1234-5678-9",
	StartTime:      time.Date(2015, 01, 02, 3, 4, 5, 0, time.UTC),
	EndTime:        time.Date(2015, 01, 02, 4, 5, 6, 0, time.UTC),
	OrganizerName:  "kiwih",
	OrganizerEmail: "testing@example.com",
	Location:       "Auckland, NZ",
	Summary:        "This is a test event. All good things in life have been thoroughly tested.",
}

var testingIcal *Ical

var expectedOutput = `BEGIN:VCALENDAR
VERSION:2.0
CALSCALE:GREGORIAN
BEGIN:VEVENT
UID:abcd-1234-5678-9
ORGANIZER;CN=kiwih:MAILTO=testing@example.com
DTSTART:20150102T030405Z
DTEND:20150102T040506Z
SUMMARY:This is a test event. All good things in life have been thoroughly tested.
LOCATION:Auckland, NZ
END:VEVENT
END:VCALENDAR`

func TestNew(t *testing.T) {
	testingIcal = NewIcal()
	if testingIcal == nil {
		t.Fatalf("Failed to make New Ical")
	}
}

func TestAddEvent(t *testing.T) {
	testingIcal.AddEvent(&testEvent)
	if testingIcal.Events[0].GetUId() != testEvent.UId {
		t.Fatalf("Failed to append ical event")
	}
}

func TestWrite(t *testing.T) {
	testingBuffer := &bytes.Buffer{}
	if err := testingIcal.Write(testingBuffer); err != nil {
		t.Fatalf("Failed to write out ical")
	}
	output := string(testingBuffer.Bytes())
	for i := 0; i < len(output); i++ {
		if i > len(expectedOutput) {
			t.Fatalf("Output longer than expected output (i=%v)", i)
		}
		if output[i] != expectedOutput[i] {
			t.Fatalf("Output of test does not match expected at i=%v\n\nExpected:\n%s\n\nOutput:\n%s\n\n", i, expectedOutput, output)
		}
	}

}

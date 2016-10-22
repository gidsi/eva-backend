package main

import (
	"github.com/gidsi/ics-golang"
	"log"
)

func getCalendars() {
	spaceDataArray := readSpacedata()

	for _, spaceData := range spaceDataArray {
		parser := ics.New()
		parserChan := parser.GetInputChan()
		parserChan <- spaceData.Feeds.Calendar.Url
		outputChan := parser.GetOutputChan()
		calendar := Calendar{}
		calendar.Space = spaceData.Space
		events := []Event{}
		go func() {
			for event := range outputChan {
				events = append(events, mapEventObject(event))
			}
		}()
		parser.Wait()
		calendar.Events = events
		writeCalendar(calendar)
		log.Println("calendar write done for " + spaceData.Space)
	}
}

func mapEventObject(event *ics.Event) Event {
	eventData := Event{}

	eventData.Start = event.GetStart()
	eventData.ImportedId = event.GetImportedID()
	eventData.Status = event.GetStatus()
	eventData.Description = event.GetDescription()
	eventData.Location = event.GetLocation()
	eventData.Summary = event.GetSummary()
	eventData.Rrule = event.GetRRule()
	eventData.Url = event.GetUrl()
	eventData.Class = event.GetClass()
	eventData.Sequence = event.GetSequence()
	eventData.WholeDayEvent = event.GetWholeDayEvent()

	return eventData
}
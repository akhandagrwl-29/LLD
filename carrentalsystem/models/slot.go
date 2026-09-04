package models

import "time"

type Slot struct {
	Id         string
	StartTime  time.Time
	EndTime    time.Time
	IsOccupied bool
}

func NewSlot(id string, startTime, endTime time.Time) *Slot {
	return &Slot{
		Id:         id,
		StartTime:  startTime,
		EndTime:    endTime,
		IsOccupied: false,
	}
}

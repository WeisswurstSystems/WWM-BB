package adapter

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/meeting/application"
)

type MeetingInteractor interface {
	CloseMeeting(e application.CloseMeeting) error
	CreateMeeting(e application.CreateMeeting) error
	PutProduct(e application.PutProduct) error
	RemoveProduct(e application.RemoveProduct) error
	SetBuyer(e application.SetBuyer) error
	SetPlace(e application.SetPlace) error
}

type CommandHandler struct {
	MeetingInteractor MeetingInteractor
}

type QueryHandler struct {
	MeetingStore meeting.Store
}

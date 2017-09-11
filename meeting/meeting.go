package meeting

import (
	"github.com/WeisswurstSystems/WWM-BB/wwm"
	"net/http"
	"time"
	"github.com/WeisswurstSystems/WWM-BB/util"
)

type MeetingID string

type Meeting struct {
	ID        MeetingID `json:"id"`
	Place     string    `json:"place"`
	Creator   string    `json:"creator"`
	Buyer     string    `json:"buyer"`
	Date      time.Time `json:"date"`
	CloseDate time.Time `json:"closeDate"`
	Closed    bool      `json:"closed"`
	Orders    []Order   `json:"orders"`
	Offer     Offer     `json:"offer"`
}

type DetailedMeeting struct {
	ID         MeetingID       `json:"id"`
	Place      string          `json:"place"`
	Creator    string          `json:"creator"`
	Buyer      string          `json:"buyer"`
	Date       time.Time       `json:"date"`
	CloseDate  time.Time       `json:"closeDate"`
	Closed     bool            `json:"closed"`
	Orders     []DetailedOrder `json:"orders"`
	Offer      Offer           `json:"offer"`
	TotalPrice float32         `json:"totalPrice"`
	TotalItems []OrderItem     `json:"totalItems"`
}

type ReducedMeeting struct {
	ID        MeetingID `json:"id"`
	Place     string    `json:"place"`
	Date      time.Time `json:"date"`
	CloseDate time.Time `json:"closeDate"`
	Closed    bool      `json:"closed"`
}

func (m Meeting) Reduced() ReducedMeeting {
	return ReducedMeeting{
		ID:        m.ID,
		Place:     m.Place,
		Date:      m.Date,
		CloseDate: m.CloseDate,
		Closed:    m.Closed,
	}
}

func AllReduced(meetings []Meeting) []ReducedMeeting {
	list := make([]ReducedMeeting, 0, len(meetings))
	for _, v := range meetings {
		list = append(list, v.Reduced())
	}
	return list
}

func ToDetailedMeeting(m Meeting, paypalLink string) DetailedMeeting {
	detailedMeeting := DetailedMeeting{
		ID:        m.ID,
		Place:     m.Place,
		Creator:   m.Creator,
		Buyer:     m.Buyer,
		Date:      m.Date,
		CloseDate: m.CloseDate,
		Closed:    m.Closed,
		Orders:    ToDetailedOrders(m.Orders, m.Offer, paypalLink),
		Offer:     m.Offer,
	}

	var totalPrice float32
	var totalItems []OrderItem

	for _, order := range detailedMeeting.Orders {
		totalPrice += totalPrice + order.TotalPrice
		for _, item := range order.Items {
			index := util.IndexOf(len(totalItems), func(i int) bool {
				return totalItems[i].ItemName == item.ItemName
			})
			if index == -1 {
				totalItems = append(totalItems, OrderItem{
					ItemName: item.ItemName,
					Amount: item.Amount,
				})
			} else {
				totalItems[index].Amount = totalItems[index].Amount + item.Amount
			}
		}
	}

	detailedMeeting.TotalPrice = totalPrice
	detailedMeeting.TotalItems = totalItems

	return detailedMeeting
}

type ReadStore interface {
	FindAll() ([]Meeting, error)
	FindAllReduced() ([]ReducedMeeting, error)
	FindOne(id MeetingID) (Meeting, error)
}

type Store interface {
	ReadStore
	WriteStore
}

type WriteStore interface {
	Save(meeting Meeting) error
}

var (
	ErrMeetingNotFound = wwm.Error{Code: http.StatusNotFound, Message: "The meeting does not exist"}
	ErrNotAllowed      = wwm.Error{Code: http.StatusUnauthorized, Message: "Not allowed on this meeting"}
)

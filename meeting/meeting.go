package meeting

import (
	"net/http"
	"time"

	"github.com/WeisswurstSystems/WWM-BB/wwm"
)

// MeetingID uniquely identifies a meeting.
type MeetingID string

// Meeting describes a planned meeting for multiple persons.
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

// ReadStore can query meetings.
type ReadStore interface {
	FindAll() ([]Meeting, error)
	FindAllReduced() ([]ReducedMeeting, error)
	FindOne(id MeetingID) (Meeting, error)
}

// Store can read and write meetings.
type Store interface {
	ReadStore
	WriteStore
}

// WriteStore can save (or change) a meeting.
type WriteStore interface {
	Save(meeting Meeting) error
}

var (
	// ErrMeetingNotFound if a meeting is not found in a store.
	ErrMeetingNotFound = wwm.Error{Code: http.StatusNotFound, Message: "The meeting does not exist"}
	// ErrNotAllowed if a user is not allowed to do something on a meeting.
	ErrNotAllowed = wwm.Error{Code: http.StatusUnauthorized, Message: "Not allowed on this meeting"}
)

// FindOrderByCustomer in the order collection. If not found return a new Order for the customer.
func (m *Meeting) FindOrderByCustomer(customer CustomerMail) (index int, order Order, found bool) {
	for i, order := range m.Orders {
		if order.Customer == customer {
			return i, order, true
		}
	}
	return -1, Order{Customer: customer}, false
}

// AddOrderItemForCustomer in the order collection. If no order for the customer exists, a new one is created.
func (m *Meeting) AddOrderItemForCustomer(item OrderItem, customer CustomerMail) {
	i, order, found := m.FindOrderByCustomer(customer)
	if !found {
		order.AddItem(item)
		m.Orders = append(m.Orders, order)
		return
	}
	m.Orders[i].AddItem(item)
}

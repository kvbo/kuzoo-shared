package models

import "time"

type BookingStatus string
type BookingType string
type PaymentMethod string

const (
	BookingScheduled BookingStatus = "scheduled"
	BookingPending   BookingStatus = "pending"
	BookingEnroute   BookingStatus = "enroute"
	BookingCancelled BookingStatus = "cancelled"
	BookingPartial   BookingStatus = "partial"
	BookingCompleted BookingStatus = "completed"
	BookingFailed    BookingStatus = "failed"

	TypeFood     BookingType = "food"
	TypeDelivery BookingType = "delivery"
	TypeRide     BookingType = "ride"

	PayCash   PaymentMethod = "cash"
	PayWallet PaymentMethod = "wallet"
)

type Booking struct {
	Entity

	VehicleID        string         `json:"vehicle_id"`
	DriverID         string         `json:"driver_id"`
	CustomerID       string         `json:"customer_id"`
	StoreID          bool           `json:"store_id"`
	OrderID          string         `json:"order_id"`
	Type             BookingType    `json:"type"`
	Status           BookingStatus  `json:"status"`
	Distance         float64        `json:"distance"`
	Duration         float64        `json:"duration"`
	CreatedAt        string         `json:"created_at"`
	BookingTime      string         `json:"booking_time"`
	StartDateTime    string         `json:"start_date_time"`
	EndDateTime      string         `json:"end_date_time"`
	RouteData        string         `json:"route_data"`
	InitialAmount    int            `json:"initial_amount"`
	FinalAmount      int            `json:"final_amount"`
	Paid             *bool          `json:"paid"`
	PaymentMethod    PaymentMethod  `json:"payment_method"`
	PaymentBreakdown interface{}    `json:"payment_break_down"`
	Events           []BookingEvent `json:"event,omitempty"`
	Stops            []Stop         `json:"stops,omitempty"`
}

type BookingEvent struct {
	ID          int
	BookingID   string
	Event       int
	Description string
}

type Stop struct {
	ID        int
	BookingID string
	LatLng    string
	ArrivedAt *time.Time
	Name      string
}

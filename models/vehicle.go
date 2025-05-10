package models

type VehicleType string
type VehicleStatus string

const (
	TypeBicycle   VehicleType = "bicycle"
	TypeMotorBike VehicleType = "motor_bike"
	TypeSaloon    VehicleType = "sedan"
	TypeVan       VehicleType = "van"
	TypeTruck     VehicleType = "truck"

	StatusPending  VehicleStatus = "pending"
	StatusApproved VehicleStatus = "approved"
	StatusBlocked  VehicleStatus = "blocked"
	StatusRejected VehicleStatus = "rejected"
)

type Vehicle struct {
	Entity

	PlateNumber      string        `json:"plate_number"`
	Make             string        `json:"make"`
	Brand            string        `json:"brand"`
	FrontImage       string        `json:"front_image"`
	RearImage        string        `json:"rear_image"`
	SideImageRight   string        `json:"side_image_right"`
	SideImageLeft    string        `json:"side_image_left"`
	Type             VehicleType   `json:"type"`
	FleetID          string        `json:"fleet_id"`
	Status           VehicleStatus `json:"status"`
	DeliveryClassRef string        `json:"delivery_class_ref"`
	RideClassRef     string        `json:"ride_class_ref"`
	FoodClassRef     string        `json:"food_class_ref"`
	TaxImage         string        `json:"tax_image"`
	InsuranceImage   string        `json:"insurance_image"`
	CanDeliver       bool          `json:"can_deliver"`
	CanRide          bool          `json:"can_ride"`
	CanFood          bool          `json:"can_food"`
}

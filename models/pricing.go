package models

type ClassType string

const (
	Ride     ClassType = "ride"
	Delivery ClassType = "delivery"
	Food     ClassType = "food"
)

type Class struct {
	Entity
	Image        string    `json:"image"`
	VehicleTypes []string  `json:"vehicle_types"` // from JSON array
	For          ClassType `json:"for"`
}

type Zone struct {
	Entity
	Name            string
	Shape           string
	TimeMultipliers []TimeMultiplier
}

type TimeMultiplier struct {
	Entity
	StartTime string
	EndTime   string
	Value     float64
	RRule     string
	Operator  Operator
	Zones     []Zone
}

type Feature string

const (
	Toll Feature = "toll"
	POI  Feature = "poi"
)

type FeatureAction string

const (
	IntersectsRoute FeatureAction = "intersects_route"
	NearSource      FeatureAction = "near_source"
	NearDestination FeatureAction = "near_destination"
	NearStop        FeatureAction = "near_stop"
	NearAny         FeatureAction = "near_any"
)

type FeatureMultiplier struct {
	Entity
	Value     float64
	Operator  Operator
	Type      Feature
	ApplyWhen FeatureAction
}

type ClassZones struct {
	Entity
	ZoneID            *string
	Zone              *Zone
	Class             *Class
	ClassID           *string
	BasePerStop       *int
	RatePerKM         *int
	RatePerMinute     *int
	RatePerWaitMinute *int
	BaseDistance      *int
	MinFare           *int
}

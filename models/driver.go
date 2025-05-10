package models

type DriverStatus string

const (
	DriverApproved DriverStatus = "approved"
	DriverRejected DriverStatus = "rejected"
	DriverPending  DriverStatus = "pending"
	DriverBlocked  DriverStatus = "blocked"
)

type Driver struct {
	Entity

	FirstName        *string      `json:"first_name"`
	LastName         *string      `json:"last_name"`
	Avatar           *string      `json:"avatar"`
	NationalID       *string      `json:"national_id"`
	DOB              *string      `json:"dob"`
	UserRef          *string      `json:"user_external_ref"`
	IsOnline         *bool        `json:"is_online"`
	Status           DriverStatus `json:"status"`
	CurrentVehicleID *string      `json:"current_vehicle_ref"`
	LicenseFront     *string      `json:"license_front"`
	LicenseBack      *string      `json:"licence_back"`
	TotalRatings     *int         `json:"total_ratings"`
	TotalUsersRated  *int         `json:"total_user_count"`
	Email            *string      `json:"email"`
	Phone            *string      `json:"phone"`
	IDFront          *string      `json:"id_front"`
	IDBack           *string      `json:"id_back"`
	LicenseNumber    *string      `json:"licence_number"`
}

type DriverFleetRole string

const (
	RoleOwner  DriverFleetRole = "owner"
	RoleDriver DriverFleetRole = "driver"
)

type DriverFleet struct {
	DriverID *string         `json:"driver_id"`
	FleetID  *string         `json:"fleet_id"`
	Role     DriverFleetRole `json:"role"`
}

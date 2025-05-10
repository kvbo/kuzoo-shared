package models

type CustomerStatus string

const (
	CustomerActive  CustomerStatus = "active"
	CustomerBlocked CustomerStatus = "blocked"
)

type Customer struct {
	IsActive bool `json:"is_active"`

	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Phone     string         `json:"phone"`
	UserRef   string         `json:"user_external_ref"`
	WalletRef string         `json:"external_wallet_ref"`
	Email     string         `json:"email"`
	Status    CustomerStatus `json:"status"`
}

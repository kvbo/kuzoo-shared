package models

type Fleet struct {
	Entity

	ExternalFloatAccountRef string `json:"external_float_account_ref"`
	IsActive                bool   `json:"is_active"`
	FloatBalance            int    `json:"float_balance"`
	WalletAccountRef        int    `json:"external_wallet_account_ref"`
	WalletBalance           int    `json:"wallet_balance"`
}

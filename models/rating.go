package models

type Rating struct {
	Entity

	RaterType string          `json:"rater_type,omitempty"`
	RaterID   string          `json:"rater_id,omitempty"`
	RateeID   string          `json:"ratee_id,omitempty"`
	RateeType string          `json:"ratee_type,omitempty"`
	Details   []RatingDetails `json:"details,omitempty"`
}

type RatingDetails struct {
	Rating   *int
	Comment  *int
	RatingID string
}

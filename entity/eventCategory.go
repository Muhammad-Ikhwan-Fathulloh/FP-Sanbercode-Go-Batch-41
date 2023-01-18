package entity

import "time"

type EventCategory struct {
	EventCategoryId          int       `json:"event_category_id"`
	EventCategoryName        string    `json:"event_category_name"`
	EventCategoryDescription string    `json:"event_category_description"`
	CreatedAt                time.Time `json:"created_at,omitempty"`
	CreatedBy                int       `json:"created_by"`
	UpdatedAt                time.Time `json:"updated_at,omitempty"`
	UpdatedBy                int       `json:"updated_by"`
}

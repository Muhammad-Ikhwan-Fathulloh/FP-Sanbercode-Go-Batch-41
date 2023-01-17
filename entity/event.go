package entity

import "time"

type Event struct {
	EventId          int64     `json:"event_id"`
	CommunityId      int64     `json:"community_id"`
	EventCategoryId  int64     `json:"event_category_id"`
	EventName        string    `json:"event_name"`
	EventLogo        string    `json:"event_logo"`
	EventImage       string    `json:"event_image"`
	EventDescription string    `json:"event_description"`
	EventAddress     string    `json:"event_address"`
	EventLatitude    string    `json:"event_latitude"`
	EventLongitude   string    `json:"event_longitude"`
	EventStartDate   string    `json:"event_start_date"`
	EventFinishDate  string    `json:"event_finish_date"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	CreatedBy        int       `json:"created_by"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	UpdatedBy        int       `json:"updated_by"`
}

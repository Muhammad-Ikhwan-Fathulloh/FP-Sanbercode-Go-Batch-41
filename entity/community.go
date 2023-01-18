package entity

import "time"

type Community struct {
	CommunityId          int       `json:"community_id"`
	CommunityName        string    `json:"community_name"`
	CommunityLogo        string    `json:"community_logo"`
	CommunityImage       string    `json:"community_image"`
	CommunityDescription string    `json:"community_description"`
	CommunityAddress     string    `json:"community_address"`
	CommunityLatitude    string    `json:"community_latitude"`
	CommunityLongitude   string    `json:"community_longitude"`
	CommunityInfo        string    `json:"community_info"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
	CreatedBy            int       `json:"created_by"`
	UpdatedAt            time.Time `json:"updated_at,omitempty"`
	UpdatedBy            int       `json:"updated_by"`
}

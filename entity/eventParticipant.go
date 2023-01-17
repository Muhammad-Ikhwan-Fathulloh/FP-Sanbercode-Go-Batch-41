package entity

import "time"

type EventParticipant struct {
	EventParticipantId        int64     `json:"event_participant_id"`
	EventId                   int64     `json:"event_id"`
	EventParticipantName      string    `json:"event_participant_name"`
	EventParticipantEmail     string    `json:"event_participant_email"`
	EventParticipantNoHp      string    `json:"event_participant_no_hp"`
	EventParticipantCommunity string    `json:"event_participant_community"`
	EventParticipantStatus    string    `json:"event_participant_status"`
	CreatedAt                 time.Time `json:"created_at,omitempty"`
	UpdatedAt                 time.Time `json:"updated_at,omitempty"`
}

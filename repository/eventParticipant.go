package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	_ "fmt"
	_ "log"
	"time"
)

func GetAllEventParticipant(db *sql.DB) []entity.EventParticipant {
	var results = []entity.EventParticipant{}
	sql := `SELECT * FROM event_participants`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var eventParticipant = entity.EventParticipant{}

		err = rows.Scan(eventParticipant.EventParticipantId, eventParticipant.EventId, eventParticipant.EventParticipantName, eventParticipant.EventParticipantEmail, eventParticipant.EventParticipantNoHp, eventParticipant.EventParticipantCommunity, eventParticipant.EventParticipantStatus)
		if err != nil {
			panic(err)
		}

		results = append(results, eventParticipant)
	}

	return results
}

func InsertEventParticipant(db *sql.DB, eventParticipant entity.EventParticipant) (err error) {
	sql := `INSERT INTO event_participants (event_participant_id, event_id, event_participant_name, event_participant_email, event_participant_no_hp, event_participant_community, event_participant_status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	errs := db.QueryRow(sql, eventParticipant.EventParticipantId, eventParticipant.EventId, eventParticipant.EventParticipantName, eventParticipant.EventParticipantEmail, eventParticipant.EventParticipantNoHp, eventParticipant.EventParticipantCommunity, eventParticipant.EventParticipantStatus, time.Now(), time.Now())

	return errs.Err()
}

func UpdateEventParticipant(db *sql.DB, eventParticipant entity.EventParticipant) (err error) {
	sql := `UPDATE event_participants SET event_id = $1, event_participant_name = $2, event_participant_email = $3, event_participant_no_hp = $4, event_participant_community = $5, event_participant_status = $6, updated_at = $7 WHERE event_participant_id = $8`

	errs := db.QueryRow(sql, eventParticipant.EventId, eventParticipant.EventParticipantName, eventParticipant.EventParticipantEmail, eventParticipant.EventParticipantNoHp, eventParticipant.EventParticipantCommunity, eventParticipant.EventParticipantStatus, time.Now(), eventParticipant.EventParticipantId)

	return errs.Err()
}

func DeleteEventParticipant(db *sql.DB, eventParticipant entity.EventParticipant) (err error) {
	sql := `DELETE FROM event_participants WHERE event_participant_id = $1`

	errs := db.QueryRow(sql, eventParticipant.EventParticipantId)

	return errs.Err()
}

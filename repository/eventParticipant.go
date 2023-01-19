package repository

import (
	"FP-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	"errors"
	_ "fmt"
	_ "log"
	"time"
)

func GeteventParticipantById(db *sql.DB, eventParticipantId int) (eventParticipantResponse entity.EventParticipant, err error) {
	sqlStatement := "SELECT * FROM event_participants WHERE event_participant_id = $1"

	err = db.QueryRow(sqlStatement, eventParticipantId).Scan(&eventParticipantResponse.EventParticipantId, &eventParticipantResponse.EventId, &eventParticipantResponse.EventParticipantName, &eventParticipantResponse.EventParticipantEmail, &eventParticipantResponse.EventParticipantNoHp, &eventParticipantResponse.EventParticipantCommunity, &eventParticipantResponse.EventParticipantStatus, &eventParticipantResponse.CreatedAt, &eventParticipantResponse.UpdatedAt)

	return
}

func GetAlleventParticipantByEvent(db *sql.DB, eventId int) (eventParticipantResponse entity.EventParticipant, err error) {
	sqlStatement := "SELECT * FROM event_participants WHERE event_id = $1"

	err = db.QueryRow(sqlStatement, eventId).Scan(&eventParticipantResponse.EventParticipantId, &eventParticipantResponse.EventId, &eventParticipantResponse.EventParticipantName, &eventParticipantResponse.EventParticipantEmail, &eventParticipantResponse.EventParticipantNoHp, &eventParticipantResponse.EventParticipantCommunity, &eventParticipantResponse.EventParticipantStatus, &eventParticipantResponse.CreatedAt, &eventParticipantResponse.UpdatedAt)

	return
}

func GetAllEventParticipant(db *sql.DB) (results []entity.EventParticipant, err error) {
	sqlStatement := `SELECT * FROM event_participants`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var eventParticipant = entity.EventParticipant{}

		err = rows.Scan(&eventParticipant.EventParticipantId, &eventParticipant.EventId, &eventParticipant.EventParticipantName, &eventParticipant.EventParticipantEmail, &eventParticipant.EventParticipantNoHp, &eventParticipant.EventParticipantCommunity, &eventParticipant.EventParticipantStatus, &eventParticipant.CreatedAt, &eventParticipant.UpdatedAt)

		results = append(results, eventParticipant)
	}

	return
}

func InsertEventParticipant(db *sql.DB, eventParticipant entity.EventParticipant) (err error) {
	sqlStatement := `INSERT INTO event_participants (event_id, event_participant_name, event_participant_email, event_participant_no_hp, event_participant_community, event_participant_status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	errs := db.QueryRow(sqlStatement, &eventParticipant.EventId, &eventParticipant.EventParticipantName, &eventParticipant.EventParticipantEmail, &eventParticipant.EventParticipantNoHp, &eventParticipant.EventParticipantCommunity, &eventParticipant.EventParticipantStatus, time.Now(), time.Now())

	return errs.Err()
}

func UpdateEventParticipant(db *sql.DB, eventParticipant entity.EventParticipant) (err error) {
	sqlStatement := `UPDATE event_participants SET event_id = $2, event_participant_name = $3, event_participant_email = $4, event_participant_no_hp = $5, event_participant_community = $6, event_participant_status = $7, updated_at = $8 WHERE event_participant_id = $1`

	res, err := db.Exec(sqlStatement, &eventParticipant.EventParticipantId, &eventParticipant.EventId, &eventParticipant.EventParticipantName, &eventParticipant.EventParticipantEmail, &eventParticipant.EventParticipantNoHp, &eventParticipant.EventParticipantCommunity, &eventParticipant.EventParticipantStatus, time.Now())

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

func DeleteEventParticipant(db *sql.DB, eventParticipant entity.EventParticipant) (err error) {
	sqlStatement := `DELETE FROM event_participants WHERE event_participant_id = $1`

	res, err := db.Exec(sqlStatement, &eventParticipant.EventParticipantId)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

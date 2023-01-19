package repository

import (
	"FP-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	"errors"
	_ "fmt"
	_ "log"
	"time"
)

func GetAllEvent(db *sql.DB) (results []entity.Event, err error) {
	sqlStatement := `SELECT * FROM events`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var event = entity.Event{}

		err = rows.Scan(event.EventId, &event.EventName, &event.EventLogo, &event.EventImage, &event.EventDescription, &event.EventAddress, &event.EventLatitude, &event.EventLongitude, &event.EventStartDate, &event.EventFinishDate, &event.CreatedAt, &event.CreatedBy, &event.UpdatedAt, &event.UpdatedBy)

		results = append(results, event)
	}

	return
}

func GetEventById(db *sql.DB, eventId int) (eventResponse entity.Event, err error) {
	sqlStatement := `SELECT * FROM events WHERE event_id={$1}`

	rows, err := db.Query(sqlStatement, eventId)
	if err != nil {
		panic(err)
	}

	err = rows.Scan(&eventResponse.EventId, &eventResponse.EventName, &eventResponse.EventLogo, &eventResponse.EventImage, &eventResponse.EventDescription, &eventResponse.EventAddress, &eventResponse.EventLatitude, &eventResponse.EventLongitude, &eventResponse.EventStartDate, &eventResponse.EventFinishDate, &eventResponse.CreatedAt, &eventResponse.CreatedBy, &eventResponse.UpdatedAt, &eventResponse.UpdatedBy)
	if err != nil {
		panic(err)
	}

	return
}

func GetAllEventByCommunity(db *sql.DB, CommunityId int) (eventResponse entity.Event, err error) {
	sqlStatement := `SELECT * FROM events WHERE community_id={$1}`

	rows, err := db.Query(sqlStatement, CommunityId)
	if err != nil {
		panic(err)
	}

	err = rows.Scan(&eventResponse.EventId, &eventResponse.EventName, &eventResponse.EventLogo, &eventResponse.EventImage, &eventResponse.EventDescription, &eventResponse.EventAddress, &eventResponse.EventLatitude, &eventResponse.EventLongitude, &eventResponse.EventStartDate, &eventResponse.EventFinishDate, &eventResponse.CreatedAt, &eventResponse.CreatedBy, &eventResponse.UpdatedAt, &eventResponse.UpdatedBy)
	if err != nil {
		panic(err)
	}

	return
}

func GetAllEventByCategory(db *sql.DB, eventCategoryId int) (eventResponse entity.Event, err error) {
	sqlStatement := `SELECT * FROM events WHERE event_category_id={$1}`

	rows, err := db.Query(sqlStatement, eventCategoryId)
	if err != nil {
		panic(err)
	}

	err = rows.Scan(&eventResponse.EventId, &eventResponse.EventName, &eventResponse.EventLogo, &eventResponse.EventImage, &eventResponse.EventDescription, &eventResponse.EventAddress, &eventResponse.EventLatitude, &eventResponse.EventLongitude, &eventResponse.EventStartDate, &eventResponse.EventFinishDate, &eventResponse.CreatedAt, &eventResponse.CreatedBy, &eventResponse.UpdatedAt, &eventResponse.UpdatedBy)
	if err != nil {
		panic(err)
	}

	return
}

func InsertEvent(db *sql.DB, event entity.Event) (err error) {
	sqlStatement := `INSERT INTO events (community_id, event_category_id, event_name, event_logo, event_image, event_description, event_address, event_latitude, event_longitude, event_start_date, event_finish_date, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	errs := db.QueryRow(sqlStatement, &event.CommunityId, &event.EventCategoryId, &event.EventName, &event.EventLogo, &event.EventImage, &event.EventDescription, &event.EventAddress, &event.EventLatitude, &event.EventLongitude, &event.EventStartDate, &event.EventFinishDate, time.Now(), &event.CreatedBy, time.Now(), &event.UpdatedBy)

	return errs.Err()
}

func UpdateEvent(db *sql.DB, event entity.Event) (err error) {
	sqlStatement := `UPDATE events SET community_id = $2, event_category_id = $3, event_name = $4, event_logo = $5, event_image = $6, event_description = $7, event_address = $8, event_latitude = $9, event_longitude = $10, event_start_date = $11, event_finish_date = $12, , updated_at = $13, updated_by = $14 WHERE event_id = $1`

	res, err := db.Exec(sqlStatement, &event.EventId, &event.CommunityId, &event.EventCategoryId, &event.EventName, &event.EventLogo, &event.EventImage, &event.EventDescription, &event.EventAddress, &event.EventLatitude, &event.EventLongitude, &event.EventStartDate, &event.EventFinishDate, time.Now(), &event.UpdatedBy)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

func DeleteEvent(db *sql.DB, event entity.Event) (err error) {
	sqlStatement := `DELETE FROM events WHERE event_id = $1`

	res, err := db.Exec(sqlStatement, &event.EventId)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

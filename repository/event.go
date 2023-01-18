package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	_ "fmt"
	_ "log"
	"time"
)

func GetAllEvent(db *sql.DB) []entity.Event {
	var results = []entity.Event{}
	sql := `SELECT * FROM events`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var event = entity.Event{}

		err = rows.Scan(event.EventId, event.EventName, event.EventLogo, event.EventImage, event.EventDescription, event.EventAddress, event.EventLatitude, event.EventLongitude, event.EventStartDate, event.EventFinishDate)
		if err != nil {
			panic(err)
		}

		results = append(results, event)
	}

	return results
}

func GetAllEventByCategory(id int, db *sql.DB, results []entity.Event) (err error) {
	sql := `SELECT * FROM events WHERE event_category_id={$1}`

	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	var event = entity.Event{}
	err = rows.Scan(&event.EventId, &event.EventName, &event.EventLogo, &event.EventImage, &event.EventDescription, &event.EventAddress, &event.EventLatitude, &event.EventLongitude, &event.EventStartDate, &event.EventFinishDate, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		panic(err)
	}

	return
}

func InsertEvent(db *sql.DB, event entity.Event) (err error) {
	sql := `INSERT INTO events (event_id, community_id, event_category_id, event_name, event_logo, event_image, event_description, event_address, event_latitude, event_longitude, event_start_date, event_finish_date, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`

	errs := db.QueryRow(sql, event.EventId, event.CommunityId, event.EventCategoryId, event.EventName, event.EventLogo, event.EventImage, event.EventDescription, event.EventAddress, event.EventLatitude, event.EventLongitude, event.EventStartDate, event.EventFinishDate, time.Now(), event.CreatedBy, time.Now(), event.UpdatedBy)

	return errs.Err()
}

func UpdateEvent(db *sql.DB, event entity.Event) (err error) {
	sql := `UPDATE events SET community_id = $1, event_category_id = $2, event_name = $3, event_logo = $4, event_image = $5, event_description = $6, event_address = $7, event_latitude = $8, event_longitude = $9, event_start_date = $10, event_finish_date = $11, , updated_at = $12, updated_by = $13 WHERE user_id = $14`

	errs := db.QueryRow(sql, event.CommunityId, event.EventCategoryId, event.EventName, event.EventLogo, event.EventImage, event.EventDescription, event.EventAddress, event.EventLatitude, event.EventLongitude, event.EventStartDate, event.EventFinishDate, time.Now(), event.UpdatedBy, event.EventId)

	return errs.Err()
}

func DeleteEvent(db *sql.DB, event entity.Event) (err error) {
	sql := `DELETE FROM events WHERE event_id = $1`

	errs := db.QueryRow(sql, event.EventId)

	return errs.Err()
}

package repository

import (
	"FP-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	"errors"
	_ "fmt"
	_ "log"
	"time"
)

func GeteventCategoryById(db *sql.DB, eventCategoryById int) (eventCategoryResponse entity.EventCategory, err error) {
	sqlStatement := "SELECT * FROM event_categories WHERE event_category_id = $1"

	err = db.QueryRow(sqlStatement, eventCategoryById).Scan(&eventCategoryResponse.EventCategoryId, &eventCategoryResponse.EventCategoryName, &eventCategoryResponse.EventCategoryDescription, &eventCategoryResponse.CreatedAt, &eventCategoryResponse.CreatedBy, &eventCategoryResponse.UpdatedAt, &eventCategoryResponse.UpdatedBy)

	return
}

func GetAllEventCategory(db *sql.DB) (results []entity.EventCategory, err error) {
	sqlStatement := `SELECT * FROM event_categories`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var eventCategory = entity.EventCategory{}

		err = rows.Scan(&eventCategory.EventCategoryId, &eventCategory.EventCategoryName, &eventCategory.EventCategoryDescription, &eventCategory.CreatedAt, &eventCategory.CreatedBy, &eventCategory.UpdatedAt, &eventCategory.UpdatedBy)

		results = append(results, eventCategory)
	}

	return
}

func InsertEventCategory(db *sql.DB, eventCategory entity.EventCategory) (err error) {
	sqlStatement := `INSERT INTO event_categories (event_category_name, event_category_description, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6)`

	errs := db.QueryRow(sqlStatement, &eventCategory.EventCategoryName, &eventCategory.EventCategoryDescription, time.Now(), &eventCategory.CreatedBy, time.Now(), &eventCategory.UpdatedBy)

	return errs.Err()
}

func UpdateEventCategory(db *sql.DB, eventCategory entity.EventCategory) (err error) {
	sqlStatement := `UPDATE event_categories SET event_category_name = $2, event_category_description = $3, updated_at = $4, updated_by = $5 WHERE event_category_id = $1`

	res, err := db.Exec(sqlStatement, &eventCategory.EventCategoryId, &eventCategory.EventCategoryName, &eventCategory.EventCategoryDescription, time.Now(), &eventCategory.UpdatedBy)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

func DeleteEventCategory(db *sql.DB, eventCategory entity.EventCategory) (err error) {
	sqlStatement := `DELETE FROM event_categories WHERE event_category_id = $1`

	res, err := db.Exec(sqlStatement, &eventCategory.EventCategoryId)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

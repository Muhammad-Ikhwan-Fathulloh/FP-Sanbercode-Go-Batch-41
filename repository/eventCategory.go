package repository

import (
	"FP-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	_ "fmt"
	_ "log"
	"time"
)

func GetAllEventCategory(db *sql.DB) []entity.EventCategory {
	var results = []entity.EventCategory{}
	sql := `SELECT * FROM event_categories`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var eventCategory = entity.EventCategory{}

		err = rows.Scan(&eventCategory.EventCategoryId, &eventCategory.EventCategoryName, &eventCategory.EventCategoryDescription, &eventCategory.CreatedAt, &eventCategory.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, eventCategory)
	}

	return results
}

func InsertEventCategory(db *sql.DB, eventCategory entity.EventCategory) (err error) {
	sql := `INSERT INTO event_categories (event_category_id, event_category_name, event_category_description, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	errs := db.QueryRow(sql, eventCategory.EventCategoryId, eventCategory.EventCategoryName, eventCategory.EventCategoryDescription, time.Now(), eventCategory.CreatedBy, time.Now(), eventCategory.UpdatedBy)

	return errs.Err()
}

func UpdateEventCategory(db *sql.DB, eventCategory entity.EventCategory) (err error) {
	sql := `UPDATE event_categories SET event_category_name = $1, event_category_description = $2, updated_at = $3, updated_by = $4 WHERE event_category_id = $5`

	errs := db.QueryRow(sql, eventCategory.EventCategoryName, eventCategory.EventCategoryDescription, time.Now(), eventCategory.UpdatedBy, eventCategory.EventCategoryId)

	return errs.Err()
}

func DeleteEventCategory(db *sql.DB, eventCategory entity.EventCategory) (err error) {
	sql := `DELETE FROM event_categories WHERE event_category_id = $1`

	errs := db.QueryRow(sql, eventCategory.EventCategoryId)

	return errs.Err()
}

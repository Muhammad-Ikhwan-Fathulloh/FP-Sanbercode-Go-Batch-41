package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	_ "fmt"
	_ "log"
	"time"
)

func GetAllCommunity(db *sql.DB) []entity.Community {
	var results = []entity.Community{}
	sql := `SELECT * FROM communities`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var community = entity.Community{}

		err = rows.Scan(community.CommunityId, community.CommunityName, community.CommunityLogo, community.CommunityImage, community.CommunityDescription, community.CommunityAddress, community.CommunityLatitude, community.CommunityLongitude, community.CommunityInfo)
		if err != nil {
			panic(err)
		}

		results = append(results, community)
	}

	return results
}

func InsertCommunity(db *sql.DB, community entity.Community) (err error) {
	sql := `INSERT INTO communities (community_id, community_name, community_logo, community_image, community_description, community_address, community_latitude, community_longitude, community_info, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	errs := db.QueryRow(sql, community.CommunityId, community.CommunityName, community.CommunityLogo, community.CommunityImage, community.CommunityDescription, community.CommunityAddress, community.CommunityLatitude, community.CommunityLongitude, community.CommunityInfo, time.Now(), community.CreatedBy, time.Now(), community.UpdatedBy)

	return errs.Err()
}

func UpdateCommunity(db *sql.DB, community entity.Community) (err error) {
	sql := `UPDATE communities SET community_name = $1, community_logo = $2, community_image = $3, community_description = $4, community_description = $5, community_address = $6, community_latitude = $7, community_longitude = $8, community_info = $9, updated_at = $10, updated_by = $11 WHERE user_id = $12`

	errs := db.QueryRow(sql, community.CommunityName, community.CommunityLogo, community.CommunityImage, community.CommunityDescription, community.CommunityAddress, community.CommunityLatitude, community.CommunityLongitude, community.CommunityInfo, time.Now(), community.UpdatedBy, community.CommunityId)

	return errs.Err()
}

func DeleteCommunity(db *sql.DB, community entity.Community) (err error) {
	sql := `DELETE FROM communities WHERE community_id = $1`

	errs := db.QueryRow(sql, community.CommunityId)

	return errs.Err()
}

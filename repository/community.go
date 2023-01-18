package repository

import (
	"FP-Sanbercode-Go-Batch-41/entity"
	"database/sql"
	"errors"
	_ "fmt"
	_ "log"
	"time"
)

func GetCommunityById(db *sql.DB, communityId int) (communityResponse entity.Community, err error) {
	sqlStatement := `SELECT * FROM communities WHERE community_id={$1}`

	rows, err := db.Query(sqlStatement, communityId)
	if err != nil {
		panic(err)
	}

	err = rows.Scan(&communityResponse.CommunityId, &communityResponse.CommunityName, &communityResponse.CommunityLogo, &communityResponse.CommunityImage, &communityResponse.CommunityDescription, &communityResponse.CommunityAddress, &communityResponse.CommunityLatitude, &communityResponse.CommunityLongitude, &communityResponse.CommunityInfo, &communityResponse.CreatedAt, &communityResponse.CreatedBy, &communityResponse.UpdatedAt, &communityResponse.UpdatedBy)
	if err != nil {
		panic(err)
	}

	return
}

func GetAllCommunity(db *sql.DB) (results []entity.Community, err error) {
	sqlStatement := `SELECT * FROM communities`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var community = entity.Community{}

		err = rows.Scan(&community.CommunityId, &community.CommunityName, &community.CommunityLogo, &community.CommunityImage, &community.CommunityDescription, &community.CommunityAddress, &community.CommunityLatitude, &community.CommunityLongitude, &community.CommunityInfo, &community.CreatedAt, &community.CreatedBy, &community.UpdatedAt, &community.UpdatedBy)

		results = append(results, community)
	}

	return
}

func InsertCommunity(db *sql.DB, community entity.Community) (err error) {
	sqlStatement := `INSERT INTO communities (community_name, community_logo, community_image, community_description, community_address, community_latitude, community_longitude, community_info, created_at, created_by, updated_at, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	errs := db.QueryRow(sqlStatement, &community.CommunityName, &community.CommunityLogo, &community.CommunityImage, &community.CommunityDescription, &community.CommunityAddress, &community.CommunityLatitude, &community.CommunityLongitude, &community.CommunityInfo, time.Now(), &community.CreatedBy, time.Now(), &community.UpdatedBy)

	return errs.Err()
}

func UpdateCommunity(db *sql.DB, community entity.Community) (err error) {
	sqlStatement := `UPDATE communities SET community_name = $2, community_logo = $3, community_image = $4, community_description = $5, community_address = $6, community_latitude = $7, community_longitude = $8, community_info = $9, updated_at = $10, updated_by = $11 WHERE community_id = $1`

	res, err := db.Exec(sqlStatement, &community.CommunityId, &community.CommunityName, &community.CommunityLogo, &community.CommunityImage, &community.CommunityDescription, &community.CommunityAddress, &community.CommunityLatitude, &community.CommunityLongitude, &community.CommunityInfo, time.Now(), &community.UpdatedBy)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

func DeleteCommunity(db *sql.DB, community entity.Community) (err error) {
	sqlStatement := `DELETE FROM communities WHERE community_id = $1`

	res, err := db.Exec(sqlStatement, &community.CommunityId)

	if err != nil {
		return
	}

	count, err := res.RowsAffected()
	if err == nil && count == 0 {
		err = errors.New("rows is empty")
	}

	return
}

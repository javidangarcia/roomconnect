package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SearchUsers(queryFields [][2]string) []*UserProfile {
	var queryValues []any

	sqlQuery := `
		SELECT 	display_name,
				gender,
				profile_pic,
				clean_tier,
				budget_tier,
				loud_tier,
				coed
		FROM users
		`

	// Add AND WHERE clauses in SQL string
	for idx, slice := range queryFields {
		field := slice[0]
		value := slice[1]

		// build clauses dynamically
		clause := "WHERE"
		if idx > 0 {
			clause = "AND"
		}

		sqlQuery += fmt.Sprintf("%s %s=$%d ", clause, field, idx+1)

		// store values in order
		queryValues = append(queryValues, value)
	}

	// spread values in order into query function
	rows, err := db.Query(context.Background(), sqlQuery, queryValues...)

	if err != nil {
		fmt.Println("Error searching for users in field. Error:", err.Error())
	}

	usersFound, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[UserProfile])

	if err != nil {
		fmt.Println("Error collecting rows in SearchUsers. Error:", err.Error())
		return []*UserProfile{}
	}

	return usersFound
}
package query

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func QueryCount(db *sqlx.DB, query string) (string, error) {
	var queryResult string
	err := db.QueryRow(query).Scan(&queryResult)
	if err != nil {
		return "", err
	}
	return queryResult, nil
}

func QueryCountInt(db *sqlx.DB, query string) (int, error) {
	var queryResult string
	err := db.QueryRow(query).Scan(&queryResult)
	if err != nil {
		return 0, err
	}
	if strings.Contains(queryResult, ".") {
		queryResult = strings.Split(queryResult, ".")[0]
	}

	queryResultNumber, err := strconv.Atoi(queryResult)
	if err != nil {
		return 0, err
	}

	return queryResultNumber, nil
}

func QueryJson(db *sqlx.DB, resultSlice interface{}, query string) ([]byte, error) {
	err := db.Select(resultSlice, query)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(resultSlice)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

func QueryStruct(db *sqlx.DB, resultSlice interface{}, query string) (interface{}, error) {
	err := db.Select(resultSlice, query)
	if err != nil {
		return nil, err
	}

	return resultSlice, nil
}

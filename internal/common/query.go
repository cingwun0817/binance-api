package common

import (
	"context"
	"database/sql"
	"log"
)

func GetSymbols(ctx context.Context, db *sql.DB) []string {
	var result []string

	rows, err := db.QueryContext(ctx, "SELECT `name` FROM `binance`.`symbol` WHERE `status` = 'TRADING';")
	// rows, err := db.QueryContext(ctx, "SELECT `name` FROM `binance`.`symbol` WHERE `status` = 'TRADING' AND `name` = 'BTCUSDT';")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}

		if name[len(name)-4:len(name)] == "USDT" {
			result = append(result, name)
		}
	}

	return result
}

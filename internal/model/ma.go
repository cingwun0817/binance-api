package model

import (
	"context"
	"database/sql"
)

type CoinMaFound struct {
	Symbol  string
	Found1h bool
	Found2h bool
	Found4h bool
	Found8h bool
	Found1d bool
}

func GetFullMaFound(ctx context.Context, db *sql.DB, symbol string) CoinMaFound {
	var windowSizes = []string{"1h", "2h", "4h", "8h", "1d"}

	var result CoinMaFound
	result.Symbol = symbol

	for _, windowSize := range windowSizes {
		var date string
		var hour int
		var ma30, ma45, ma60 float64
		db.QueryRowContext(ctx, "SELECT date, hour, ma30, ma45, ma60 FROM `binance`.`ticker_full` WHERE `symbol` = ? AND `windowSize` = ? ORDER BY date DESC LIMIT 1;", symbol, windowSize).Scan(&date, &hour, &ma30, &ma45, &ma60)

		if ma30 != 0 && ma45 != 0 && ma60 != 0 {
			switch windowSize {
			case "1h":
				if ma30 > ma60 || ma45 > ma60 {
					result.Found1h = true
				}
			case "2h":
				if ma30 > ma60 || ma45 > ma60 {
					result.Found2h = true
				}
			case "4h":
				if ma30 > ma60 || ma45 > ma60 {
					result.Found4h = true
				}
			case "8h":
				if ma30 > ma60 || ma45 > ma60 {
					result.Found8h = true
				}
			case "1d":
				if ma30 > ma60 || ma45 > ma60 {
					result.Found1d = true
				}
			}
		}
	}

	return result
}

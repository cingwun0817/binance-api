package api

import (
	"binance-api/internal/common"
	"context"
	"database/sql"
	"fmt"
	"sort"

	"github.com/gofiber/fiber/v2"
)

type TickerFull struct {
	// Date               string  `json:"date"`
	// Hour               int     `json:"hour"`
	Symbol string `json:"symbol"`
	// PriceChangePercent float64 `json:"priceChangePercent"`
	Items                  []TickerFullItem `json:"items"`
	LastPriceChangePercent float64          `json:"lastPriceChangePercent"`
}

type TickerFullItem struct {
	Name               string  `json:"name"`
	PriceChangePercent float64 `json:"priceChangePercent"`
}

func GetTickerFull(c *fiber.Ctx) error {
	ctx := context.Background()
	var data []TickerFull
	mapData := make(map[string]*TickerFull)

	// query
	query := c.Queries()
	windowSize := query["windowSize"]
	pDate := query["date"]
	pHour := query["hour"]

	var rows *sql.Rows
	var err error
	if pHour == "" {
		rows, err = common.Db.QueryContext(ctx, "SELECT `date`, `hour`, `symbol`, `priceChangePercent` FROM `binance`.`ticker_full` WHERE `windowSize` = ? AND `date` >= ?;", windowSize, pDate)
	} else {
		rows, err = common.Db.QueryContext(ctx, "SELECT `date`, `hour`, `symbol`, `priceChangePercent` FROM `binance`.`ticker_full` WHERE `windowSize` = ? AND `date` >= ? AND `hour` = ?;", windowSize, pDate, pHour)
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": err.Error(),
		})
	}

	var symbols []string
	for rows.Next() {
		var date, symbol string
		var hour int
		var priceChangePercent float64
		if err := rows.Scan(&date, &hour, &symbol, &priceChangePercent); err != nil {
			return err
		}

		if _, ok := mapData[symbol]; !ok {
			mapData[symbol] = &TickerFull{
				Symbol: symbol,
				Items:  []TickerFullItem{},
			}

			symbols = append(symbols, symbol)
		}

		var item TickerFullItem

		if windowSize == "1d" {
			item.Name = date
		} else {
			item.Name = fmt.Sprintf("%s.%d", date, hour)
		}
		item.PriceChangePercent = priceChangePercent

		mapData[symbol].Items = append(mapData[symbol].Items, item)
		mapData[symbol].LastPriceChangePercent = priceChangePercent
	}

	for _, symbol := range symbols {
		data = append(data, *mapData[symbol])
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].LastPriceChangePercent > data[j].LastPriceChangePercent
	})

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

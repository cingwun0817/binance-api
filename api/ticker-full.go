package api

import (
	"binance-api/internal/common"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type TickerFull struct {
	// Date               string  `json:"date"`
	// Hour               int     `json:"hour"`
	Symbol string `json:"symbol"`
	// PriceChangePercent float64 `json:"priceChangePercent"`
	Items []TickerFullItem
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

	rows, err := common.Db.QueryContext(ctx, "SELECT `date`, `hour`, `symbol`, `priceChangePercent` FROM `binance`.`ticker_full` WHERE `windowSize` = ? AND `date` >= ?", windowSize, pDate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": err.Error(),
		})
	}

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
		}

		var item TickerFullItem

		if windowSize == "1d" {
			item.Name = date
		} else {
			item.Name = fmt.Sprintf("%d", hour)
		}
		item.PriceChangePercent = priceChangePercent

		mapData[symbol].Items = append(mapData[symbol].Items, item)
	}

	for _, rowdata := range mapData {
		data = append(data, *rowdata)
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

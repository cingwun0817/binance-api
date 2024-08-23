package api

import (
	"binance-api/internal/common"
	"context"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Id   int
	Name string
}

func GetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	var data []User

	rows, err := common.Db.QueryContext(ctx, "SELECT `id`, `name` FROM `binance`.`user`")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": err.Error(),
		})
	}

	for rows.Next() {
		var userId int
		var name string
		if err := rows.Scan(&userId, &name); err != nil {
			return err
		}

		data = append(data, User{userId, name})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

type UserPosition struct {
	Symbol                string  `json:"symbol"`
	OpenPrice             float64 `json:"openPrice"`
	OpenQuantity          float64 `json:"openQuantity"`
	Leverage              float64 `json:"leverage"`
	UnRealizedProfit      float64 `json:"unRealizedProfit"`
	UserId                int     `json:"userId"`
	StopOrderStatus       int     `json:"stopOrderStatus"`
	TakeProfitOrderStatus int     `json:"takeProfitOrderStatus"`
}

func GetUserPosition(c *fiber.Ctx) error {
	ctx := context.Background()
	var data []UserPosition

	// query
	query := c.Queries()
	userId := query["userId"]

	rows, err := common.Db.QueryContext(ctx, "SELECT `symbol`, `openPrice`, `openQuantity`, `leverage`, `unRealizedProfit`, `userId`, `stopOrderStatus`, `takeProfitOrderStatus` FROM `binance`.`user_position` WHERE `userId` = ?", userId)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"code":    500,
			"message": err.Error(),
		})
	}

	for rows.Next() {
		var symbol string
		var userId, stopOrderStatus, takeProfitOrderStatus int
		var openPrice, openQuantity, leverage, unRealizedProfit float64
		if err := rows.Scan(&symbol, &openPrice, &openQuantity, &leverage, &unRealizedProfit, &userId, &stopOrderStatus, &takeProfitOrderStatus); err != nil {
			return err
		}

		data = append(data, UserPosition{symbol, openPrice, openQuantity, leverage, unRealizedProfit, userId, stopOrderStatus, takeProfitOrderStatus})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

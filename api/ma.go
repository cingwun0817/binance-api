package api

import (
	"binance-api/internal/common"
	"binance-api/internal/model"
	"context"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CoinMa struct {
	Symbol   string
	MaResult []string
}

func GetMa(c *fiber.Ctx) error {
	ctx := context.Background()
	var data []CoinMa

	symbols := common.GetSymbols(ctx, common.Db)

	for _, symbol := range symbols {
		maFound := model.GetFullMaFound(ctx, common.Db, symbol)

		if maFound.Found1h || maFound.Found2h || maFound.Found4h || maFound.Found8h || maFound.Found1d {
			var coin CoinMa

			coin.Symbol = symbol[0:strings.Index(symbol, "USDT")]

			if maFound.Found1d {
				coin.MaResult = append(coin.MaResult, "1d")
			}

			if maFound.Found8h {
				coin.MaResult = append(coin.MaResult, "8h")
			}

			if maFound.Found4h {
				coin.MaResult = append(coin.MaResult, "4h")
			}

			if maFound.Found2h {
				coin.MaResult = append(coin.MaResult, "2h")
			}

			if maFound.Found1h {
				coin.MaResult = append(coin.MaResult, "1h")
			}

			data = append(data, coin)
		}
	}

	sort.Slice(data, func(i, j int) bool {
		return len(data[i].MaResult) > len(data[j].MaResult)
	})

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "success",
		"data":    data,
	})
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type Profit struct {
	MaximumProfit int
	MinimumPrice  int
	MaximumPrice  int
	TimeBuy       int
	TimeSell      int
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println(`Reading File ....`)
		fileBytes, err := ioutil.ReadFile("adwadw.txt")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		sliceData := strings.Split(string(fileBytes), " ")
		profit := findProfit(sliceData)
		fmt.Println(`Done ....`)
		return c.JSON(http.StatusOK, profit)
	})

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}

func findProfit(data []string) Profit {
	var profit Profit
	si := make([]int, 0, len(data))
	min, err := strconv.Atoi(data[0])
	max, err := strconv.Atoi(data[0])
	var idxMax [1]int
	var idxMin [1]int
	if err != nil {
		// Add code here to handle the error!
	}
	for idx, a := range data {
		i, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println(err)
		}
		if max < i {
			idxMax[0] = idx + 1
			max = i
		}
		if min > i {
			idxMin[0] = idx + 1
			min = i
			max = i
		}
		si = append(si, i)
	}
	profit.MaximumProfit = max - min
	profit.MinimumPrice = min
	profit.MaximumPrice = max
	profit.TimeBuy = idxMin[0]
	profit.TimeSell = idxMax[0]
	return profit
}

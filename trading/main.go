package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

// - preprocessing input

func convertStringToInt(val string) (int64, error) {
	number, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return int64(number), nil
}

func linesFromReader(r io.Reader) ([]int64, error) {
	var input []int64
	reader := bufio.NewReader(r)

	for {
		line, err := reader.ReadString(' ')
		if err == io.EOF {
			break
		}

		// - convert from string into integer
		value, err := convertStringToInt(strings.Replace(line, " ", "", -1))
		if err != nil {
			return input, err
		}

		input = append(input, value)
	}

	return input, nil
}

func urlToLines(url string) ([]int64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return linesFromReader(resp.Body)
}

func maxProfit(prices []int64) (int64, int64) {
	var maxProfit int64 = 0

	const MaxUint = ^uint(0)
	minPrice := int64(MaxUint >> 1)

	if len(prices) < 2 {
		return minPrice, maxProfit
	}

	for x := 0; x < len(prices); x++ {
		if prices[x] < minPrice {
			minPrice = prices[x]
		} else if (prices[x] - minPrice) > maxProfit {
			maxProfit = prices[x] - minPrice
		}
	}

	return minPrice, maxProfit
}

func main() {
	// - raw data
	url := "https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt"

	// - preprocessing data
	prepResult, err := urlToLines(url)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	// - max profit logic
	minPrice, maxProfit := maxProfit(prepResult)

	fmt.Printf("BuyAt : %d\n", minPrice)
	fmt.Printf("SellAt : %d\n", maxProfit)
	fmt.Printf("Maximum Profit : %d\n", maxProfit-minPrice)
}

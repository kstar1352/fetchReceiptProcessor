package server

import (
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")

func JsonWriter(rw http.ResponseWriter, jsonData interface{}) {
	data, err := json.Marshal(jsonData)
	if err != nil {
		rw.Write([]byte(err.Error()))
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

func CalcPoints(receipt Receipt) int {
	name := strings.Trim(receipt.Retailer, " ")
	items := receipt.Items
	date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	t, _ := time.Parse("15:04", receipt.PurchaseTime)
	total := receipt.Total
	points := 0
	points = calcName(name, points)
	points = calcTotal(total, points)
	points = calcItems(items, points)
	points = calcDateTime(date, t, points)
	return points

}

func calcDateTime(date, t time.Time, points int) int {
	startT, _ := time.Parse("15:04", "14:00")
	endT, _ := time.Parse("15:04", "16:00")
	if date.Day()%2 == 1 {
		points += 6
	}
	if t.After(startT) && t.Before(endT) {
		points += 10
	}
	return points
}

func calcName(name string, points int) int {
	for i := range name {
		letter := name[i]
		if isAlphanumeric(string(letter)) {
			points += 1
		}
	}
	return points
}

func calcTotal(total string, points int) int {
	cents := strings.Split(total, ".")[1]
	if cents == "00" {
		points += 50
	}
	if cents == "00" || cents == "25" || cents == "50" || cents == "75" {
		points += 25
	}
	return points
}

func calcItems(items []Item, points int) int {
	itemLen := len(items)
	pairs := itemLen / 2
	newPoints := 5 * pairs
	points += newPoints
	for i := range items {
		item := items[i]
		trimmedName := strings.Trim(item.ShortDescription, " ")
		if len(trimmedName)%3 == 0 {
			floatPrice, _ := strconv.ParseFloat(item.Price, 64)
			addAmt := int(math.Ceil(floatPrice * 0.2))
			points += addAmt
		}
	}
	return points
}

func isAlphanumeric(s string) bool {
	return alphanumeric.MatchString(s)
}

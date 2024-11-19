package models

import (
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

type Receipt struct {
	Retailer     string        `json:"retailer" validate:"required"`
	PurchaseDate string        `json:"purchaseDate" validate:"required,date"`
	PurchaseTime string        `json:"purchaseTime" validate:"required,time"`
	Items        []ReceiptItem `json:"items"`
	Total        string        `json:"total" validate:"required,numeric"`
}

func IsDate(fl validator.FieldLevel) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

func IsTime(fl validator.FieldLevel) bool {
	layout := "15:04"
	_, err := time.Parse(layout, fl.Field().String())
	return err == nil
}

func ClearString(str string) string {
	rx := regexp.MustCompile(`[^a-zA-Z0-9]+`) // alphanumeric, not space
	return rx.ReplaceAllString(str, "")
}

func HasNoCents(str string) bool {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return val == float64(int(val))
}

func IsMultipleOfQuarter(str string) bool {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return num*4 == float64(int(num*4))
}

func IsSpecificDate(str string) bool {
	layout := "2006-01-02"
	tt, err := time.Parse(layout, str)
	if err != nil {
		return false
	}
	return tt.Day()%2 == 1
}

func IsSpecificTime(str string) bool {
	layout := "15:04"
	tt, err := time.Parse(layout, str)
	if err != nil {
		return false
	}
	startTime, _ := time.Parse(layout, "14:00")
	endTime, _ := time.Parse(layout, "16:00")
	return tt.Hour() >= startTime.Hour() && tt.Hour() < endTime.Hour()
}

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,numeric"`
}

func RoundPrice(amount string) int {
	num, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0
	}
	return int(math.Ceil(num * 0.2))
}

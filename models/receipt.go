package models

import (
	"regexp"

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
	rx := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return rx.MatchString(fl.Field().String())
}

func IsTime(fl validator.FieldLevel) bool {
	rx := regexp.MustCompile(`^\d{2}:\d{2}$`)
	return rx.MatchString(fl.Field().String())
}

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required,numeric"`
}

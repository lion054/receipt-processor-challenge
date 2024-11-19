package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"receipt-processor-challenge/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

/*
 * POST /receipts/process
 *
 * Store a receipt
 */

type StoreReceiptResp struct {
	Id string `json:"id"`
}

func (s *Server) StoreReceipt(c *gin.Context) {
	// validate payload
	dec := json.NewDecoder(c.Request.Body)
	dec.DisallowUnknownFields()
	var params models.Receipt
	err := dec.Decode(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	validate := validator.New()
	validate.RegisterValidation("date", models.IsDate)
	validate.RegisterValidation("time", models.IsTime)
	err = validate.Struct(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// insert data to map
	id := uuid.NewString()
	s.DataMap[id] = params

	// return response
	c.JSON(http.StatusOK, StoreReceiptResp{
		Id: id,
	})
}

/*
 * GET /receipts/:key/points
 *
 * Calculate the points of a receipt
 */

type CalcPointsResp struct {
	Points int `json:"points"`
}

func (s *Server) CalcPoints(c *gin.Context) {
	// validate params
	key := c.Param("key")
	record, ok := s.DataMap[key]
	if !ok {
		c.JSON(http.StatusNotFound, errors.New("this company does not exist"))
		return
	}

	var points = 0
	// One point for every alphanumeric character in the retailer name.
	retailer := models.ClearString(record.Retailer)
	points += len(retailer)
	// 50 points if the total is a round dollar amount with no cents.
	if models.HasNoCents(record.Total) {
		points += 50
	}
	// 25 points if the total is a multiple of 0.25.
	if models.IsMultipleOfQuarter(record.Total) {
		points += 25
	}
	// 5 points for every two items on the receipt.
	itemCount := len(record.Items)
	points += (itemCount / 2) * 5
	// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range record.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			points += models.RoundPrice(item.Price)
		}
	}
	// 6 points if the day in the purchase date is odd.
	if models.IsSpecificDate(record.PurchaseDate) {
		points += 6
	}
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	if models.IsSpecificTime(record.PurchaseTime) {
		points += 10
	}

	// return response
	c.JSON(http.StatusOK, CalcPointsResp{
		Points: points,
	})
}

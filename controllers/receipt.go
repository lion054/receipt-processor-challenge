package controllers

import (
	"encoding/json"
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
	params.Retailer = strings.TrimSpace(params.Retailer)
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

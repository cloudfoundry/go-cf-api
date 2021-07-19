package controllers

import (
	"context"
	"net/http"

	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/config"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/logging"
	mysqlModels "github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/sqlboiler/mysql"
	psqlModels "github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/sqlboiler/postgres"
	"github.com/FloThinksPi/golang-vuejs-bootstrap/internal/app/gopilot/storage/db"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
)

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /accounts [get]
func ListPilots(c echo.Context) error {
	db := db.GetConnection()

	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true))
	pilots, err := mysqlModels.Pilots(qm.Limit(5)).All(ctx, db)
	if err != nil {
		zap.L().Error("Couldn't select", zap.Error(err))
	}
	return c.JSON(http.StatusOK, pilots)
}

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Failure default {object} httputil.DefaultError
// @Router /accounts/{id} [get]
func InsertPilot(c echo.Context) error {
	db := db.GetConnection()
	var p1 interface{}
	if config.Get().DB.Type == "postgres" {
		p1 = psqlModels.Pilot{}
	} else {
		p1 = mysqlModels.Pilot{}
	}

	p1.Name = "Larry"
	ctx := boil.WithDebugWriter(boil.WithDebug(context.Background(), true), logging.NewBoilLogger(true))
	err := p1.Insert(ctx, db, boil.Infer())
	if err != nil {
		zap.L().Error("Couldn't insert", zap.Error(err))
	}
	return c.JSON(http.StatusOK, p1)
}

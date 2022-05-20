package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"timey/context"
	"timey/model"
	"timey/service"
)

type SOWController struct {
	repo service.StatementOfWorkRepository
}

func InitSOWs() {
	e, err := context.Get[echo.Echo]("echo")
	if err != nil {
		logrus.Panic(err)
	}

	repo, err := context.Get[service.StatementOfWorkRepository](service.SowRepoQualifier)
	if err != nil {
		logrus.Panic(err)
	}

	controller := SOWController{*repo}

	e.GET("/api/customers/:id/sows", controller.GetAllSOWs)
	e.POST("/api/customers/:id/sows", controller.CreateSOW)
	e.GET("/api/customers/:id/sows/:sowId", controller.GetSOW)
}

func (c SOWController) GetAllSOWs(ctx echo.Context) error {
	customerId := ctx.Param("id")

	sows, err := c.repo.GetByCustomerID(customerId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, sows)
}

func (c SOWController) GetSOW(ctx echo.Context) error {
	customerId := ctx.Param("id")
	sowId := ctx.Param("sowId")

	sow, err := c.repo.GetByCustomerIDAndSowID(customerId, sowId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, sow)
}

func (c SOWController) CreateSOW(ctx echo.Context) (err error) {
	customerId := ctx.Param("id")

	var sow model.StatementOfWork
	err = ctx.Bind(&sow)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	sow.CustomerID = customerId
	_, err = c.repo.Create(sow.ID, &sow)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	ctx.Response().Header().Set("Location", fmt.Sprintf("/api/customers/%s/sows/%s", customerId, sow.ID))
	ctx.Response().WriteHeader(http.StatusCreated)

	return nil
}

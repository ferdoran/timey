package api

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
	"timey/context"
	"timey/model"
	"timey/service"
)

type CustomersController struct {
	repo service.CRUDRepository[string, model.Customer]
}

func InitCustomers() {
	e, err := context.Get[echo.Echo]("echo")
	if err != nil {
		logrus.Panic(err)
	}

	repo, err := context.Get[service.CRUDRepository[string, model.Customer]](service.CustomerRepoQualifier)
	if err != nil {
		logrus.Panic(err)
	}

	controller := CustomersController{*repo}

	e.GET("/api/customers", controller.GetAllCustomers)
	e.POST("/api/customers", controller.CreateCustomer)
	e.GET("/api/customers/:id", controller.GetCustomer)
	e.DELETE("/api/customers/:id", controller.DeleteCustomer)
}

func (cc CustomersController) GetAllCustomers(c echo.Context) error {
	customers, err := cc.repo.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, customers)
}

func (cc CustomersController) CreateCustomer(c echo.Context) error {
	var customer model.Customer
	if err := c.Bind(&customer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if _, err := cc.repo.Create(customer.ID, &customer); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Location", "/api/customers/"+customer.ID)
	c.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (cc CustomersController) DeleteCustomer(c echo.Context) error {
	id := c.Param("id")

	if err := cc.repo.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

func (cc CustomersController) GetCustomer(c echo.Context) error {
	id := c.Param("id")
	var customer model.Customer
	customer, err := cc.repo.Get(id)
	if err != nil {
		if errors.Is(err, service.ErrEntryNotFound) {
			return c.NoContent(http.StatusNotFound)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, customer)
}

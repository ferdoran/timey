package api

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"timey/context"
	"timey/model"
	"timey/service"
)

var (
	customer = model.Customer{
		ID:   "test",
		Name: "test",
	}
	customerRepo = service.NewInMemoryRepository[string, model.Customer]()
	controller   = CustomersController{&customerRepo}
)

func TestInitCustomers(t *testing.T) {
	t.Run("customer controller initialises", func(t *testing.T) {
		context.Bind("echo", echo.New())
		var repo service.CRUDRepository[string, model.Customer] = &customerRepo
		context.Bind[service.CRUDRepository[string, model.Customer]](service.CustomerRepoQualifier, &repo)
		InitCustomers()
	})
}

func TestCustomersController_CreateCustomer(t *testing.T) {
	e := echo.New()
	customerJson := `{"id":"1","name":"test"}`
	req := httptest.NewRequest(http.MethodPost, "/api/customers", strings.NewReader(customerJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("creates customer successfully", func(t *testing.T) {
		err := controller.CreateCustomer(c)
		if err != nil {
			t.Error(err)
		}

		if rec.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, rec.Code)
		}

		if location := rec.Header().Get("Location"); location != "/api/customers/1" {
			t.Errorf("expected Location header to contain %s but got %s", "/api/customers/1", location)
		}
	})
}

func TestCustomersController_GetCustomer(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/customers/"+customer.ID, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/customers/:id")
	c.SetParamNames("id")
	c.SetParamValues(customer.ID)
	_, err := customerRepo.Create(customer.ID, &customer)
	if err != nil {
		t.Error("failed to create customer")
	}

	t.Run("gets customer successfully", func(t *testing.T) {

		err = controller.GetCustomer(c)

		if rec.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusCreated, rec.Code)
		}

		var result model.Customer
		err = json.Unmarshal(rec.Body.Bytes(), &result)
		if err != nil {
			t.Error(err)
		}

		if result.ID != customer.ID || result.Name != customer.Name {
			t.Errorf("Expected %v but got %v", customer, result)
		}
	})
}

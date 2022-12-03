package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/endyApina/golang-meetup/http/helpers"
)

type ListModel struct {
	Status string          `json:"status"`
	Data   []EmployeeModel `json:"data"`
}

type EmployeeModel struct {
	ID             int    `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

type API struct {
	URL string
}

func NewClientHandler(url string) *API {
	return &API{
		URL: url,
	}
}

func (api *API) GetEmployees() (*ListModel, error) {
	employees := &ListModel{}

	to := time.Duration(10)
	opt := &helpers.HttpOptions{
		Ctx:    context.Background(),
		Url:    api.URL + "/api/v1/employees",
		TO:     &to,
		Method: http.MethodGet,
	}

	_, err := helpers.DoRequest(opt, employees)
	return employees, err
}

func GetEmployeeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	w.Header().Set("Content-Type", "application/json")
	clientHandlers := NewClientHandler("http://dummy.restapiexample.com/")
	models, err := clientHandlers.GetEmployees()
	if err != nil {
		errorMesaage := fmt.Sprintf("%v", err.Error())
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(errorMesaage)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models)
}

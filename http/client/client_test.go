package client

import (
	"net/http"
	"testing"

	"github.com/endyApina/golang-meetup/http/helpers"
)

func TestAPIGetEmployees(t *testing.T) {
	resp := &ListModel{
		Status: "success",
		Data: []EmployeeModel{
			{
				ID:             1,
				EmployeeName:   "Tiger Nixon",
				EmployeeSalary: 320800,
				EmployeeAge:    61,
				ProfileImage:   "",
			},
		},
	}

	srv := helpers.HttpMock("/api/v1/employees", http.StatusOK, resp)
	defer srv.Close()

	api := API{URL: srv.URL}

	employees, err := api.GetEmployees()
	if err != nil {
		t.Error(err)
	}

	if err != nil {
		t.Error("expected", nil, "got", err.Error())
	}
	if employees.Status != "success" {
		t.Error("expected status success got:", employees.Status)
	}
	if len(employees.Data) != 1 {
		t.Error("expected 1 data got", len(employees.Data))
	}
}

func TestAPIFailGetEmployees(t *testing.T) {
	resp := &ListModel{
		Status: "success",
		Data: []EmployeeModel{
			{
				ID:             1,
				EmployeeName:   "Tiger Nixon",
				EmployeeSalary: 320800,
				EmployeeAge:    61,
				ProfileImage:   "",
			},
		},
	}

	srv := helpers.HttpMock("/invalid", http.StatusOK, resp)
	defer srv.Close()

	api := API{URL: srv.URL}

	employees, err := api.GetEmployees()
	if err == nil {
		t.Error(err)
	}

	if err == nil {
		t.Error("expected", err.Error(), "got", nil)
	}
	if employees.Status == "success" {
		t.Error("expected status error got success")
	}
	if len(employees.Data) == 1 {
		t.Error("expected 0 data got", len(employees.Data))
	}
}

package tests

import (
	"asira_lender/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestBankServiceList(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+adminBasicToken)
	})

	adminToken := getLenderAdminToken(e, auth)

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	// valid response
	auth.GET("/admin/bank_services").
		Expect().
		Status(http.StatusOK).JSON().Object()

	// test query found
	obj := auth.GET("/admin/bank_services").WithQuery("service_id", "1").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("total_data").ValueEqual("total_data", 1)
	obj = auth.GET("/admin/bank_services").WithQuery("bank_id", "1").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("total_data").ValueEqual("total_data", 3)

	// test query not found
	obj = auth.GET("/admin/bank_services").WithQuery("bank_id", "99").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("total_data").ValueEqual("total_data", 0)
}

func TestNewBankService(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+adminBasicToken)
	})

	adminToken := getLenderAdminToken(e, auth)

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	// normal scenario
	payload := map[string]interface{}{
		"service_id": "1",
		"bank_id":    "1",
		"image":      "base64 encoded long string of image",
		"status":     "active",
	}
	auth.POST("/admin/bank_services").WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()
	payload = map[string]interface{}{
		"service_id": "1",
		"bank_id":    "1",
		"image":      "base64 encoded long string of image",
		"status":     "inactive",
	}
	auth.POST("/admin/bank_services").WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	// invalids
	payload = map[string]interface{}{
		"service_id": "1",
		"bank_id":    "1",
		"image":      "base64 encoded long string of image",
		"status":     "not valid",
	}
	auth.POST("/admin/bank_services").WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()

	payload = map[string]interface{}{
		"service_id": "99",
		"bank_id":    "1",
		"image":      "base64 encoded long string of image",
		"status":     "active",
	}
	auth.POST("/admin/bank_services").WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()

	payload = map[string]interface{}{
		"service_id": "1",
		"bank_id":    "99",
		"image":      "base64 encoded long string of image",
		"status":     "active",
	}
	auth.POST("/admin/bank_services").WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()
}

func TestGetBankServicebyID(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+adminBasicToken)
	})

	adminToken := getLenderAdminToken(e, auth)

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	// valid response
	obj := auth.GET("/admin/bank_services/1").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("id").ValueEqual("id", 1)

	// not found
	auth.GET("/admin/bank_services/9999").
		Expect().
		Status(http.StatusNotFound).JSON().Object()
}

func TestPatchBankService(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+adminBasicToken)
	})

	adminToken := getLenderAdminToken(e, auth)

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	payload := map[string]interface{}{
		"status": "inactive",
	}

	// valid response
	obj := auth.PATCH("/admin/bank_services/1").WithJSON(payload).
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("status").ValueEqual("status", "inactive")

	// invalid status
	payload = map[string]interface{}{
		"status": "not valid",
	}
	auth.PATCH("/admin/bank_services/1").WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()

	// test invalid token
	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer wrong token")
	})
	auth.PATCH("/admin/bank_services/1").WithJSON(payload).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object()
}

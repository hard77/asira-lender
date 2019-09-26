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

	// test query not found
	obj = auth.GET("/admin/bank_services").WithQuery("bank_id", "99").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("total_data").ValueEqual("total_data", 0)
}

// func TestNewService(t *testing.T) {
// 	RebuildData()

// 	api := router.NewRouter()

// 	server := httptest.NewServer(api)

// 	defer server.Close()

// 	e := httpexpect.New(t, server.URL)

// 	auth := e.Builder(func(req *httpexpect.Request) {
// 		req.WithHeader("Authorization", "Basic "+adminBasicToken)
// 	})

// 	adminToken := getLenderAdminToken(e, auth)

// 	auth = e.Builder(func(req *httpexpect.Request) {
// 		req.WithHeader("Authorization", "Bearer "+adminToken)
// 	})

// 	payload := map[string]interface{}{
// 		"name":   "Test New Bank Service",
// 		"status": "active",
// 	}

// 	// normal scenario
// 	obj := auth.POST("/admin/services").WithJSON(payload).
// 		Expect().
// 		Status(http.StatusCreated).JSON().Object()
// 	obj.ContainsKey("name").ValueEqual("name", "Test New Bank Service")

// 	// invalid status
// 	payload = map[string]interface{}{
// 		"name":   "Test New Bank Service",
// 		"status": "not valid",
// 	}
// 	auth.PATCH("/admin/services/1").WithJSON(payload).
// 		Expect().
// 		Status(http.StatusUnprocessableEntity).JSON().Object()

// 	// test invalid
// 	payload = map[string]interface{}{
// 		"name": "",
// 	}
// 	auth.POST("/admin/services").WithJSON(payload).
// 		Expect().
// 		Status(http.StatusUnprocessableEntity).JSON().Object()
// }

// func TestGetServicebyID(t *testing.T) {
// 	RebuildData()

// 	api := router.NewRouter()

// 	server := httptest.NewServer(api)

// 	defer server.Close()

// 	e := httpexpect.New(t, server.URL)

// 	auth := e.Builder(func(req *httpexpect.Request) {
// 		req.WithHeader("Authorization", "Basic "+adminBasicToken)
// 	})

// 	adminToken := getLenderAdminToken(e, auth)

// 	auth = e.Builder(func(req *httpexpect.Request) {
// 		req.WithHeader("Authorization", "Bearer "+adminToken)
// 	})

// 	// valid response
// 	obj := auth.GET("/admin/services/1").
// 		Expect().
// 		Status(http.StatusOK).JSON().Object()
// 	obj.ContainsKey("id").ValueEqual("id", 1)

// 	// not found
// 	auth.GET("/admin/services/9999").
// 		Expect().
// 		Status(http.StatusNotFound).JSON().Object()
// }

// func TestPatchService(t *testing.T) {
// 	RebuildData()

// 	api := router.NewRouter()

// 	server := httptest.NewServer(api)

// 	defer server.Close()

// 	e := httpexpect.New(t, server.URL)

// 	auth := e.Builder(func(req *httpexpect.Request) {
// 		req.WithHeader("Authorization", "Basic "+adminBasicToken)
// 	})

// 	adminToken := getLenderAdminToken(e, auth)

// 	auth = e.Builder(func(req *httpexpect.Request) {
// 		req.WithHeader("Authorization", "Bearer "+adminToken)
// 	})

// 	payload := map[string]interface{}{
// 		"name": "Test Service Patch",
// 	}

// 	// valid response
// 	obj := auth.PATCH("/admin/services/1").WithJSON(payload).
// 		Expect().
// 		Status(http.StatusOK).JSON().Object()
// 	obj.ContainsKey("name").ValueEqual("name", "Test Service Patch")

// 	// invalid status
// 	payload = map[string]interface{}{
// 		"status": "not valid",
// 	}
// 	auth.PATCH("/admin/services/1").WithJSON(payload).
// 		Expect().
// 		Status(http.StatusUnprocessableEntity).JSON().Object()

// 	// test invalid token
// 	auth = e.Builder(func(req *httpexpect.Request) {
// 		req.WithHeader("Authorization", "Bearer wrong token")
// 	})
// 	auth.PATCH("/admin/services/1").WithJSON(payload).
// 		Expect().
// 		Status(http.StatusUnauthorized).JSON().Object()
// }

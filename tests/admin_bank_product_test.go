package tests

import (
	"asira_lender/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestBankProductList(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+clientBasicToken)
	})

	adminToken := getAdminLoginToken(e, auth, "1")

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	// valid response
	auth.GET("/admin/bank_products").
		Expect().
		Status(http.StatusOK).JSON().Object()

	// test query found
	obj := auth.GET("/admin/bank_products").WithQuery("product_id", "1").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("total_data").ValueEqual("total_data", 1)

	//with part of name
	obj = auth.GET("/admin/bank_products").WithQuery("assurance", "assu").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("total_data").ValueEqual("total_data", 5)

	// test query invalid
	obj = auth.GET("/admin/bank_products").WithQuery("assurance", "should not found this").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("total_data").ValueEqual("total_data", 0)
}

func TestNewBankProduct(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+clientBasicToken)
	})

	adminToken := getAdminLoginToken(e, auth, "1")

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	payload := map[string]interface{}{
		"product_id":      1,
		"bank_service_id": 1,
		"min_timespan":    3,
		"max_timespan":    12,
		"interest":        5,
		"min_loan":        1000000,
		"max_loan":        10000000,
		"fees": []interface{}{
			map[string]interface{}{
				"description": "Admin Fee",
				"amount":      "1%",
			},
			map[string]interface{}{
				"description": "Convenience Fee",
				"amount":      "2%",
			},
		},
		"collaterals":      []string{"Surat Tanah", "BPKB"},
		"financing_sector": []string{"Pendidikan"},
		"assurance":        "assuransi",
		"status":           "active",
	}

	// normal scenario
	auth.POST("/admin/bank_products").WithJSON(payload).
		Expect().
		Status(http.StatusCreated).JSON().Object()

	// test invalid
	payload = map[string]interface{}{
		"name": "",
	}
	auth.POST("/admin/bank_products").WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()
}

func TestGetBankProductbyID(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+clientBasicToken)
	})

	adminToken := getAdminLoginToken(e, auth, "1")

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	// valid response
	obj := auth.GET("/admin/bank_products/1").
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("id").ValueEqual("id", 1)

	// not found
	auth.GET("/admin/bank_products/9999").
		Expect().
		Status(http.StatusNotFound).JSON().Object()
}

func TestPatchBankProduct(t *testing.T) {
	RebuildData()

	api := router.NewRouter()

	server := httptest.NewServer(api)

	defer server.Close()

	e := httpexpect.New(t, server.URL)

	auth := e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Basic "+clientBasicToken)
	})

	adminToken := getAdminLoginToken(e, auth, "1")

	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+adminToken)
	})

	payload := map[string]interface{}{
		"status": "inactive",
	}

	// valid response
	obj := auth.PATCH("/admin/bank_products/1").WithJSON(payload).
		Expect().
		Status(http.StatusOK).JSON().Object()
	obj.ContainsKey("status").ValueEqual("status", "inactive")

	// valid response
	payload = map[string]interface{}{
		"status": "invalid",
	}
	auth.PATCH("/admin/bank_products/1").WithJSON(payload).
		Expect().
		Status(http.StatusUnprocessableEntity).JSON().Object()

	// test invalid token
	auth = e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer wrong token")
	})
	auth.PATCH("/admin/bank_products/1").WithJSON(payload).
		Expect().
		Status(http.StatusUnauthorized).JSON().Object()
}

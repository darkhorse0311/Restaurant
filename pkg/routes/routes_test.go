package routes

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var port = "9001"

func TestGetServerIsUp(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:"+port, nil)

	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rec := httptest.NewRecorder()
	GetServerIsUp(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response %v", err)
	}

	if string(b) != "server is live" {
		t.Fatalf("expected \"server is live\"; got \"%s\"", b)
	}
}

func TestRouteNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:"+port+"/dfsfsdfsdfsdfsdfs", nil)

	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rec := httptest.NewRecorder()
	RouteNotFound(rec, req)

	res := rec.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected status NotFound; got %v", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response %v", err)
	}

	if string(b) != "route not found" {
		t.Fatalf("expected \"route not found\"; got \"%s\"", b)
	}
}

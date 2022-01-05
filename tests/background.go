package tests

import (
	"flag"
	_ "github.com/lib/pq"
	"io"
	"net/http"
	"net/http/httptest"
	"test/internal/app/github/interfaces/apis"
	"test/tests/pkg/testcontainer"
)

var resourcesPath = flag.String("f", "../.", "set resources file which viper will loading.")

func setUp() *testcontainer.Background {
	flag.Parse()
	background, clean, err := CreateBackground(*resourcesPath)
	defer clean()
	if err != nil {
		panic(err)
	}
	return background
}

func callAPI(api apis.API, method, url string, body io.Reader) *http.Response {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, body)
	api.GetInfraContext().GetRoute().ServeHTTP(w, r)
	return w.Result()
}

func NewMockAPI() *apis.API {
	return &apis.API{}
}

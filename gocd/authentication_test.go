package gocd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAuthentication_Login(t *testing.T) {

	setup()
	defer teardown()

	mockCookie := "JSESSIONID=hash;Path=/go;Expires=Mon, 15-Jun-2015 10:16:20 GMT"

	mux.HandleFunc("/api/api/agents", func(w http.ResponseWriter, r *http.Request) {
		testStringInSlice(t, r.Header["Accept"], "application/vnd.go.cd.v2+json")
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)

		w.Header().Set("Set-Cookie", mockCookie)

		j, _ := ioutil.ReadFile("test/resources/agents.0.json")
		fmt.Fprint(w, string(j))
	})

	client.Login(context.Background())

	if client.cookie != mockCookie {
		t.Errorf("Expected '%s'. Got '%s'.", mockCookie, client.cookie)
	}

}

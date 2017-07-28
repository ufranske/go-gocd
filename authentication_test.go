package gocd

import (
	"testing"
	"context"
	"io/ioutil"
	"fmt"
	"net/http"
)

func TestAuthentication_Login(t *testing.T) {

	setup()
	defer teardown()

	mock_cookie := "JSESSIONID=hash;Path=/go;Expires=Mon, 15-Jun-2015 10:16:20 GMT"

	mux.HandleFunc("/api/api/agents", func(w http.ResponseWriter, r *http.Request) {
		testStringInSlice(t, r.Header["Accept"], "application/vnd.go.cd.v2+json")
		testMethod(t, r, "GET")
		testAuth(t, r, mockAuthorization)

		w.Header().Set("Set-Cookie", mock_cookie, )

		j, _ := ioutil.ReadFile("test/resources/agents.0.json")
		fmt.Fprint(w, string(j))
	})

	client.Login(context.Background())

	if client.cookie != mock_cookie {
		t.Errorf("Expected '%s'. Got '%s'.", mock_cookie, client.cookie)
	}

}

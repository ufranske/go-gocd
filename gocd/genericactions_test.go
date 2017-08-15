package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenericActions(t *testing.T) {
	setup()
	defer teardown()

	t.Run("HeadSuccess", funcTestGenericHeadActionSuccess)
	t.Run("HeadFail", funcTestGenericHeadActionFail)
	t.Run("Post", funcTestGenericPost)
}
func funcTestGenericPost(t *testing.T) {
	mux.HandleFunc("/api/mock-post", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		fmt.Fprint(w, "")
	})
	_, _, err := client.postAction(context.Background(), &APIClientRequest{
		Path: "mock-post",
	})
	assert.Nil(t, err)
}

func funcTestGenericHeadActionFail(t *testing.T) {
	mux.HandleFunc("/api/mock-head-fail", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "HEAD", "Unexpected HTTP method")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "")
	})

	head, resp, err := client.genericHeadAction(context.Background(), "mock-head-fail", apiV1)
	if err != nil {
		t.Error(err)
	}
	assert.False(t, head)
	assert.NotNil(t, resp)
}

func funcTestGenericHeadActionSuccess(t *testing.T) {

	mux.HandleFunc("/api/mock-head-success", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "HEAD", "Unexpected HTTP method")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "")
	})

	head, resp, err := client.genericHeadAction(context.Background(), "mock-head-success", apiV1)
	assert.EqualError(t, err, "Received HTTP Status '400 Bad Request': ''")

	assert.True(t, head)
	assert.NotNil(t, resp)

}

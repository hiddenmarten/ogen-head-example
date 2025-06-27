package internal

import (
	"context"
	"errors"
	"github.com/hiddenmarten/ogen-head-example/api/client"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHead200(t *testing.T) {
	go func() {
		panic(http.ListenAndServe(":8080", NewServeMux()))
	}()
	c, err := client.NewClient("http://localhost:8080")
	require.NoError(t, err)
	resp, err := c.HeadFilesByFile(context.Background(), client.HeadFilesByFileParams{
		File: "1.txt",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, client.OptString{Value: "bytes", Set: true}, resp.AcceptRanges)
	require.Equal(t, client.OptString{Value: "Tue, 22 Apr 2025 19:01:37 GMT", Set: true}, resp.LastModified)
}

func TestHead404(t *testing.T) {
	go func() {
		panic(http.ListenAndServe(":8080", NewServeMux()))
	}()
	c, err := client.NewClient("http://localhost:8080")
	require.NoError(t, err)
	_, err = c.HeadFilesByFile(context.Background(), client.HeadFilesByFileParams{
		File: "2.txt",
	})

	var errorModel *client.ErrorModelStatusCode
	ok := errors.As(errors.Unwrap(errors.Unwrap(err)), &errorModel)
	if !ok {
		t.Errorf("expected UnsuccessfulResponse, got %T: %v", err, err)
		return
	}
	require.Equal(t, errorModel.StatusCode, http.StatusNotFound)
}

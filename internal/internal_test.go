package internal

import (
	"context"
	"github.com/hiddenmarten/ogen-head-example/api/client"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHead204Happy(t *testing.T) {
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

func TestHead404Panic(t *testing.T) {
	go func() {
		panic(http.ListenAndServe(":8080", NewServeMux()))
	}()
	c, err := client.NewClient("http://localhost:8080")
	require.NoError(t, err)
	resp, err := c.HeadFilesByFile(context.Background(), client.HeadFilesByFileParams{
		File: "2.txt",
	})

	// ATTENTION!!!
	// ogen can't decode a response with an empty body from a head request
	// despite having no fields being required in "#/components/schemas/ErrorModel".
	// error: decode ErrorModel: "{" expected
	require.NoError(t, err)
	require.NotNil(t, resp)
}

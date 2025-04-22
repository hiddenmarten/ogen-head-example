package internal

import (
	"context"
	"net/http"
	"net/http/httptest"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	_ = NewHumaAPI(context.Background(), mux)
	return mux
}

func NewHumaAPI(ctx context.Context, mux *http.ServeMux) huma.API {
	api := humago.New(mux, huma.DefaultConfig("Static file informer", "1.0.0"))
	huma.Head(api, "/files/{file}", head)
	return api
}

func head(ctx context.Context, in *HeadInput) (*HeadOutput, error) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodHead, "/"+in.File, nil)
	http.FileServer(http.Dir("files")).ServeHTTP(w, req)
	if w.Code == http.StatusOK {
		return &HeadOutput{
			ContentType:   w.Header().Get("Content-Type"),
			ContentLength: w.Header().Get("Content-Length"),
			AcceptRanges:  w.Header().Get("Accept-Ranges"),
			LastModified:  w.Header().Get("Last-Modified"),
		}, nil
	} else if w.Code == http.StatusNotFound {
		return nil, huma.Error404NotFound("File not found")
	}
	return nil, huma.Error500InternalServerError("Internal server error")
}

type HeadInput struct {
	File string `path:"file" description:"File name" example:"1.txt" required:"true"`
}

type HeadOutput struct {
	ContentLength string `header:"Content-Length"`
	ContentType   string `header:"Content-Type"`
	AcceptRanges  string `header:"Accept-Ranges"`
	LastModified  string `header:"Last-Modified"`
}

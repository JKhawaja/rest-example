// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "GitHub SSH Keys": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/JKhawaja/replicated/design
// --out=$(GOPATH)\src\github.com\JKhawaja\replicated
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// ListKeysContext provides the keys list action context.
type ListKeysContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload ListKeysPayload
}

// NewListKeysContext parses the incoming request URL and body, performs validations and creates the
// context used by the keys controller list action.
func NewListKeysContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListKeysContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListKeysContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// ListKeysPayload is the keys list action payload.
type ListKeysPayload []string

// OK sends a HTTP response with status code 200.
func (ctx *ListKeysContext) OK(r UserCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	if r == nil {
		r = UserCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListKeysContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListKeysContext) Unauthorized(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListKeysContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListKeysContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

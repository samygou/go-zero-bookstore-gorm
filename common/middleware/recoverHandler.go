package middleware

import (
	"net/http"
	"time"

	"github.com/rs/xid"

	"go-zero-bookstore/common/errorx"
	"go-zero-bookstore/common/logx"
	"go-zero-bookstore/common/result"
)

type RecoverMiddleware struct {
	HTTPCtx *result.Context
}

var RecoverMiddlewareHandler *RecoverMiddleware

func NewRecoverMiddleware() *RecoverMiddleware {
	return &RecoverMiddleware{
		HTTPCtx: result.NewContext(),
	}
}

// Handle returns a middleware that recovers if panic happens.
func (m *RecoverMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.HTTPCtx.RequestId = xid.New().String()
		m.HTTPCtx.StartTime = time.Now().UnixNano() / 1000000
		m.HTTPCtx.HTTPCtx = r.Context()
		m.HTTPCtx.RW = w

		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case errorx.Status:
					logx.Error(err)
				default:
					logx.Error(err)
					m.HTTPCtx.ResponseJson(&result.JsonResponse{Err: errorx.NewError(errorx.CodeUnknown, "unknown err")})
				}
			}
		}()

		next.ServeHTTP(w, r)
	}
}

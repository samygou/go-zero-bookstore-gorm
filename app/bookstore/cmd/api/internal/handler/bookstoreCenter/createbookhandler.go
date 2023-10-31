package bookstoreCenter

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-bookstore/app/bookstore/cmd/api/internal/logic/bookstoreCenter"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/api/internal/types"
	"go-zero-bookstore/common/errorx"
	"go-zero-bookstore/common/result"
)

func CreateBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBookReq
		if err := httpx.Parse(r, &req); err != nil {
			svcCtx.RecoverMiddlewareCtx.HTTPCtx.ResponseJson(&result.JsonResponse{Err: errorx.NewError(errorx.CodeInvalidArgument, err.Error())})
			return
		}

		l := bookstoreCenter.NewCreateBookLogic(r.Context(), svcCtx)
		resp, err := l.CreateBook(&req)
		if err != nil {
			svcCtx.RecoverMiddlewareCtx.HTTPCtx.ResponseJson(&result.JsonResponse{Err: err})
		} else {
			svcCtx.RecoverMiddlewareCtx.HTTPCtx.ResponseJson(&result.JsonResponse{Value: result.M{
				"data": resp,
			}})
		}
	}
}

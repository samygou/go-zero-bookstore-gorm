package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-bookstore/app/usercenter/cmd/api/internal/logic/user"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/api/internal/types"
	"go-zero-bookstore/common/errorx"
	"go-zero-bookstore/common/result"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			svcCtx.RecoverMiddlewareCtx.HTTPCtx.ResponseJson(&result.JsonResponse{Err: errorx.NewError(errorx.CodeInvalidArgument, err.Error())})
			return
		}

		l := user.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update(&req)
		if err != nil {
			svcCtx.RecoverMiddlewareCtx.HTTPCtx.ResponseJson(&result.JsonResponse{Err: err})
			return
		} else {
			svcCtx.RecoverMiddlewareCtx.HTTPCtx.ResponseJson(&result.JsonResponse{Value: result.M{}})
			return
		}
	}
}

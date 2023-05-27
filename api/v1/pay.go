package v1

import (
	"github.com/CocaineCong/gin-mall/pkg/utils/ctl"
	"github.com/CocaineCong/gin-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/gin-mall/pkg/utils/log"
	"github.com/CocaineCong/gin-mall/types"
)

func OrderPaymentHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.PaymentDownReq
		if err := ctx.ShouldBind(&req); err != nil {
			// 参数校验
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(ctx, err))
			return
		}

		l := service.GetPaymentSrv()
		resp, err := l.PayDown(ctx.Request.Context(), &req)
		if err != nil {
			log.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ctx, err))
			return
		}
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, resp))
		return
	}
}

package v1

import (
	"github.com/gin-gonic/gin"
	"mall/consts"
	util "mall/pkg/utils"
	"mall/service"
)

func ListPsConfig(c *gin.Context) {
	listPisConfigService := service.PsConfigService{}
	if err := c.ShouldBind(&listPisConfigService); err == nil {
		res := listPisConfigService.List(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

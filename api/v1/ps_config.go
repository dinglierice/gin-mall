package v1

import (
	"github.com/gin-gonic/gin"
	"mall/consts"
	util "mall/pkg/utils"
	"mall/service"
)

// ListPsConfig  godoc
// @Summary      List psConfigs
// @Description  Get a list of psConfig
// @Tags         psConfigs
// @Accept       json
// @Produce      json
// @Success      200  {object}  serializer.Response
// @Router       /api/v1/ps_configs [get]
func ListPsConfig(c *gin.Context) {
	listPisConfigService := service.PsConfigService{}
	if err := c.ShouldBind(&listPisConfigService); err == nil {
		res := listPisConfigService.List2(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

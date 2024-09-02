package v1

import (
	"github.com/gin-gonic/gin"
	"mall/consts"
	util "mall/pkg/utils"
	"mall/service"
)

// ListPsConfigs  godoc
// @Summary      List psConfigs
// @Description  Get a list of psConfig
// @Tags         psConfigs
// @Accept       json
// @Produce      json
// @Success      200  {object}  serializer.Response
// @Router       /api/v1/ps_configs [get]
func ListPsConfigs(c *gin.Context) {
	listPisConfigService := service.PsConfigService{}
	if err := c.ShouldBind(&listPisConfigService); err == nil {
		res := listPisConfigService.List(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowPsConfig godoc
// @Summary      Get a psConfig
// @Description  Get a single psConfig by ID
// @Tags         psConfigs
// @Accept       json
// @Produce      json
// @Param        id   path    string  true  "PsConfig ID"
// @Success      200  {object}  serializer.Response
// @Router       /api/v1/ps_config/{id} [get]
func ShowPsConfig(c *gin.Context) {
	showPsConfigService := service.PsConfigService{}
	if err := c.ShouldBind(&showPsConfigService); err == nil {
		res := showPsConfigService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

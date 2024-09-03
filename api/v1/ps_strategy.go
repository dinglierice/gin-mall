package v1

import (
	"github.com/gin-gonic/gin"
	"mall/consts"
	util "mall/pkg/utils"
	"mall/service"
)

// ListPsStrategies godoc
// @Summary List PS strategies
// @Description Get a list of all PS strategies
// @Tags psStrategies
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
// @Success 200 {object} serializer.Response "Successful operation"
// @Failure 400 {object} serializer.Response "Bad request"
// @Router /api/v1/ps_strategies [get]
func ListPsStrategies(c *gin.Context) {
	listPsStrategyService := service.PsStrategyService{}
	if err := c.ShouldBind(&listPsStrategyService); err == nil {
		res := listPsStrategyService.List(c.Request.Context())
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowPsStrategy godoc
// @Summary Show a PS strategy
// @Description Get details of a specific PS strategy by ID
// @Tags psStrategies
// @Accept json
// @Produce json
// @Param id path string true "PS Strategy ID"
// @Success 200 {object} serializer.Response "Successful operation"
// @Failure 400 {object} serializer.Response "Bad request"
// @Failure 404 {object} serializer.Response "PS Strategy not found"
// @Router /api/v1/ps_strategy/{id} [get]
func ShowPsStrategy(c *gin.Context) {
	showPsStrategyService := service.PsStrategyService{}
	if err := c.ShouldBind(&showPsStrategyService); err == nil {
		res := showPsStrategyService.Show(c.Request.Context(), c.Param("id"))
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func CreatePsStrategy(c *gin.Context) {
	var createStrategyService service.PsStrategyService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createStrategyService); err == nil {
		res := createStrategyService.Create(c.Request.Context(), claims.ID)
		c.JSON(consts.StatusOK, res)
	} else {
		c.JSON(consts.IlleageRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

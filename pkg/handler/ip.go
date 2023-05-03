package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary		BlockIp
// @Tags			ip
// @Description	block ip
// @ID				block-ip
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.BlockIp	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/ip/block-ip [post]
func (h *Handler) blockIp(c *gin.Context) {
	type Body struct {
		IpAddress string `json:"ip_address"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resBlockIp, statusCode, err := h.services.BlockIp(body.IpAddress)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resBlockIp,
		})
		return
	}
	if resBlockIp {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "успешная блокировка ip адреса",
			"result":  resBlockIp,
		})
	}
	if !resBlockIp {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "ошибка блокировки ip адреса",
			"result":  resBlockIp,
		})
	}
}

// @Summary		CheckIpBlock
// @Tags			ip
// @Description	check ip block
// @ID				check-ip-block
// @Accept			json
// @Produce		json
// @Param			input	body		appl_row.CheckIpBlock	true	"account info"
// @Success		200		{integer}	integer				1
// @Failure		400,404	{object}	error
// @Failure		500		{object}	error
// @Failure		default	{object}	error
// @Router			/ip/check-ip-block [post]
func (h *Handler) checkIpBlock(c *gin.Context) {
	type Body struct {
		IpAddress string `json:"ip_address"`
	}
	var body Body
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "некорректно переданы данные в body",
		})
		return
	}
	resCheckIpBlock, statusCode, err := h.services.CheckIpBlock(body.IpAddress)
	if err != nil {
		c.JSON(statusCode, map[string]interface{}{
			"status":  statusCode,
			"message": err.Error(),
			"result":  resCheckIpBlock,
		})
		return
	}
	if resCheckIpBlock {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "ip адрес находится в блок листе",
			"result":  resCheckIpBlock,
		})
	}
	if !resCheckIpBlock {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  http.StatusOK,
			"message": "ip адрес не находится в блок листе",
			"result":  resCheckIpBlock,
		})
	}
}

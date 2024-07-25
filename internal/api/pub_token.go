package api

import (
	"app/internal/consts"
	"app/internal/help"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateTokenHandler(c *gin.Context) {

	// TODO:: User authentication
	token, err := help.Libs.Token.GenerateTokenData(123456)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  consts.StatusError,
			"message": "error when generate token data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  consts.StatusSuccess,
		"message": consts.StatusSuccess,
		"data":    token,
	})
}

func RefreshTokenHandler(c *gin.Context) {

	// 1.
	argv := struct {
		UserId       int64  `json:"user_id,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}{}
	err := c.ShouldBind(&argv)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  consts.StatusInvalidRequest,
			"message": consts.StatusInvalidRequest,
		})
		return
	}

	// 2. check refresh_token
	if !help.Libs.Token.CheckRefreshToken(argv.UserId, argv.RefreshToken) {
		c.JSON(http.StatusOK, gin.H{
			"status":  consts.StatusError,
			"message": "error refresh token",
		})
		return
	}

	// new access_token
	token, err := help.Libs.Token.GenerateTokenData(argv.UserId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  consts.StatusServerError,
			"message": consts.StatusServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  consts.StatusSuccess,
		"message": consts.StatusSuccess,
		"data":    token,
	})
}

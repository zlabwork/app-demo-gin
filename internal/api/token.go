package api

import (
	"app/internal/entity"
	"app/internal/msg"
	"app/internal/repo/cache"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
)

const (
	accessTokenTimeout  = time.Hour * 72
	refreshTokenTimeout = time.Hour * 24 * 20
	refreshTokenKey     = "_RT:%d"
)

func generateTokenData(userId int64) (*entity.Token, error) {

	audience := ""
	expiresAt := time.Now().Add(accessTokenTimeout)
	claims := &entity.TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			Subject:   "jwt",
			Audience:  []string{audience},
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-5 * time.Minute)),
			ID:        uuid.New().String(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	at, err := token.SignedString([]byte(os.Getenv("APP_KEY")))
	if err != nil {
		return nil, err
	}

	// Refresh Token
	rt := uuid.New().String()
	key := fmt.Sprintf(refreshTokenKey, userId)
	_ = cache.Set(context.TODO(), key, []byte(rt), refreshTokenTimeout)

	return &entity.Token{
		UserId:       userId,
		AccessToken:  at,
		RefreshToken: rt,
		ExpiresAt:    expiresAt,
	}, nil
}

func GenerateTokenHandler(c *gin.Context) {

	// TODO:: User authentication

	token, err := generateTokenData(123456)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  msg.StatusError,
			"message": "error when generate token data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  msg.StatusSuccess,
		"message": msg.StatusSuccess,
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
			"status":  msg.StatusInvalidRequest,
			"message": msg.StatusInvalidRequest,
		})
		return
	}

	// get refresh_token
	key := fmt.Sprintf(refreshTokenKey, argv.UserId)
	bs, err := cache.Get(context.TODO(), key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  msg.StatusServerError,
			"message": msg.StatusServerError,
		})
		return
	}

	// check
	if string(bs) != argv.RefreshToken {
		c.JSON(http.StatusOK, gin.H{
			"status":  msg.StatusUnauthorized,
			"message": msg.StatusUnauthorized,
		})
		return
	}

	// new access_token
	token, err := generateTokenData(argv.UserId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  msg.StatusServerError,
			"message": msg.StatusServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  msg.StatusSuccess,
		"message": msg.StatusSuccess,
		"data":    token,
	})
}

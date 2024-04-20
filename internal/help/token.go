package help

import (
	"app/internal/entity"
	"app/internal/repo/cache"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type tokenHelp struct {
	appName         string
	key             []byte
	accessTimeout   time.Duration
	refreshTimeout  time.Duration
	refreshCacheKey string
}

func newTokenHelp() *tokenHelp {
	return &tokenHelp{
		appName:         os.Getenv("APP_NAME"),
		key:             []byte(os.Getenv("APP_KEY")),
		accessTimeout:   time.Hour * 72,
		refreshTimeout:  time.Hour * 24 * 20,
		refreshCacheKey: "_RT:%d",
	}
}

func (tkh *tokenHelp) ParseTokenString(tokenString string) (*entity.TokenClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &entity.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tkh.key, nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*entity.TokenClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

func (tkh *tokenHelp) GenerateTokenData(userId int64) (*entity.Token, error) {

	exp := time.Now().Add(tkh.accessTimeout)

	// 1. Create Token
	at, err := tkh.generateTokenString(userId, exp)
	if err != nil {
		return nil, err
	}

	// 2. Refresh Token
	rt := uuid.New().String()
	key := fmt.Sprintf(tkh.refreshCacheKey, userId)
	err = cache.Set(context.TODO(), key, []byte(rt), tkh.refreshTimeout)
	if err != nil {
		log.Println(err)
	}

	return &entity.Token{
		UserId:       userId,
		AccessToken:  at,
		RefreshToken: rt,
		ExpiresAt:    exp,
	}, nil
}

func (tkh *tokenHelp) CheckRefreshToken(userId int64, refreshKey string) bool {
	key := fmt.Sprintf(tkh.refreshCacheKey, userId)
	bs, err := cache.Get(context.TODO(), key)
	if err != nil {
		return false
	}
	if string(bs) != refreshKey {
		return false
	}
	return true
}

func (tkh *tokenHelp) generateTokenString(userId int64, expiresAt time.Time) (string, error) {
	audience := ""
	claims := &entity.TokenClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tkh.appName,
			Subject:   "jwt",
			Audience:  []string{audience},
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-5 * time.Minute)),
			ID:        uuid.New().String(),
		},
	}

	// Create token with claims
	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return jwtClaims.SignedString(tkh.key)
}

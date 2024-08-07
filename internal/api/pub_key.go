package api

import (
	"app/internal/consts"
	"encoding/base64"
	"encoding/pem"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

var _publicKey []byte

type publicKey struct {
	Format    string `json:"format,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

func PublicKeyHandler(c *gin.Context) {

	if _publicKey == nil {

		// 1.
		f, err := os.Open("config/public.pem")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  consts.StatusMaintenance,
				"message": consts.StatusMaintenance,
			})
			return
		}
		bs, err := io.ReadAll(f)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  consts.StatusError,
				"message": consts.StatusError,
			})
			return
		}

		// 2.
		block, _ := pem.Decode(bs)
		_publicKey = block.Bytes
	}

	key := base64.StdEncoding.EncodeToString(_publicKey)
	c.JSON(http.StatusOK, gin.H{
		"status":  consts.StatusSuccess,
		"message": consts.StatusSuccess,
		"data": &publicKey{
			Format:    "pkcs1",
			PublicKey: key,
		},
	})
}

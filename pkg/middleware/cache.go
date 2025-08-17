package middleware

import (
	"inorder/pkg/utils"
	"time"
)

var authTokenCache *utils.CacheController

func init() {
	authTokenCache = utils.NewCacheController(time.Minute * 10)
}

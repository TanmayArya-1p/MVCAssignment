package models

import (
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"time"
)

var MenuCache *utils.SingleCache[[]*types.Item]
var ItemsCache *utils.CacheController
var TagsCache *utils.SingleCache[[]*types.Tag]

func init() {
	MenuCache = utils.NewSingleCache[[]*types.Item]()
	ItemsCache = utils.NewCacheController(time.Minute * 10)
	TagsCache = utils.NewSingleCache[[]*types.Tag]()
}

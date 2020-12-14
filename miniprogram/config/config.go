//Package config 小程序config配置
package config

import (
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/credential"
)

//Config config for 小程序
type Config struct {
	AppID     string `json:"app_id"`     //appid
	AppSecret string `json:"app_secret"` //appsecret
	Cache     cache.Cache
	AkHandle  *credential.AccessTokenHandle
}

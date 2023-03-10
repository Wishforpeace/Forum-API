package config

import "Forum-API/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 应用名称
			"name": config.Env("APP_NAME", "Forum"),

			// 当前环境，用以区分多环境，一般为local，stage，production，test
			"env": config.Env("APP_ENV", "production"),

			// 是否进去调试模式
			"debug": config.Env("APP_DEBUG", false),

			// 应用服务端口
			"port": config.Env("APP_PORT", "3000"),

			// 加密会话，JWT加密
			"key": config.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Env("APP_URL", "127.0.0.1:3000"),

			// 设置时区，JWT内会使用，日志记录也会使用
			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}

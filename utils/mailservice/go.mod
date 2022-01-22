module github.com/vfluxus/mailservice

go 1.16

replace github.com/uc-cdis/go-authutils => github.com/uc-cdis/go-authutils v0.0.0-20201026165355-17b5f353bf4f // lock version

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.7.2
	github.com/lib/pq v1.10.2
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/uc-cdis/go-authutils v0.0.0-00010101000000-000000000000
	github.com/vfluxus/dvergr v0.2.2
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
)

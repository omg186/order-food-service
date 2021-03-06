package main

import (
	"flag"
	"order-food-service/dao"
	"order-food-service/global"
	"order-food-service/routers"
)

var (
	SERVER_PORT          = flag.String("server_port", "9990", "server port")
	MYSQL_USER_NAME      = flag.String("username", "root", "mysql user name")
	MYSQL_PASSWORD       = flag.String("password", "12345678", "mysql password")
	MYSQL_URL            = flag.String("url", "localhost", "mysql url")
	MYSQL_PORT           = flag.String("port", "3306", "mysql port")
	MYSQL_DATABASE       = flag.String("database", "test", "mysql database")
	MYSQL_MAX_IDLE_CONNS = flag.Int("maxidle", 500, "mysql max idle conns")
	MYSQL_MAX_OPEN_CONNS = flag.Int("maxopen", 500, "mysql max open conns")

	REDIS_ADDR      = flag.String("redis_addr", "218.62.64.24:8993", "redis addr. 218.62.64.24:8993")
	REDIS_PASSWORD  = flag.String("redis_password", "yungengny", "redis password")
	REDIS_DB        = flag.Int("redis_db", 0, "redis db")
	REDIS_POOL_SIZE = flag.Int("redis_pool_size", 100, "redis pool size")
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://razeen.me

// @contact.name Razeen
// @contact.url https://razeen.me
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
// @tag.name Login
// @tag.description 登录
// @tag.name User
// @tag.description 用户
func main() {
	flag.Parse()
	//初始化MySQL
	dao.DB = dao.Init(*MYSQL_USER_NAME, *MYSQL_PASSWORD, *MYSQL_URL, *MYSQL_PORT, *MYSQL_DATABASE, *MYSQL_MAX_IDLE_CONNS, *MYSQL_MAX_OPEN_CONNS)

	//初始化Redis
	// redis.InitRedisClient(*REDIS_ADDR, *REDIS_PASSWORD, *REDIS_DB, *REDIS_POOL_SIZE)

	global.AppLog.Info("shrimp map service starting......")

	r := routers.Router()
	r.Run(":" + *SERVER_PORT)
}

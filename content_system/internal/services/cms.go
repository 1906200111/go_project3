package services

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	"go_project2/internal/api/operate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsAPP struct {
	db               *gorm.DB
	rdb              *redis.Client
	operateAppClient operate.AppClient
}

func connOperateAppClient(app *CmsAPP) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
	)
	if err != nil {
		panic(err)
	}
	client := operate.NewAppClient(conn)
	app.operateAppClient = client
}

func NewCmsApp() *CmsAPP {
	app := &CmsAPP{} //创建结构体实例
	connDB(app)      //给实例加上mysql
	connRdb(app)     //给实例加上redis
	return app
}

func connDB(app *CmsAPP) {
	//user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	mysqlDB, er := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"))
	if er != nil {
		panic(er)
	}
	//拿到mysqlDB的实例
	db, err := mysqlDB.DB()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100) //最大连接数
	db.SetMaxIdleConns(50)  //最大空闲连接数，一般为最大连接数/2
	//if env == "test" {
	//	mysqlDB = mysqlDB.Debug()
	//}
	app.db = mysqlDB
	//return mysqlDB
}

func connRdb(app *CmsAPP) {
	//redis-cli
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil && err != redis.Nil {
		panic(err)
	}
	app.rdb = rdb
}

package services

import (
	"content_system/internal/api/operate"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/redis/go-redis/v9"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type CmsAPP struct {
	db                *gorm.DB
	rdb               *redis.Client
	operateAppClient  operate.AppClient
	operateUserClient operate.UserClient
}

func NewCmsApp() *CmsAPP {
	app := &CmsAPP{}           //创建结构体实例
	connDB(app)                //给实例加上mysql
	connRdb(app)               //给实例加上redis
	connOperateAppClient(app)  //给实例加上内容grpc服务
	connOperateUserClient(app) //给实例加上用户grpc服务
	return app
}

func connOperateAppClient(app *CmsAPP) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"}, //从etcd中，服务的发现
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///content_manage" //把etcd的Name标识符拿过来，找到对应的服务ip
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中
		//我们的服务，现在微服务化了，有很多的分布式节点，需要etcd帮我们存储现在所有的分布式节点
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), //服务的发现
	)
	if err != nil {
		panic(err)
	}
	appclient := operate.NewAppClient(conn)
	app.operateAppClient = appclient
}

func connOperateUserClient(app *CmsAPP) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"}, //从etcd中，服务的发现
	})
	if err != nil {
		panic(err)
	}
	// new dis with etcd client
	dis := etcd.New(client)

	//endpoint := "discovery:///provider"
	endpoint := "discovery:///user_manage" //把etcd的Name标识符拿过来，找到对应的服务ip
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		//上面127.0.0.1:9000写死的，后面引入etcd服务发现。把服务注册到etcd中,端口自动分配
		//我们的服务，现在微服务化了，有很多的分布式节点，需要etcd帮我们存储现在所有的分布式节点
		//实现负载均衡的能力，指定请求的节点，每个请求路由到不同的机器节点上
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(dis), //服务的发现
	)
	if err != nil {
		panic(err)
	}
	appclient := operate.NewUserClient(conn)
	app.operateUserClient = appclient
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

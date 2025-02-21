package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
	"user_manage/api/operate"
)

func main() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := operate.NewUserClient(conn)

	//增
	//reply, err := client.CreateUser(context.Background(), &operate.CreateUserRequest{
	//	User: &operate.UserInfo{
	//		Userid:   "1",
	//		Password: "1",
	//		Nickname: "1",
	//	},
	//})
	//改
	//reply, err := client.UpdateUser(context.Background(), &operate.UpdateUserRequest{
	//	User: &operate.UserInfo{
	//		Id:       1,
	//		Userid:   "123456",
	//		Password: "123456789",
	//		Nickname: "user_lyx1",
	//	},
	//})
	//删
	//reply, err := client.DeleteUser(context.Background(), &operate.DeleteUserRequest{
	//	Id: 2,
	//})
	//查
	reply, err := client.GetUser(context.Background(), &operate.GetUserRequest{
		//Id:       1,
		Page:     1,
		PageSize: 2,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] user %+v\n", reply)
}

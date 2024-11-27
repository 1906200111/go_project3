package main

import (
	"content_manage/api/operate"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"log"
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
	client := operate.NewAppClient(conn)
	//增
	//reply, err := client.CreateContent(context.Background(), &operate.CreateContentReq{
	//	Content: &operate.Content{
	//		Title:       "test content_manage create",
	//		VideoUrl:    "https://example.com/video.mp4",
	//		Author:      "lucky",
	//		Description: "test create",
	//	},
	//})
	//改
	//reply, err := client.UpdateContent(context.Background(), &operate.UpdateContentReq{
	//	Content: &operate.Content{
	//		Id:          9,
	//		Title:       "test content_manage create",
	//		VideoUrl:    "https://example.com/video.mp4",
	//		Author:      "lucky",
	//		Description: "test update",
	//	},
	//})
	//删
	//reply, err := client.DeleteContent(context.Background(), &operate.DeleteContentReq{
	//	Id: 11,
	//})
	//查
	reply, err := client.FindContent(context.Background(), &operate.FindContentReq{
		//Id:       7,
		Page:     2,
		PageSize: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] Content %+v\n", reply)
}

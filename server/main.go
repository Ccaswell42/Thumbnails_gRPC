package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"thumb/cache"
	"thumb/proto"
	"time"
)

type Object struct {
	proto.UnimplementedGetPictureServer
	*cache.Redis
}

func (o *Object) GetThumbnail(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	thumb := &proto.Response{Pic: nil}
	pic, err := o.Redis.Get(req.Id)
	if err == nil {
		log.Println("get pic from cache")
		thumb.Pic = pic
		return thumb, nil
	} else {
		log.Println("no thumbnails in cache", err)
	}
	httpClient := &http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := httpClient.Get(req.Url)
	if err != nil {
		return thumb, errors.New("get request's error")
	}
	defer res.Body.Close()
	respBody, err := io.ReadAll(res.Body)
	thumb.Pic = respBody
	err = o.Redis.Set(thumb.Pic, req.Id)
	if err != nil {
		log.Println("can't save pic in DB", err)
	}
	log.Println(req.Id, "successfully download")
	return thumb, nil
}

func main() {
	var obj Object
	obj.Redis = cache.NewRedisCLi()
	fmt.Println(obj.Rdb.Ping(context.Background()))
	lis, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatalln("cant listen port", err)
	}

	server := grpc.NewServer()
	proto.RegisterGetPictureServer(server, &obj)
	fmt.Println("starting server at :8081")
	go server.Serve(lis)

	time.Sleep(3 * time.Second)
	//graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	server.GracefulStop()
	log.Println("Shutdown completed successfully")

}

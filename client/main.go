package main

import (
	"bufio"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
	"sync"
	"thumb/proto"
)

// https://img.youtube.com/vi/<insert-youtube-video-id-here>/hqdefault.jpg

const (
	YouTubeLinkTemplate string = "https://www.youtube.com/watch?v="
	BeforeDownloadUrl   string = "https://img.youtube.com/vi/"
	AfterDownloadUrl    string = "/hqdefault.jpg"
)

func parser(links []string) []string {
	var id []string
	for _, val := range links {
		_, after, ok := strings.Cut(val, YouTubeLinkTemplate)
		if !ok {
			continue
		}
		before, _, ok := strings.Cut(after, "&")
		if !ok {
			id = append(id, after)
		} else {
			id = append(id, before)
		}
	}
	return id
}

func pictures(id []string) []proto.Request {
	var pictures []proto.Request
	for _, val := range id {
		image := BeforeDownloadUrl + val + AfterDownloadUrl
		req := proto.Request{Id: val, Url: image}
		pictures = append(pictures, req)
	}
	return pictures
}

func grpcClient(req proto.Request, wg *sync.WaitGroup) {
	grcpConn, err := grpc.Dial(
		"0.0.0.0:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("can't connect to grpc")
	}
	defer grcpConn.Close()
	c := proto.NewGetPictureClient(grcpConn)
	ctx := context.Background()
	respBody, err := c.GetThumbnail(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	fileName := req.Id + ".jpg"
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write(respBody.Pic)
	log.Println(req.Id, "successfully download")
	wg.Done()
}

func main() {
	links := strings.Split(os.Args[1], " ")
	wg := &sync.WaitGroup{}
	wg.Add(len(links))
	id := parser(links)
	reqs := pictures(id)

	err := os.Chdir("Downloads")
	if err != nil {
		log.Println("can't change dir")
	}
	for _, val := range reqs {
		go grpcClient(val, wg)
	}
	wg.Wait()
}

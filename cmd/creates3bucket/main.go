package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"time"
)

const endpoint = "http://localhost:4566/"

func main() {
	log.Println("start")
	defer func(){
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("exit")
	}()

	endpointResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           endpoint,
			SigningRegion: region,
			HostnameImmutable: true,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithEndpointResolver(endpointResolver),
	)
	if err != nil {
		panic(fmt.Errorf("config.LoadDefaultConfig failed:%v", err))
	}

	client := s3.NewFromConfig(cfg)
	createBucketInput := &s3.CreateBucketInput{
		Bucket: aws.String(fmt.Sprintf("test-%s", time.Now().Format("20060102-150405"))),
	}
	createBucketOutput, err := client.CreateBucket(context.TODO(), createBucketInput)
	if err != nil {
		panic(fmt.Errorf("s3.CreateBucket failed:%v", err))
	}
	buf, _ := json.MarshalIndent(&createBucketOutput, "", "    ")
	log.Println(string(buf))

	listBucketOutput, err := client.ListBuckets(context.TODO(), nil)
	if err != nil {
		panic(fmt.Errorf("s3.ListBuckets failed:%v", err))
	}
	buf, _ = json.MarshalIndent(&listBucketOutput, "", "    ")
	log.Println(string(buf))
}

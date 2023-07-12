package utils

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var s3Session = session.New()

func UtilAWS() *session.Session {

	region := os.Getenv("REGION")
	if region == "" {
		log.Println("region vazia")
	}
	key := os.Getenv("KEY")
	if key == "" {
		log.Println("key vazia")
	}
	secretpass := os.Getenv("SECRETPASS")
	if secretpass == "" {
		log.Println("secretpass vazia")
	}
	s3Config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secretpass, ""),
	}

	s3Session = session.New(s3Config)
	return session.New(s3Config.WithRegion(region))

}

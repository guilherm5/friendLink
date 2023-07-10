package utils

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
)

var s3Session = session.New()

func UtilAWS() *session.Session {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	region := os.Getenv("REGION")
	key := os.Getenv("KEY")
	secretpass := os.Getenv("SECRETPASS")
	s3Config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secretpass, ""),
	}

	s3Session = session.New(s3Config)
	return session.New(s3Config.WithRegion("us-east-1"))

}

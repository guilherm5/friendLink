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
	//Variaveis de login para o aws
	region := os.Getenv("REGION")
	key := os.Getenv("KEY")
	secretpass := os.Getenv("SECRETPASS")
	if region == "" || key == "" || secretpass == "" {
		log.Println("Variavel de ambiente não pode ser null")
	}

	s3Config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(key, secretpass, ""),
	}

	s3Session = session.New(s3Config)
	return session.New(s3Config.WithRegion(region))
}

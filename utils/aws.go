package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var s3Session = session.New()

func UtilAWS() *session.Session {
	s3Config := &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("AKIAVBCXT23G42GKHFC2", "EneitgstlP+4Ue5kMq3gCIvCtV2g//nPuJf6SeXC", ""),
	}

	s3Session = session.New(s3Config)
	return session.New(s3Config.WithRegion("us-east-1"))

}

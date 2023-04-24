package db

import(
 "github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Connectdb()(*dynamodb.DynamoDB, error){
		// Creating session for client
sess := session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

// Create DynamoDB client
svc := dynamodb.New(sess)

return svc, nil
}
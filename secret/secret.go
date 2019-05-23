package secret

import (
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"encoding/json"
	"fmt"
)

// parameters: 
// access config: 
//     region: for example, "us-east-2"
//     access_key_id:
//     secret_access_key: 
//     ASK YOUR MANAGER IF YOU DON'T KNOW THE ABOVE 3 VALUES
// secret config: 
//     secret_name, predefined secret name in AWS, for example "test/SongPrivKeyERC" 
//     key_name, predefined key name in AWS, for example "my_private_key_1" 
//     ASK YOUR MANAGER IF YOU DON'T KNOW THE ABOVE TWO VALUES
func GetSecret(region, access_key_id, secret_access_key, secret_name, key_name string) (string, error) {
	var cred = credentials.NewStaticCredentials(access_key_id, secret_access_key, "")

	svc := secretsmanager.New(session.New(&aws.Config{
                Credentials: cred}), aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secret_name),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
				case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

				case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

				case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

				case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

				case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())

				case cloudwatchlogs.ErrCodeUnrecognizedClientException:
				// UnrecognizedClientException: The security token included in the request is invalid.
				fmt.Println(cloudwatchlogs.ErrCodeUnrecognizedClientException, aerr.Error())

				default:
				fmt.Println(aerr.Code(), ",", aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return "", err
	}

	// Decrypts secret using the associated KMS CMK.
	// in our case, it will be string.
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	if secretString != "" {
		var dat map[string]interface{}
		if err := json.Unmarshal([]byte(secretString), &dat); err != nil {
			panic(err)
		}

		secretString = dat[key_name].(string)
	}

	return secretString, nil
}

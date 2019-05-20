package secret_test

import (
	"testing"
	_"fmt"

	secretmanager "github.com/Ankr-network/dccn-common/secret"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

const (
	MAINNET_PRIVATE_KEY = "MAINNET_PRIVATE_KEY"
	MAINNET_ADDRESS = "MAINNET_ADDRESS"
	BEP_PRIVATE_KEY = "BEP_PRIVATE_KEY"
	ERC_FROM_KEY = "ERC_FROM_KEY"
	ERC_FROM_PASSWORD = "ERC_FROM_PASSWORD"
)

func TestGetSecret(t *testing.T) {
        _, err := secretmanager.GetSecret("us-east-2", // region
		"XXXXXXXXXXXXXXXXXXXX", // access key id
		"AaaaAaaAaaaAaaaaaAaAaaaaAAaAAAAAaAAaaAAA", //secret access key
		"test/SongPrivKeyERC", // secret name
		"ERC_FROM_PASSWORD") // key name of the key/value for our private key

        if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			if aerr.Code() == cloudwatchlogs.ErrCodeUnrecognizedClientException {
				t.Logf("expected failure.\n")
				return
			}
		}
		t.Error("unexpected failure.\n")
        } else {
		t.Error(err)
        }
}


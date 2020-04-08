package main

import (
	"context"
	"fmt"
	"os"

	"github.com/emilaasa/ip2asn"

	"github.com/aws/aws-lambda-go/lambda"
)

type s3Event struct {
	Records []struct {
		S3 struct {
			Bucket struct {
				Name string `json:"name"`
				Arn  string `json:"arn"`
			} `json:"bucket"`
			Object struct {
				Key       string `json:"key"`
				Size      string `json:"size"`
				ETag      string `json:"eTag"`
				VersionID string `json:"versionId"`
				Sequencer string `json:"sequencer"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

var asnDB ip2asn.AsnDB

func init() {
	f, _ := os.Open("IPASN.DAT")
	asnDB := ip2asn.NewLookuperFromFile(f)
}

// HandleRequest is called once for every event a lambda func recieves
func HandleRequest(ctx context.Context, ev s3Event) (string, error) {
	s3Event := ev.Records[0].S3
	// s3.read
	// gzip.read
	// ip2asn.read
	/// s3.write
	return fmt.Sprintf("Hello %s!", s3Event), nil
}

func s3() {
	//do the s3 thing
}

func ulf() {
	asnDB.Lookup()
}

func main() {
	lambda.Start(HandleRequest)
}

// S3 Event for reference and debugging, delete when it works
// Records []struct {
// 	EventVersion string    `json:"eventVersion"`
// 	EventSource  string    `json:"eventSource"`
// 	AwsRegion    string    `json:"awsRegion"`
// 	EventTime    time.Time `json:"eventTime"`
// 	EventName    string    `json:"eventName"`
// 	UserIdentity struct {
// 		PrincipalID string `json:"principalId"`
// 	} `json:"userIdentity"`
// 	RequestParameters struct {
// 		SourceIPAddress string `json:"sourceIPAddress"`
// 	} `json:"requestParameters"`
// 	ResponseElements struct {
// 		XAmzRequestID string `json:"x-amz-request-id"`
// 		XAmzID2       string `json:"x-amz-id-2"`
// 	} `json:"responseElements"`
// 	S3 struct {
// 		S3SchemaVersion string `json:"s3SchemaVersion"`
// 		ConfigurationID string `json:"configurationId"`
// 		Bucket          struct {
// 			Name          string `json:"name"`
// 			OwnerIdentity struct {
// 				PrincipalID string `json:"principalId"`
// 			} `json:"ownerIdentity"`
// 			Arn string `json:"arn"`
// 		} `json:"bucket"`
// 		Object struct {
// 			Key       string `json:"key"`
// 			Size      string `json:"size"`
// 			ETag      string `json:"eTag"`
// 			VersionID string `json:"versionId"`
// 			Sequencer string `json:"sequencer"`
// 		} `json:"object"`
// 	} `json:"s3"`
// 	GlacierEventData struct {
// 		RestoreEventData struct {
// 			LifecycleRestorationExpiryTime time.Time `json:"lifecycleRestorationExpiryTime"`
// 			LifecycleRestoreStorageClass   string    `json:"lifecycleRestoreStorageClass"`
// 		} `json:"restoreEventData"`
// 	} `json:"glacierEventData"`
// } `json:"Records"`

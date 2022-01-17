package dynamodbkinsumer

import (
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// StreamRecord is unmarshalled dynamodbstreams.StreamRecord
type StreamRecord struct {
	ApproximateCreationDateTime *time.Time
	Keys                        map[string]*dynamodb.AttributeValue
	NewImage                    map[string]*dynamodb.AttributeValue
	OldImage                    map[string]*dynamodb.AttributeValue
	SequenceNumber              *string
	SizeBytes                   *int64
	StreamViewType              *string
}

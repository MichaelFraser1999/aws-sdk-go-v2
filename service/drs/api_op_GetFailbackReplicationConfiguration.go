// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all Failback ReplicationConfigurations, filtered by Recovery Instance ID.
func (c *Client) GetFailbackReplicationConfiguration(ctx context.Context, params *GetFailbackReplicationConfigurationInput, optFns ...func(*Options)) (*GetFailbackReplicationConfigurationOutput, error) {
	if params == nil {
		params = &GetFailbackReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFailbackReplicationConfiguration", params, optFns, c.addOperationGetFailbackReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFailbackReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFailbackReplicationConfigurationInput struct {

	// The ID of the Recovery Instance whose failback replication configuration should
	// be returned.
	//
	// This member is required.
	RecoveryInstanceID *string

	noSmithyDocumentSerde
}

type GetFailbackReplicationConfigurationOutput struct {

	// The ID of the Recovery Instance.
	//
	// This member is required.
	RecoveryInstanceID *string

	// Configure bandwidth throttling for the outbound data transfer rate of the
	// Recovery Instance in Mbps.
	BandwidthThrottling int64

	// The name of the Failback Replication Configuration.
	Name *string

	// Whether to use Private IP for the failback replication of the Recovery Instance.
	UsePrivateIP *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFailbackReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFailbackReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFailbackReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetFailbackReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFailbackReplicationConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetFailbackReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "GetFailbackReplicationConfiguration",
	}
}
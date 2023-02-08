// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// When you delete Amazon Security Lake from your account, Security Lake is
// disabled in all Amazon Web Services Regions. Also, this API automatically takes
// steps to remove the account from Security Lake . This operation disables
// security data collection from sources, deletes data stored, and stops making
// data accessible to subscribers. Security Lake also deletes all the existing
// settings and resources that it stores or maintains for your Amazon Web Services
// account in the current Region, including security log and event data. The
// DeleteDatalake operation does not delete the Amazon S3 bucket, which is owned by
// your Amazon Web Services account. For more information, see the Amazon Security
// Lake User Guide
// (https://docs.aws.amazon.com/security-lake/latest/userguide/disable-security-lake.html).
func (c *Client) DeleteDatalake(ctx context.Context, params *DeleteDatalakeInput, optFns ...func(*Options)) (*DeleteDatalakeOutput, error) {
	if params == nil {
		params = &DeleteDatalakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDatalake", params, optFns, c.addOperationDeleteDatalakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDatalakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDatalakeInput struct {
	noSmithyDocumentSerde
}

type DeleteDatalakeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDatalakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDatalake{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDatalake{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDatalake(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDatalake(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "DeleteDatalake",
	}
}
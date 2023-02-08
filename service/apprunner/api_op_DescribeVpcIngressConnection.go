// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Return a full description of an App Runner VPC Ingress Connection resource.
func (c *Client) DescribeVpcIngressConnection(ctx context.Context, params *DescribeVpcIngressConnectionInput, optFns ...func(*Options)) (*DescribeVpcIngressConnectionOutput, error) {
	if params == nil {
		params = &DescribeVpcIngressConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcIngressConnection", params, optFns, c.addOperationDescribeVpcIngressConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcIngressConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcIngressConnectionInput struct {

	// The Amazon Resource Name (ARN) of the App Runner VPC Ingress Connection that you
	// want a description for.
	//
	// This member is required.
	VpcIngressConnectionArn *string

	noSmithyDocumentSerde
}

type DescribeVpcIngressConnectionOutput struct {

	// A description of the App Runner VPC Ingress Connection that you specified in
	// this request.
	//
	// This member is required.
	VpcIngressConnection *types.VpcIngressConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcIngressConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeVpcIngressConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeVpcIngressConnection{}, middleware.After)
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
	if err = addOpDescribeVpcIngressConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcIngressConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcIngressConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "DescribeVpcIngressConnection",
	}
}
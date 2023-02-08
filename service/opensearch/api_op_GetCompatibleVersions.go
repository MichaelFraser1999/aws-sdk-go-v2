// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a map of OpenSearch or Elasticsearch versions and the versions you can
// upgrade them to.
func (c *Client) GetCompatibleVersions(ctx context.Context, params *GetCompatibleVersionsInput, optFns ...func(*Options)) (*GetCompatibleVersionsOutput, error) {
	if params == nil {
		params = &GetCompatibleVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCompatibleVersions", params, optFns, c.addOperationGetCompatibleVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCompatibleVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the request parameters to GetCompatibleVersions operation.
type GetCompatibleVersionsInput struct {

	// The name of an existing domain. Provide this parameter to limit the results to a
	// single domain.
	DomainName *string

	noSmithyDocumentSerde
}

// Container for the response returned by the GetCompatibleVersions operation.
type GetCompatibleVersionsOutput struct {

	// A map of OpenSearch or Elasticsearch versions and the versions you can upgrade
	// them to.
	CompatibleVersions []types.CompatibleVersionsMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCompatibleVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCompatibleVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCompatibleVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCompatibleVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCompatibleVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "GetCompatibleVersions",
	}
}
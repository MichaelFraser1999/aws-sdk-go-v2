// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initializes an Amazon Security Lake instance with the provided (or default)
// configuration. You can enable Security Lake in Amazon Web Services Regions with
// customized settings before enabling log collection in Regions. You can either
// use the enableAll parameter to specify all Regions or specify the Regions where
// you want to enable Security Lake. To specify particular Regions, use the Regions
// parameter and then configure these Regions using the configurations parameter.
// If you have already enabled Security Lake in a Region when you call this
// command, the command will update the Region if you provide new configuration
// parameters. If you have not already enabled Security Lake in the Region when you
// call this API, it will set up the data lake in the Region with the specified
// configurations. When you enable Security Lake, it starts ingesting security data
// after the CreateAwsLogSource call. This includes ingesting security data from
// sources, storing data, and making data accessible to subscribers. Security Lake
// also enables all the existing settings and resources that it stores or maintains
// for your Amazon Web Services account in the current Region, including security
// log and event data. For more information, see the Amazon Security Lake User
// Guide
// (https://docs.aws.amazon.com/security-lake/latest/userguide/what-is-security-lake.html).
func (c *Client) CreateDatalake(ctx context.Context, params *CreateDatalakeInput, optFns ...func(*Options)) (*CreateDatalakeOutput, error) {
	if params == nil {
		params = &CreateDatalakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatalake", params, optFns, c.addOperationCreateDatalakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatalakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatalakeInput struct {

	// Specify the Region or Regions that will contribute data to the rollup region.
	Configurations map[string]types.LakeConfigurationRequest

	// Enable Security Lake in all Regions.
	EnableAll *bool

	// The Amazon Resource Name (ARN) used to create and update the Glue table. This
	// table contains partitions generated by the ingestion and normalization of Amazon
	// Web Services log sources and custom sources.
	MetaStoreManagerRoleArn *string

	// Enable Security Lake in the specified Regions. To enable Security Lake in
	// specific Amazon Web Services Regions, such as us-east-1 or ap-northeast-3,
	// provide the Region codes. For a list of Region codes, see Amazon Security Lake
	// endpoints (https://docs.aws.amazon.com/general/latest/gr/securitylake.html) in
	// the Amazon Web Services General Reference.
	Regions []types.Region

	noSmithyDocumentSerde
}

type CreateDatalakeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatalakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDatalake{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDatalake{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatalake(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatalake(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateDatalake",
	}
}
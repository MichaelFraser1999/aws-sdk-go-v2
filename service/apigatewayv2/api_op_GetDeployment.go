// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDeploymentRequest
type GetDeploymentInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// DeploymentId is a required field
	DeploymentId *string `location:"uri" locationName:"deploymentId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDeploymentInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.DeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApiId != nil {
		v := *s.ApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "apiId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "deploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDeploymentResponse
type GetDeploymentOutput struct {
	_ struct{} `type:"structure"`

	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp" timestampFormat:"unix"`

	// The identifier.
	DeploymentId *string `locationName:"deploymentId" type:"string"`

	// Represents a deployment status.
	DeploymentStatus DeploymentStatus `locationName:"deploymentStatus" type:"string" enum:"true"`

	DeploymentStatusMessage *string `locationName:"deploymentStatusMessage" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`
}

// String returns the string representation
func (s GetDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreatedDate != nil {
		v := *s.CreatedDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "createdDate", protocol.TimeValue{V: v, Format: protocol.UnixTimeFormat}, metadata)
	}
	if s.DeploymentId != nil {
		v := *s.DeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.DeploymentStatus) > 0 {
		v := s.DeploymentStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.DeploymentStatusMessage != nil {
		v := *s.DeploymentStatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "deploymentStatusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetDeployment = "GetDeployment"

// GetDeploymentRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Gets a Deployment.
//
//    // Example sending a request using GetDeploymentRequest.
//    req := client.GetDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/GetDeployment
func (c *Client) GetDeploymentRequest(input *GetDeploymentInput) GetDeploymentRequest {
	op := &aws.Operation{
		Name:       opGetDeployment,
		HTTPMethod: "GET",
		HTTPPath:   "/v2/apis/{apiId}/deployments/{deploymentId}",
	}

	if input == nil {
		input = &GetDeploymentInput{}
	}

	req := c.newRequest(op, input, &GetDeploymentOutput{})
	return GetDeploymentRequest{Request: req, Input: input, Copy: c.GetDeploymentRequest}
}

// GetDeploymentRequest is the request type for the
// GetDeployment API operation.
type GetDeploymentRequest struct {
	*aws.Request
	Input *GetDeploymentInput
	Copy  func(*GetDeploymentInput) GetDeploymentRequest
}

// Send marshals and sends the GetDeployment API request.
func (r GetDeploymentRequest) Send(ctx context.Context) (*GetDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDeploymentResponse{
		GetDeploymentOutput: r.Request.Data.(*GetDeploymentOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDeploymentResponse is the response type for the
// GetDeployment API operation.
type GetDeploymentResponse struct {
	*GetDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDeployment request.
func (r *GetDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
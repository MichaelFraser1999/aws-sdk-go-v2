// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateWorkforceInput struct {
	_ struct{} `type:"structure"`

	// A list of one to four worker IP address ranges (CIDRs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html))
	// that can be used to access tasks assigned to this workforce.
	//
	// Maximum: Four CIDR values
	SourceIpConfig *SourceIpConfig `type:"structure"`

	// The name of the private workforce whose access you want to restrict. WorkforceName
	// is automatically set to default when a workforce is created and cannot be
	// modified.
	//
	// WorkforceName is a required field
	WorkforceName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateWorkforceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWorkforceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateWorkforceInput"}

	if s.WorkforceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("WorkforceName"))
	}
	if s.WorkforceName != nil && len(*s.WorkforceName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkforceName", 1))
	}
	if s.SourceIpConfig != nil {
		if err := s.SourceIpConfig.Validate(); err != nil {
			invalidParams.AddNested("SourceIpConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateWorkforceOutput struct {
	_ struct{} `type:"structure"`

	// A single private workforce, which is automatically created when you create
	// your first private work team. You can create one private work force in each
	// AWS Region. By default, any workforce-related API operation used in a specific
	// region will apply to the workforce created in that region. To learn how to
	// create a private workforce, see Create a Private Workforce (https://docs.aws.amazon.com/sagemaker/latest/dg/sms-workforce-create-private.html).
	//
	// Workforce is a required field
	Workforce *Workforce `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateWorkforceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateWorkforce = "UpdateWorkforce"

// UpdateWorkforceRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Restricts access to tasks assigned to workers in the specified workforce
// to those within specific ranges of IP addresses. You specify allowed IP addresses
// by creating a list of up to four CIDRs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html).
//
// By default, a workforce isn't restricted to specific IP addresses. If you
// specify a range of IP addresses, workers who attempt to access tasks using
// any IP address outside the specified range are denied access and get a Not
// Found error message on the worker portal. After restricting access with this
// operation, you can see the allowed IP values for a private workforce with
// the operation.
//
// This operation applies only to private workforces.
//
//    // Example sending a request using UpdateWorkforceRequest.
//    req := client.UpdateWorkforceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateWorkforce
func (c *Client) UpdateWorkforceRequest(input *UpdateWorkforceInput) UpdateWorkforceRequest {
	op := &aws.Operation{
		Name:       opUpdateWorkforce,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateWorkforceInput{}
	}

	req := c.newRequest(op, input, &UpdateWorkforceOutput{})

	return UpdateWorkforceRequest{Request: req, Input: input, Copy: c.UpdateWorkforceRequest}
}

// UpdateWorkforceRequest is the request type for the
// UpdateWorkforce API operation.
type UpdateWorkforceRequest struct {
	*aws.Request
	Input *UpdateWorkforceInput
	Copy  func(*UpdateWorkforceInput) UpdateWorkforceRequest
}

// Send marshals and sends the UpdateWorkforce API request.
func (r UpdateWorkforceRequest) Send(ctx context.Context) (*UpdateWorkforceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateWorkforceResponse{
		UpdateWorkforceOutput: r.Request.Data.(*UpdateWorkforceOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateWorkforceResponse is the response type for the
// UpdateWorkforce API operation.
type UpdateWorkforceResponse struct {
	*UpdateWorkforceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateWorkforce request.
func (r *UpdateWorkforceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
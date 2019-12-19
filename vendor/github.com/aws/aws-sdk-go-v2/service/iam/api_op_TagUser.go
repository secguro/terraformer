// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type TagUserInput struct {
	_ struct{} `type:"structure"`

	// The list of tags that you want to attach to the user. Each tag consists of
	// a key name and an associated value.
	//
	// Tags is a required field
	Tags []Tag `type:"list" required:"true"`

	// The name of the user that you want to add tags to.
	//
	// This parameter accepts (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that consist of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: =,.@-
	//
	// UserName is a required field
	UserName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TagUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TagUserInput"}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TagUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagUserOutput) String() string {
	return awsutil.Prettify(s)
}

const opTagUser = "TagUser"

// TagUserRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Adds one or more tags to an IAM user. If a tag with the same key name already
// exists, then that tag is overwritten with the new value.
//
// A tag consists of a key name and an associated value. By assigning tags to
// your resources, you can do the following:
//
//    * Administrative grouping and discovery - Attach tags to resources to
//    aid in organization and search. For example, you could search for all
//    resources with the key name Project and the value MyImportantProject.
//    Or search for all resources with the key name Cost Center and the value
//    41200.
//
//    * Access control - Reference tags in IAM user-based and resource-based
//    policies. You can use tags to restrict access to only an IAM requesting
//    user or to a role that has a specified tag attached. You can also restrict
//    access to only those resources that have a certain tag attached. For examples
//    of policies that show how to use tags to control access, see Control Access
//    Using IAM Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html)
//    in the IAM User Guide.
//
//    * Cost allocation - Use tags to help track which individuals and teams
//    are using which AWS resources.
//
//    * Make sure that you have no invalid tags and that you do not exceed the
//    allowed number of tags per role. In either case, the entire request fails
//    and no tags are added to the role.
//
//    * AWS always interprets the tag Value as a single string. If you need
//    to store an array, you can store comma-separated values in the string.
//    However, you must interpret the value in your code.
//
// For more information about tagging, see Tagging IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
// in the IAM User Guide.
//
//    // Example sending a request using TagUserRequest.
//    req := client.TagUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/TagUser
func (c *Client) TagUserRequest(input *TagUserInput) TagUserRequest {
	op := &aws.Operation{
		Name:       opTagUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagUserInput{}
	}

	req := c.newRequest(op, input, &TagUserOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return TagUserRequest{Request: req, Input: input, Copy: c.TagUserRequest}
}

// TagUserRequest is the request type for the
// TagUser API operation.
type TagUserRequest struct {
	*aws.Request
	Input *TagUserInput
	Copy  func(*TagUserInput) TagUserRequest
}

// Send marshals and sends the TagUser API request.
func (r TagUserRequest) Send(ctx context.Context) (*TagUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagUserResponse{
		TagUserOutput: r.Request.Data.(*TagUserOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagUserResponse is the response type for the
// TagUser API operation.
type TagUserResponse struct {
	*TagUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagUser request.
func (r *TagUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
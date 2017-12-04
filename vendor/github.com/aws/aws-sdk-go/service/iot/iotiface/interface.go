// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotiface provides an interface to enable mocking the AWS IoT service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iot"
)

// IoTAPI provides an interface to enable mocking the
// iot.IoT service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT.
//    func myFunc(svc iotiface.IoTAPI) bool {
//        // Make svc.AcceptCertificateTransfer request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iot.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTClient struct {
//        iotiface.IoTAPI
//    }
//    func (m *mockIoTClient) AcceptCertificateTransfer(input *iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IoTAPI interface {
	AcceptCertificateTransfer(*iot.AcceptCertificateTransferInput) (*iot.AcceptCertificateTransferOutput, error)
	AcceptCertificateTransferWithContext(aws.Context, *iot.AcceptCertificateTransferInput, ...request.Option) (*iot.AcceptCertificateTransferOutput, error)
	AcceptCertificateTransferRequest(*iot.AcceptCertificateTransferInput) (*request.Request, *iot.AcceptCertificateTransferOutput)

	AddThingToThingGroup(*iot.AddThingToThingGroupInput) (*iot.AddThingToThingGroupOutput, error)
	AddThingToThingGroupWithContext(aws.Context, *iot.AddThingToThingGroupInput, ...request.Option) (*iot.AddThingToThingGroupOutput, error)
	AddThingToThingGroupRequest(*iot.AddThingToThingGroupInput) (*request.Request, *iot.AddThingToThingGroupOutput)

	AssociateTargetsWithJob(*iot.AssociateTargetsWithJobInput) (*iot.AssociateTargetsWithJobOutput, error)
	AssociateTargetsWithJobWithContext(aws.Context, *iot.AssociateTargetsWithJobInput, ...request.Option) (*iot.AssociateTargetsWithJobOutput, error)
	AssociateTargetsWithJobRequest(*iot.AssociateTargetsWithJobInput) (*request.Request, *iot.AssociateTargetsWithJobOutput)

	AttachPolicy(*iot.AttachPolicyInput) (*iot.AttachPolicyOutput, error)
	AttachPolicyWithContext(aws.Context, *iot.AttachPolicyInput, ...request.Option) (*iot.AttachPolicyOutput, error)
	AttachPolicyRequest(*iot.AttachPolicyInput) (*request.Request, *iot.AttachPolicyOutput)

	AttachPrincipalPolicy(*iot.AttachPrincipalPolicyInput) (*iot.AttachPrincipalPolicyOutput, error)
	AttachPrincipalPolicyWithContext(aws.Context, *iot.AttachPrincipalPolicyInput, ...request.Option) (*iot.AttachPrincipalPolicyOutput, error)
	AttachPrincipalPolicyRequest(*iot.AttachPrincipalPolicyInput) (*request.Request, *iot.AttachPrincipalPolicyOutput)

	AttachThingPrincipal(*iot.AttachThingPrincipalInput) (*iot.AttachThingPrincipalOutput, error)
	AttachThingPrincipalWithContext(aws.Context, *iot.AttachThingPrincipalInput, ...request.Option) (*iot.AttachThingPrincipalOutput, error)
	AttachThingPrincipalRequest(*iot.AttachThingPrincipalInput) (*request.Request, *iot.AttachThingPrincipalOutput)

	CancelCertificateTransfer(*iot.CancelCertificateTransferInput) (*iot.CancelCertificateTransferOutput, error)
	CancelCertificateTransferWithContext(aws.Context, *iot.CancelCertificateTransferInput, ...request.Option) (*iot.CancelCertificateTransferOutput, error)
	CancelCertificateTransferRequest(*iot.CancelCertificateTransferInput) (*request.Request, *iot.CancelCertificateTransferOutput)

	CancelJob(*iot.CancelJobInput) (*iot.CancelJobOutput, error)
	CancelJobWithContext(aws.Context, *iot.CancelJobInput, ...request.Option) (*iot.CancelJobOutput, error)
	CancelJobRequest(*iot.CancelJobInput) (*request.Request, *iot.CancelJobOutput)

	ClearDefaultAuthorizer(*iot.ClearDefaultAuthorizerInput) (*iot.ClearDefaultAuthorizerOutput, error)
	ClearDefaultAuthorizerWithContext(aws.Context, *iot.ClearDefaultAuthorizerInput, ...request.Option) (*iot.ClearDefaultAuthorizerOutput, error)
	ClearDefaultAuthorizerRequest(*iot.ClearDefaultAuthorizerInput) (*request.Request, *iot.ClearDefaultAuthorizerOutput)

	CreateAuthorizer(*iot.CreateAuthorizerInput) (*iot.CreateAuthorizerOutput, error)
	CreateAuthorizerWithContext(aws.Context, *iot.CreateAuthorizerInput, ...request.Option) (*iot.CreateAuthorizerOutput, error)
	CreateAuthorizerRequest(*iot.CreateAuthorizerInput) (*request.Request, *iot.CreateAuthorizerOutput)

	CreateCertificateFromCsr(*iot.CreateCertificateFromCsrInput) (*iot.CreateCertificateFromCsrOutput, error)
	CreateCertificateFromCsrWithContext(aws.Context, *iot.CreateCertificateFromCsrInput, ...request.Option) (*iot.CreateCertificateFromCsrOutput, error)
	CreateCertificateFromCsrRequest(*iot.CreateCertificateFromCsrInput) (*request.Request, *iot.CreateCertificateFromCsrOutput)

	CreateJob(*iot.CreateJobInput) (*iot.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *iot.CreateJobInput, ...request.Option) (*iot.CreateJobOutput, error)
	CreateJobRequest(*iot.CreateJobInput) (*request.Request, *iot.CreateJobOutput)

	CreateKeysAndCertificate(*iot.CreateKeysAndCertificateInput) (*iot.CreateKeysAndCertificateOutput, error)
	CreateKeysAndCertificateWithContext(aws.Context, *iot.CreateKeysAndCertificateInput, ...request.Option) (*iot.CreateKeysAndCertificateOutput, error)
	CreateKeysAndCertificateRequest(*iot.CreateKeysAndCertificateInput) (*request.Request, *iot.CreateKeysAndCertificateOutput)

	CreatePolicy(*iot.CreatePolicyInput) (*iot.CreatePolicyOutput, error)
	CreatePolicyWithContext(aws.Context, *iot.CreatePolicyInput, ...request.Option) (*iot.CreatePolicyOutput, error)
	CreatePolicyRequest(*iot.CreatePolicyInput) (*request.Request, *iot.CreatePolicyOutput)

	CreatePolicyVersion(*iot.CreatePolicyVersionInput) (*iot.CreatePolicyVersionOutput, error)
	CreatePolicyVersionWithContext(aws.Context, *iot.CreatePolicyVersionInput, ...request.Option) (*iot.CreatePolicyVersionOutput, error)
	CreatePolicyVersionRequest(*iot.CreatePolicyVersionInput) (*request.Request, *iot.CreatePolicyVersionOutput)

	CreateRoleAlias(*iot.CreateRoleAliasInput) (*iot.CreateRoleAliasOutput, error)
	CreateRoleAliasWithContext(aws.Context, *iot.CreateRoleAliasInput, ...request.Option) (*iot.CreateRoleAliasOutput, error)
	CreateRoleAliasRequest(*iot.CreateRoleAliasInput) (*request.Request, *iot.CreateRoleAliasOutput)

	CreateThing(*iot.CreateThingInput) (*iot.CreateThingOutput, error)
	CreateThingWithContext(aws.Context, *iot.CreateThingInput, ...request.Option) (*iot.CreateThingOutput, error)
	CreateThingRequest(*iot.CreateThingInput) (*request.Request, *iot.CreateThingOutput)

	CreateThingGroup(*iot.CreateThingGroupInput) (*iot.CreateThingGroupOutput, error)
	CreateThingGroupWithContext(aws.Context, *iot.CreateThingGroupInput, ...request.Option) (*iot.CreateThingGroupOutput, error)
	CreateThingGroupRequest(*iot.CreateThingGroupInput) (*request.Request, *iot.CreateThingGroupOutput)

	CreateThingType(*iot.CreateThingTypeInput) (*iot.CreateThingTypeOutput, error)
	CreateThingTypeWithContext(aws.Context, *iot.CreateThingTypeInput, ...request.Option) (*iot.CreateThingTypeOutput, error)
	CreateThingTypeRequest(*iot.CreateThingTypeInput) (*request.Request, *iot.CreateThingTypeOutput)

	CreateTopicRule(*iot.CreateTopicRuleInput) (*iot.CreateTopicRuleOutput, error)
	CreateTopicRuleWithContext(aws.Context, *iot.CreateTopicRuleInput, ...request.Option) (*iot.CreateTopicRuleOutput, error)
	CreateTopicRuleRequest(*iot.CreateTopicRuleInput) (*request.Request, *iot.CreateTopicRuleOutput)

	DeleteAuthorizer(*iot.DeleteAuthorizerInput) (*iot.DeleteAuthorizerOutput, error)
	DeleteAuthorizerWithContext(aws.Context, *iot.DeleteAuthorizerInput, ...request.Option) (*iot.DeleteAuthorizerOutput, error)
	DeleteAuthorizerRequest(*iot.DeleteAuthorizerInput) (*request.Request, *iot.DeleteAuthorizerOutput)

	DeleteCACertificate(*iot.DeleteCACertificateInput) (*iot.DeleteCACertificateOutput, error)
	DeleteCACertificateWithContext(aws.Context, *iot.DeleteCACertificateInput, ...request.Option) (*iot.DeleteCACertificateOutput, error)
	DeleteCACertificateRequest(*iot.DeleteCACertificateInput) (*request.Request, *iot.DeleteCACertificateOutput)

	DeleteCertificate(*iot.DeleteCertificateInput) (*iot.DeleteCertificateOutput, error)
	DeleteCertificateWithContext(aws.Context, *iot.DeleteCertificateInput, ...request.Option) (*iot.DeleteCertificateOutput, error)
	DeleteCertificateRequest(*iot.DeleteCertificateInput) (*request.Request, *iot.DeleteCertificateOutput)

	DeletePolicy(*iot.DeletePolicyInput) (*iot.DeletePolicyOutput, error)
	DeletePolicyWithContext(aws.Context, *iot.DeletePolicyInput, ...request.Option) (*iot.DeletePolicyOutput, error)
	DeletePolicyRequest(*iot.DeletePolicyInput) (*request.Request, *iot.DeletePolicyOutput)

	DeletePolicyVersion(*iot.DeletePolicyVersionInput) (*iot.DeletePolicyVersionOutput, error)
	DeletePolicyVersionWithContext(aws.Context, *iot.DeletePolicyVersionInput, ...request.Option) (*iot.DeletePolicyVersionOutput, error)
	DeletePolicyVersionRequest(*iot.DeletePolicyVersionInput) (*request.Request, *iot.DeletePolicyVersionOutput)

	DeleteRegistrationCode(*iot.DeleteRegistrationCodeInput) (*iot.DeleteRegistrationCodeOutput, error)
	DeleteRegistrationCodeWithContext(aws.Context, *iot.DeleteRegistrationCodeInput, ...request.Option) (*iot.DeleteRegistrationCodeOutput, error)
	DeleteRegistrationCodeRequest(*iot.DeleteRegistrationCodeInput) (*request.Request, *iot.DeleteRegistrationCodeOutput)

	DeleteRoleAlias(*iot.DeleteRoleAliasInput) (*iot.DeleteRoleAliasOutput, error)
	DeleteRoleAliasWithContext(aws.Context, *iot.DeleteRoleAliasInput, ...request.Option) (*iot.DeleteRoleAliasOutput, error)
	DeleteRoleAliasRequest(*iot.DeleteRoleAliasInput) (*request.Request, *iot.DeleteRoleAliasOutput)

	DeleteThing(*iot.DeleteThingInput) (*iot.DeleteThingOutput, error)
	DeleteThingWithContext(aws.Context, *iot.DeleteThingInput, ...request.Option) (*iot.DeleteThingOutput, error)
	DeleteThingRequest(*iot.DeleteThingInput) (*request.Request, *iot.DeleteThingOutput)

	DeleteThingGroup(*iot.DeleteThingGroupInput) (*iot.DeleteThingGroupOutput, error)
	DeleteThingGroupWithContext(aws.Context, *iot.DeleteThingGroupInput, ...request.Option) (*iot.DeleteThingGroupOutput, error)
	DeleteThingGroupRequest(*iot.DeleteThingGroupInput) (*request.Request, *iot.DeleteThingGroupOutput)

	DeleteThingType(*iot.DeleteThingTypeInput) (*iot.DeleteThingTypeOutput, error)
	DeleteThingTypeWithContext(aws.Context, *iot.DeleteThingTypeInput, ...request.Option) (*iot.DeleteThingTypeOutput, error)
	DeleteThingTypeRequest(*iot.DeleteThingTypeInput) (*request.Request, *iot.DeleteThingTypeOutput)

	DeleteTopicRule(*iot.DeleteTopicRuleInput) (*iot.DeleteTopicRuleOutput, error)
	DeleteTopicRuleWithContext(aws.Context, *iot.DeleteTopicRuleInput, ...request.Option) (*iot.DeleteTopicRuleOutput, error)
	DeleteTopicRuleRequest(*iot.DeleteTopicRuleInput) (*request.Request, *iot.DeleteTopicRuleOutput)

	DeleteV2LoggingLevel(*iot.DeleteV2LoggingLevelInput) (*iot.DeleteV2LoggingLevelOutput, error)
	DeleteV2LoggingLevelWithContext(aws.Context, *iot.DeleteV2LoggingLevelInput, ...request.Option) (*iot.DeleteV2LoggingLevelOutput, error)
	DeleteV2LoggingLevelRequest(*iot.DeleteV2LoggingLevelInput) (*request.Request, *iot.DeleteV2LoggingLevelOutput)

	DeprecateThingType(*iot.DeprecateThingTypeInput) (*iot.DeprecateThingTypeOutput, error)
	DeprecateThingTypeWithContext(aws.Context, *iot.DeprecateThingTypeInput, ...request.Option) (*iot.DeprecateThingTypeOutput, error)
	DeprecateThingTypeRequest(*iot.DeprecateThingTypeInput) (*request.Request, *iot.DeprecateThingTypeOutput)

	DescribeAuthorizer(*iot.DescribeAuthorizerInput) (*iot.DescribeAuthorizerOutput, error)
	DescribeAuthorizerWithContext(aws.Context, *iot.DescribeAuthorizerInput, ...request.Option) (*iot.DescribeAuthorizerOutput, error)
	DescribeAuthorizerRequest(*iot.DescribeAuthorizerInput) (*request.Request, *iot.DescribeAuthorizerOutput)

	DescribeCACertificate(*iot.DescribeCACertificateInput) (*iot.DescribeCACertificateOutput, error)
	DescribeCACertificateWithContext(aws.Context, *iot.DescribeCACertificateInput, ...request.Option) (*iot.DescribeCACertificateOutput, error)
	DescribeCACertificateRequest(*iot.DescribeCACertificateInput) (*request.Request, *iot.DescribeCACertificateOutput)

	DescribeCertificate(*iot.DescribeCertificateInput) (*iot.DescribeCertificateOutput, error)
	DescribeCertificateWithContext(aws.Context, *iot.DescribeCertificateInput, ...request.Option) (*iot.DescribeCertificateOutput, error)
	DescribeCertificateRequest(*iot.DescribeCertificateInput) (*request.Request, *iot.DescribeCertificateOutput)

	DescribeDefaultAuthorizer(*iot.DescribeDefaultAuthorizerInput) (*iot.DescribeDefaultAuthorizerOutput, error)
	DescribeDefaultAuthorizerWithContext(aws.Context, *iot.DescribeDefaultAuthorizerInput, ...request.Option) (*iot.DescribeDefaultAuthorizerOutput, error)
	DescribeDefaultAuthorizerRequest(*iot.DescribeDefaultAuthorizerInput) (*request.Request, *iot.DescribeDefaultAuthorizerOutput)

	DescribeEndpoint(*iot.DescribeEndpointInput) (*iot.DescribeEndpointOutput, error)
	DescribeEndpointWithContext(aws.Context, *iot.DescribeEndpointInput, ...request.Option) (*iot.DescribeEndpointOutput, error)
	DescribeEndpointRequest(*iot.DescribeEndpointInput) (*request.Request, *iot.DescribeEndpointOutput)

	DescribeEventConfigurations(*iot.DescribeEventConfigurationsInput) (*iot.DescribeEventConfigurationsOutput, error)
	DescribeEventConfigurationsWithContext(aws.Context, *iot.DescribeEventConfigurationsInput, ...request.Option) (*iot.DescribeEventConfigurationsOutput, error)
	DescribeEventConfigurationsRequest(*iot.DescribeEventConfigurationsInput) (*request.Request, *iot.DescribeEventConfigurationsOutput)

	DescribeIndex(*iot.DescribeIndexInput) (*iot.DescribeIndexOutput, error)
	DescribeIndexWithContext(aws.Context, *iot.DescribeIndexInput, ...request.Option) (*iot.DescribeIndexOutput, error)
	DescribeIndexRequest(*iot.DescribeIndexInput) (*request.Request, *iot.DescribeIndexOutput)

	DescribeJob(*iot.DescribeJobInput) (*iot.DescribeJobOutput, error)
	DescribeJobWithContext(aws.Context, *iot.DescribeJobInput, ...request.Option) (*iot.DescribeJobOutput, error)
	DescribeJobRequest(*iot.DescribeJobInput) (*request.Request, *iot.DescribeJobOutput)

	DescribeJobExecution(*iot.DescribeJobExecutionInput) (*iot.DescribeJobExecutionOutput, error)
	DescribeJobExecutionWithContext(aws.Context, *iot.DescribeJobExecutionInput, ...request.Option) (*iot.DescribeJobExecutionOutput, error)
	DescribeJobExecutionRequest(*iot.DescribeJobExecutionInput) (*request.Request, *iot.DescribeJobExecutionOutput)

	DescribeRoleAlias(*iot.DescribeRoleAliasInput) (*iot.DescribeRoleAliasOutput, error)
	DescribeRoleAliasWithContext(aws.Context, *iot.DescribeRoleAliasInput, ...request.Option) (*iot.DescribeRoleAliasOutput, error)
	DescribeRoleAliasRequest(*iot.DescribeRoleAliasInput) (*request.Request, *iot.DescribeRoleAliasOutput)

	DescribeThing(*iot.DescribeThingInput) (*iot.DescribeThingOutput, error)
	DescribeThingWithContext(aws.Context, *iot.DescribeThingInput, ...request.Option) (*iot.DescribeThingOutput, error)
	DescribeThingRequest(*iot.DescribeThingInput) (*request.Request, *iot.DescribeThingOutput)

	DescribeThingGroup(*iot.DescribeThingGroupInput) (*iot.DescribeThingGroupOutput, error)
	DescribeThingGroupWithContext(aws.Context, *iot.DescribeThingGroupInput, ...request.Option) (*iot.DescribeThingGroupOutput, error)
	DescribeThingGroupRequest(*iot.DescribeThingGroupInput) (*request.Request, *iot.DescribeThingGroupOutput)

	DescribeThingRegistrationTask(*iot.DescribeThingRegistrationTaskInput) (*iot.DescribeThingRegistrationTaskOutput, error)
	DescribeThingRegistrationTaskWithContext(aws.Context, *iot.DescribeThingRegistrationTaskInput, ...request.Option) (*iot.DescribeThingRegistrationTaskOutput, error)
	DescribeThingRegistrationTaskRequest(*iot.DescribeThingRegistrationTaskInput) (*request.Request, *iot.DescribeThingRegistrationTaskOutput)

	DescribeThingType(*iot.DescribeThingTypeInput) (*iot.DescribeThingTypeOutput, error)
	DescribeThingTypeWithContext(aws.Context, *iot.DescribeThingTypeInput, ...request.Option) (*iot.DescribeThingTypeOutput, error)
	DescribeThingTypeRequest(*iot.DescribeThingTypeInput) (*request.Request, *iot.DescribeThingTypeOutput)

	DetachPolicy(*iot.DetachPolicyInput) (*iot.DetachPolicyOutput, error)
	DetachPolicyWithContext(aws.Context, *iot.DetachPolicyInput, ...request.Option) (*iot.DetachPolicyOutput, error)
	DetachPolicyRequest(*iot.DetachPolicyInput) (*request.Request, *iot.DetachPolicyOutput)

	DetachPrincipalPolicy(*iot.DetachPrincipalPolicyInput) (*iot.DetachPrincipalPolicyOutput, error)
	DetachPrincipalPolicyWithContext(aws.Context, *iot.DetachPrincipalPolicyInput, ...request.Option) (*iot.DetachPrincipalPolicyOutput, error)
	DetachPrincipalPolicyRequest(*iot.DetachPrincipalPolicyInput) (*request.Request, *iot.DetachPrincipalPolicyOutput)

	DetachThingPrincipal(*iot.DetachThingPrincipalInput) (*iot.DetachThingPrincipalOutput, error)
	DetachThingPrincipalWithContext(aws.Context, *iot.DetachThingPrincipalInput, ...request.Option) (*iot.DetachThingPrincipalOutput, error)
	DetachThingPrincipalRequest(*iot.DetachThingPrincipalInput) (*request.Request, *iot.DetachThingPrincipalOutput)

	DisableTopicRule(*iot.DisableTopicRuleInput) (*iot.DisableTopicRuleOutput, error)
	DisableTopicRuleWithContext(aws.Context, *iot.DisableTopicRuleInput, ...request.Option) (*iot.DisableTopicRuleOutput, error)
	DisableTopicRuleRequest(*iot.DisableTopicRuleInput) (*request.Request, *iot.DisableTopicRuleOutput)

	EnableTopicRule(*iot.EnableTopicRuleInput) (*iot.EnableTopicRuleOutput, error)
	EnableTopicRuleWithContext(aws.Context, *iot.EnableTopicRuleInput, ...request.Option) (*iot.EnableTopicRuleOutput, error)
	EnableTopicRuleRequest(*iot.EnableTopicRuleInput) (*request.Request, *iot.EnableTopicRuleOutput)

	GetEffectivePolicies(*iot.GetEffectivePoliciesInput) (*iot.GetEffectivePoliciesOutput, error)
	GetEffectivePoliciesWithContext(aws.Context, *iot.GetEffectivePoliciesInput, ...request.Option) (*iot.GetEffectivePoliciesOutput, error)
	GetEffectivePoliciesRequest(*iot.GetEffectivePoliciesInput) (*request.Request, *iot.GetEffectivePoliciesOutput)

	GetIndexingConfiguration(*iot.GetIndexingConfigurationInput) (*iot.GetIndexingConfigurationOutput, error)
	GetIndexingConfigurationWithContext(aws.Context, *iot.GetIndexingConfigurationInput, ...request.Option) (*iot.GetIndexingConfigurationOutput, error)
	GetIndexingConfigurationRequest(*iot.GetIndexingConfigurationInput) (*request.Request, *iot.GetIndexingConfigurationOutput)

	GetJobDocument(*iot.GetJobDocumentInput) (*iot.GetJobDocumentOutput, error)
	GetJobDocumentWithContext(aws.Context, *iot.GetJobDocumentInput, ...request.Option) (*iot.GetJobDocumentOutput, error)
	GetJobDocumentRequest(*iot.GetJobDocumentInput) (*request.Request, *iot.GetJobDocumentOutput)

	GetLoggingOptions(*iot.GetLoggingOptionsInput) (*iot.GetLoggingOptionsOutput, error)
	GetLoggingOptionsWithContext(aws.Context, *iot.GetLoggingOptionsInput, ...request.Option) (*iot.GetLoggingOptionsOutput, error)
	GetLoggingOptionsRequest(*iot.GetLoggingOptionsInput) (*request.Request, *iot.GetLoggingOptionsOutput)

	GetPolicy(*iot.GetPolicyInput) (*iot.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *iot.GetPolicyInput, ...request.Option) (*iot.GetPolicyOutput, error)
	GetPolicyRequest(*iot.GetPolicyInput) (*request.Request, *iot.GetPolicyOutput)

	GetPolicyVersion(*iot.GetPolicyVersionInput) (*iot.GetPolicyVersionOutput, error)
	GetPolicyVersionWithContext(aws.Context, *iot.GetPolicyVersionInput, ...request.Option) (*iot.GetPolicyVersionOutput, error)
	GetPolicyVersionRequest(*iot.GetPolicyVersionInput) (*request.Request, *iot.GetPolicyVersionOutput)

	GetRegistrationCode(*iot.GetRegistrationCodeInput) (*iot.GetRegistrationCodeOutput, error)
	GetRegistrationCodeWithContext(aws.Context, *iot.GetRegistrationCodeInput, ...request.Option) (*iot.GetRegistrationCodeOutput, error)
	GetRegistrationCodeRequest(*iot.GetRegistrationCodeInput) (*request.Request, *iot.GetRegistrationCodeOutput)

	GetTopicRule(*iot.GetTopicRuleInput) (*iot.GetTopicRuleOutput, error)
	GetTopicRuleWithContext(aws.Context, *iot.GetTopicRuleInput, ...request.Option) (*iot.GetTopicRuleOutput, error)
	GetTopicRuleRequest(*iot.GetTopicRuleInput) (*request.Request, *iot.GetTopicRuleOutput)

	GetV2LoggingOptions(*iot.GetV2LoggingOptionsInput) (*iot.GetV2LoggingOptionsOutput, error)
	GetV2LoggingOptionsWithContext(aws.Context, *iot.GetV2LoggingOptionsInput, ...request.Option) (*iot.GetV2LoggingOptionsOutput, error)
	GetV2LoggingOptionsRequest(*iot.GetV2LoggingOptionsInput) (*request.Request, *iot.GetV2LoggingOptionsOutput)

	ListAttachedPolicies(*iot.ListAttachedPoliciesInput) (*iot.ListAttachedPoliciesOutput, error)
	ListAttachedPoliciesWithContext(aws.Context, *iot.ListAttachedPoliciesInput, ...request.Option) (*iot.ListAttachedPoliciesOutput, error)
	ListAttachedPoliciesRequest(*iot.ListAttachedPoliciesInput) (*request.Request, *iot.ListAttachedPoliciesOutput)

	ListAuthorizers(*iot.ListAuthorizersInput) (*iot.ListAuthorizersOutput, error)
	ListAuthorizersWithContext(aws.Context, *iot.ListAuthorizersInput, ...request.Option) (*iot.ListAuthorizersOutput, error)
	ListAuthorizersRequest(*iot.ListAuthorizersInput) (*request.Request, *iot.ListAuthorizersOutput)

	ListCACertificates(*iot.ListCACertificatesInput) (*iot.ListCACertificatesOutput, error)
	ListCACertificatesWithContext(aws.Context, *iot.ListCACertificatesInput, ...request.Option) (*iot.ListCACertificatesOutput, error)
	ListCACertificatesRequest(*iot.ListCACertificatesInput) (*request.Request, *iot.ListCACertificatesOutput)

	ListCertificates(*iot.ListCertificatesInput) (*iot.ListCertificatesOutput, error)
	ListCertificatesWithContext(aws.Context, *iot.ListCertificatesInput, ...request.Option) (*iot.ListCertificatesOutput, error)
	ListCertificatesRequest(*iot.ListCertificatesInput) (*request.Request, *iot.ListCertificatesOutput)

	ListCertificatesByCA(*iot.ListCertificatesByCAInput) (*iot.ListCertificatesByCAOutput, error)
	ListCertificatesByCAWithContext(aws.Context, *iot.ListCertificatesByCAInput, ...request.Option) (*iot.ListCertificatesByCAOutput, error)
	ListCertificatesByCARequest(*iot.ListCertificatesByCAInput) (*request.Request, *iot.ListCertificatesByCAOutput)

	ListIndices(*iot.ListIndicesInput) (*iot.ListIndicesOutput, error)
	ListIndicesWithContext(aws.Context, *iot.ListIndicesInput, ...request.Option) (*iot.ListIndicesOutput, error)
	ListIndicesRequest(*iot.ListIndicesInput) (*request.Request, *iot.ListIndicesOutput)

	ListJobExecutionsForJob(*iot.ListJobExecutionsForJobInput) (*iot.ListJobExecutionsForJobOutput, error)
	ListJobExecutionsForJobWithContext(aws.Context, *iot.ListJobExecutionsForJobInput, ...request.Option) (*iot.ListJobExecutionsForJobOutput, error)
	ListJobExecutionsForJobRequest(*iot.ListJobExecutionsForJobInput) (*request.Request, *iot.ListJobExecutionsForJobOutput)

	ListJobExecutionsForThing(*iot.ListJobExecutionsForThingInput) (*iot.ListJobExecutionsForThingOutput, error)
	ListJobExecutionsForThingWithContext(aws.Context, *iot.ListJobExecutionsForThingInput, ...request.Option) (*iot.ListJobExecutionsForThingOutput, error)
	ListJobExecutionsForThingRequest(*iot.ListJobExecutionsForThingInput) (*request.Request, *iot.ListJobExecutionsForThingOutput)

	ListJobs(*iot.ListJobsInput) (*iot.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *iot.ListJobsInput, ...request.Option) (*iot.ListJobsOutput, error)
	ListJobsRequest(*iot.ListJobsInput) (*request.Request, *iot.ListJobsOutput)

	ListOutgoingCertificates(*iot.ListOutgoingCertificatesInput) (*iot.ListOutgoingCertificatesOutput, error)
	ListOutgoingCertificatesWithContext(aws.Context, *iot.ListOutgoingCertificatesInput, ...request.Option) (*iot.ListOutgoingCertificatesOutput, error)
	ListOutgoingCertificatesRequest(*iot.ListOutgoingCertificatesInput) (*request.Request, *iot.ListOutgoingCertificatesOutput)

	ListPolicies(*iot.ListPoliciesInput) (*iot.ListPoliciesOutput, error)
	ListPoliciesWithContext(aws.Context, *iot.ListPoliciesInput, ...request.Option) (*iot.ListPoliciesOutput, error)
	ListPoliciesRequest(*iot.ListPoliciesInput) (*request.Request, *iot.ListPoliciesOutput)

	ListPolicyPrincipals(*iot.ListPolicyPrincipalsInput) (*iot.ListPolicyPrincipalsOutput, error)
	ListPolicyPrincipalsWithContext(aws.Context, *iot.ListPolicyPrincipalsInput, ...request.Option) (*iot.ListPolicyPrincipalsOutput, error)
	ListPolicyPrincipalsRequest(*iot.ListPolicyPrincipalsInput) (*request.Request, *iot.ListPolicyPrincipalsOutput)

	ListPolicyVersions(*iot.ListPolicyVersionsInput) (*iot.ListPolicyVersionsOutput, error)
	ListPolicyVersionsWithContext(aws.Context, *iot.ListPolicyVersionsInput, ...request.Option) (*iot.ListPolicyVersionsOutput, error)
	ListPolicyVersionsRequest(*iot.ListPolicyVersionsInput) (*request.Request, *iot.ListPolicyVersionsOutput)

	ListPrincipalPolicies(*iot.ListPrincipalPoliciesInput) (*iot.ListPrincipalPoliciesOutput, error)
	ListPrincipalPoliciesWithContext(aws.Context, *iot.ListPrincipalPoliciesInput, ...request.Option) (*iot.ListPrincipalPoliciesOutput, error)
	ListPrincipalPoliciesRequest(*iot.ListPrincipalPoliciesInput) (*request.Request, *iot.ListPrincipalPoliciesOutput)

	ListPrincipalThings(*iot.ListPrincipalThingsInput) (*iot.ListPrincipalThingsOutput, error)
	ListPrincipalThingsWithContext(aws.Context, *iot.ListPrincipalThingsInput, ...request.Option) (*iot.ListPrincipalThingsOutput, error)
	ListPrincipalThingsRequest(*iot.ListPrincipalThingsInput) (*request.Request, *iot.ListPrincipalThingsOutput)

	ListRoleAliases(*iot.ListRoleAliasesInput) (*iot.ListRoleAliasesOutput, error)
	ListRoleAliasesWithContext(aws.Context, *iot.ListRoleAliasesInput, ...request.Option) (*iot.ListRoleAliasesOutput, error)
	ListRoleAliasesRequest(*iot.ListRoleAliasesInput) (*request.Request, *iot.ListRoleAliasesOutput)

	ListTargetsForPolicy(*iot.ListTargetsForPolicyInput) (*iot.ListTargetsForPolicyOutput, error)
	ListTargetsForPolicyWithContext(aws.Context, *iot.ListTargetsForPolicyInput, ...request.Option) (*iot.ListTargetsForPolicyOutput, error)
	ListTargetsForPolicyRequest(*iot.ListTargetsForPolicyInput) (*request.Request, *iot.ListTargetsForPolicyOutput)

	ListThingGroups(*iot.ListThingGroupsInput) (*iot.ListThingGroupsOutput, error)
	ListThingGroupsWithContext(aws.Context, *iot.ListThingGroupsInput, ...request.Option) (*iot.ListThingGroupsOutput, error)
	ListThingGroupsRequest(*iot.ListThingGroupsInput) (*request.Request, *iot.ListThingGroupsOutput)

	ListThingGroupsForThing(*iot.ListThingGroupsForThingInput) (*iot.ListThingGroupsForThingOutput, error)
	ListThingGroupsForThingWithContext(aws.Context, *iot.ListThingGroupsForThingInput, ...request.Option) (*iot.ListThingGroupsForThingOutput, error)
	ListThingGroupsForThingRequest(*iot.ListThingGroupsForThingInput) (*request.Request, *iot.ListThingGroupsForThingOutput)

	ListThingPrincipals(*iot.ListThingPrincipalsInput) (*iot.ListThingPrincipalsOutput, error)
	ListThingPrincipalsWithContext(aws.Context, *iot.ListThingPrincipalsInput, ...request.Option) (*iot.ListThingPrincipalsOutput, error)
	ListThingPrincipalsRequest(*iot.ListThingPrincipalsInput) (*request.Request, *iot.ListThingPrincipalsOutput)

	ListThingRegistrationTaskReports(*iot.ListThingRegistrationTaskReportsInput) (*iot.ListThingRegistrationTaskReportsOutput, error)
	ListThingRegistrationTaskReportsWithContext(aws.Context, *iot.ListThingRegistrationTaskReportsInput, ...request.Option) (*iot.ListThingRegistrationTaskReportsOutput, error)
	ListThingRegistrationTaskReportsRequest(*iot.ListThingRegistrationTaskReportsInput) (*request.Request, *iot.ListThingRegistrationTaskReportsOutput)

	ListThingRegistrationTasks(*iot.ListThingRegistrationTasksInput) (*iot.ListThingRegistrationTasksOutput, error)
	ListThingRegistrationTasksWithContext(aws.Context, *iot.ListThingRegistrationTasksInput, ...request.Option) (*iot.ListThingRegistrationTasksOutput, error)
	ListThingRegistrationTasksRequest(*iot.ListThingRegistrationTasksInput) (*request.Request, *iot.ListThingRegistrationTasksOutput)

	ListThingTypes(*iot.ListThingTypesInput) (*iot.ListThingTypesOutput, error)
	ListThingTypesWithContext(aws.Context, *iot.ListThingTypesInput, ...request.Option) (*iot.ListThingTypesOutput, error)
	ListThingTypesRequest(*iot.ListThingTypesInput) (*request.Request, *iot.ListThingTypesOutput)

	ListThings(*iot.ListThingsInput) (*iot.ListThingsOutput, error)
	ListThingsWithContext(aws.Context, *iot.ListThingsInput, ...request.Option) (*iot.ListThingsOutput, error)
	ListThingsRequest(*iot.ListThingsInput) (*request.Request, *iot.ListThingsOutput)

	ListThingsInThingGroup(*iot.ListThingsInThingGroupInput) (*iot.ListThingsInThingGroupOutput, error)
	ListThingsInThingGroupWithContext(aws.Context, *iot.ListThingsInThingGroupInput, ...request.Option) (*iot.ListThingsInThingGroupOutput, error)
	ListThingsInThingGroupRequest(*iot.ListThingsInThingGroupInput) (*request.Request, *iot.ListThingsInThingGroupOutput)

	ListTopicRules(*iot.ListTopicRulesInput) (*iot.ListTopicRulesOutput, error)
	ListTopicRulesWithContext(aws.Context, *iot.ListTopicRulesInput, ...request.Option) (*iot.ListTopicRulesOutput, error)
	ListTopicRulesRequest(*iot.ListTopicRulesInput) (*request.Request, *iot.ListTopicRulesOutput)

	ListV2LoggingLevels(*iot.ListV2LoggingLevelsInput) (*iot.ListV2LoggingLevelsOutput, error)
	ListV2LoggingLevelsWithContext(aws.Context, *iot.ListV2LoggingLevelsInput, ...request.Option) (*iot.ListV2LoggingLevelsOutput, error)
	ListV2LoggingLevelsRequest(*iot.ListV2LoggingLevelsInput) (*request.Request, *iot.ListV2LoggingLevelsOutput)

	RegisterCACertificate(*iot.RegisterCACertificateInput) (*iot.RegisterCACertificateOutput, error)
	RegisterCACertificateWithContext(aws.Context, *iot.RegisterCACertificateInput, ...request.Option) (*iot.RegisterCACertificateOutput, error)
	RegisterCACertificateRequest(*iot.RegisterCACertificateInput) (*request.Request, *iot.RegisterCACertificateOutput)

	RegisterCertificate(*iot.RegisterCertificateInput) (*iot.RegisterCertificateOutput, error)
	RegisterCertificateWithContext(aws.Context, *iot.RegisterCertificateInput, ...request.Option) (*iot.RegisterCertificateOutput, error)
	RegisterCertificateRequest(*iot.RegisterCertificateInput) (*request.Request, *iot.RegisterCertificateOutput)

	RegisterThing(*iot.RegisterThingInput) (*iot.RegisterThingOutput, error)
	RegisterThingWithContext(aws.Context, *iot.RegisterThingInput, ...request.Option) (*iot.RegisterThingOutput, error)
	RegisterThingRequest(*iot.RegisterThingInput) (*request.Request, *iot.RegisterThingOutput)

	RejectCertificateTransfer(*iot.RejectCertificateTransferInput) (*iot.RejectCertificateTransferOutput, error)
	RejectCertificateTransferWithContext(aws.Context, *iot.RejectCertificateTransferInput, ...request.Option) (*iot.RejectCertificateTransferOutput, error)
	RejectCertificateTransferRequest(*iot.RejectCertificateTransferInput) (*request.Request, *iot.RejectCertificateTransferOutput)

	RemoveThingFromThingGroup(*iot.RemoveThingFromThingGroupInput) (*iot.RemoveThingFromThingGroupOutput, error)
	RemoveThingFromThingGroupWithContext(aws.Context, *iot.RemoveThingFromThingGroupInput, ...request.Option) (*iot.RemoveThingFromThingGroupOutput, error)
	RemoveThingFromThingGroupRequest(*iot.RemoveThingFromThingGroupInput) (*request.Request, *iot.RemoveThingFromThingGroupOutput)

	ReplaceTopicRule(*iot.ReplaceTopicRuleInput) (*iot.ReplaceTopicRuleOutput, error)
	ReplaceTopicRuleWithContext(aws.Context, *iot.ReplaceTopicRuleInput, ...request.Option) (*iot.ReplaceTopicRuleOutput, error)
	ReplaceTopicRuleRequest(*iot.ReplaceTopicRuleInput) (*request.Request, *iot.ReplaceTopicRuleOutput)

	SearchIndex(*iot.SearchIndexInput) (*iot.SearchIndexOutput, error)
	SearchIndexWithContext(aws.Context, *iot.SearchIndexInput, ...request.Option) (*iot.SearchIndexOutput, error)
	SearchIndexRequest(*iot.SearchIndexInput) (*request.Request, *iot.SearchIndexOutput)

	SetDefaultAuthorizer(*iot.SetDefaultAuthorizerInput) (*iot.SetDefaultAuthorizerOutput, error)
	SetDefaultAuthorizerWithContext(aws.Context, *iot.SetDefaultAuthorizerInput, ...request.Option) (*iot.SetDefaultAuthorizerOutput, error)
	SetDefaultAuthorizerRequest(*iot.SetDefaultAuthorizerInput) (*request.Request, *iot.SetDefaultAuthorizerOutput)

	SetDefaultPolicyVersion(*iot.SetDefaultPolicyVersionInput) (*iot.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionWithContext(aws.Context, *iot.SetDefaultPolicyVersionInput, ...request.Option) (*iot.SetDefaultPolicyVersionOutput, error)
	SetDefaultPolicyVersionRequest(*iot.SetDefaultPolicyVersionInput) (*request.Request, *iot.SetDefaultPolicyVersionOutput)

	SetLoggingOptions(*iot.SetLoggingOptionsInput) (*iot.SetLoggingOptionsOutput, error)
	SetLoggingOptionsWithContext(aws.Context, *iot.SetLoggingOptionsInput, ...request.Option) (*iot.SetLoggingOptionsOutput, error)
	SetLoggingOptionsRequest(*iot.SetLoggingOptionsInput) (*request.Request, *iot.SetLoggingOptionsOutput)

	SetV2LoggingLevel(*iot.SetV2LoggingLevelInput) (*iot.SetV2LoggingLevelOutput, error)
	SetV2LoggingLevelWithContext(aws.Context, *iot.SetV2LoggingLevelInput, ...request.Option) (*iot.SetV2LoggingLevelOutput, error)
	SetV2LoggingLevelRequest(*iot.SetV2LoggingLevelInput) (*request.Request, *iot.SetV2LoggingLevelOutput)

	SetV2LoggingOptions(*iot.SetV2LoggingOptionsInput) (*iot.SetV2LoggingOptionsOutput, error)
	SetV2LoggingOptionsWithContext(aws.Context, *iot.SetV2LoggingOptionsInput, ...request.Option) (*iot.SetV2LoggingOptionsOutput, error)
	SetV2LoggingOptionsRequest(*iot.SetV2LoggingOptionsInput) (*request.Request, *iot.SetV2LoggingOptionsOutput)

	StartThingRegistrationTask(*iot.StartThingRegistrationTaskInput) (*iot.StartThingRegistrationTaskOutput, error)
	StartThingRegistrationTaskWithContext(aws.Context, *iot.StartThingRegistrationTaskInput, ...request.Option) (*iot.StartThingRegistrationTaskOutput, error)
	StartThingRegistrationTaskRequest(*iot.StartThingRegistrationTaskInput) (*request.Request, *iot.StartThingRegistrationTaskOutput)

	StopThingRegistrationTask(*iot.StopThingRegistrationTaskInput) (*iot.StopThingRegistrationTaskOutput, error)
	StopThingRegistrationTaskWithContext(aws.Context, *iot.StopThingRegistrationTaskInput, ...request.Option) (*iot.StopThingRegistrationTaskOutput, error)
	StopThingRegistrationTaskRequest(*iot.StopThingRegistrationTaskInput) (*request.Request, *iot.StopThingRegistrationTaskOutput)

	TestAuthorization(*iot.TestAuthorizationInput) (*iot.TestAuthorizationOutput, error)
	TestAuthorizationWithContext(aws.Context, *iot.TestAuthorizationInput, ...request.Option) (*iot.TestAuthorizationOutput, error)
	TestAuthorizationRequest(*iot.TestAuthorizationInput) (*request.Request, *iot.TestAuthorizationOutput)

	TestInvokeAuthorizer(*iot.TestInvokeAuthorizerInput) (*iot.TestInvokeAuthorizerOutput, error)
	TestInvokeAuthorizerWithContext(aws.Context, *iot.TestInvokeAuthorizerInput, ...request.Option) (*iot.TestInvokeAuthorizerOutput, error)
	TestInvokeAuthorizerRequest(*iot.TestInvokeAuthorizerInput) (*request.Request, *iot.TestInvokeAuthorizerOutput)

	TransferCertificate(*iot.TransferCertificateInput) (*iot.TransferCertificateOutput, error)
	TransferCertificateWithContext(aws.Context, *iot.TransferCertificateInput, ...request.Option) (*iot.TransferCertificateOutput, error)
	TransferCertificateRequest(*iot.TransferCertificateInput) (*request.Request, *iot.TransferCertificateOutput)

	UpdateAuthorizer(*iot.UpdateAuthorizerInput) (*iot.UpdateAuthorizerOutput, error)
	UpdateAuthorizerWithContext(aws.Context, *iot.UpdateAuthorizerInput, ...request.Option) (*iot.UpdateAuthorizerOutput, error)
	UpdateAuthorizerRequest(*iot.UpdateAuthorizerInput) (*request.Request, *iot.UpdateAuthorizerOutput)

	UpdateCACertificate(*iot.UpdateCACertificateInput) (*iot.UpdateCACertificateOutput, error)
	UpdateCACertificateWithContext(aws.Context, *iot.UpdateCACertificateInput, ...request.Option) (*iot.UpdateCACertificateOutput, error)
	UpdateCACertificateRequest(*iot.UpdateCACertificateInput) (*request.Request, *iot.UpdateCACertificateOutput)

	UpdateCertificate(*iot.UpdateCertificateInput) (*iot.UpdateCertificateOutput, error)
	UpdateCertificateWithContext(aws.Context, *iot.UpdateCertificateInput, ...request.Option) (*iot.UpdateCertificateOutput, error)
	UpdateCertificateRequest(*iot.UpdateCertificateInput) (*request.Request, *iot.UpdateCertificateOutput)

	UpdateEventConfigurations(*iot.UpdateEventConfigurationsInput) (*iot.UpdateEventConfigurationsOutput, error)
	UpdateEventConfigurationsWithContext(aws.Context, *iot.UpdateEventConfigurationsInput, ...request.Option) (*iot.UpdateEventConfigurationsOutput, error)
	UpdateEventConfigurationsRequest(*iot.UpdateEventConfigurationsInput) (*request.Request, *iot.UpdateEventConfigurationsOutput)

	UpdateIndexingConfiguration(*iot.UpdateIndexingConfigurationInput) (*iot.UpdateIndexingConfigurationOutput, error)
	UpdateIndexingConfigurationWithContext(aws.Context, *iot.UpdateIndexingConfigurationInput, ...request.Option) (*iot.UpdateIndexingConfigurationOutput, error)
	UpdateIndexingConfigurationRequest(*iot.UpdateIndexingConfigurationInput) (*request.Request, *iot.UpdateIndexingConfigurationOutput)

	UpdateRoleAlias(*iot.UpdateRoleAliasInput) (*iot.UpdateRoleAliasOutput, error)
	UpdateRoleAliasWithContext(aws.Context, *iot.UpdateRoleAliasInput, ...request.Option) (*iot.UpdateRoleAliasOutput, error)
	UpdateRoleAliasRequest(*iot.UpdateRoleAliasInput) (*request.Request, *iot.UpdateRoleAliasOutput)

	UpdateThing(*iot.UpdateThingInput) (*iot.UpdateThingOutput, error)
	UpdateThingWithContext(aws.Context, *iot.UpdateThingInput, ...request.Option) (*iot.UpdateThingOutput, error)
	UpdateThingRequest(*iot.UpdateThingInput) (*request.Request, *iot.UpdateThingOutput)

	UpdateThingGroup(*iot.UpdateThingGroupInput) (*iot.UpdateThingGroupOutput, error)
	UpdateThingGroupWithContext(aws.Context, *iot.UpdateThingGroupInput, ...request.Option) (*iot.UpdateThingGroupOutput, error)
	UpdateThingGroupRequest(*iot.UpdateThingGroupInput) (*request.Request, *iot.UpdateThingGroupOutput)

	UpdateThingGroupsForThing(*iot.UpdateThingGroupsForThingInput) (*iot.UpdateThingGroupsForThingOutput, error)
	UpdateThingGroupsForThingWithContext(aws.Context, *iot.UpdateThingGroupsForThingInput, ...request.Option) (*iot.UpdateThingGroupsForThingOutput, error)
	UpdateThingGroupsForThingRequest(*iot.UpdateThingGroupsForThingInput) (*request.Request, *iot.UpdateThingGroupsForThingOutput)
}

var _ IoTAPI = (*iot.IoT)(nil)
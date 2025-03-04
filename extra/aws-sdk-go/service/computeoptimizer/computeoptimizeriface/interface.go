// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package computeoptimizeriface provides an interface to enable mocking the AWS Compute Optimizer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package computeoptimizeriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
)

// ComputeOptimizerAPI provides an interface to enable mocking the
// computeoptimizer.ComputeOptimizer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Compute Optimizer.
//    func myFunc(svc computeoptimizeriface.ComputeOptimizerAPI) bool {
//        // Make svc.DeleteRecommendationPreferences request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := computeoptimizer.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockComputeOptimizerClient struct {
//        computeoptimizeriface.ComputeOptimizerAPI
//    }
//    func (m *mockComputeOptimizerClient) DeleteRecommendationPreferences(input *computeoptimizer.DeleteRecommendationPreferencesInput) (*computeoptimizer.DeleteRecommendationPreferencesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockComputeOptimizerClient{}
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
type ComputeOptimizerAPI interface {
	DeleteRecommendationPreferences(*computeoptimizer.DeleteRecommendationPreferencesInput) (*computeoptimizer.DeleteRecommendationPreferencesOutput, error)
	DeleteRecommendationPreferencesWithContext(aws.Context, *computeoptimizer.DeleteRecommendationPreferencesInput, ...request.Option) (*computeoptimizer.DeleteRecommendationPreferencesOutput, error)
	DeleteRecommendationPreferencesRequest(*computeoptimizer.DeleteRecommendationPreferencesInput) (*request.Request, *computeoptimizer.DeleteRecommendationPreferencesOutput)

	DescribeRecommendationExportJobs(*computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
	DescribeRecommendationExportJobsWithContext(aws.Context, *computeoptimizer.DescribeRecommendationExportJobsInput, ...request.Option) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
	DescribeRecommendationExportJobsRequest(*computeoptimizer.DescribeRecommendationExportJobsInput) (*request.Request, *computeoptimizer.DescribeRecommendationExportJobsOutput)

	DescribeRecommendationExportJobsPages(*computeoptimizer.DescribeRecommendationExportJobsInput, func(*computeoptimizer.DescribeRecommendationExportJobsOutput, bool) bool) error
	DescribeRecommendationExportJobsPagesWithContext(aws.Context, *computeoptimizer.DescribeRecommendationExportJobsInput, func(*computeoptimizer.DescribeRecommendationExportJobsOutput, bool) bool, ...request.Option) error

	ExportAutoScalingGroupRecommendations(*computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
	ExportAutoScalingGroupRecommendationsWithContext(aws.Context, *computeoptimizer.ExportAutoScalingGroupRecommendationsInput, ...request.Option) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
	ExportAutoScalingGroupRecommendationsRequest(*computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*request.Request, *computeoptimizer.ExportAutoScalingGroupRecommendationsOutput)

	ExportEBSVolumeRecommendations(*computeoptimizer.ExportEBSVolumeRecommendationsInput) (*computeoptimizer.ExportEBSVolumeRecommendationsOutput, error)
	ExportEBSVolumeRecommendationsWithContext(aws.Context, *computeoptimizer.ExportEBSVolumeRecommendationsInput, ...request.Option) (*computeoptimizer.ExportEBSVolumeRecommendationsOutput, error)
	ExportEBSVolumeRecommendationsRequest(*computeoptimizer.ExportEBSVolumeRecommendationsInput) (*request.Request, *computeoptimizer.ExportEBSVolumeRecommendationsOutput)

	ExportEC2InstanceRecommendations(*computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
	ExportEC2InstanceRecommendationsWithContext(aws.Context, *computeoptimizer.ExportEC2InstanceRecommendationsInput, ...request.Option) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
	ExportEC2InstanceRecommendationsRequest(*computeoptimizer.ExportEC2InstanceRecommendationsInput) (*request.Request, *computeoptimizer.ExportEC2InstanceRecommendationsOutput)

	ExportECSServiceRecommendations(*computeoptimizer.ExportECSServiceRecommendationsInput) (*computeoptimizer.ExportECSServiceRecommendationsOutput, error)
	ExportECSServiceRecommendationsWithContext(aws.Context, *computeoptimizer.ExportECSServiceRecommendationsInput, ...request.Option) (*computeoptimizer.ExportECSServiceRecommendationsOutput, error)
	ExportECSServiceRecommendationsRequest(*computeoptimizer.ExportECSServiceRecommendationsInput) (*request.Request, *computeoptimizer.ExportECSServiceRecommendationsOutput)

	ExportLambdaFunctionRecommendations(*computeoptimizer.ExportLambdaFunctionRecommendationsInput) (*computeoptimizer.ExportLambdaFunctionRecommendationsOutput, error)
	ExportLambdaFunctionRecommendationsWithContext(aws.Context, *computeoptimizer.ExportLambdaFunctionRecommendationsInput, ...request.Option) (*computeoptimizer.ExportLambdaFunctionRecommendationsOutput, error)
	ExportLambdaFunctionRecommendationsRequest(*computeoptimizer.ExportLambdaFunctionRecommendationsInput) (*request.Request, *computeoptimizer.ExportLambdaFunctionRecommendationsOutput)

	GetAutoScalingGroupRecommendations(*computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
	GetAutoScalingGroupRecommendationsWithContext(aws.Context, *computeoptimizer.GetAutoScalingGroupRecommendationsInput, ...request.Option) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
	GetAutoScalingGroupRecommendationsRequest(*computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*request.Request, *computeoptimizer.GetAutoScalingGroupRecommendationsOutput)

	GetEBSVolumeRecommendations(*computeoptimizer.GetEBSVolumeRecommendationsInput) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error)
	GetEBSVolumeRecommendationsWithContext(aws.Context, *computeoptimizer.GetEBSVolumeRecommendationsInput, ...request.Option) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error)
	GetEBSVolumeRecommendationsRequest(*computeoptimizer.GetEBSVolumeRecommendationsInput) (*request.Request, *computeoptimizer.GetEBSVolumeRecommendationsOutput)

	GetEC2InstanceRecommendations(*computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
	GetEC2InstanceRecommendationsWithContext(aws.Context, *computeoptimizer.GetEC2InstanceRecommendationsInput, ...request.Option) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
	GetEC2InstanceRecommendationsRequest(*computeoptimizer.GetEC2InstanceRecommendationsInput) (*request.Request, *computeoptimizer.GetEC2InstanceRecommendationsOutput)

	GetEC2RecommendationProjectedMetrics(*computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
	GetEC2RecommendationProjectedMetricsWithContext(aws.Context, *computeoptimizer.GetEC2RecommendationProjectedMetricsInput, ...request.Option) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
	GetEC2RecommendationProjectedMetricsRequest(*computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*request.Request, *computeoptimizer.GetEC2RecommendationProjectedMetricsOutput)

	GetECSServiceRecommendationProjectedMetrics(*computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput) (*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput, error)
	GetECSServiceRecommendationProjectedMetricsWithContext(aws.Context, *computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput, ...request.Option) (*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput, error)
	GetECSServiceRecommendationProjectedMetricsRequest(*computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput) (*request.Request, *computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput)

	GetECSServiceRecommendations(*computeoptimizer.GetECSServiceRecommendationsInput) (*computeoptimizer.GetECSServiceRecommendationsOutput, error)
	GetECSServiceRecommendationsWithContext(aws.Context, *computeoptimizer.GetECSServiceRecommendationsInput, ...request.Option) (*computeoptimizer.GetECSServiceRecommendationsOutput, error)
	GetECSServiceRecommendationsRequest(*computeoptimizer.GetECSServiceRecommendationsInput) (*request.Request, *computeoptimizer.GetECSServiceRecommendationsOutput)

	GetEffectiveRecommendationPreferences(*computeoptimizer.GetEffectiveRecommendationPreferencesInput) (*computeoptimizer.GetEffectiveRecommendationPreferencesOutput, error)
	GetEffectiveRecommendationPreferencesWithContext(aws.Context, *computeoptimizer.GetEffectiveRecommendationPreferencesInput, ...request.Option) (*computeoptimizer.GetEffectiveRecommendationPreferencesOutput, error)
	GetEffectiveRecommendationPreferencesRequest(*computeoptimizer.GetEffectiveRecommendationPreferencesInput) (*request.Request, *computeoptimizer.GetEffectiveRecommendationPreferencesOutput)

	GetEnrollmentStatus(*computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error)
	GetEnrollmentStatusWithContext(aws.Context, *computeoptimizer.GetEnrollmentStatusInput, ...request.Option) (*computeoptimizer.GetEnrollmentStatusOutput, error)
	GetEnrollmentStatusRequest(*computeoptimizer.GetEnrollmentStatusInput) (*request.Request, *computeoptimizer.GetEnrollmentStatusOutput)

	GetEnrollmentStatusesForOrganization(*computeoptimizer.GetEnrollmentStatusesForOrganizationInput) (*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, error)
	GetEnrollmentStatusesForOrganizationWithContext(aws.Context, *computeoptimizer.GetEnrollmentStatusesForOrganizationInput, ...request.Option) (*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, error)
	GetEnrollmentStatusesForOrganizationRequest(*computeoptimizer.GetEnrollmentStatusesForOrganizationInput) (*request.Request, *computeoptimizer.GetEnrollmentStatusesForOrganizationOutput)

	GetEnrollmentStatusesForOrganizationPages(*computeoptimizer.GetEnrollmentStatusesForOrganizationInput, func(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, bool) bool) error
	GetEnrollmentStatusesForOrganizationPagesWithContext(aws.Context, *computeoptimizer.GetEnrollmentStatusesForOrganizationInput, func(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, bool) bool, ...request.Option) error

	GetLambdaFunctionRecommendations(*computeoptimizer.GetLambdaFunctionRecommendationsInput) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error)
	GetLambdaFunctionRecommendationsWithContext(aws.Context, *computeoptimizer.GetLambdaFunctionRecommendationsInput, ...request.Option) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error)
	GetLambdaFunctionRecommendationsRequest(*computeoptimizer.GetLambdaFunctionRecommendationsInput) (*request.Request, *computeoptimizer.GetLambdaFunctionRecommendationsOutput)

	GetLambdaFunctionRecommendationsPages(*computeoptimizer.GetLambdaFunctionRecommendationsInput, func(*computeoptimizer.GetLambdaFunctionRecommendationsOutput, bool) bool) error
	GetLambdaFunctionRecommendationsPagesWithContext(aws.Context, *computeoptimizer.GetLambdaFunctionRecommendationsInput, func(*computeoptimizer.GetLambdaFunctionRecommendationsOutput, bool) bool, ...request.Option) error

	GetRecommendationPreferences(*computeoptimizer.GetRecommendationPreferencesInput) (*computeoptimizer.GetRecommendationPreferencesOutput, error)
	GetRecommendationPreferencesWithContext(aws.Context, *computeoptimizer.GetRecommendationPreferencesInput, ...request.Option) (*computeoptimizer.GetRecommendationPreferencesOutput, error)
	GetRecommendationPreferencesRequest(*computeoptimizer.GetRecommendationPreferencesInput) (*request.Request, *computeoptimizer.GetRecommendationPreferencesOutput)

	GetRecommendationPreferencesPages(*computeoptimizer.GetRecommendationPreferencesInput, func(*computeoptimizer.GetRecommendationPreferencesOutput, bool) bool) error
	GetRecommendationPreferencesPagesWithContext(aws.Context, *computeoptimizer.GetRecommendationPreferencesInput, func(*computeoptimizer.GetRecommendationPreferencesOutput, bool) bool, ...request.Option) error

	GetRecommendationSummaries(*computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error)
	GetRecommendationSummariesWithContext(aws.Context, *computeoptimizer.GetRecommendationSummariesInput, ...request.Option) (*computeoptimizer.GetRecommendationSummariesOutput, error)
	GetRecommendationSummariesRequest(*computeoptimizer.GetRecommendationSummariesInput) (*request.Request, *computeoptimizer.GetRecommendationSummariesOutput)

	GetRecommendationSummariesPages(*computeoptimizer.GetRecommendationSummariesInput, func(*computeoptimizer.GetRecommendationSummariesOutput, bool) bool) error
	GetRecommendationSummariesPagesWithContext(aws.Context, *computeoptimizer.GetRecommendationSummariesInput, func(*computeoptimizer.GetRecommendationSummariesOutput, bool) bool, ...request.Option) error

	PutRecommendationPreferences(*computeoptimizer.PutRecommendationPreferencesInput) (*computeoptimizer.PutRecommendationPreferencesOutput, error)
	PutRecommendationPreferencesWithContext(aws.Context, *computeoptimizer.PutRecommendationPreferencesInput, ...request.Option) (*computeoptimizer.PutRecommendationPreferencesOutput, error)
	PutRecommendationPreferencesRequest(*computeoptimizer.PutRecommendationPreferencesInput) (*request.Request, *computeoptimizer.PutRecommendationPreferencesOutput)

	UpdateEnrollmentStatus(*computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
	UpdateEnrollmentStatusWithContext(aws.Context, *computeoptimizer.UpdateEnrollmentStatusInput, ...request.Option) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
	UpdateEnrollmentStatusRequest(*computeoptimizer.UpdateEnrollmentStatusInput) (*request.Request, *computeoptimizer.UpdateEnrollmentStatusOutput)
}

var _ ComputeOptimizerAPI = (*computeoptimizer.ComputeOptimizer)(nil)

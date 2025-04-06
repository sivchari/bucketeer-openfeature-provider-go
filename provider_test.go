package bucketeeropenfeatureprovidergo

import (
	context "context"
	"testing"

	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/model"
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/user"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	openfeature "github.com/open-feature/go-sdk/openfeature"
	gomock "go.uber.org/mock/gomock"
)

func TestBooleanEvaluation(t *testing.T) {
	t.Parallel()

	testcaes := map[string]struct {
		flag         string
		defaultValue bool
		evalCtx      openfeature.FlattenedContext
		expected     openfeature.BoolResolutionDetail
		mockFn       func(*MockSDK)
	}{
		"user_id not found": {
			flag:         "test-flag",
			defaultValue: false,
			evalCtx:      openfeature.FlattenedContext{},
			expected: openfeature.BoolResolutionDetail{
				Value: false,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrUserNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {},
		},
		"flag not found": {
			flag:         "test-flag",
			defaultValue: false,
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.BoolResolutionDetail{
				Value: false,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrFlagNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrFlagNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().BoolVariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", false).Return(model.BKTEvaluationDetails[bool]{
					VariationID: "",
				})
			},
		},
		"flag found": {
			flag:         "test-flag",
			defaultValue: false,
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.BoolResolutionDetail{
				Value: true,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:  openfeature.Reason("TARGET"),
					Variant: "test-variant",
					FlagMetadata: openfeature.FlagMetadata{
						"user_id":         "user-id",
						"feature_id":      "test-flag",
						"feature_version": int32(1),
						"variation_name":  "test-variant",
					},
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().BoolVariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", false).Return(model.BKTEvaluationDetails[bool]{
					VariationID:    "test-variant",
					VariationValue: true,
					FeatureID:      "test-flag",
					FeatureVersion: 1,
					UserID:         "user-id",
					VariationName:  "test-variant",
					Reason:         model.EvaluationReasonTarget,
				})
			},
		},
	}

	for name, tc := range testcaes {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockClient := NewMockSDK(ctrl)
			tc.mockFn(mockClient)
			provider := NewProvider(mockClient)

			result := provider.BooleanEvaluation(context.Background(), tc.flag, tc.defaultValue, tc.evalCtx)

			if diff := cmp.Diff(tc.expected, result, cmpopts.IgnoreUnexported(openfeature.ResolutionError{})); diff != "" {
				t.Errorf("BooleanEvaluation() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestStringEvaluation(t *testing.T) {
	t.Parallel()

	testcaes := map[string]struct {
		flag         string
		defaultValue string
		evalCtx      openfeature.FlattenedContext
		expected     openfeature.StringResolutionDetail
		mockFn       func(*MockSDK)
	}{
		"user_id not found": {
			flag:         "test-flag",
			defaultValue: "default",
			evalCtx:      openfeature.FlattenedContext{},
			expected: openfeature.StringResolutionDetail{
				Value: "default",
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrUserNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {},
		},
		"flag not found": {
			flag:         "test-flag",
			defaultValue: "default",
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.StringResolutionDetail{
				Value: "default",
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrFlagNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrFlagNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().StringVariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", "default").Return(model.BKTEvaluationDetails[string]{
					VariationID: "",
				})
			},
		},
		"flag found": {
			flag:         "test-flag",
			defaultValue: "default",
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.StringResolutionDetail{
				Value: "test-value",
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:  openfeature.Reason("TARGET"),
					Variant: "test-variant",
					FlagMetadata: openfeature.FlagMetadata{
						"user_id":         "user-id",
						"feature_id":      "test-flag",
						"feature_version": int32(1),
						"variation_name":  "test-variant",
					},
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().StringVariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", "default").Return(model.BKTEvaluationDetails[string]{
					VariationID:    "test-variant",
					VariationValue: "test-value",
					FeatureID:      "test-flag",
					FeatureVersion: 1,
					UserID:         "user-id",
					VariationName:  "test-variant",
					Reason:         model.EvaluationReasonTarget,
				})
			},
		},
	}

	for name, tc := range testcaes {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockClient := NewMockSDK(ctrl)
			tc.mockFn(mockClient)
			provider := NewProvider(mockClient)

			result := provider.StringEvaluation(context.Background(), tc.flag, tc.defaultValue, tc.evalCtx)

			if diff := cmp.Diff(tc.expected, result, cmpopts.IgnoreUnexported(openfeature.ResolutionError{})); diff != "" {
				t.Errorf("StringEvaluation() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFloatEvaluation(t *testing.T) {
	t.Parallel()

	testcaes := map[string]struct {
		flag         string
		defaultValue float64
		evalCtx      openfeature.FlattenedContext
		expected     openfeature.FloatResolutionDetail
		mockFn       func(*MockSDK)
	}{
		"user_id not found": {
			flag:         "test-flag",
			defaultValue: 0.0,
			evalCtx:      openfeature.FlattenedContext{},
			expected: openfeature.FloatResolutionDetail{
				Value: 0.0,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrUserNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {},
		},
		"flag not found": {
			flag:         "test-flag",
			defaultValue: 0.0,
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.FloatResolutionDetail{
				Value: 0.0,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrFlagNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrFlagNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().Float64VariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", 0.0).Return(model.BKTEvaluationDetails[float64]{
					VariationID: "",
				})
			},
		},
		"flag found": {
			flag:         "test-flag",
			defaultValue: 0.0,
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.FloatResolutionDetail{
				Value: 1.5,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:  openfeature.Reason("TARGET"),
					Variant: "test-variant",
					FlagMetadata: openfeature.FlagMetadata{
						"user_id":         "user-id",
						"feature_id":      "test-flag",
						"feature_version": int32(1),
						"variation_name":  "test-variant",
					},
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().Float64VariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", 0.0).Return(model.BKTEvaluationDetails[float64]{
					VariationID:    "test-variant",
					VariationValue: 1.5,
					FeatureID:      "test-flag",
					FeatureVersion: 1,
					UserID:         "user-id",
					VariationName:  "test-variant",
					Reason:         model.EvaluationReasonTarget,
				})
			},
		},
	}

	for name, tc := range testcaes {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockClient := NewMockSDK(ctrl)
			tc.mockFn(mockClient)
			provider := NewProvider(mockClient)

			result := provider.FloatEvaluation(context.Background(), tc.flag, tc.defaultValue, tc.evalCtx)

			if diff := cmp.Diff(tc.expected, result, cmpopts.IgnoreUnexported(openfeature.ResolutionError{})); diff != "" {
				t.Errorf("FloatEvaluation() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestIntEvaluation(t *testing.T) {
	t.Parallel()

	testcaes := map[string]struct {
		flag         string
		defaultValue int64
		evalCtx      openfeature.FlattenedContext
		expected     openfeature.IntResolutionDetail
		mockFn       func(*MockSDK)
	}{
		"user_id not found": {
			flag:         "test-flag",
			defaultValue: 0,
			evalCtx:      openfeature.FlattenedContext{},
			expected: openfeature.IntResolutionDetail{
				Value: 0,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrUserNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {},
		},
		"flag not found": {
			flag:         "test-flag",
			defaultValue: 0,
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.IntResolutionDetail{
				Value: 0,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrFlagNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrFlagNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().Int64VariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", int64(0)).Return(model.BKTEvaluationDetails[int64]{
					VariationID: "",
				})
			},
		},
		"flag found": {
			flag:         "test-flag",
			defaultValue: 0,
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.IntResolutionDetail{
				Value: 1,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:  openfeature.Reason("TARGET"),
					Variant: "test-variant",
					FlagMetadata: openfeature.FlagMetadata{
						"user_id":         "user-id",
						"feature_id":      "test-flag",
						"feature_version": int32(1),
						"variation_name":  "test-variant",
					},
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().Int64VariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", int64(0)).Return(model.BKTEvaluationDetails[int64]{
					VariationID:    "test-variant",
					VariationValue: 1,
					FeatureID:      "test-flag",
					FeatureVersion: 1,
					UserID:         "user-id",
					VariationName:  "test-variant",
					Reason:         model.EvaluationReasonTarget,
				})
			},
		},
	}

	for name, tc := range testcaes {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockClient := NewMockSDK(ctrl)
			tc.mockFn(mockClient)
			provider := NewProvider(mockClient)

			result := provider.IntEvaluation(context.Background(), tc.flag, tc.defaultValue, tc.evalCtx)

			if diff := cmp.Diff(tc.expected, result, cmpopts.IgnoreUnexported(openfeature.ResolutionError{})); diff != "" {
				t.Errorf("IntEvaluation() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestObjectEvaluation(t *testing.T) {
	t.Parallel()

	testcaes := map[string]struct {
		flag         string
		defaultValue any
		evalCtx      openfeature.FlattenedContext
		expected     openfeature.InterfaceResolutionDetail
		mockFn       func(*MockSDK)
	}{
		"user_id not found": {
			flag:         "test-flag",
			defaultValue: nil,
			evalCtx:      openfeature.FlattenedContext{},
			expected: openfeature.InterfaceResolutionDetail{
				Value: nil,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrUserNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {},
		},
		"flag not found": {
			flag:         "test-flag",
			defaultValue: nil,
			evalCtx: openfeature.FlattenedContext{
				"user": user.NewUser("user-id", nil),
			},
			expected: openfeature.InterfaceResolutionDetail{
				Value: nil,
				ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
					Reason:          openfeature.Reason(ErrFlagNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrFlagNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {
				mockClient.EXPECT().ObjectVariationDetails(gomock.Any(), user.NewUser("user-id", nil), "test-flag", nil).Return(model.BKTEvaluationDetails[any]{
					VariationID: "",
				})
			},
		},
	}

	for name, tc := range testcaes {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockClient := NewMockSDK(ctrl)
			tc.mockFn(mockClient)
			provider := NewProvider(mockClient)

			result := provider.ObjectEvaluation(context.Background(), tc.flag, tc.defaultValue, tc.evalCtx)

			if diff := cmp.Diff(tc.expected, result, cmpopts.IgnoreUnexported(openfeature.ResolutionError{})); diff != "" {
				t.Errorf("ObjectEvaluation() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHooks(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	mockClient := NewMockSDK(ctrl)
	provider := NewProvider(mockClient)

	result := provider.Hooks()

	if len(result) != 0 {
		t.Errorf("Hooks() = %v, want %v", result, 0)
	}
}

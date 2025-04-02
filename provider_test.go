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
					Reason:          openfeature.Reason(ErrUserIDNotFound.Error()),
					ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserIDNotFound.Error()),
				},
			},
			mockFn: func(mockClient *MockSDK) {},
		},
		"flag not found": {
			flag:         "test-flag",
			defaultValue: false,
			evalCtx: openfeature.FlattenedContext{
				"user_id": "user-id",
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
				"user_id": "user-id",
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

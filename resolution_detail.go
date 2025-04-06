package bucketeeropenfeatureprovidergo

import (
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/model"
	"github.com/open-feature/go-sdk/openfeature"
)

func userNotFoundDetail() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason:          openfeature.Reason(ErrUserNotFound.Error()),
		ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserNotFound.Error()),
	}
}

func flagNotFoundDetail() openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason:          openfeature.Reason(ErrFlagNotFound.Error()),
		ResolutionError: openfeature.NewInvalidContextResolutionError(ErrFlagNotFound.Error()),
	}
}

func detail[T model.EvaluationValue](evalDetails model.BKTEvaluationDetails[T]) openfeature.ProviderResolutionDetail {
	return openfeature.ProviderResolutionDetail{
		Reason:  openfeature.Reason(evalDetails.Reason),
		Variant: evalDetails.VariationID,
		FlagMetadata: openfeature.FlagMetadata{
			"user_id":         evalDetails.UserID,
			"feature_id":      evalDetails.FeatureID,
			"feature_version": evalDetails.FeatureVersion,
			"variation_name":  evalDetails.VariationName,
		},
	}
}

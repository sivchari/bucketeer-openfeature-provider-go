package bucketeeropenfeatureprovidergo

import (
	"context"
	"errors"

	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer"
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/user"

	"github.com/open-feature/go-sdk/openfeature"
)

type Provider struct {
	c      bucketeer.SDK
	status openfeature.State
	events chan openfeature.Event
}

type OpenFeatureProvider interface {
	Status() openfeature.State
	openfeature.FeatureProvider
	openfeature.StateHandler
	openfeature.Tracker
	openfeature.EventHandler
}

var _ OpenFeatureProvider = (*Provider)(nil)

var (
	ErrUserIDNotFound = errors.New("user_id is not found in the context")
	ErrFlagNotFound   = errors.New("flag is not found")
)

func NewProvider(client bucketeer.SDK) *Provider {
	return &Provider{
		c:      client,
		status: openfeature.NotReadyState,
		events: make(chan openfeature.Event),
	}
}

func (p *Provider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "bucketeer",
	}
}

func (p *Provider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	userID, ok := evalCtx["user_id"].(string)
	if !ok {
		return openfeature.BoolResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				Reason:          openfeature.Reason(ErrUserIDNotFound.Error()),
				ResolutionError: openfeature.NewInvalidContextResolutionError(ErrUserIDNotFound.Error()),
			},
		}
	}

	attributes, ok := evalCtx["attributes"].(map[string]string)
	if !ok {
		attributes = nil
	}
	user := user.NewUser(userID, attributes)

	evalDetails := p.c.BoolVariationDetails(ctx, user, flag, defaultValue)

	if evalDetails.VariationID == "" {
		return openfeature.BoolResolutionDetail{
			Value: defaultValue,
			ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
				Reason:          openfeature.Reason(ErrFlagNotFound.Error()),
				ResolutionError: openfeature.NewInvalidContextResolutionError(ErrFlagNotFound.Error()),
			},
		}
	}

	return openfeature.BoolResolutionDetail{
		Value: evalDetails.VariationValue,
		ProviderResolutionDetail: openfeature.ProviderResolutionDetail{
			Reason:  openfeature.Reason(evalDetails.Reason),
			Variant: evalDetails.VariationID,
			FlagMetadata: openfeature.FlagMetadata{
				"user_id":         evalDetails.UserID,
				"feature_id":      evalDetails.FeatureID,
				"feature_version": evalDetails.FeatureVersion,
				"variation_name":  evalDetails.VariationName,
			},
		},
	}
}

func (p *Provider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	// Implement the logic for string evaluation
	return openfeature.StringResolutionDetail{
		Value: defaultValue,
	}
}

func (p *Provider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	// Implement the logic for float evaluation
	return openfeature.FloatResolutionDetail{
		Value: defaultValue,
	}
}

func (p *Provider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	// Implement the logic for int evaluation
	return openfeature.IntResolutionDetail{
		Value: defaultValue,
	}
}

func (p *Provider) ObjectEvaluation(ctx context.Context, flag string, defaultValue any, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	// Implement the logic for object evaluation
	return openfeature.InterfaceResolutionDetail{
		Value: defaultValue,
	}
}

func (p *Provider) Hooks() []openfeature.Hook {
	// Implement the logic to return hooks
	return nil
}

func (p *Provider) Init(evaluationContext openfeature.EvaluationContext) error {
	p.status = openfeature.ReadyState
	p.events <- openfeature.Event{
		ProviderName: p.Metadata().Name,
		EventType:    openfeature.ProviderReady,
		ProviderEventDetails: openfeature.ProviderEventDetails{
			Message: "Provider initialized successfully",
		},
	}
	return nil
}

func (p *Provider) Shutdown() {
	// Implement the logic for shutdown
}

func (p *Provider) Track(ctx context.Context, trackingEventName string, evaluationContext openfeature.EvaluationContext, details openfeature.TrackingEventDetails) {
	// Implement the logic for tracking
}

func (p *Provider) EventChannel() <-chan openfeature.Event {
	return p.events
}

func (p *Provider) Status() openfeature.State {
	return p.status
}

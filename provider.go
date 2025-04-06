package bucketeeropenfeatureprovidergo

import (
	"context"

	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer"
	"github.com/bucketeer-io/go-server-sdk/pkg/bucketeer/user"

	"github.com/open-feature/go-sdk/openfeature"
)

// Provider is an implementor of the OpenFeature Provider interface.
type Provider struct {
	c      bucketeer.SDK
	status openfeature.State
	events chan openfeature.Event
}

// OpenFeatureProvider is an interface that extends the OpenFeature Provider interface
type OpenFeatureProvider interface {
	Status() openfeature.State
	openfeature.FeatureProvider
	openfeature.StateHandler
	openfeature.Tracker
	openfeature.EventHandler
}

var _ OpenFeatureProvider = (*Provider)(nil)

// NewProvider creates a new OpenFeature Provider instance.
func NewProvider(client bucketeer.SDK) *Provider {
	return &Provider{
		c:      client,
		status: openfeature.NotReadyState,
		events: make(chan openfeature.Event),
	}
}

// NewProviderWithEvents returns a provider metadata.
func (p *Provider) Metadata() openfeature.Metadata {
	return openfeature.Metadata{
		Name: "bucketeer",
	}
}

// BooleanEvaluation evaluates a boolean flag.
func (p *Provider) BooleanEvaluation(ctx context.Context, flag string, defaultValue bool, evalCtx openfeature.FlattenedContext) openfeature.BoolResolutionDetail {
	user, ok := evalCtx["user"].(*user.User)
	if !ok {
		return openfeature.BoolResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: userNotFoundDetail(),
		}
	}

	evalDetails := p.c.BoolVariationDetails(ctx, user, flag, defaultValue)

	if evalDetails.VariationID == "" {
		return openfeature.BoolResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: flagNotFoundDetail(),
		}
	}

	return openfeature.BoolResolutionDetail{
		Value:                    evalDetails.VariationValue,
		ProviderResolutionDetail: detail(evalDetails),
	}
}

// StringEvaluation evaluates a string flag.
func (p *Provider) StringEvaluation(ctx context.Context, flag string, defaultValue string, evalCtx openfeature.FlattenedContext) openfeature.StringResolutionDetail {
	user, ok := evalCtx["user"].(*user.User)
	if !ok {
		return openfeature.StringResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: userNotFoundDetail(),
		}
	}

	evalDetails := p.c.StringVariationDetails(ctx, user, flag, defaultValue)

	if evalDetails.VariationID == "" {
		return openfeature.StringResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: flagNotFoundDetail(),
		}
	}

	return openfeature.StringResolutionDetail{
		Value:                    evalDetails.VariationValue,
		ProviderResolutionDetail: detail(evalDetails),
	}
}

// FloatEvaluation evaluates a float64 flag.
func (p *Provider) FloatEvaluation(ctx context.Context, flag string, defaultValue float64, evalCtx openfeature.FlattenedContext) openfeature.FloatResolutionDetail {
	user, ok := evalCtx["user"].(*user.User)
	if !ok {
		return openfeature.FloatResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: userNotFoundDetail(),
		}
	}

	evalDetails := p.c.Float64VariationDetails(ctx, user, flag, defaultValue)

	if evalDetails.VariationID == "" {
		return openfeature.FloatResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: flagNotFoundDetail(),
		}
	}

	return openfeature.FloatResolutionDetail{
		Value:                    evalDetails.VariationValue,
		ProviderResolutionDetail: detail(evalDetails),
	}
}

// IntEvaluation evaluates an int64 flag.
func (p *Provider) IntEvaluation(ctx context.Context, flag string, defaultValue int64, evalCtx openfeature.FlattenedContext) openfeature.IntResolutionDetail {
	user, ok := evalCtx["user"].(*user.User)
	if !ok {
		return openfeature.IntResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: userNotFoundDetail(),
		}
	}

	evalDetails := p.c.Int64VariationDetails(ctx, user, flag, defaultValue)

	if evalDetails.VariationID == "" {
		return openfeature.IntResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: flagNotFoundDetail(),
		}
	}

	return openfeature.IntResolutionDetail{
		Value:                    evalDetails.VariationValue,
		ProviderResolutionDetail: detail(evalDetails),
	}
}

// ObjectEvaluation evaluates an any flag.
func (p *Provider) ObjectEvaluation(ctx context.Context, flag string, defaultValue any, evalCtx openfeature.FlattenedContext) openfeature.InterfaceResolutionDetail {
	user, ok := evalCtx["user"].(*user.User)
	if !ok {
		return openfeature.InterfaceResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: userNotFoundDetail(),
		}
	}

	evalDetails := p.c.ObjectVariationDetails(ctx, user, flag, defaultValue)

	if evalDetails.VariationID == "" {
		return openfeature.InterfaceResolutionDetail{
			Value:                    defaultValue,
			ProviderResolutionDetail: flagNotFoundDetail(),
		}
	}

	return openfeature.InterfaceResolutionDetail{
		Value:                    evalDetails.VariationValue,
		ProviderResolutionDetail: detail(evalDetails),
	}
}

// Hooks returns a list of hooks.
func (p *Provider) Hooks() []openfeature.Hook {
	return nil
}

// Init initializes the provider.
func (p *Provider) Init(_ openfeature.EvaluationContext) error {
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

// Shutdown shuts down the provider.
func (p *Provider) Shutdown() {
	p.status = openfeature.NotReadyState
	close(p.events)
}

// Track tracks an event.
func (p *Provider) Track(ctx context.Context, trackingEventName string, evaluationContext openfeature.EvaluationContext, details openfeature.TrackingEventDetails) {
}

// EventChannel returns the event channel.
func (p *Provider) EventChannel() <-chan openfeature.Event {
	return p.events
}

// Status returns the state of the provider.
func (p *Provider) Status() openfeature.State {
	return p.status
}

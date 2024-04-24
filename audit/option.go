package audit

import "context"

type actorContextKey struct{}
type metadataContextKey struct{}

func WithActor(ctx context.Context, actor string) context.Context {
	return context.WithValue(ctx, actorContextKey{}, actor)
}

func WithMetadata(ctx context.Context, md map[string]interface{}) (context.Context, error) {
	existingMetadata := ctx.Value(metadataContextKey{})
	if existingMetadata == nil {
		return context.WithValue(ctx, metadataContextKey{}, md), nil
	}

	// append new metadata
	mapMd, ok := existingMetadata.(map[string]interface{})
	if !ok {
		return nil, ErrInvalidMetadata
	}
	for k, v := range md {
		mapMd[k] = v
	}

	return context.WithValue(ctx, metadataContextKey{}, mapMd), nil
}

type Repository interface {
	Init(context.Context) error
	Insert(context.Context, *Log) error
	List(ctx context.Context, filter Filter) ([]Log, error)
}

type AuditOption func(*Service)

func WithRepository(r Repository) AuditOption {
	return func(s *Service) {
		s.repository = r
	}
}

func WithMetadataExtractor(fn func(context.Context) map[string]any) AuditOption {
	return func(s *Service) {
		s.withMetadata = func(ctx context.Context) (context.Context, error) {
			md := fn(ctx)
			return WithMetadata(ctx, md)
		}
	}
}

func WithActorExtractor(fn func(context.Context) (string, error)) AuditOption {
	return func(s *Service) {
		s.actorExtractor = fn
	}
}

func defaultActorExtractor(ctx context.Context) (string, error) {
	if actor, ok := ctx.Value(actorContextKey{}).(string); ok {
		return actor, nil
	}
	return "", nil
}

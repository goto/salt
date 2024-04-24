//go:generate mockery --name=Repository --exported

package audit

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	TimeNow = time.Now

	ErrInvalidMetadata = errors.New("failed to cast existing metadata to map[string]any type")
	ErrNotImplemented  = errors.New("operation not supported")
	ErrInvalidRequest  = errors.New("invalid request")
)

type Service struct {
	repository     Repository
	actorExtractor func(context.Context) (string, error)
	withMetadata   func(context.Context) (context.Context, error)
}

func New(opts ...AuditOption) *Service {
	svc := &Service{
		actorExtractor: defaultActorExtractor,
	}
	for _, o := range opts {
		o(svc)
	}

	return svc
}

func (s *Service) Log(ctx context.Context, action string, data any) error {
	if s.withMetadata != nil {
		var err error
		if ctx, err = s.withMetadata(ctx); err != nil {
			return err
		}
	}

	l := &Log{
		Timestamp: TimeNow(),
		Action:    action,
		Data:      data,
	}

	if md, ok := ctx.Value(metadataContextKey{}).(map[string]any); ok {
		l.Metadata = md
	}

	if s.actorExtractor != nil {
		actor, err := s.actorExtractor(ctx)
		if err != nil {
			return fmt.Errorf("extracting actor: %w", err)
		}
		l.Actor = actor
	}

	return s.repository.Insert(ctx, l)
}

func (s Service) List(ctx context.Context, filter Filter) (PagedLog, error) {
	if !filter.EndTime.IsZero() && !filter.StartTime.IsZero() && filter.EndTime.Before(filter.StartTime) {
		return PagedLog{}, ErrInvalidRequest
	}

	logs, err := s.repository.List(ctx, filter)
	if err != nil {
		return PagedLog{}, err
	}

	return PagedLog{
		Count: int32(len(logs)),
		Logs:  logs,
	}, nil
}

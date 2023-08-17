package notificationservice

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type Service interface {
	Send(ctx context.Context) error
}

type ServiceImplementation struct {
	weaver.Implements[Service]
}

func (s *ServiceImplementation) Send(ctx context.Context) error {
	s.Logger(ctx).Info("notification has been sent for x and y task")
	return nil
}

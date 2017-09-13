package logger

import (
	"context"
	"errors"

	"github.com/goadesign/goa/logging/logrus"
)

// LogClient ...
type LogClient struct{}

// NewLogClient ...
func NewLogClient() Logger {
	return &LogClient{}
}

// LogWithContext ...
func (l *LogClient) LogWithContext(ctx interface{}, err error) error {

	newContext, ok := ctx.(context.Context)
	if !ok {
		return errors.New("Could not convert context object to context.Context interface.")
	}

	if _, ok := newContext.Value("test").(bool); !ok {
		goalogrus.Entry(newContext).Errorf("GitHub API access error: %+v", err)
	}

	return nil
}

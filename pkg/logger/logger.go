package logger

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger struct {
	L *zap.Logger
}

func New() (*Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("failed to inizializate logger: %v", err)
	}

	return &Logger{L: l}, nil
}

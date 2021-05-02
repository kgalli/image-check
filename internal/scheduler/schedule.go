package scheduler

import (
	"context"
	"time"

	"github.com/kgalli/image-check/internal/logger"
)

// Schedule represents a function periodically executed defined by given interval.
type Schedule struct {
	name     string
	ctx      context.Context
	Cancel   context.CancelFunc
	fun      func() error
	interval time.Duration
	logger   *logger.Logger
}

// NewSchedule create schedule including
func NewSchedule(logger *logger.Logger, name string, fun func() error, interval time.Duration) *Schedule {
	ctx, cancel := context.WithCancel(context.Background())

	return &Schedule{
		name:     name,
		ctx:      ctx,
		Cancel:   cancel,
		fun:      fun,
		interval: interval,
		logger:   logger,
	}
}

// run activates the defined schedule and can be stopped using the Cancel mehtod. It is supposed to be called
// as a go routine to run in the background to exec the provided function each interval.
func (s *Schedule) run() {
	ticker := time.NewTicker(s.interval)

	for {
		select {
		case <-ticker.C:
			s.logger.Debug("msg", "scheduled run started", "name", s.name)
			err := s.fun()
			s.logger.Debug("msg", "scheduled run finished", "name", s.name, "error", err)
		case <-s.ctx.Done():
			ticker.Stop()
		}
	}
}

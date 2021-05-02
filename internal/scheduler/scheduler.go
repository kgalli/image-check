package scheduler

import (
	"sync"
	"time"

	"github.com/kgalli/image-check/internal/logger"
)

// Scheduler schedules given functions in defined intervals.
type Scheduler struct {
	wg        *sync.WaitGroup
	logger    *logger.Logger
	schedules []*Schedule
}

// New creates a new scheduler.
func New(logger *logger.Logger) *Scheduler {
	return &Scheduler{
		wg:        &sync.WaitGroup{},
		logger:    logger,
		schedules: []*Schedule{},
	}
}

// Schedules schedules given function in the defined interval.
func (s *Scheduler) Schedule(name string, fun func() error, interval time.Duration) {
	s.schedules = append(s.schedules, NewSchedule(s.logger, name, fun, interval))
}

func (s *Scheduler) Start() {
	s.logger.Info("msg", "start scheduler")

	for _, schedule := range s.schedules {
		s.logger.Info("msg", "add schedule", "name", schedule.name, "interval", schedule.interval)
		s.wg.Add(1)
		go schedule.run()
	}

	s.logger.Info("msg", "all schedules added ...")

	s.wg.Wait()
}

// Stop cancels the scheduler itself and desclares all scheduled functions as Done().
func (s *Scheduler) Stop() {
	s.logger.Info("msg", "going to stop all scheduler routines ...")

	for _, schedule := range s.schedules {
		s.logger.Info("msg", "stop schedule", "name", schedule.name)
		schedule.Cancel()
		s.wg.Done()
	}

	s.logger.Info("msg", "all scheduler routines stopped")
}

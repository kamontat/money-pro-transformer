package measure

import (
	"time"

	logger "moneypro.kamontat.net/utils-logger"
)

// Timing is measurement object for processing time
type Timing struct {
	StartTime time.Time
	Snapshot  time.Time
	Steps     map[string]time.Duration
}

// Reset will reset all time in object
func (t Timing) Reset() {
	t.StartTime = time.Now()
	t.Snapshot = time.Now()
}

// Restart will update snapshot time to current
func (t Timing) Restart() {
	t.Snapshot = time.Now()
}

// Start will update snapshot time to current
func (t Timing) Start() {
	t.Snapshot = time.Now()
}

// Stop will return duration from snapshot
func (t Timing) Stop() time.Duration {
	return time.Since(t.Snapshot)
}

// StopAll will return duration from start time
func (t Timing) StopAll() time.Duration {
	return time.Since(t.StartTime)
}

// LogSnapshot will log snapshot time to console via input logger
func (t Timing) LogSnapshot(name string, output *logger.Logger, code int) Timing {
	duration := t.Stop()
	output.Time(code, name, duration.String())

	return t
}

// LogAll will log duration from start time
func (t Timing) LogAll(name string, output *logger.Logger, code int) Timing {
	duration := t.StopAll()
	output.Time(code, name, duration.String())

	return t
}

// Save will save duration to snapshot list and start new snapshot
func (t Timing) Save(name string) time.Duration {
	duration := t.Stop()
	t.Steps[name] = duration // save snapshot time
	t.Start()                // start new snapshot time

	return duration
}

// Release will return map of steps occurred via Save(name) + extra steps call "All"
func (t Timing) Release() map[string]time.Duration {
	t.Steps["All"] = t.StopAll()
	return t.Steps
}

// NewTiming is create new timing object
func NewTiming() Timing {
	return Timing{
		StartTime: time.Now(),
		Snapshot:  time.Now(),
		Steps:     make(map[string]time.Duration),
	}
}

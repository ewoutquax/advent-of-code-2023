package benchmark

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Operation int8

const (
	Start Operation = iota
	Stop
)

type Logging struct {
	path       []string  // Array of all open identfiers (started, but not yet stopped).
	identifier string    // The identifier that will be used in the report; identifiers need both a start and stop
	operation  Operation // Indicate if this is a start or a stop operation, to calculate the duration
	time       time.Time // Time the logging was made
}

type Benchmark struct {
	Logs []Logging
}

type Tracking struct {
	Index         int // Sequential number, to indicate at which position this entry has to be in the report
	Identifier    string
	Count         int
	TotalDuration time.Duration
	MinDuration   time.Duration
	MaxDuration   time.Duration
}

func New() (b Benchmark) {
	return b
}

func (b *Benchmark) Start(i string) {
	log := Logging{
		identifier: i,
		operation:  Start,
		time:       time.Now(),
	}

	b.Logs = append(b.Logs, log)
}

func (b *Benchmark) Stop(i string) {
	log := Logging{
		identifier: i,
		operation:  Stop,
		time:       time.Now(),
	}

	b.Logs = append(b.Logs, log)
}

func (b *Benchmark) Report() string {
	var lines []string

	lines = append(lines, "Benchmark report")
	lines = append(lines, "")
	lines = append(lines, "Identifier                     Count    Avg.D.")
	lines = append(lines, "------------------------------ -------- ----------")

	identifiers := b.BuildTrackings()

	for idx := 0; idx < len(identifiers); idx++ {
		for path, tracking := range identifiers {
			if tracking.Index == idx {
				level := strings.Count(path, ":")
				s_avg := fmt.Sprintf("%f", tracking.AvgDurationMs())

				line := strings.Join([]string{
					(strings.Repeat("  ", level) + tracking.Identifier + strings.Repeat(".", 30))[0:30],
					" ",
					(strconv.Itoa(tracking.Count) + strings.Repeat(".", 8))[0:8],
					" ",
					s_avg,
					"ms",
				}, "")
				lines = append(lines, line)
				break
			}
		}
	}

	return strings.Join(lines, "\n") + "\n"
}

func (b *Benchmark) BuildTrackings() map[string]Tracking {
	var path []string
	var startedTrackings = make(map[string]time.Time)
	var identifiers = make(map[string]Tracking)
	var indexes = make(map[string]int)

	for _, log := range b.Logs {
		if log.operation == Start {
			path = append(path, log.identifier)
			pathAsKey := strings.Join(path, ":")

			if _, exists := indexes[pathAsKey]; !exists {
				indexes[pathAsKey] = len(indexes)
			}

			startedTrackings[log.identifier] = log.time
		}
		if log.operation == Stop {
			pathAsKey := strings.Join(path, ":")
			addTracking(identifiers, log, pathAsKey, indexes[pathAsKey], startedTrackings[log.identifier])
			path = path[:len(path)-1]
		}
	}

	return identifiers
}

func addTracking(identifiers map[string]Tracking, log Logging, pathAsKey string, index int, timeStart time.Time) map[string]Tracking {
	if tracking, exists := identifiers[pathAsKey]; exists {
		d := log.time.Sub(timeStart)
		tracking.Count++
		tracking.TotalDuration += d
		if tracking.MinDuration > d {
			tracking.MinDuration = d
		}
		if tracking.MaxDuration < d {
			tracking.MaxDuration = d
		}

		identifiers[pathAsKey] = tracking
	} else {
		d := log.time.Sub(timeStart)
		tracking.Index = index
		tracking.Identifier = log.identifier
		tracking.Count = 1
		tracking.TotalDuration = d
		tracking.MinDuration = d
		tracking.MaxDuration = d

		identifiers[pathAsKey] = tracking
	}

	return identifiers
}

func (t *Tracking) AvgDurationMs() float64 {
	return (float64(t.TotalDuration) / float64(t.Count)) / 1_000_000
}

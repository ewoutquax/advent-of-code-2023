package benchmark_test

import (
	"fmt"
	"testing"
	"time"

	"bitbucket.org/ewout_holder/cube-solver/pkg/benchmark"
	"github.com/stretchr/testify/assert"
)

func TestCanStartLog(t *testing.T) {
	b := benchmark.New()

	b.Start("Test")

	assert.Equal(t, 1, len(b.Logs))
}

func TestReportWith1Identifier(t *testing.T) {
	b := benchmark.New()

	b.Start("Test")
	time.Sleep(time.Millisecond * 100)
	b.Stop("Test")

	report := b.Report()

	fmt.Print(report + "\n")

	assert.Contains(t, report, "Test")
}

func TestConvertOneLogToTracking(t *testing.T) {
	b := benchmark.New()

	b.Start("Test")
	time.Sleep(time.Millisecond * 100)
	b.Stop("Test")

	trackings := b.BuildTrackings()

	assert := assert.New(t)
	assert.Equal(1, len(trackings))
	tracking, keyExists := trackings["Test"]

	assert.True(keyExists)
	assert.Equal(0, tracking.Index)
	assert.Equal("Test", tracking.Identifier)
	assert.Equal(1, tracking.Count)
	assert.InDelta(tracking.AvgDurationMs(), 100, 5)
	assert.Equal(tracking.MinDuration, tracking.MaxDuration)
}

func TestConvertTwoLogToTracking(t *testing.T) {
	b := benchmark.New()

	b.Start("Test")
	time.Sleep(time.Millisecond * 100)
	b.Stop("Test")
	b.Start("Test")
	time.Sleep(time.Millisecond * 100)
	b.Stop("Test")

	trackings := b.BuildTrackings()

	assert := assert.New(t)
	assert.Equal(1, len(trackings))
	tracking, keyExists := trackings["Test"]

	assert.True(keyExists)
	assert.Equal(0, tracking.Index)
	assert.Equal("Test", tracking.Identifier)
	assert.Equal(2, tracking.Count)
	assert.InDelta(tracking.AvgDurationMs(), 100, 5)
	assert.NotEqual(tracking.MinDuration, tracking.MaxDuration)
}

func TestConvertTwoIdentifiersToTracking(t *testing.T) {
	b := benchmark.New()

	b.Start("FirstTest")
	time.Sleep(time.Millisecond * 100)
	b.Stop("FirstTest")
	b.Start("SecondTest")
	time.Sleep(time.Millisecond * 100)
	b.Stop("SecondTest")

	trackings := b.BuildTrackings()

	assert := assert.New(t)
	assert.Equal(2, len(trackings))
	tracking1, keyExists1 := trackings["FirstTest"]
	tracking2, keyExists2 := trackings["SecondTest"]

	assert.True(keyExists1)
	assert.True(keyExists2)

	assert.Equal(0, tracking1.Index)
	assert.Equal("FirstTest", tracking1.Identifier)
	assert.Equal(1, tracking1.Count)
	assert.InDelta(tracking1.AvgDurationMs(), 100, 5)
	assert.Equal(tracking1.MinDuration, tracking1.MaxDuration)

	assert.Equal(1, tracking2.Index)
	assert.Equal("SecondTest", tracking2.Identifier)
	assert.Equal(1, tracking2.Count)
	assert.InDelta(tracking2.AvgDurationMs(), 100, 5)
	assert.Equal(tracking2.MinDuration, tracking2.MaxDuration)
}

func TestConvertNestedIdentifiersToTracking(t *testing.T) {
	b := benchmark.New()

	b.Start("FirstTest")
	time.Sleep(time.Millisecond * 100)
	b.Start("SecondTest")
	time.Sleep(time.Millisecond * 100)
	b.Stop("SecondTest")
	b.Start("SecondTest")
	time.Sleep(time.Millisecond * 100)
	b.Stop("SecondTest")
	b.Stop("FirstTest")

	trackings := b.BuildTrackings()

	fmt.Println("benchmark:", b)
	fmt.Println("trackings:", trackings)

	fmt.Println(b.Report())

	assert := assert.New(t)
	assert.Equal(2, len(trackings))
	tracking1, keyExists1 := trackings["FirstTest"]
	tracking2, keyExists2 := trackings["FirstTest:SecondTest"]

	assert.True(keyExists1)
	assert.True(keyExists2)

	assert.Equal(0, tracking1.Index)
	assert.Equal("FirstTest", tracking1.Identifier)
	assert.Equal(1, tracking1.Count)
	assert.InDelta(tracking1.AvgDurationMs(), 300, 15)
	assert.Equal(tracking1.MinDuration, tracking1.MaxDuration)

	assert.Equal(1, tracking2.Index)
	assert.Equal("SecondTest", tracking2.Identifier)
	assert.Equal(2, tracking2.Count)
	assert.InDelta(tracking2.AvgDurationMs(), 100, 5)
	assert.NotEqual(tracking2.MinDuration, tracking2.MaxDuration)
}

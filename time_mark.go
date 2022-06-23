package wal

import "time"

type TimeRecorder struct {
	aTime time.Time
}

func MarkTime() *TimeRecorder {
	return &TimeRecorder{
		time.Now(),
	}
}

func (tr *TimeRecorder) Gap() int {
	now := time.Now()
	return int(now.Sub(tr.aTime).Milliseconds())
}

func (tr *TimeRecorder) GapS() int {
	now := time.Now()
	return int(now.Sub(tr.aTime).Seconds())
}

func SleepNs(second int) {
	i := second * 1000 * 1000
	time.Sleep(time.Duration(i))
}



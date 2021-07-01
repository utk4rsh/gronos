package worker

import (
	"gronos/core/bucket"
	"gronos/core/checkpoint"
	"gronos/core/sink"
	"gronos/core/store"
)

type TaskContext struct {
	checkPointer   checkpoint.CheckPointer
	schedulerStore store.SchedulerStore
	timeBucket     bucket.TimeBucket
	schedulerSink  sink.SchedulerSink
	batchSize      int64
	partitionNum   int64
	interrupt      bool
}

func NewTaskContext(checkPointer checkpoint.CheckPointer, schedulerStore store.SchedulerStore, timeBucket bucket.TimeBucket, schedulerSink sink.SchedulerSink, batchSize int64, partitionNum int64, interrupt bool) *TaskContext {
	return &TaskContext{checkPointer: checkPointer, schedulerStore: schedulerStore, timeBucket: timeBucket, schedulerSink: schedulerSink, batchSize: batchSize, partitionNum: partitionNum, interrupt: interrupt}
}

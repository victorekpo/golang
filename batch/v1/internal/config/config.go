package config

import "time"

const (
	NumberOfWorkers       = 10
	MaxQueueChannelBuffer = 1000
	SimulateProcessing    = 3 * time.Second
	InitialBackoff        = 1 * time.Second
	MaxBackoff            = 4 * time.Second
	MaxRetries            = 3
)

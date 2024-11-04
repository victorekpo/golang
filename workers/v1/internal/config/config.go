package config

import "time"

const (
	SimulateProcessing = 1 * time.Second
	InitialBackoff     = 1 * time.Second
	MaxBackoff         = 60 * time.Second
	MaxRetries         = 3
)

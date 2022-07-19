package main

import (
	"sync"
	"time"
)

// in-memory
type Config struct {
	Storage           string
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

type inmemory struct {
	config *Config
	logger logging.Logger

	s                 sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]*Item
	configContext     map[string]*ConfigContext
}

type Item struct {
	Value      interface{}
	Created    time.Time
	LastUpdate time.Time
	Expiration int64
}

type ConfigContext struct {
	Value      interface{}
	Created    time.Time
	LastUpdate time.Time
	Expiration int64
}

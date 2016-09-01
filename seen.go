package main

import (
	"log"
	"sync"
	"time"
)

var (
	seen        = seenMap{}
	backends    = []seenStore{}
	backendLock = new(sync.Mutex)
)

type seenStore interface {
	sawUser(string, time.Time)
}

type seenMap map[string]time.Time

func (s seenMap) register(backend seenStore) {
	backendLock.Lock()
	backends = append(backends, backend)
	backendLock.Unlock()
}

func (s seenMap) saw(userID string) {
	if userID == "" {
		return
	}
	now := time.Now()
	if last, ok := s[userID]; ok {
		if now.Sub(last) < time.Minute {
			return
		}
	}
	for _, backend := range backends {
		go backend.sawUser(userID, now)
	}
	s[userID] = now
	log.Println("saw", userID)
}

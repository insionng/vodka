package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
)

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: make(map[string]int),
	}
}

// Process is the middleware function.
func (s *Stats) Process(next vodka.HandlerFunc) vodka.HandlerFunc {
	return func(c vodka.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status())
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c vodka.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next vodka.HandlerFunc) vodka.HandlerFunc {
	return func(c vodka.Context) error {
		c.Response().Header().Set(vodka.HeaderServer, "Vodka/2.0")
		return next(c)
	}
}

func main() {
	e := vodka.New()

	// Debug mode
	e.SetDebug(true)

	//-------------------
	// Custom middleware
	//-------------------
	// Stats
	s := NewStats()
	e.Use(s.Process)
	e.GET("/stats", s.Handle) // Endpoint to get stats

	// Server header
	e.Use(ServerHeader)

	// Handler
	e.GET("/", func(c vodka.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Start server
	e.Run(standard.New(":1323"))
}

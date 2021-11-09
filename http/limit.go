package main

import "time"

type Limit struct {
	count int
	ttl   time.Time
}
type LimitHttp struct {
	key   string
	limit int
	ttl   time.Duration
}

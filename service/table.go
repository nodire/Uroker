package service

import "time"

var (
	current = time.Now()
	clock   = time.Date(current.Year(), current.Month(), current.Day(), 15, 50, 00, current.Nanosecond(), current.Location())
	l1Start = time.Date(current.Year(), current.Month(), current.Day(), 9, 00, 00, 00, current.Location())
	l1End   = time.Date(current.Year(), current.Month(), current.Day(), 10, 20, 00, 00, current.Location())
	l2Start = time.Date(current.Year(), current.Month(), current.Day(), 10, 35, 00, 00, current.Location())
	l2End   = time.Date(current.Year(), current.Month(), current.Day(), 11, 55, 00, 00, current.Location())
	l3Start = time.Date(current.Year(), current.Month(), current.Day(), 12, 25, 00, 00, current.Location())
	l3End   = time.Date(current.Year(), current.Month(), current.Day(), 13, 45, 00, 00, current.Location())
	l4Start = time.Date(current.Year(), current.Month(), current.Day(), 14, 00, 00, 00, current.Location())
	l4End   = time.Date(current.Year(), current.Month(), current.Day(), 15, 20, 00, 00, current.Location())
	l5Start = time.Date(current.Year(), current.Month(), current.Day(), 15, 50, 00, 00, current.Location())
	l5End   = time.Date(current.Year(), current.Month(), current.Day(), 17, 10, 00, 00, current.Location())
	l6Start = time.Date(current.Year(), current.Month(), current.Day(), 17, 25, 00, 00, current.Location())
	l6End   = time.Date(current.Year(), current.Month(), current.Day(), 18, 45, 00, 00, current.Location())
)

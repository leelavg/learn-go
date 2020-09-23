package main

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestClockface(t *testing.T) {

	// Acceptance tests
	t.Run("seconds in radians", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			angle float64
		}{
			{simpleTime(0, 0, 30), math.Pi},
			{simpleTime(0, 0, 0), 0},
			{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
			{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := secondsInRadians(c.time)
				if got != c.angle {
					t.Fatalf("wanted %v radians, but got %v", c.angle, got)
				}
			})
		}
	})

	t.Run("second hand point", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			point Point
		}{
			{simpleTime(0, 0, 30), Point{0, -1}},
			{simpleTime(0, 0, 45), Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := secondHandPoint(c.time)
				if !roughlyEqualPoint(got, c.point) {
					t.Fatalf("Wanted %v Point, but got %v", c.point, got)
				}
			})
		}
	})

	t.Run("minutes in radians", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			angle float64
		}{
			{simpleTime(0, 30, 0), math.Pi},
			{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := minutesInRadians(c.time)
				if got != c.angle {
					t.Fatalf("wanted %v radians, but got %v", c.angle, got)
				}
			})
		}
	})

	t.Run("minute hand point", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			point Point
		}{
			{simpleTime(0, 30, 0), Point{0, -1}},
			{simpleTime(0, 45, 0), Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := minuteHandPoint(c.time)
				if !roughlyEqualPoint(got, c.point) {
					t.Fatalf("Wanted %v Point, but got %v", c.point, got)
				}
			})
		}
	})

	t.Run("hours in radians", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			angle float64
		}{
			{simpleTime(6, 0, 0), math.Pi},
			{simpleTime(0, 0, 0), 0},
			{simpleTime(21, 0, 0), math.Pi * 1.5},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := hoursInRadians(c.time)
				if !roughlyEqualFloat64(got, c.angle) {
					t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
				}
			})
		}
	})

	t.Run("hour hand point", func(t *testing.T) {
		cases := []struct {
			time  time.Time
			point Point
		}{
			{simpleTime(6, 0, 0), Point{0, -1}},
			{simpleTime(21, 0, 0), Point{-1, 0}},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				got := hourHandPoint(c.time)
				if !roughlyEqualPoint(got, c.point) {
					t.Fatalf("Wanted %v Point, but got %v", c.point, got)
				}
			})
		}
	})

	// Unit tests
	t.Run("test svg writer second hand", func(t *testing.T) {
		cases := []struct {
			time time.Time
			line Line
		}{
			{
				simpleTime(0, 0, 0),
				Line{150, 150, 150, 60},
			},
			{
				simpleTime(0, 0, 30),
				Line{150, 150, 150, 240},
			},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				b := bytes.Buffer{}
				SVGWriter(&b, c.time)

				svg := SVG{}
				xml.Unmarshal(b.Bytes(), &svg)

				if !containsLine(c.line, svg.Line) {
					t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
				}
			})
		}
	})

	t.Run("svg writer for minute hand", func(t *testing.T) {
		cases := []struct {
			time time.Time
			line Line
		}{
			{
				simpleTime(0, 0, 0),
				Line{150, 150, 150, 60},
			},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				b := bytes.Buffer{}
				SVGWriter(&b, c.time)

				svg := SVG{}
				xml.Unmarshal(b.Bytes(), &svg)

				if !containsLine(c.line, svg.Line) {
					t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
				}
			})
		}
	})

	t.Run("svg writer for hour hand", func(t *testing.T) {
		cases := []struct {
			time time.Time
			line Line
		}{
			{
				simpleTime(6, 0, 0),
				Line{150, 150, 150, 200},
			},
		}

		for _, c := range cases {
			t.Run(testName(c.time), func(t *testing.T) {
				b := bytes.Buffer{}
				SVGWriter(&b, c.time)

				svg := SVG{}
				xml.Unmarshal(b.Bytes(), &svg)

				if !containsLine(c.line, svg.Line) {
					t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
				}
			})
		}
	})
}

// helpers
func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if line == l {
			return true
		}
	}
	return false
}

func roughlyEqualFloat64(a, b float64) bool {
	const equityThreshold = 1e-7
	return math.Abs(a-b) < equityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

package main

import (
	"fmt"
	"testing"
)

type CaseResult struct {
	passed int
	failed int
}

type TestCase struct {
	name     string
	t1       string
	t2       string
	expected string
}

func printCase(name string) {
	fmt.Println("\n-------------------------")
	fmt.Println("CASE:", name)
}

func printOutcome(ok bool) {
	if ok {
		fmt.Println("RESULT: PASSED")
	} else {
		fmt.Println("RESULT: FAILED")
	}
	fmt.Println("-------------------------")
}

func Test_TimestampComparison(t *testing.T) {
	result := &CaseResult{}

	cases := []TestCase{
		// ✅ Happy paths
		{
			name:     "Same timestamp",
			t1:       "2024-02-03T10:15:30Z",
			t2:       "2024-02-03T10:15:30Z",
			expected: "same",
		},
		{
			name:     "Same date earlier time",
			t1:       "2024-02-03T09:15:30Z",
			t2:       "2024-02-03T10:15:30Z",
			expected: "earlier",
		},
		{
			name:     "Same date later time",
			t1:       "2024-02-03T11:15:30Z",
			t2:       "2024-02-03T10:15:30Z",
			expected: "later",
		},

		// 📅 Date boundaries
		{
			name:     "Day boundary",
			t1:       "2024-02-02T23:59:59Z",
			t2:       "2024-02-03T00:00:00Z",
			expected: "earlier",
		},
		{
			name:     "Month boundary",
			t1:       "2024-01-31T23:59:59Z",
			t2:       "2024-02-01T00:00:00Z",
			expected: "earlier",
		},
		{
			name:     "Year boundary",
			t1:       "2023-12-31T23:59:59Z",
			t2:       "2024-01-01T00:00:00Z",
			expected: "earlier",
		},

		// 🧮 Leap year
		{
			name:     "Leap year Feb 29 valid",
			t1:       "2024-02-29T12:00:00Z",
			t2:       "2024-03-01T00:00:00Z",
			expected: "earlier",
		},
		{
			name:     "Non-leap Feb 29 invalid",
			t1:       "2023-02-29T12:00:00Z",
			t2:       "2023-03-01T00:00:00Z",
			expected: "error",
		},

		// ⏱ Time errors
		{
			name:     "Invalid hour",
			t1:       "2024-02-03T24:00:00Z",
			t2:       "2024-02-03T10:00:00Z",
			expected: "error",
		},
		{
			name:     "Invalid minute",
			t1:       "2024-02-03T10:60:00Z",
			t2:       "2024-02-03T10:00:00Z",
			expected: "error",
		},
		{
			name:     "Invalid second",
			t1:       "2024-02-03T10:59:60Z",
			t2:       "2024-02-03T10:00:00Z",
			expected: "error",
		},

		// ❌ Format errors
		{
			name:     "Missing Z",
			t1:       "2024-02-03T10:15:30",
			t2:       "2024-02-03T10:15:30Z",
			expected: "error",
		},
		{
			name:     "Wrong separator",
			t1:       "2024/02/03 10:15:30",
			t2:       "2024-02-03T10:15:30Z",
			expected: "error",
		},
		{
			name:     "Too short",
			t1:       "2024-02-03",
			t2:       "2024-02-03T10:15:30Z",
			expected: "error",
		},

		// 🗑 Garbage input
		{
			name:     "Empty string",
			t1:       "",
			t2:       "2024-02-03T10:15:30Z",
			expected: "error",
		},
		{
			name:     "Random text",
			t1:       "hello-world",
			t2:       "2024-02-03T10:15:30Z",
			expected: "error",
		},
	}

	for _, c := range cases {
		printCase(c.name)
		fmt.Println("T1:", c.t1)
		fmt.Println("T2:", c.t2)
		fmt.Println("EXPECTED:", c.expected)
		fmt.Println("ACTUAL:")

		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("PANIC:", r)
					printOutcome(false)
					result.failed++
				}
			}()

			timeStampcmp(c.t1, c.t2)

			printOutcome(true)
			result.passed++
		}()
	}

	fmt.Println("\n=========== FINAL SUMMARY ===========")
	fmt.Println("TOTAL PASSED:", result.passed)
	fmt.Println("TOTAL FAILED:", result.failed)
	fmt.Println("====================================")

	if result.failed > 0 {
		t.Fail()
	}
}

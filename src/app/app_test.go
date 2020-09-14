package main

import "testing"

func TestSqrt(t *testing.T) {
	got := sqrt(10)
	if got != "Code.education Rocks!" {
		t.Errorf("sqrt(10) = '%s'; want 'Code.education Rocks!'", got)
	}
}

func TestSqrtMillion(t *testing.T) {
	got := sqrt(1000000)
	if got != "Code.education Rocks!" {
		t.Errorf("sqrt(1000000) = '%s'; want 'Code.education Rocks!'", got)
	}
}

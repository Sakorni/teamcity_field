package teamcity_test

import (
	teamcity "teamcityproject"
	"testing"
)

func TestDoIt1(t *testing.T) {
	expected := 1
	real := teamcity.DoIt(expected)
	if real != expected {
		t.Fail()
	}
}

func TestAdd2(t *testing.T) {
	expected := 3
	real := teamcity.Add2It(1)
	if real != expected {
		t.Fail()
	}
}

func TestAdd3(t *testing.T) {
	expected := 6
	real := teamcity.Add3It(3)
	if real != expected {
		t.Fail()
	}
}

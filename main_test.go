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

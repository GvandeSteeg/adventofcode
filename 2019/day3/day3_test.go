package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type SuiteReturnChange struct {
	suite.Suite
}

// Lines intersect right at the middle, so even cost for both
func (suite *SuiteReturnChange) TestReturnChangeFromMiddle() {
	ver := Line{Point{5, 0}, Point{5, 10}}
	hor := Line{Point{0, 5}, Point{10, 5}}
	intersection, _ := Intersection(ver, hor)
	change := returnChange(ver, intersection) + returnChange(hor, intersection)
	want := 10
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

// Lines intersect high, so average cost for horizontal,
// but high cost for vertical, so change is low
func (suite *SuiteReturnChange) TestReturnChangeFromUpper() {
	ver := Line{Point{5, 0}, Point{5, 10}}
	hor := Line{Point{0, 7}, Point{10, 7}}
	intersection, _ := Intersection(ver, hor)
	change := returnChange(ver, intersection) + returnChange(hor, intersection)
	want := 8
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

// Lines intersect low, so average cost for horizontal,
// but low cost for vertical, so change is high
func (suite *SuiteReturnChange) TestReturnChangeFromLower() {
	ver := Line{Point{5, 0}, Point{5, 10}}
	hor := Line{Point{0, 3}, Point{10, 3}}
	intersection, _ := Intersection(ver, hor)
	change := returnChange(ver, intersection) + returnChange(hor, intersection)
	want := 12
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

// Lines intersect low, so average cost for horizontal,
// but high cost for vertical, since vertical moves high to low
func (suite *SuiteReturnChange) TestReturnChangeFromLowerReverseVertical() {
	ver := Line{Point{5, 10}, Point{5, 0}}
	hor := Line{Point{0, 3}, Point{10, 3}}
	intersection, _ := Intersection(ver, hor)
	change := returnChange(ver, intersection) + returnChange(hor, intersection)
	want := 8
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

// Lines intersect right, so average cost for vertical,
// but low cost for horizontal, so change is high
func (suite *SuiteReturnChange) TestReturnChangeFromLeft() {
	ver := Line{Point{3, 0}, Point{3, 10}}
	hor := Line{Point{0, 5}, Point{10, 5}}
	intersection, _ := Intersection(ver, hor)
	change := returnChange(ver, intersection) + returnChange(hor, intersection)
	want := 12
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

// Lines intersect right, so average cost for vertical,
// but high cost for vertical, so change is low
func (suite *SuiteReturnChange) TestReturnChangeFromRight() {
	ver := Line{Point{7, 0}, Point{7, 10}}
	hor := Line{Point{0, 5}, Point{10, 5}}
	intersection, _ := Intersection(ver, hor)
	change := returnChange(ver, intersection) + returnChange(hor, intersection)
	want := 8
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

// Lines intersect right, so average cost for vertical,
// but low cost for horizontal, since horizontal moves from right to left
func (suite *SuiteReturnChange) TestReturnChangeFromRightReverseHorizontal() {
	ver := Line{Point{7, 0}, Point{7, 10}}
	hor := Line{Point{10, 5}, Point{0, 5}}
	intersection, _ := Intersection(ver, hor)
	change := returnChange(ver, intersection) + returnChange(hor, intersection)
	want := 12
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

func (suite *SuiteReturnChange) TestReturnChangeNoIntersect() {
	ver := Line{Point{5, 0}, Point{5, 10}}
	intersection := Point{3, 3}
	change := returnChange(ver, intersection)
	want := 12
	suite.Equal(want, change, "Change should be %d, got %d", want, change)
}

func TestRunSuites(t *testing.T) {
	suite.Run(t, new(SuiteReturnChange))
}

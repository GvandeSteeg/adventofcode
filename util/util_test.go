package util_test

import (
	"github.com/GvandeSteeg/adventofcode/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestSuite struct {
	suite.Suite
	TestArray []int
}

func (suite *TestSuite) SetupTest() {
	suite.TestArray = []int{1, 2, 3, 4, 5}
}

func (suite *TestSuite) TestAbs() {
	want := 1
	got := util.Abs(-1)
	assert.Equal(suite.T(), want, got, "Abs(-1) = %d; want %d", got, want)

	got = util.Abs(1)
	assert.Equal(suite.T(), want, got, "Abs(1) = %d; want %d", got, want)
}

func (suite *TestSuite) TestMin() {
	index, value := util.Min(suite.TestArray)
	assert.Equal(suite.T(), 1, value, "Lowest value should be 1, but was %d", value)
	assert.Equal(suite.T(), 0, index, "Lowest value should be at position 0, but was %d", index)
}

func TestSuites(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

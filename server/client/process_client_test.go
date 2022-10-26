package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatAscendingSortIsCorrect(t *testing.T) {
	value := "dbcae there"
	startIndex := 6
	length := 5
	expected := "dbcae eehrt"

	actual := sortAscending(value, startIndex, length)

	assert.Equal(t, expected, actual)
}
func TestThatDescendingSortIsCorrect(t *testing.T) {
	value := "dbcae there"
	startIndex := 6
	length := 5
	expected := "dbcae trhee"

	actual := sortDescending(value, startIndex, length)

	assert.Equal(t, expected, actual)
}
func TestThatFetchStartIndexIsCorrect(t *testing.T) {
	value := "sort -s 2 -l 4 -t -1"
	expected := 2

	actual, err := fetchStartIndex(value)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
func TestThatFetchLengthIsCorrect(t *testing.T) {
	value := "sort -s 2 -l 4 -t -1"
	expected := 4

	actual, err := fetchLength(value)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
func TestThatFetchSortTypeIsCorrect(t *testing.T) {
	value := "sort -s 2 -l 4 -t -1"
	expected := -1

	actual, err := fetchSortType(value)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func TestThatSortRequestIsCorrect(t *testing.T) {
	value := `sort -s 6 -l 5 -t -1 "dbcae there"`
	expected := "dbcae trhee"

	actual, err := SortRequest(value)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

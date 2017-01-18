package yelp

import "testing"

// TestCoordinateOptions will check using location options with bounding coordinates.
func TestCoordinateOptions(t *testing.T) {
	client := getClient(t)
	options := SearchOptions{
		CoordinateOptions: &CoordinateOptions{
			FloatFrom(37.9),
			FloatFrom(-122.5),
			FloatFrom(0),
			FloatFrom(0),
			FloatFrom(0),
		},
	}
	result, err := client.DoSearch(options)
	check(t, err)
	assert(t, len(result.Businesses) > 0, containsResults)
}

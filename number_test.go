package number_test

import (
	"reflect"
	"testing"

	number "github.com/go-random/number"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized number should not be 0")
}

func TestIntInRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := 10
	max := 20
	randomNumber := randomizer.IntInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized number should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized number should be less than max")
}

func TestIntTypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int", randomNumberType.Name(), "randomized number should be of type int")
}

func TestIntWithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized number should not be 0")
}

func TestIntInRangeWithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := 153243212
	max := 153243213
	randomNumber := randomizer.IntInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized number should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized number should be less than max")
}

func TestIntTypeCheckWithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int", randomNumberType.Name(), "randomized number should be of type int")
}

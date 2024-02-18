package number_test

import (
	"reflect"
	"testing"

	number "github.com/go-random/number"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomNumber(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.GenerateRandomNumber()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "Generated number should not be 0")
}

func TestGenerateRandomNumberInRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := 10
	max := 20
	randomNumber := randomizer.GenerateRandomNumberInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "Generated number should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "Generated number should be less than max")
}

func TestGenerateRandomNumberTypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.GenerateRandomNumber()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int", randomNumberType.Name(), "Generated number should be of type int")
}

func TestGenerateRandomNumberWithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.GenerateRandomNumber()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "Generated number should not be 0")
}

func TestGenerateRandomNumberInRangeWithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := 153243212
	max := 153243213
	randomNumber := randomizer.GenerateRandomNumberInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "Generated number should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "Generated number should be less than max")
}

func TestGenerateRandomNumberTypeCheckWithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.GenerateRandomNumber()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int", randomNumberType.Name(), "Generated number should be of type int")
}

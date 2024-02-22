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
	assert.NotEqual(t, 0, randomNumber, "randomized signed integer should not be 0")
}

func TestIntInRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := 10
	max := 20
	randomNumber := randomizer.IntInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized signed integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized signed integer should be less than max")
}

func TestInt_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int", randomNumberType.Name(), "randomized signed integer should be of type int")
}

func TestInt_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized signed integer should not be 0")
}

func TestIntInRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := 153243212
	max := 153243213
	randomNumber := randomizer.IntInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized signed integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized signed integer should be less than max")
}

func TestInt_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int", randomNumberType.Name(), "randomized signed integer should be of type int")
}

func TestInt32(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int32()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized signed integer should not be 0")
}

func TestInt32InRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := int32(10)
	max := int32(20)
	randomNumber := randomizer.Int32InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized signed integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized signed integer should be less than max")
}

func TestInt32_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int32()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int32", randomNumberType.Name(), "randomized signed integer should be of type int32")
}

func TestInt32_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int32()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized signed integer should not be 0")
}

func TestInt32InRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := int32(153243212)
	max := int32(153243213)
	randomNumber := randomizer.Int32InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized signed integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized signed integer should be less than max")
}

func TestInt32_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int32()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int32", randomNumberType.Name(), "randomized signed integer should be of type int32")
}

func TestInt64(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int64()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized signed integer should not be 0")
}

func TestInt64InRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := int64(10)
	max := int64(20)
	randomNumber := randomizer.Int64InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized signed integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized signed integer should be less than max")
}

func TestInt64_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Int64()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int64", randomNumberType.Name(), "randomized signed integer should be of type int64")
}

func TestInt64_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int64()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized signed integer should not be 0")
}

func TestInt64InRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := int64(153243212)
	max := int64(153243213)
	randomNumber := randomizer.Int64InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized signed integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized signed integer should be less than max")
}

func TestInt64_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Int64()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "int64", randomNumberType.Name(), "randomized signed integer should be of type int64")
}

func TestUInt(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.UInt()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized unsigned integer should not be 0")
}

func TestUIntInRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := uint(10)
	max := uint(20)
	randomNumber := randomizer.UIntInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized unsigned integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized unsigned integer should be less than max")
}

func TestUInt_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.UInt()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "uint", randomNumberType.Name(), "randomized unsigned integer should be of type uint")
}

func TestUInt_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.UInt()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized unsigned integer should not be 0")
}

func TestUIntInRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := uint(153243212)
	max := uint(153243213)
	randomNumber := randomizer.UIntInRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized unsigned integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized unsigned integer should be less than max")
}

func TestUInt_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.UInt()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "uint", randomNumberType.Name(), "randomized unsigned integer should be of type uint")
}

func TestUInt32(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.UInt32()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized unsigned integer should not be 0")
}

func TestUInt32InRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := uint32(10)
	max := uint32(20)
	randomNumber := randomizer.UInt32InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized unsigned integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized unsigned integer should be less than max")
}

func TestUInt32_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.UInt32()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "uint32", randomNumberType.Name(), "randomized unsigned integer should be of type uint32")
}

func TestUInt32_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.UInt32()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized unsigned integer should not be 0")
}

func TestUInt32InRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := uint32(153243212)
	max := uint32(153243213)
	randomNumber := randomizer.UInt32InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized unsigned integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized unsigned integer should be less than max")
}

func TestUInt32_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.UInt32()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "uint32", randomNumberType.Name(), "randomized unsigned integer should be of type uint32")
}

func TestUInt64(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.UInt64()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized unsigned integer should not be 0")
}

func TestUInt64InRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := uint64(10)
	max := uint64(20)
	randomNumber := randomizer.UInt64InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized unsigned integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized unsigned integer should be less than max")
}

func TestUInt64_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.UInt64()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "uint64", randomNumberType.Name(), "randomized unsigned integer should be of type uint64")
}

func TestUInt64_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.UInt64()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized unsigned integer should not be 0")
}

func TestUInt64InRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := uint64(153243212)
	max := uint64(153243213)
	randomNumber := randomizer.UInt64InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized unsigned integer should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized unsigned integer should be less than max")
}

func TestUInt64_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.UInt64()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "uint64", randomNumberType.Name(), "randomized unsigned integer should be of type uint64")
}

func TestFloat32(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Float32()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized float should not be 0")
	assert.NotEqual(t, 1, randomNumber, "randomized float should not be 1")
}

func TestFloat32InRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := float32(1.0)
	max := float32(10.0)
	randomNumber := randomizer.Float32InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized float should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized float should be less than max")
}

func TestFloat32_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Float32()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "float32", randomNumberType.Name(), "randomized float should be of type float32")
}

func TestFloat32_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Float32()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized float should not be 0")
}

func TestFloat32InRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := float32(154.0)
	max := float32(160.0)
	randomNumber := randomizer.Float32InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized float should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized float should be less than max")
}

func TestFloat32_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Float32()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "float32", randomNumberType.Name(), "randomized float should be of type float32")
}

func TestFloat64(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Float64()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized float should not be 0")
	assert.NotEqual(t, 1, randomNumber, "randomized float should not be 1")
}

func TestFloat64InRange(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	min := float64(1.0)
	max := float64(10.0)
	randomNumber := randomizer.Float64InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized float should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized float should be less than max")
}

func TestFloat64_TypeCheck(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.Float64()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "float64", randomNumberType.Name(), "randomized float should be of type float64")
}

func TestFloat64_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Float64()

	// Assert
	assert.NotEqual(t, 0, randomNumber, "randomized float should not be 0")
}

func TestFloat64InRange_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	min := float64(154.0)
	max := float64(160.0)
	randomNumber := randomizer.Float64InRange(min, max)

	// Assert
	assert.GreaterOrEqual(t, randomNumber, min, "randomized float should be greater than or equal to min")
	assert.Less(t, randomNumber, max, "randomized float should be less than max")
}

func TestFloat64_TypeCheck_WithSeed(t *testing.T) {
	// Setup
	seed := int64(31 * 41 * 53)
	randomizer := number.NewRandomizerWithSeed(seed)
	randomNumber := randomizer.Float64()
	randomNumberType := reflect.TypeOf(randomNumber)

	// Assert
	assert.Equal(t, "float64", randomNumberType.Name(), "randomized float should be of type float64")
}

func TestRandomPlusOrMinusOne(t *testing.T) {
	// Setup
	randomizer := number.NewRandomizer()
	randomNumber := randomizer.RandomPlusOrMinusOne()

	// Assert
	assert.Contains(t, []int{-1, 1}, randomNumber, "randomized number should be either -1 or 1")
}

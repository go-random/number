package number

// Int generates a random integer using the full range of the randomizer.
func (nr *Randomizer) Int() int {
	return nr.Rand.Int()
}

// IntInRange generates a random integer between min (inclusive) and max (exclusive).
func (nr *Randomizer) IntInRange(min, max int) int {
	checkRange(min, max)
	return nr.Rand.Intn(max-min) + min
}

// UInt generates a random unsigned integer using the full range of the randomizer.
func (nr *Randomizer) UInt() uint {
	return uint(nr.Rand.Uint64())
}

// UIntInRange generates a random unsigned integer between min (inclusive) and max (exclusive).
func (nr *Randomizer) UIntInRange(min, max int) uint {
	checkRange(min, max)
	return uint(nr.Rand.Uint64())%uint(max-min) + uint(min)
}

// UInt32 generates a random unsigned integer 32 using the full range of the randomizer.
func (nr *Randomizer) UInt32() uint32 {
	return nr.Rand.Uint32()
}

// UInt32InRange generates a random unsigned integer 32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) UInt32InRange(min, max int) uint32 {
	checkRange(min, max)
	return nr.Rand.Uint32()%uint32(max-min) + uint32(min)
}

// UInt64 generates a random unsigned integer 64 using the full range of the randomizer.
func (nr *Randomizer) UInt64() uint64 {
	return nr.Rand.Uint64()
}

// UInt64InRange generates a random unsigned integer 64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) UInt64InRange(min, max int) uint64 {
	checkRange(min, max)
	return nr.Rand.Uint64()%uint64(max-min) + uint64(min)
}

// Int32 generates a random integer 32 using the full range of the randomizer.
func (nr *Randomizer) Int32() int32 {
	return nr.Rand.Int31()
}

// Int32InRange generates a random integer 32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Int32InRange(min, max int) int32 {
	checkRange(min, max)
	return nr.Rand.Int31()%int32(max-min) + int32(min)
}

// Int64 generates a random integer 64 using the full range of the randomizer.
func (nr *Randomizer) Int64() int64 {
	return nr.Rand.Int63()
}

// Int64InRange generates a random integer 64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Int64InRange(min, max int) int64 {
	checkRange(min, max)
	return nr.Rand.Int63()%int64(max-min) + int64(min)
}

// checkRange ensures that min is less than max; otherwise, it panics.
func checkRange(min, max int) {
	if min >= max {
		panic("min must be less than max")
	}
}

// Float32 generates a random float32 between 0.0 (inclusive) and 1.0 (exclusive).
func (nr *Randomizer) Float32() float32 {
	return nr.Rand.Float32()
}

// Float32InRange generates a random float32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Float32InRange(min, max float32) float32 {
	checkFloatRange(float64(min), float64(max))
	return float32(nr.Rand.Float64())*(max-min) + min
}

// Float64 generates a random float64 between 0.0 (inclusive) and 1.0 (exclusive).
func (nr *Randomizer) Float64() float64 {
	return nr.Rand.Float64()
}

// Float64InRange generates a random float64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Float64InRange(min, max float64) float64 {
	checkFloatRange(min, max)
	return nr.Rand.Float64()*(max-min) + min
}

// checkFloatRange ensures that min is less than max; otherwise, it panics.
func checkFloatRange(min, max float64) {
	if min >= max {
		panic("min must be less than max")
	}
}

// PositiveOrNegative generates a random value of -1 or 1.
func (nr *Randomizer) PositiveOrNegative() int {
	if nr.Rand.Intn(2) == 0 {
		return -1
	}
	return 1
}

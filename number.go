package number

// Int generates a random integer using the full range of the randomizer.
func (nr *Randomizer) Int() int {
	return nr.Rand.Int()
}

// IntInRange generates a random integer between min (inclusive) and max (exclusive).
func (nr *Randomizer) IntInRange(min, max int) int {
	checkIntRange(min, max)
	return int(nr.Rand.Int63()%int64(max-min) + int64(min))
}

// UInt generates a random unsigned integer using the full range of the randomizer.
func (nr *Randomizer) UInt() uint {
	return uint(nr.Rand.Uint64())
}

// UIntInRange generates a random unsigned integer between min (inclusive) and max (exclusive).
func (nr *Randomizer) UIntInRange(min, max uint) uint {
	checkUIntRange(min, max)
	return uint(nr.Rand.Uint64()%(uint64(max)-uint64(min)) + uint64(min))
}

// UInt32 generates a random unsigned integer 32 using the full range of the randomizer.
func (nr *Randomizer) UInt32() uint32 {
	return nr.Rand.Uint32()
}

// UInt32InRange generates a random unsigned integer 32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) UInt32InRange(min, max uint32) uint32 {
	checkUInt32Range(min, max)
	return nr.Rand.Uint32()%(max-min) + min
}

// UInt64 generates a random unsigned integer 64 using the full range of the randomizer.
func (nr *Randomizer) UInt64() uint64 {
	return nr.Rand.Uint64()
}

// UInt64InRange generates a random unsigned integer 64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) UInt64InRange(min, max uint64) uint64 {
	checkUInt64Range(min, max)
	return nr.Rand.Uint64()%(max-min) + min
}

// Int32 generates a random integer 32 using the full range of the randomizer.
func (nr *Randomizer) Int32() int32 {
	return nr.Rand.Int31()
}

// Int32InRange generates a random integer 32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Int32InRange(min, max int32) int32 {
	checkInt32Range(min, max)
	return nr.Rand.Int31()%(max-min) + min
}

// Int64 generates a random integer 64 using the full range of the randomizer.
func (nr *Randomizer) Int64() int64 {
	return nr.Rand.Int63()
}

// Int64InRange generates a random integer 64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Int64InRange(min, max int64) int64 {
	checkInt64Range(min, max)
	return nr.Rand.Int63()%(max-min) + min
}

// Float32 generates a random float32 between 0.0 (inclusive) and 1.0 (exclusive).
func (nr *Randomizer) Float32() float32 {
	return nr.Rand.Float32()
}

// Float32InRange generates a random float32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Float32InRange(min, max float32) float32 {
	checkFloat32Range(min, max)
	return float32(nr.Rand.Float64())*(max-min) + min
}

// Float64 generates a random float64 between 0.0 (inclusive) and 1.0 (exclusive).
func (nr *Randomizer) Float64() float64 {
	return nr.Rand.Float64()
}

// Float64InRange generates a random float64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Float64InRange(min, max float64) float64 {
	checkFloat64Range(min, max)
	return nr.Rand.Float64()*(max-min) + min
}

// checkIntRange ensures that min is less than max; otherwise, it panics.
func checkIntRange(min, max int) {
	if min >= max {
		panic("min must be less than max")
	}
}

// checkInt32Range ensures that min is less than max; otherwise, it panics.
func checkInt32Range(min, max int32) {
	if min >= max {
		panic("min must be less than max")
	}
}

// checkInt64Range ensures that min is less than max; otherwise, it panics.
func checkInt64Range(min, max int64) {
	if min >= max {
		panic("min must be less than max")
	}
}

// checkUIntRange ensures that min is less than max; otherwise, it panics.
func checkUIntRange(min, max uint) {
	if min >= max {
		panic("min must be less than max")
	}
}

// checkUInt32Range ensures that min is less than max; otherwise, it panics.
func checkUInt32Range(min, max uint32) {
	if min >= max {
		panic("min must be less than max")
	}
}

// checkUInt64Range ensures that min is less than max; otherwise, it panics.
func checkUInt64Range(min, max uint64) {
	if min >= max {
		panic("min must be less than max")
	}
}

// checkFloat32Range ensures that min is less than max; otherwise, it panics.
func checkFloat32Range(min, max float32) {
	if min >= max {
		panic("min must be less than max")
	}
}

// checkFloat64Range ensures that min is less than max; otherwise, it panics.
func checkFloat64Range(min, max float64) {
	if min >= max {
		panic("min must be less than max")
	}
}

// RandomPlusOrMinusOne generates a random value of -1 or 1.
func (nr *Randomizer) RandomPlusOrMinusOne() int {
	if nr.Rand.Intn(2) == 0 {
		return -1
	}
	return 1
}

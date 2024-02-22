package number

// Int generates a random integer using the full range of the randomizer.
func (nr *Randomizer) Int() int {
	return nr.Rand.Int()
}

// IntInRange generates a random integer between min (inclusive) and max (exclusive).
func (nr *Randomizer) IntInRange(min, max int) int {
	if min >= max {
		panic("min must be less than max in IntInRange")
	}
	return nr.Rand.Intn(max-min) + min
}

// UInt generates a random unsigned integer using the full range of the randomizer.
func (nr *Randomizer) UInt() uint {
	return uint(nr.Rand.Uint64())
}

// UIntInRange generates a random unsigned integer between min (inclusive) and max (exclusive).
func (nr *Randomizer) UIntInRange(min, max int) uint {
	if min >= max {
		panic("min must be less than max in UIntInRange")
	}
	return uint(nr.Rand.Uint64())%uint(max-min) + uint(min)
}

// UInt32 generates a random unsigned integer 32 using the full range of the randomizer.
func (nr *Randomizer) UInt32() uint32 {
	return nr.Rand.Uint32()
}

// UInt32InRange generates a random unsigned integer 32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) UInt32InRange(min, max int) uint32 {
	if min >= max {
		panic("min must be less than max in UInt32InRange")
	}
	return nr.Rand.Uint32()%uint32(max-min) + uint32(min)
}

// UInt64 generates a random unsigned integer 64 using the full range of the randomizer.
func (nr *Randomizer) UInt64() uint64 {
	return nr.Rand.Uint64()
}

// UInt64InRange generates a random unsigned integer 64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) UInt64InRange(min, max int) uint64 {
	if min >= max {
		panic("min must be less than max in UInt64InRange")
	}
	return nr.Rand.Uint64()%uint64(max-min) + uint64(min)
}

// Int32 generates a random integer 32 using the full range of the randomizer.
func (nr *Randomizer) Int32() int32 {
	return nr.Rand.Int31()
}

// Int32InRange generates a random integer 32 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Int32InRange(min, max int) int32 {
	if min >= max {
		panic("min must be less than max in Int32InRange")
	}
	return nr.Rand.Int31()%int32(max-min) + int32(min)
}

// Int32 generates a random integer 64 using the full range of the randomizer.
func (nr *Randomizer) Int64() int64 {
	return nr.Rand.Int63()
}

// Int64InRange generates a random integer 64 between min (inclusive) and max (exclusive).
func (nr *Randomizer) Int64InRange(min, max int) int64 {
	if min >= max {
		panic("min must be less than max in Int64InRange")
	}
	return nr.Rand.Int63()%int64(max-min) + int64(min)
}

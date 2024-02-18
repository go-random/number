package number

// GenerateRandomNumber generates a random number using the full range of the randomizer.
func (nr *Randomizer) GenerateRandomNumber() int {
	return nr.Rand.Int()
}

// GenerateRandomNumberInRange generates a random number between min (inclusive) and max (exclusive).
func (nr *Randomizer) GenerateRandomNumberInRange(min, max int) int {
	if min >= max {
		panic("min must be less than max in GenerateRandomNumberInRange")
	}
	return nr.Rand.Intn(max-min) + min
}

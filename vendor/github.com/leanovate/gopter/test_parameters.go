package gopter

import (
	"math/rand"
	"time"
)

// TestParameters to run property tests
type TestParameters struct {
	MinSuccessfulTests int
	// MinSize is an (inclusive) lower limit on the size of the parameters
	MinSize int
	// MaxSize is an (exclusive) upper limit on the size of the parameters
	MaxSize         int
	MaxShrinkCount  int
	seed            int64
	Rng             *rand.Rand
	Workers         int
	MaxDiscardRatio float64
}

func (t *TestParameters) Seed() int64 {
	return t.seed
}

func (t *TestParameters) SetSeed(seed int64) {
	t.seed = seed
	t.Rng.Seed(seed)
}

// DefaultTestParameterWithSeeds creates reasonable default Parameters for most cases based on a fixed RNG-seed
func DefaultTestParametersWithSeed(seed int64) *TestParameters {
	return &TestParameters{
		MinSuccessfulTests: 100,
		MinSize:            0,
		MaxSize:            100,
		MaxShrinkCount:     1000,
		seed:               seed,
		Rng:                rand.New(NewLockedSource(seed)),
		Workers:            1,
		MaxDiscardRatio:    5,
	}
}

// DefaultTestParameterWithSeeds creates reasonable default Parameters for most cases with an undefined RNG-seed
func DefaultTestParameters() *TestParameters {
	return DefaultTestParametersWithSeed(time.Now().UnixNano())
}

package app

import (
	"math"
)

// Linear represents a D3-like linear scale.
type Linear struct {
	d0, d1 float64 // domain
	r0, r1 float64 // range
	clamp  bool
}

// NewLinear creates a new linear scale with default domain [0,1] and range [0,1].
func NewLinear() *Linear {
	return &Linear{d0: 0, d1: 1, r0: 0, r1: 1, clamp: false}
}

// Domain sets the domain [d0, d1] and returns the scale for chaining.
func (s *Linear) Domain(d0, d1 float64) *Linear {
	s.d0, s.d1 = d0, d1
	return s
}

// Range sets the range [r0, r1] and returns the scale for chaining.
func (s *Linear) Range(r0, r1 float64) *Linear {
	s.r0, s.r1 = r0, r1
	return s
}

// Clamp enables or disables clamping of input values to the domain.
func (s *Linear) Clamp(enabled bool) *Linear {
	s.clamp = enabled
	return s
}

// Copy returns a shallow copy of the scale.
func (s *Linear) Copy() *Linear {
	return &Linear{s.d0, s.d1, s.r0, s.r1, s.clamp}
}

// Scale maps a domain value x to the range.
func (s *Linear) Scale(x float64) float64 {
	// optionally clamp input
	if s.clamp {
		if s.d0 < s.d1 {
			if x < s.d0 {
				x = s.d0
			} else if x > s.d1 {
				x = s.d1
			}
		} else {
			// reversed domain
			if x > s.d0 {
				x = s.d0
			} else if x < s.d1 {
				x = s.d1
			}
		}
	}

	// handle zero domain span
	if s.d1 == s.d0 {
		return (s.r0 + s.r1) / 2
	}

	t := (x - s.d0) / (s.d1 - s.d0)
	// mapping t to range (respect possible reversed range)
	return s.r0 + t*(s.r1-s.r0)
}

// Invert maps a range value y back to the domain.
func (s *Linear) Invert(y float64) float64 {
	if s.r1 == s.r0 {
		return (s.d0 + s.d1) / 2
	}
	t := (y - s.r0) / (s.r1 - s.r0)
	return s.d0 + t*(s.d1-s.d0)
}

// Ticks returns a slice of nice tick values with approximately count ticks.
// This uses a "nice" step algorithm similar to d3-scale's tickStep.
func (s *Linear) Ticks(count int) []float64 {
	if count <= 0 {
		return nil
	}
	start, stop := s.d0, s.d1
	// if domain reversed, swap so ticks are ascending then reverse at the end
	reversed := false
	if start > stop {
		start, stop = stop, start
		reversed = true
	}

	span := stop - start
	if span == 0 {
		return []float64{start}
	}

	step := tickStep(start, stop, count)
	// compute start and end aligned to step
	min := math.Ceil(start/step) * step
	max := math.Floor(stop/step) * step

	// avoid floating imprecision by rounding values to sensible precision
	precision := -int(math.Floor(math.Log10(step)))
	if precision < 0 {
		precision = 0
	}
	round := func(v float64) float64 {
		pow := math.Pow(10, float64(precision))
		return math.Round(v*pow) / pow
	}

	var ticks []float64
	for v := min; v <= max+1e-9; v += step {
		ticks = append(ticks, round(v))
	}

	if reversed {
		// reverse ticks to match reversed domains
		for i, j := 0, len(ticks)-1; i < j; i, j = i+1, j-1 {
			ticks[i], ticks[j] = ticks[j], ticks[i]
		}
	}
	return ticks
}

// Nice expands the domain to "nice" round numbers based on tick spacing.
// After calling Nice you will often call Ticks to get aligned ticks.
func (s *Linear) Nice(count int) *Linear {
	if count <= 0 {
		return s
	}
	start, stop := s.d0, s.d1
	reverse := false
	if start > stop {
		start, stop = stop, start
		reverse = true
	}
	step := tickStep(start, stop, count)
	if step == 0 {
		return s
	}
	niceStart := math.Floor(start/step) * step
	niceStop := math.Ceil(stop/step) * step
	if reverse {
		s.d0, s.d1 = niceStop, niceStart
	} else {
		s.d0, s.d1 = niceStart, niceStop
	}
	return s
}

// tickStep returns a "nice" step between ticks for given start, stop and count.
func tickStep(start, stop float64, count int) float64 {
	if count <= 0 {
		return math.NaN()
	}
	step := (stop - start) / float64(count)
	if step == 0 {
		return 0
	}
	// get magnitude (10^floor(log10(step)))
	pow := math.Floor(math.Log10(step))
	mag := math.Pow(10, pow)
	// normalized step between 1 and 10
	norm := step / mag

	// pick multiplier from set {1,2,5,10} based on norm
	var mult float64
	// Use thresholds that give 1,2,5,10 progression.
	// These thresholds are a common heuristic (similar to d3).
	if norm >= 7.5 {
		mult = 10
	} else if norm >= 3.5 {
		mult = 5
	} else if norm >= 1.5 {
		mult = 2
	} else {
		mult = 1
	}
	return mult * mag
}

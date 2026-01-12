package farm

import (
	"fmt"
)

// Size represents a farm size value object
type Size struct {
	value float64
}

// NewSize creates a new Size value object with validation
func NewSize(value float64) (Size, error) {
	if value <= 0 {
		return Size{}, ErrInvalidFarmSize
	}
	if value > 1000000 {
		return Size{}, ErrInvalidFarmSize
	}
	return Size{value: value}, nil
}

// Value returns the size value
func (s Size) Value() float64 {
	return s.value
}

// String returns a formatted string representation
func (s Size) String() string {
	return fmt.Sprintf("%.2f hectares", s.value)
}

// Equals checks if two sizes are equal
func (s Size) Equals(other Size) bool {
	return s.value == other.value
}

// IsLargerThan checks if this size is larger than another
func (s Size) IsLargerThan(other Size) bool {
	return s.value > other.value
}

// IsSmallerThan checks if this size is smaller than another
func (s Size) IsSmallerThan(other Size) bool {
	return s.value < other.value
}

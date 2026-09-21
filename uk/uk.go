// Package uk is the UK AHT + B2B subset seat.
//
// UK = AHT + B2B subset. Not a customs/IOSS/MRN pack and not a 3PL suite.
package uk

// Seat is the locked package seat name.
const Seat = "uk"

// Marketplace is locked to UK for this subset.
const Marketplace = "UK"

// Subset is the UK AHT + B2B hours view.
type Subset struct {
	Marketplace string `json:"marketplace"`
	AHTDays     int    `json:"aht_days"`
	B2BOpen     int    `json:"b2b_open"`
	B2BClose    int    `json:"b2b_close"`
}

// B2BWindow is close minus open in hours. Invalid windows return 0.
func (s Subset) B2BWindow() int {
	if s.B2BClose <= s.B2BOpen {
		return 0
	}
	return s.B2BClose - s.B2BOpen
}

// Fixture is a deterministic UK subset row for smoke tests (no network).
func Fixture() Subset {
	return Subset{
		Marketplace: Marketplace,
		AHTDays:     2,
		B2BOpen:     9,
		B2BClose:    17,
	}
}

// Package flip is the FBM↔FBA margin flip seat.
//
// Margin flip is a signal, not a profit suite and not Helium10/Keepa.
package flip

// Seat is the locked package seat name.
const Seat = "flip"

const (
	PreferFBM = "fbm"
	PreferFBA = "fba"
)

// Margin is one SKU FBM vs FBA contribution row.
type Margin struct {
	Marketplace string  `json:"marketplace"`
	SKU         string  `json:"sku"`
	FBM         float64 `json:"fbm"`
	FBA         float64 `json:"fba"`
}

// Prefer is the higher-margin channel. Ties stay FBM (thin-SKU default).
func (m Margin) Prefer() string {
	if m.FBA > m.FBM {
		return PreferFBA
	}
	return PreferFBM
}

// Flip reports whether FBA is strictly better than FBM.
func (m Margin) Flip() bool {
	return m.Prefer() == PreferFBA
}

// Fixture is a deterministic DE flip row for smoke tests (no network).
func Fixture() Margin {
	return Margin{
		Marketplace: "DE",
		SKU:         "thin-sku-1",
		FBM:         1.20,
		FBA:         2.40,
	}
}

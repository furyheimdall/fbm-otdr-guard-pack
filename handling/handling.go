// Package handling is the stated-vs-actual handling delta seat.
//
// Handling lies on the clock. Delta + AHT auto-enable candidates
// are findings, not a prep/label tool.
package handling

// Seat is the locked package seat name.
const Seat = "handling"

// Delta is one SKU stated-vs-actual handling row.
type Delta struct {
	Marketplace   string `json:"marketplace"`
	SKU           string `json:"sku"`
	StatedDays    int    `json:"stated_days"`
	ActualDays    int    `json:"actual_days"`
	AHTAutoEnable bool   `json:"aht_auto_enable"`
}

// Days is actual minus stated. Positive means slower than stated.
func (d Delta) Days() int {
	return d.ActualDays - d.StatedDays
}

// Late reports whether actual handling exceeded stated handling.
func (d Delta) Late() bool {
	return d.Days() > 0
}

// Fixture is a deterministic DE row for smoke tests (no network).
func Fixture() Delta {
	return Delta{
		Marketplace:   "DE",
		SKU:           "thin-sku-1",
		StatedDays:    1,
		ActualDays:    3,
		AHTAutoEnable: true,
	}
}

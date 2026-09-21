// Package otdr is the OTDR risk scorecard seat.
//
// Thin-SKU EU FBM ops guard. OTDR is a scorecard finding, not a
// profit suite and not a prep stack.
package otdr

// Seat is the locked package seat name.
const Seat = "otdr"

// DefaultThreshold is the fixture OTDR percent threshold (1.0%).
const DefaultThreshold = 1.0

// Risk is a scorecard flag. It is not a profit verdict.
const (
	RiskOK   = "ok"
	RiskWarn = "warn"
	RiskHot  = "hot"
)

// Scorecard is one marketplace SKU OTDR row.
type Scorecard struct {
	Marketplace string  `json:"marketplace"`
	SKU         string  `json:"sku"`
	Rate        float64 `json:"rate"`
	Threshold   float64 `json:"threshold"`
}

// Risk flags rate vs threshold. Hot is at/over threshold; warn is half or more.
func (s Scorecard) Risk() string {
	if s.Threshold <= 0 {
		return RiskOK
	}
	if s.Rate >= s.Threshold {
		return RiskHot
	}
	if s.Rate >= s.Threshold/2 {
		return RiskWarn
	}
	return RiskOK
}

// Fixture is a deterministic DE row for smoke tests (no network).
func Fixture() Scorecard {
	return Scorecard{
		Marketplace: "DE",
		SKU:         "thin-sku-1",
		Rate:        1.2,
		Threshold:   DefaultThreshold,
	}
}

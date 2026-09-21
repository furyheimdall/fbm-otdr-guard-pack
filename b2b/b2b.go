// Package b2b is the B2B hours segment seat.
//
// B2B hours are a segment, not a B2B suite and not a prep stack.
package b2b

// Seat is the locked package seat name.
const Seat = "b2b"

// Hours is one marketplace Amazon Business hours window.
type Hours struct {
	Marketplace string `json:"marketplace"`
	Segment     string `json:"segment"`
	OpenHour    int    `json:"open_hour"`
	CloseHour   int    `json:"close_hour"`
}

// Window is close minus open in hours. Invalid windows return 0.
func (h Hours) Window() int {
	if h.CloseHour <= h.OpenHour {
		return 0
	}
	return h.CloseHour - h.OpenHour
}

// Fixture is a deterministic DE B2B window for smoke tests (no network).
func Fixture() Hours {
	return Hours{
		Marketplace: "DE",
		Segment:     "b2b",
		OpenHour:    8,
		CloseHour:   16,
	}
}

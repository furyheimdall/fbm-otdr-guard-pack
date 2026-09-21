package b2b

import "testing"

func TestSeat(t *testing.T) {
	if Seat != "b2b" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestFixtureWindow(t *testing.T) {
	h := Fixture()
	if h.Marketplace != "DE" || h.Segment != "b2b" {
		t.Fatalf("fixture = %+v", h)
	}
	if got := h.Window(); got != 8 {
		t.Fatalf("Window = %d, want 8", got)
	}
}

func TestInvalidWindow(t *testing.T) {
	h := Hours{OpenHour: 16, CloseHour: 8}
	if got := h.Window(); got != 0 {
		t.Fatalf("Window = %d, want 0", got)
	}
}

package flip

import "testing"

func TestSeat(t *testing.T) {
	if Seat != "flip" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestFixtureFlip(t *testing.T) {
	m := Fixture()
	if m.Marketplace != "DE" {
		t.Fatalf("marketplace = %q", m.Marketplace)
	}
	if !m.Flip() || m.Prefer() != PreferFBA {
		t.Fatalf("fixture prefer = %q flip=%t", m.Prefer(), m.Flip())
	}
}

func TestStayFBM(t *testing.T) {
	m := Margin{FBM: 3, FBA: 3}
	if m.Flip() || m.Prefer() != PreferFBM {
		t.Fatalf("tie should stay FBM: %+v", m)
	}
}

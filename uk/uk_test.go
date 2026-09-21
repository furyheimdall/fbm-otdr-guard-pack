package uk

import "testing"

func TestSeat(t *testing.T) {
	if Seat != "uk" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestFixtureUK(t *testing.T) {
	s := Fixture()
	if s.Marketplace != Marketplace {
		t.Fatalf("marketplace = %q, want %q", s.Marketplace, Marketplace)
	}
	if s.AHTDays != 2 {
		t.Fatalf("AHTDays = %d", s.AHTDays)
	}
	if got := s.B2BWindow(); got != 8 {
		t.Fatalf("B2BWindow = %d, want 8", got)
	}
}

func TestInvalidWindow(t *testing.T) {
	s := Subset{B2BOpen: 17, B2BClose: 9}
	if got := s.B2BWindow(); got != 0 {
		t.Fatalf("B2BWindow = %d, want 0", got)
	}
}

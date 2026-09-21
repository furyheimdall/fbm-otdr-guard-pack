package handling

import "testing"

func TestSeat(t *testing.T) {
	if Seat != "handling" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestFixtureDelta(t *testing.T) {
	d := Fixture()
	if d.Marketplace != "DE" {
		t.Fatalf("marketplace = %q", d.Marketplace)
	}
	if got := d.Days(); got != 2 {
		t.Fatalf("Days = %d, want 2", got)
	}
	if !d.Late() {
		t.Fatal("expected fixture to be late")
	}
}

func TestOnTime(t *testing.T) {
	d := Delta{StatedDays: 2, ActualDays: 2}
	if d.Late() || d.Days() != 0 {
		t.Fatalf("on-time delta = %+v", d)
	}
}

package alerts

import "testing"

func TestSeat(t *testing.T) {
	if Seat != "alerts" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestFixtureBHDR(t *testing.T) {
	s := FixtureBHDR()
	if s.Kind != KindBHDR || s.Marketplace != "DE" {
		t.Fatalf("fixture = %+v", s)
	}
	got := Format(ChannelSlack, s)
	if got != "slack: bhdr_cliff DE thin-sku-1" {
		t.Fatalf("Format = %q", got)
	}
	if Format(ChannelEmail, s) == "" {
		t.Fatal("email format empty")
	}
}

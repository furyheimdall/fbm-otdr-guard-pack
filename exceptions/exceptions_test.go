package exceptions

import "testing"

func TestSeat(t *testing.T) {
	if Seat != "exceptions" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestFixtureOpenCount(t *testing.T) {
	c := Fixture()
	if c.Marketplace != "DE" {
		t.Fatalf("marketplace = %q", c.Marketplace)
	}
	if got := c.OpenCount(); got != 2 {
		t.Fatalf("OpenCount = %d, want 2", got)
	}
	if len(c.Items) != 3 {
		t.Fatalf("items = %d", len(c.Items))
	}
	if c.Items[0].Kind != KindSSA || c.Items[1].Kind != KindAHT || c.Items[2].Kind != KindBuyShipping {
		t.Fatalf("kinds = %+v", c.Items)
	}
}

func TestOpenCountEmpty(t *testing.T) {
	if got := (Checklist{}).OpenCount(); got != 0 {
		t.Fatalf("OpenCount = %d, want 0", got)
	}
}

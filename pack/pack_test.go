package pack

import (
	"os"
	"strings"
	"testing"
)

func TestSeat(t *testing.T) {
	if Seat != "pack" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestSeatsLocked(t *testing.T) {
	got := Seats()
	want := []string{"otdr", "handling", "exceptions", "b2b", "flip", "uk", "pack"}
	if len(got) != len(want) {
		t.Fatalf("seats = %v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("seats[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFromFixturesRender(t *testing.T) {
	s := FromFixtures()
	if !s.NotProfit || !s.NotPrep || !s.MergeForbid {
		t.Fatalf("product flags: %+v", s)
	}
	var b strings.Builder
	if err := Render(&b, s); err != nil {
		t.Fatal(err)
	}
	out := b.String()
	for _, want := range []string{
		"FBM OTDR Guard Pack",
		"Thin-SKU EU FBM ops guard.",
		"Not a profit suite. Not a prep stack.",
		"otdr.sku:",
		"risk=hot",
		"handling.delta: 2",
		"exceptions.open: 2",
		"b2b.window: 8",
		"flip.prefer: fba",
		"uk.marketplace: UK",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("pack missing %q:\n%s", want, out)
		}
	}
}

func TestREADMESeatsLocked(t *testing.T) {
	raw, err := os.ReadFile("../README.md")
	if err != nil {
		t.Fatal(err)
	}
	body := string(raw)
	for _, seat := range []string{"otdr/", "handling/", "exceptions/", "b2b/", "flip/", "uk/", "pack/", "cmd/fogp/"} {
		if !strings.Contains(body, seat) {
			t.Fatalf("README missing seat %q", seat)
		}
	}
	for _, out := range []string{"3PL/SellerLegend", "Commingled-Exit", "Peak Fee", "customs/IOSS/MRN", "Helium10/Keepa", "ARAP/CDS/CICS/OTM/DRC/Deadbugz"} {
		if !strings.Contains(body, out) {
			t.Fatalf("README missing OUT %q", out)
		}
	}
}

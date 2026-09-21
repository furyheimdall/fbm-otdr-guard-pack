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
	want := []string{"otdr", "handling", "exceptions", "b2b", "flip", "alerts", "uk", "pack"}
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
	if !s.Not3PL || !s.NotPrep || !s.MergeForbid {
		t.Fatalf("product flags: %+v", s)
	}
	var b strings.Builder
	if err := Render(&b, s); err != nil {
		t.Fatal(err)
	}
	out := b.String()
	for _, want := range []string{
		"FBM OTDR Guard Pack",
		"OTDR slips. Handling lies on the clock.",
		"otdr.sku:",
		"risk=hot",
		"handling.delta: 2",
		"aht_auto=true",
		"exceptions.open: 2",
		"b2b.window: 8",
		"hours_gap=true",
		"flip.prefer: fba",
		"alert: bhdr_cliff",
		"uk.marketplace: UK",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("pack missing %q:\n%s", want, out)
		}
	}
}

func TestREADMECanonLocked(t *testing.T) {
	raw, err := os.ReadFile("../README.md")
	if err != nil {
		t.Fatal(err)
	}
	body := string(raw)
	for _, seat := range []string{"otdr/", "handling/", "exceptions/", "b2b/", "flip/", "alerts/", "uk/", "pack/", "cmd/fogp/"} {
		if !strings.Contains(body, seat) {
			t.Fatalf("README missing seat %q", seat)
		}
	}
	for _, want := range []string{
		"OTDR slips. Handling lies on the clock.",
		"DE/UK first",
		"Not a 3PL suite",
		"Not a prep/label tool",
		"SellerLegend = profit/analytics",
		"we = SKU compliance guard + checklist before cliffs",
		"Full 3PL / SellerLegend profit suite me-too",
		"FBA Commingled-Exit Label & Lot-Cost",
		"FBA Peak Fee Margin Guard",
		"Decision C: customs / IOSS / MRN",
		"Prep printer / Helium10/Keepa me-too",
		"Gift address / Spend-by-URL",
		"Deadbugz / DRC / OTM / ARAP / CDS / CICS",
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("README missing canon %q", want)
		}
	}
}

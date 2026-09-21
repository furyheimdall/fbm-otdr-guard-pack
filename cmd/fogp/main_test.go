package main

import (
	"strings"
	"testing"
)

func TestSeats(t *testing.T) {
	var b strings.Builder
	if code := run(nil, &b); code != 0 {
		t.Fatalf("exit %d", code)
	}
	got := b.String()
	for _, seat := range []string{"otdr", "handling", "exceptions", "b2b", "flip", "alerts", "uk", "pack"} {
		if !strings.Contains(got, seat) {
			t.Fatalf("missing seat %q in %q", seat, got)
		}
	}
}

func TestHelp(t *testing.T) {
	var b strings.Builder
	if code := run([]string{"help"}, &b); code != 0 {
		t.Fatalf("exit %d", code)
	}
	out := b.String()
	for _, want := range []string{"OTDR slips", "Handling lies on the clock", "Not a 3PL suite", "prep/label tool"} {
		if !strings.Contains(out, want) {
			t.Fatalf("help missing %q: %q", want, out)
		}
	}
}

func TestUnknown(t *testing.T) {
	var b strings.Builder
	if code := run([]string{"profit"}, &b); code != 2 {
		t.Fatalf("exit %d, want 2", code)
	}
}

func TestSeatsList(t *testing.T) {
	if len(seats()) != 8 {
		t.Fatalf("seats = %v", seats())
	}
}

func TestPack(t *testing.T) {
	var b strings.Builder
	if code := run([]string{"pack"}, &b); code != 0 {
		t.Fatalf("exit %d, out=%q", code, b.String())
	}
	out := b.String()
	for _, want := range []string{"FBM OTDR Guard Pack", "OTDR slips. Handling lies on the clock.", "otdr.sku:"} {
		if !strings.Contains(out, want) {
			t.Fatalf("pack missing %q:\n%s", want, out)
		}
	}
}

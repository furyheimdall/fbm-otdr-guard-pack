// Package pack assembles the guard seats from interfaces/fixtures.
//
// Thin OSS pack assembly only. Thin-SKU EU FBM ops guard.
// Not a profit suite. Not a prep stack. Product merge with
// ARAP/CDS/CICS/OTM/DRC/Deadbugz forbidden.
package pack

import (
	"fmt"
	"io"
	"strings"

	"github.com/furyheimdall/fbm-otdr-guard-pack/b2b"
	"github.com/furyheimdall/fbm-otdr-guard-pack/exceptions"
	"github.com/furyheimdall/fbm-otdr-guard-pack/flip"
	"github.com/furyheimdall/fbm-otdr-guard-pack/handling"
	"github.com/furyheimdall/fbm-otdr-guard-pack/otdr"
	"github.com/furyheimdall/fbm-otdr-guard-pack/uk"
)

// Seat is the locked package seat name.
const Seat = "pack"

// Seats is the locked seat list assembled by this pack.
func Seats() []string {
	return []string{
		otdr.Seat,
		handling.Seat,
		exceptions.Seat,
		b2b.Seat,
		flip.Seat,
		uk.Seat,
		Seat,
	}
}

// Snapshot is a fixture-backed assembly of the day-1 seats.
type Snapshot struct {
	OTDR        otdr.Scorecard
	Handling    handling.Delta
	Exceptions  exceptions.Checklist
	B2B         b2b.Hours
	Flip        flip.Margin
	UK          uk.Subset
	Product     string
	NotProfit   bool
	NotPrep     bool
	MergeForbid bool
}

// FromFixtures builds a snapshot with no network I/O.
func FromFixtures() Snapshot {
	return Snapshot{
		OTDR:        otdr.Fixture(),
		Handling:    handling.Fixture(),
		Exceptions:  exceptions.Fixture(),
		B2B:         b2b.Fixture(),
		Flip:        flip.Fixture(),
		UK:          uk.Fixture(),
		Product:     "FBM OTDR Guard Pack",
		NotProfit:   true,
		NotPrep:     true,
		MergeForbid: true,
	}
}

// Render writes a thin text pack. Not a SaaS landing page.
func Render(w io.Writer, s Snapshot) error {
	var b strings.Builder
	fmt.Fprintf(&b, "product: %s\n", s.Product)
	fmt.Fprintf(&b, "anchors: Thin-SKU EU FBM ops guard. Not a profit suite. Not a prep stack.\n")
	fmt.Fprintf(&b, "not: profit-suite prep-stack\n")
	fmt.Fprintf(&b, "merge_forbidden: %t (ARAP/CDS/CICS/OTM/DRC/Deadbugz side-by-side only)\n", s.MergeForbid)
	fmt.Fprintf(&b, "otdr.sku: %s marketplace=%s rate=%.2f risk=%s\n",
		s.OTDR.SKU, s.OTDR.Marketplace, s.OTDR.Rate, s.OTDR.Risk())
	fmt.Fprintf(&b, "handling.delta: %d late=%t\n", s.Handling.Days(), s.Handling.Late())
	fmt.Fprintf(&b, "exceptions.open: %d\n", s.Exceptions.OpenCount())
	fmt.Fprintf(&b, "b2b.window: %d\n", s.B2B.Window())
	fmt.Fprintf(&b, "flip.prefer: %s flip=%t\n", s.Flip.Prefer(), s.Flip.Flip())
	fmt.Fprintf(&b, "uk.marketplace: %s aht=%d b2b_window=%d\n",
		s.UK.Marketplace, s.UK.AHTDays, s.UK.B2BWindow())
	_, err := io.WriteString(w, b.String())
	return err
}

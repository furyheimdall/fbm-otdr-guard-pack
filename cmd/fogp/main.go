// Command fogp is the FBM OTDR Guard Pack CLI stub.
// It prints seat names, help, or a fixture pack. No network I/O.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/furyheimdall/fbm-otdr-guard-pack/alerts"
	"github.com/furyheimdall/fbm-otdr-guard-pack/b2b"
	"github.com/furyheimdall/fbm-otdr-guard-pack/exceptions"
	"github.com/furyheimdall/fbm-otdr-guard-pack/flip"
	"github.com/furyheimdall/fbm-otdr-guard-pack/handling"
	"github.com/furyheimdall/fbm-otdr-guard-pack/otdr"
	"github.com/furyheimdall/fbm-otdr-guard-pack/pack"
	"github.com/furyheimdall/fbm-otdr-guard-pack/uk"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout))
}

func run(args []string, w io.Writer) int {
	cmd := "seats"
	if len(args) > 0 {
		cmd = args[0]
	}
	switch cmd {
	case "help", "-h", "--help":
		return writeHelp(w)
	case "seats", "":
		return writeSeats(w)
	case "pack":
		return writePack(w)
	default:
		fmt.Fprintf(w, "unknown command %q\n", cmd)
		writeHelp(w)
		return 2
	}
}

func writeHelp(w io.Writer) int {
	fmt.Fprint(w, `FBM OTDR Guard Pack — fogp

Thin-SKU FBM guard — OTDR risk scorecard + stated-vs-actual handling delta + SSA/AHT/Buy Shipping(Veeqo) exception checklist + B2B business-hours segment + FBM↔FBA margin flip. DE/UK first.
OTDR slips. Handling lies on the clock.
Not a 3PL suite / prep/label tool.

Usage:
  fogp          print package seats
  fogp seats    print package seats
  fogp pack     render the fixture pack
  fogp help     print this help

No network. Not a 3PL suite / not a prep/label tool. Product merge with Deadbugz/DRC/OTM/ARAP/CDS/CICS forbidden.
`)
	return 0
}

func writePack(w io.Writer) int {
	if err := pack.Render(w, pack.FromFixtures()); err != nil {
		fmt.Fprintf(w, "pack: %v\n", err)
		return 1
	}
	return 0
}

func writeSeats(w io.Writer) int {
	for _, name := range seats() {
		fmt.Fprintln(w, name)
	}
	return 0
}

func seats() []string {
	return []string{
		otdr.Seat,
		handling.Seat,
		exceptions.Seat,
		b2b.Seat,
		flip.Seat,
		alerts.Seat,
		uk.Seat,
		pack.Seat,
	}
}

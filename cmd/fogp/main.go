// Command fogp is the FBM OTDR Guard Pack CLI stub.
// It prints seat names, help, or a fixture pack. No network I/O.
package main

import (
	"fmt"
	"io"
	"os"

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

Thin-SKU FBM guard — OTDR risk scorecard + stated-vs-actual handling delta + SSA/AHT/Buy Shipping(Veeqo) exception checklist + B2B hours segment + FBM↔FBA margin flip. DE/UK first.
Thin-SKU EU FBM ops guard. Not a profit suite. Not a prep stack.

Usage:
  fogp          print package seats
  fogp seats    print package seats
  fogp pack     render the fixture pack
  fogp help     print this help

No network. Not a profit suite / not a prep stack. Product merge with ARAP/CDS/CICS/OTM/DRC/Deadbugz forbidden.
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
		uk.Seat,
		pack.Seat,
	}
}

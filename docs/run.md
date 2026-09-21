# Run the pack

Thin docs only. English README canon is the single source of truth.

**Thin-SKU EU FBM ops guard.**

**Not a profit suite. Not a prep stack.**

## Test

```bash
go test ./...
```

Tests must not use the network.

## CLI

```bash
go run ./cmd/fogp
go run ./cmd/fogp help
go run ./cmd/fogp seats
go run ./cmd/fogp pack
```

`pack` assembles fixture-backed seats. It does not call Seller Central, Veeqo, 3PL, Helium10, Keepa, or any existing adjacent track.

## Day-1 IN / OUT reminders

IN:

1. OTDR risk scorecard
2. Stated-vs-actual handling delta
3. SSA / AHT / Buy Shipping (Veeqo) exception checklist
4. B2B hours segment
5. FBM↔FBA margin flip
6. UK subset (AHT + B2B)
7. Thin docs / pack assembly

OUT: 3PL/SellerLegend · Commingled-Exit · Peak Fee · customs/IOSS/MRN · Helium10/Keepa · existing tracks (ARAP/CDS/CICS/OTM/DRC/Deadbugz).

Copy rules: no heavy SaaS landing; do not position as a profit suite or prep stack; adjacent tracks stay side-by-side; product merge forbidden.

About paste and Topics: [launch-note.md](launch-note.md) (notes only — do not set GitHub Topics via API).

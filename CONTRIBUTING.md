# Contributing

English README canon is the single source of truth for scope. Keep the locked package seats. Do not open PRs for OUT scope.

## Keep these seats

| Seat | Path |
| --- | --- |
| OTDR risk scorecard | `otdr/` |
| Stated-vs-actual handling delta | `handling/` |
| SSA / AHT / Buy Shipping (Veeqo) exception checklist | `exceptions/` |
| B2B hours segment | `b2b/` |
| FBM↔FBA margin flip | `flip/` |
| UK subset (AHT + B2B) | `uk/` |
| Thin docs / pack assembly | `pack/` |
| CLI | `cmd/fogp/` |

Add implementation inside an existing seat. Do not invent sibling product surfaces.

## Do not open PRs for OUT scope

OUT of the locked MVP (see README):

- 3PL / SellerLegend
- Commingled-Exit
- Peak Fee
- customs / IOSS / MRN
- Helium10 / Keepa
- existing tracks (ARAP / CDS / CICS / OTM / DRC / Deadbugz)

PRs that add those surfaces will be closed.

## Anchors

**Thin-SKU EU FBM ops guard.**

**Not a profit suite. Not a prep stack.**

| Anchor | Meaning |
| --- | --- |
| Thin-SKU EU FBM ops guard. | Day-1 seats score OTDR, handling delta, SSA/AHT/Buy Shipping exceptions, B2B hours, and FBM↔FBA flip on DE/UK. This is not a 3PL or SellerLegend replacement. |
| Not a profit suite. | Margin flip is a signal, not a profit dashboard or Helium10/Keepa suite. |
| Not a prep stack. | Handling and exceptions stay checklists. This is not inbound prep, Commingled-Exit, Peak Fee, or customs/IOSS/MRN. |

## Copy rules

- No heavy SaaS landing
- Do not position as a profit suite or prep stack
- Adjacent tracks (ARAP/CDS/CICS/OTM/DRC/Deadbugz) = side-by-side only; product merge forbidden
- 3PL/SellerLegend, Commingled-Exit, Peak Fee, customs/IOSS/MRN, Helium10/Keepa stay OUT

About paste and Topics stay in [docs/launch-note.md](docs/launch-note.md) as notes only. Do not set GitHub Topics via API.

## Develop

```bash
go test ./...
go run ./cmd/fogp pack
```

Tests must not require live network. Pack assembly uses fixtures when other seats are still stubs. Do not add a profit suite or prep stack. See [docs/run.md](docs/run.md).

## Issues

Work is tracked under the `[Epic] FBM OTDR Guard Pack MVP` parent and its child issues. Reference the matching child in the PR body. Do not close Epic or child issues from the scaffold PR.

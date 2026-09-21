# FBM OTDR Guard Pack

Thin-SKU FBM guard — OTDR risk scorecard + stated-vs-actual handling delta + SSA/AHT/Buy Shipping(Veeqo) exception checklist + B2B hours segment + FBM↔FBA margin flip. DE/UK first.

**Thin-SKU EU FBM ops guard.**

**Not a profit suite. Not a prep stack.**

English is the single source of truth. Thin OSS Go pack — not a profit suite, not a prep stack.

## Anchors

| Anchor | Meaning |
| --- | --- |
| Thin-SKU EU FBM ops guard. | Day-1 seats score OTDR, handling delta, SSA/AHT/Buy Shipping exceptions, B2B hours, and FBM↔FBA flip on DE/UK. This is not a 3PL or SellerLegend replacement. |
| Not a profit suite. | Margin flip is a signal, not a profit dashboard or Helium10/Keepa suite. |
| Not a prep stack. | Handling and exceptions stay checklists. This is not inbound prep, Commingled-Exit, Peak Fee, or customs/IOSS/MRN. |

## ICP

- Thin-SKU EU FBM operators (DE/UK first)
- Buyer: EU FBM ops lead; agency EU marketplace ops = influencer
- Why now: OTDR risk + stated-vs-actual handling + SSA/AHT/Buy Shipping (Veeqo) exceptions + B2B hours segment + FBM↔FBA margin flip
- Adjacent (not merge): Auth Rate Audit Pack, CAPI Dedup Scorecard, Call-ID Continuity Scorecard, Offer Truth Monitor, Deposit Return Clock, Deadbugz

## Day-1 IN

1. OTDR risk scorecard
2. Stated-vs-actual handling delta
3. SSA / AHT / Buy Shipping (Veeqo) exception checklist
4. B2B hours segment
5. FBM↔FBA margin flip
6. UK subset (AHT + B2B)
7. Thin docs / pack assembly

## OUT

3PL/SellerLegend · Commingled-Exit · Peak Fee · customs/IOSS/MRN · Helium10/Keepa · existing tracks (ARAP/CDS/CICS/OTM/DRC/Deadbugz)

## Copy rules

- No heavy SaaS landing
- Do not position as a profit suite or prep stack
- Adjacent tracks (ARAP/CDS/CICS/OTM/DRC/Deadbugz) = side-by-side only; product merge forbidden
- 3PL/SellerLegend, Commingled-Exit, Peak Fee, customs/IOSS/MRN, Helium10/Keepa stay OUT

## Package seats

| Package | Seat |
| --- | --- |
| [`otdr/`](otdr/) | OTDR risk scorecard |
| [`handling/`](handling/) | Stated-vs-actual handling delta |
| [`exceptions/`](exceptions/) | SSA / AHT / Buy Shipping (Veeqo) exception checklist |
| [`b2b/`](b2b/) | B2B hours segment |
| [`flip/`](flip/) | FBM↔FBA margin flip |
| [`uk/`](uk/) | UK subset (AHT + B2B) |
| [`pack/`](pack/) | Thin docs / pack assembly — assembles the other seats from interfaces/fixtures |
| [`cmd/fogp/`](cmd/fogp/) | CLI stub (prints seat names or help; no network in tests) |

Seats are fixture-backed stubs. Tests must not use the network. Do not add OUT-scope packages. Adjacent tracks (ARAP/CDS/CICS/OTM/DRC/Deadbugz) stay side-by-side; product merge is forbidden.

## Develop

```bash
go test ./...
```

CLI (no network; pack uses fixtures):

```bash
go run ./cmd/fogp
go run ./cmd/fogp help
go run ./cmd/fogp seats
go run ./cmd/fogp pack
```

How to run the pack, plus IN/OUT reminders: [docs/run.md](docs/run.md).

About paste and Topics are launch notes only: [docs/launch-note.md](docs/launch-note.md). Do not set GitHub Topics via API.

Module: [`github.com/furyheimdall/fbm-otdr-guard-pack`](https://github.com/furyheimdall/fbm-otdr-guard-pack) · License: [MIT](LICENSE)

See [CONTRIBUTING.md](CONTRIBUTING.md).

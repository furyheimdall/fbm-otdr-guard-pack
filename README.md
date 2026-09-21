# FBM OTDR Guard Pack

Thin-SKU FBM guard — OTDR risk scorecard + stated-vs-actual handling delta + SSA/AHT/Buy Shipping(Veeqo) exception checklist + B2B business-hours segment + FBM↔FBA margin flip. **DE/UK first.**

**OTDR slips. Handling lies on the clock.**

**DE/UK first. Not a 3PL suite / prep/label tool.**

English is the single source of truth. Thin OSS Go pack — not a 3PL suite, not a prep/label tool.

## Anchors

| Anchor | Meaning |
| --- | --- |
| OTDR slips. | Per-SKU Order Defect Rate can move on a rolling window (volume-exemption aware where documented). This is a risk scorecard, not a 3PL or SellerLegend profit suite. |
| Handling lies on the clock. | Stated handling time can diverge from actual. Delta + AHT auto-enable candidates are findings. This is not a prep/label tool. |

## ICP

- Amazon FBM sellers on DE (UK for AHT + B2B hours); thin-/multi-SKU portfolios
- Buyer: seller ops / account health lead; agencies/3PL advisors = influencers
- Why now: DE OTDR 90% (~2026-07-15, listing risk ~09-01); AHT auto ~09-01; business-hours 90% ~09-30 → B2B deact risk ~10-30

## Day-1 IN

1. Per-SKU OTDR risk score (rolling window / volume exemption aware where documented)
2. Handling time stated vs actual delta; flag AHT auto-enable candidates
3. Protection-triad checklist: AHT + SSA + Buy Shipping (or Veeqo) — which SKUs/orders lack coverage
4. B2B / business-hours delivery segment + hours config gap (BHDR → deact risk)
5. At-risk SKU FBM↔FBA margin flip suggestion (profit narrative only — not auto-enrollment)
6. Slack/email before BHDR / deactivation cliffs
7. UK = AHT + B2B subset

## OUT / Hold

- Full 3PL / SellerLegend profit suite me-too
- FBA Commingled-Exit Label & Lot-Cost / FNSKU stickerless (Watch)
- FBA Peak Fee Margin Guard (Watch)
- Decision C: customs / IOSS / MRN
- Prep printer / Helium10/Keepa me-too
- Gift address / Spend-by-URL (Watch)
- Deadbugz / DRC / OTM / ARAP / CDS / CICS

## Position

SellerLegend = profit/analytics · Subke = 3PL · **we = SKU compliance guard + checklist before cliffs**. Not a prep/label tool.

## Copy rules

- No heavy SaaS landing
- Protection triad: orders on AHT+SSA+Buy Shipping not adversely counted for OTDR/BHDR (LDR still applies) — cite Amazon defs carefully
- Helium10/SellerLegend comps = internal differentiation only

## Package seats

| Package | Seat |
| --- | --- |
| [`otdr/`](otdr/) | Per-SKU OTDR risk score (rolling window / volume exemption aware where documented) |
| [`handling/`](handling/) | Handling time stated vs actual delta; flag AHT auto-enable candidates |
| [`exceptions/`](exceptions/) | Protection-triad checklist: AHT + SSA + Buy Shipping (or Veeqo) — which SKUs/orders lack coverage |
| [`b2b/`](b2b/) | B2B / business-hours delivery segment + hours config gap (BHDR → deact risk) |
| [`flip/`](flip/) | At-risk SKU FBM↔FBA margin flip suggestion (profit narrative only — not auto-enrollment) |
| [`alerts/`](alerts/) | Slack/email before BHDR / deactivation cliffs |
| [`uk/`](uk/) | UK = AHT + B2B subset |
| [`pack/`](pack/) | Thin docs / pack assembly — assembles the other seats from interfaces/fixtures |
| [`cmd/fogp/`](cmd/fogp/) | CLI stub (prints seat names or help; no network in tests) |

Seats are fixture-backed stubs. Tests must not use the network. Do not add OUT / Hold packages. Adjacent tracks (Deadbugz / DRC / OTM / ARAP / CDS / CICS) stay side-by-side; product merge is forbidden.

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

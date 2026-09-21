# Contributing

English README canon (marketing LOCK) is the single source of truth for scope. Keep the locked package seats. Do not open PRs for OUT / Hold scope.

## Keep these seats

| Seat | Path |
| --- | --- |
| Per-SKU OTDR risk score | `otdr/` |
| Handling time stated vs actual delta; AHT auto-enable candidates | `handling/` |
| Protection-triad checklist: AHT + SSA + Buy Shipping (or Veeqo) | `exceptions/` |
| B2B / business-hours segment + hours config gap (BHDR → deact risk) | `b2b/` |
| At-risk SKU FBM↔FBA margin flip (narrative only) | `flip/` |
| Slack/email before BHDR / deactivation cliffs | `alerts/` |
| UK = AHT + B2B subset | `uk/` |
| Thin docs / pack assembly | `pack/` |
| CLI | `cmd/fogp/` |

Add implementation inside an existing seat. Do not invent sibling product surfaces.

## Do not open PRs for OUT / Hold scope

OUT / Hold of the locked MVP (see README):

- Full 3PL / SellerLegend profit suite me-too
- FBA Commingled-Exit Label & Lot-Cost / FNSKU stickerless (Watch)
- FBA Peak Fee Margin Guard (Watch)
- Decision C: customs / IOSS / MRN
- Prep printer / Helium10/Keepa me-too
- Gift address / Spend-by-URL (Watch)
- Deadbugz / DRC / OTM / ARAP / CDS / CICS

PRs that add those surfaces will be closed.

## Anchors

**OTDR slips. Handling lies on the clock.**

**DE/UK first. Not a 3PL suite / prep/label tool.**

| Anchor | Meaning |
| --- | --- |
| OTDR slips. | Per-SKU Order Defect Rate can move on a rolling window (volume-exemption aware where documented). This is a risk scorecard, not a 3PL or SellerLegend profit suite. |
| Handling lies on the clock. | Stated handling time can diverge from actual. Delta + AHT auto-enable candidates are findings. This is not a prep/label tool. |

## Position

SellerLegend = profit/analytics · Subke = 3PL · **we = SKU compliance guard + checklist before cliffs**. Not a prep/label tool.

## Copy rules

- No heavy SaaS landing
- Protection triad: orders on AHT+SSA+Buy Shipping not adversely counted for OTDR/BHDR (LDR still applies) — cite Amazon defs carefully
- Helium10/SellerLegend comps = internal differentiation only

About paste and Topics stay in [docs/launch-note.md](docs/launch-note.md) as notes only. Do not set GitHub Topics via API.

## Develop

```bash
go test ./...
go run ./cmd/fogp pack
```

Tests must not require live network. Pack assembly uses fixtures when other seats are still stubs. Do not add a 3PL suite or prep/label tool. See [docs/run.md](docs/run.md).

## Issues

Work is tracked under the `[Epic] FBM OTDR Guard Pack MVP` parent and its child issues. Reference the matching child in the PR body. Do not close Epic or child issues from the scaffold PR.

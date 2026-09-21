# Run the pack

Thin docs only. English README canon (marketing LOCK) is the single source of truth.

**OTDR slips. Handling lies on the clock.**

**DE/UK first. Not a 3PL suite / prep/label tool.**

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

`pack` assembles fixture-backed seats. It does not call Seller Central, Veeqo, Slack, email, 3PL, Helium10, Keepa, or any existing adjacent track.

## Day-1 IN / OUT Hold reminders

IN:

1. Per-SKU OTDR risk score (rolling window / volume exemption aware where documented)
2. Handling time stated vs actual delta; flag AHT auto-enable candidates
3. Protection-triad checklist: AHT + SSA + Buy Shipping (or Veeqo) — which SKUs/orders lack coverage
4. B2B / business-hours delivery segment + hours config gap (BHDR → deact risk)
5. At-risk SKU FBM↔FBA margin flip suggestion (profit narrative only — not auto-enrollment)
6. Slack/email before BHDR / deactivation cliffs
7. UK = AHT + B2B subset

OUT / Hold: Full 3PL / SellerLegend profit suite me-too · FBA Commingled-Exit Label & Lot-Cost / FNSKU stickerless (Watch) · FBA Peak Fee Margin Guard (Watch) · Decision C: customs / IOSS / MRN · Prep printer / Helium10/Keepa me-too · Gift address / Spend-by-URL (Watch) · Deadbugz / DRC / OTM / ARAP / CDS / CICS.

Position: SellerLegend = profit/analytics · Subke = 3PL · we = SKU compliance guard + checklist before cliffs. Not a prep/label tool.

Copy rules: no heavy SaaS landing; protection triad cite Amazon defs carefully; Helium10/SellerLegend comps are internal differentiation only.

About paste and Topics: [launch-note.md](launch-note.md) (notes only — do not set GitHub Topics via API).

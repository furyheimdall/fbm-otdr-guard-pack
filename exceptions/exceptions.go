// Package exceptions is the protection-triad checklist seat.
//
// AHT + SSA + Buy Shipping (or Veeqo): which SKUs/orders lack
// coverage. Checklist only — not a 3PL suite.
package exceptions

// Seat is the locked package seat name.
const Seat = "exceptions"

// Kind is a locked exception class.
type Kind string

const (
	KindSSA         Kind = "ssa"
	KindAHT         Kind = "aht"
	KindBuyShipping Kind = "buy_shipping"
)

// Item is one checklist row. Open is a finding, not a remediation.
type Item struct {
	Kind Kind   `json:"kind"`
	SKU  string `json:"sku"`
	Open bool   `json:"open"`
	Note string `json:"note,omitempty"`
}

// Checklist is one marketplace exception list.
type Checklist struct {
	Marketplace string `json:"marketplace"`
	Items       []Item `json:"items"`
}

// OpenCount is the number of open exception rows.
func (c Checklist) OpenCount() int {
	n := 0
	for _, it := range c.Items {
		if it.Open {
			n++
		}
	}
	return n
}

// Fixture is a deterministic DE checklist for smoke tests (no network).
func Fixture() Checklist {
	return Checklist{
		Marketplace: "DE",
		Items: []Item{
			{Kind: KindSSA, SKU: "thin-sku-1", Open: true, Note: "ssa window miss"},
			{Kind: KindAHT, SKU: "thin-sku-1", Open: true, Note: "aht over stated"},
			{Kind: KindBuyShipping, SKU: "thin-sku-2", Open: false, Note: "veeqo buy shipping ok"},
		},
	}
}

// Package alerts is the Slack/email cliff-alert seat.
//
// Slack/email before BHDR / deactivation cliffs. Format plugs only —
// no network I/O. Not a 3PL suite.
package alerts

// Seat is the locked package seat name.
const Seat = "alerts"

// Kind is a locked cliff alert class.
type Kind string

const (
	KindBHDR  Kind = "bhdr_cliff"
	KindDeact Kind = "deactivation_cliff"
)

// Channel is a destination plug. Delivery is out of this stub.
type Channel string

const (
	ChannelSlack Channel = "slack"
	ChannelEmail Channel = "email"
)

// Signal is one evaluated cliff finding. It is not a send.
type Signal struct {
	Kind        Kind   `json:"kind"`
	Marketplace string `json:"marketplace"`
	SKU         string `json:"sku"`
	Message     string `json:"message"`
}

// Format renders a signal for Slack or email. No send, no network.
func Format(ch Channel, s Signal) string {
	return string(ch) + ": " + string(s.Kind) + " " + s.Marketplace + " " + s.SKU
}

// FixtureBHDR is a deterministic DE BHDR cliff for smoke tests (no network).
func FixtureBHDR() Signal {
	return Signal{
		Kind:        KindBHDR,
		Marketplace: "DE",
		SKU:         "thin-sku-1",
		Message:     "business-hours delivery risk before deact cliff",
	}
}

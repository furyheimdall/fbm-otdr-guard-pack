package otdr

import "testing"

func TestSeat(t *testing.T) {
	if Seat != "otdr" {
		t.Fatalf("Seat = %q", Seat)
	}
}

func TestFixtureRiskHot(t *testing.T) {
	s := Fixture()
	if s.Marketplace != "DE" {
		t.Fatalf("marketplace = %q", s.Marketplace)
	}
	if got := s.Risk(); got != RiskHot {
		t.Fatalf("Risk = %q, want %q", got, RiskHot)
	}
}

func TestRiskBands(t *testing.T) {
	cases := []struct {
		rate float64
		want string
	}{
		{0.2, RiskOK},
		{0.5, RiskWarn},
		{1.0, RiskHot},
	}
	for _, c := range cases {
		s := Scorecard{Threshold: DefaultThreshold, Rate: c.rate}
		if got := s.Risk(); got != c.want {
			t.Fatalf("rate %v: Risk = %q, want %q", c.rate, got, c.want)
		}
	}
}

package degendb

import (
	"fmt"
	"time"

	"github.com/VividCortex/ewma"
	"github.com/benleb/gloomberg/internal/style"
)

// SaLiRas is a group of saLiRa instances for different timeframes
// type SaLiRas map[time.Duration]*saLiRa.
type SaLiRas []*SaLiRa

// NewSaLiRas creates a new SaLiRas instance.
func NewSaLiRas(timeframes []time.Duration) SaLiRas {
	saliras := SaLiRas(make([]*SaLiRa, 0))

	for _, timeframe := range timeframes {
		saliras = append(saliras, &SaLiRa{
			MovingAverage: ewma.NewMovingAverage(),
			Timeframe:     timeframe,
		})
	}

	return saliras
}

// single SaLiRa instance.
type SaLiRa struct {
	ewma.MovingAverage
	Timeframe     time.Duration
	CountSales    int
	CountListings int
	Previous      float64
}

func (s *SaLiRa) Pretty(faint bool) string {
	// coloring moving average salira
	saLiRaStyle := style.TrendGreenStyle

	current := s.Value()

	if s.Previous > current {
		saLiRaStyle = style.TrendRedStyle
	}

	return fmt.Sprint(
		style.CreateTrendIndicator(s.Previous, current),
		saLiRaStyle.Faint(faint).Render(fmt.Sprintf("%4.2f", current)),
	)
}

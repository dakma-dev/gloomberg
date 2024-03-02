package chainwatcher

import (
	"github.com/benleb/gloomberg/internal/models"
)

type Subscriber interface {
	Subscribe() (chan *models.TxWithLogs, error)
}

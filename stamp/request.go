package stamp

import (
	"github.com/bmf-san/akashigo/types"
)

// StampParams is parameters for stamp.
type StampParams struct {
	Token string          `json:"token"`
	Type  types.TypeStamp `json:"type"`
	// NOTE: The actual time stamp will be the time on the AKASHI stamping server side.
	// There seems to be no point in specifying it.
	// StampedAt time.Time `json:"stampedAt"`
	Timezone string `json:"timezone"`
}

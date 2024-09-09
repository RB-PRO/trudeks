package region

import (
	"strings"

	gocouter "trudeks/pkg/go-couter"
)

func TrudFilter(meets []gocouter.Meeting) []gocouter.Meeting {
	outmeet := make([]gocouter.Meeting, 0, len(meets))
	for _, meet := range meets {
		if len(meet.Category) != 0 {
			if strings.Contains(meet.Category[0],
				"Споры, возникающие из трудовых отношений") {
				outmeet = append(outmeet, meet)
			}
		}
	}
	return outmeet
}

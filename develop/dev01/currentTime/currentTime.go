package currentTime

import (
	"os"

	"github.com/beevik/ntp"
)

func GetCurrentTime() string {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	return time.String()
}

package symo

import (
	"github.com/brutella/hc/characteristic"
)

type Power struct {
	*characteristic.Int
}

func NewPower(val int) *Power {
	p := Power{characteristic.NewInt("")}
	p.Value = val
	p.Format = characteristic.FormatUInt64
	p.Perms = characteristic.PermsRead()

	return &p
}

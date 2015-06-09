package symo

import (
	"github.com/brutella/hc/model/characteristic"
)

type Power struct {
	*characteristic.Characteristic
}

func NewPower(watt int) *Power {
	p := Power{characteristic.NewCharacteristic(watt, characteristic.FormatInt, characteristic.CharTypeUnknown, characteristic.PermsRead())}
	p.Unit = PowerUnit

	return &p
}

func (p *Power) SetPower(watt int64) {
	p.SetValue(watt)
}

func (p *Power) Power() int64 {
	return p.GetValue().(int64)
}

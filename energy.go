package symo

import (
	"github.com/brutella/hc/model/characteristic"
)

type Energy struct {
	*characteristic.Characteristic
}

func NewEnergy(kwatt int) *Energy {
	p := Energy{characteristic.NewCharacteristic(kwatt, characteristic.FormatInt, characteristic.CharTypeUnknown, characteristic.PermsRead())}
	p.Unit = EnergyUnit

	return &p
}

func (e *Energy) SetEnergy(kwatt int64) {
	e.SetValue(kwatt)
}

func (e *Energy) Energy() int64 {
	return e.GetValue().(int64)
}

package symo

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

type Accessory struct {
	*accessory.Accessory

	Inverter *Service
}

func NewAccessory(info accessory.Info) *Accessory {
	a := accessory.New(info, accessory.TypeOther)
	svc := NewService(info.Name)

	a.AddService(svc.Service)

	return &Accessory{a, svc}
}

const typeInverter = "14FA9D31-FC94-4F98-B00D-4AE878523748" // Parce Measurement Service

type Service struct {
	*service.Service

	Name    *characteristic.Name
	Current *Power
	Today   *Power
	Year    *Power
	Total   *Power
}

func NewService(name string) *Service {
	nameChar := characteristic.NewName()
	nameChar.SetValue(name)

	pow := NewPower(0)
	pow.Type = TypeTotalPower
	pow.Unit = "W"
	pow.Description = "Leistung"

	today := NewPower(0)
	today.Type = TypeTotalPower
	today.Unit = "Wh"
	today.Description = "Heute"

	year := NewPower(0)
	year.Type = TypeTotalPower
	year.Unit = "Wh"
	year.Description = "Dieses Jahr"

	total := NewPower(0)
	total.Type = TypeTotalPower
	total.Unit = "Wh"
	total.Description = "Gesamt"

	svc := service.New(typeInverter)
	svc.AddCharacteristic(pow.Characteristic)
	svc.AddCharacteristic(today.Characteristic)
	svc.AddCharacteristic(year.Characteristic)
	svc.AddCharacteristic(total.Characteristic)
	svc.AddCharacteristic(nameChar.Characteristic)

	return &Service{svc, nameChar, pow, today, year, total}
}

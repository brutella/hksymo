package symo

import (
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
)

const typeGenerator = "A1"

type genService struct {
	*service.Service

	Name   *characteristic.Name
	CurPow *Power
	Today  *Energy
	Year   *Energy // this year
	Total  *Energy
}

func newGeneratorService(name string) *genService {
	nameChar := characteristic.NewName(name)
	pow := NewPower(0)
	pow.Type = TypePower

	today := NewEnergy(0)
	today.Type = TypeTodayEnergy

	year := NewEnergy(0)
	year.Type = TypeYearEnergy

	total := NewEnergy(0)
	total.Type = TypeTotalEnergy

	svc := service.New()
	svc.Type = typeGenerator
	svc.AddCharacteristic(pow.Characteristic)
	svc.AddCharacteristic(today.Characteristic)
	svc.AddCharacteristic(year.Characteristic)
	svc.AddCharacteristic(total.Characteristic)
	svc.AddCharacteristic(nameChar.Characteristic)

	return &genService{svc, nameChar, pow, today, year, total}
}

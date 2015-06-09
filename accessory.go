package symo

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/accessory"
)

type Generator struct {
	*accessory.Accessory

	svc *genService
}

func NewGenerator(info model.Info) *Generator {
	a := accessory.New(info)
	svc := newGeneratorService(info.Name)

	a.AddService(svc.Service)

	return &Generator{a, svc}
}

func (g *Generator) SetCurrentPower(watt int64) {
	g.svc.CurPow.SetPower(watt)
}

func (g *Generator) CurrentPower() int64 {
	return g.svc.CurPow.Power()
}

func (g *Generator) SetTodayEnergy(kwatt int64) {
	g.svc.Today.SetEnergy(kwatt)
}

func (g *Generator) TodayEnergy() int64 {
	return g.svc.Today.Energy()
}

func (g *Generator) SetYearEnergy(kwatt int64) {
	g.svc.Year.SetEnergy(kwatt)
}

func (g *Generator) YearEnergy() int64 {
	return g.svc.Year.Energy()
}

func (g *Generator) SetTotalEnergy(kwatt int64) {
	g.svc.Total.SetEnergy(kwatt)
}

func (g *Generator) TotalEnergy() int64 {
	return g.svc.Total.Energy()
}

package sg

import (
	"testing"

	"github.com/InformaticsResearchCenter/croot/model"
)

func TestSetAirport(t *testing.T) {
	gis := model.Gis{
		Longitude: 103.87068815153685,
		Latitude:  1.4155690028533765,
		Name:      "seletar airport",
	}
	SetAirport(&gis)
}

package sg

import (
	"log"
	"strconv"

	"github.com/InformaticsResearchCenter/croot/model"
	"github.com/jonas-p/go-shp"
)

func SetAirport(gis *model.Gis) {
	points := []shp.Point{
		shp.Point{gis.Longitude, gis.Longitude},
	}
	fields := []shp.Field{
		// String attribute field with length 25
		shp.StringField("NAME", 25),
	}
	shape, err := shp.Create("static/shp/seletar.shp", shp.POINT)
	if err != nil {
		log.Fatal(err)
	}
	defer shape.Close()
	shape.SetFields(fields)
	for n, point := range points {
		shape.Write(&point)

		// write attribute for object n for field 0 (NAME)
		shape.WriteAttribute(n, 0, "Seletar Airport "+strconv.Itoa(n+1))
	}
}

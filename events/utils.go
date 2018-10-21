package events

import (
	"engo.io/engo"
	"strconv"
)

func PointToXYStrings(p engo.Point) (x, y string) {
	x = strconv.FormatFloat(float64(p.X), 'f', 3, 64)
	y = strconv.FormatFloat(float64(p.Y), 'f', 3, 64)
	return
}
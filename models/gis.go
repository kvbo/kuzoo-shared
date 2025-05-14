package models

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Point struct {
	X float64
	Y float64
}

func (p Point) Value() (driver.Value, error) {
	return fmt.Sprintf("POINT(%f %f)", p.X, p.Y), nil
}

func (p *Point) Scan(val interface{}) error {
	var x, y float64
	_, err := fmt.Sscanf(string(val.([]byte)), "POINT(%f %f)", &x, &y)
	if err != nil {
		return err
	}
	p.X, p.Y = x, y
	return nil
}

type LineString struct {
	Points []Point
}

func (l LineString) Value() (driver.Value, error) {
	parts := make([]string, len(l.Points))
	for i, pt := range l.Points {
		parts[i] = fmt.Sprintf("%f %f", pt.X, pt.Y)
	}
	return fmt.Sprintf("LINESTRING(%s)", strings.Join(parts, ", ")), nil
}

func (l *LineString) Scan(val interface{}) error {
	var coords []Point
	str := string(val.([]byte))
	str = strings.TrimPrefix(str, "LINESTRING(")
	str = strings.TrimSuffix(str, ")")
	pairs := strings.Split(str, ",")
	for _, p := range pairs {
		var x, y float64
		_, err := fmt.Sscanf(strings.TrimSpace(p), "%f %f", &x, &y)
		if err != nil {
			return err
		}
		coords = append(coords, Point{X: x, Y: y})
	}
	l.Points = coords
	return nil
}

type Polygon struct {
	Rings [][]Point
}

func (p Polygon) Value() (driver.Value, error) {
	var ringStrings []string
	for _, ring := range p.Rings {
		var pts []string
		for _, pt := range ring {
			pts = append(pts, fmt.Sprintf("%f %f", pt.X, pt.Y))
		}
		ringStrings = append(ringStrings, fmt.Sprintf("(%s)", strings.Join(pts, ", ")))
	}
	return fmt.Sprintf("POLYGON(%s)", strings.Join(ringStrings, ", ")), nil
}

func (p *Polygon) Scan(val interface{}) error {
	raw := string(val.([]byte))
	raw = strings.TrimPrefix(raw, "POLYGON(")
	raw = strings.TrimSuffix(raw, ")")
	rings := strings.Split(raw, "),(")
	for i := range rings {
		rings[i] = strings.TrimPrefix(rings[i], "(")
		rings[i] = strings.TrimSuffix(rings[i], ")")
	}

	var parsed [][]Point
	for _, ring := range rings {
		pairs := strings.Split(ring, ",")
		var pts []Point
		for _, pair := range pairs {
			var x, y float64
			_, err := fmt.Sscanf(strings.TrimSpace(pair), "%f %f", &x, &y)
			if err != nil {
				return err
			}
			pts = append(pts, Point{X: x, Y: y})
		}
		parsed = append(parsed, pts)
	}
	p.Rings = parsed
	return nil
}

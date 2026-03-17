package models

import (
	"encoding/json"
)

type CelestialBody struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

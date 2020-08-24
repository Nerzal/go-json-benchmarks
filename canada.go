package bench

type Canada struct {
	Features []Features `json:"features,omitempty"`
	Type     string     `json:"type,omitempty"`
}

type Features struct {
	Geometry   Geometry   `json:"geometry,omitempty"`
	Properties Properties `json:"properties,omitempty"`
	Type       string     `json:"type,omitempty"`
}

type Properties struct {
	Name string `json:"name,omitempty"`
}

type Geometry struct {
	Coordinates [][][]float64 `json:"coordinates,omitempty"`
	Type        string        `json:"type,omitempty"`
}

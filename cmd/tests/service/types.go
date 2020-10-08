package service

type distData struct {
	Version string `json:"version"`
	Build   string `json:"build"`
}

type status struct {
	API bool `json:"api"`
	DB  bool `json:"db"`
}

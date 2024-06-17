package dtos

type LookupResponse struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

type MuscleGroupLookupResponse struct {
	Id        int              `json:"id"`
	Label     string           `json:"label"`
	Exercises []LookupResponse `json:"exercises"`
}

package entities

type Topic struct {
	Name       string `json:"name"`
	Partitions []int  `json:"partitions"`
}

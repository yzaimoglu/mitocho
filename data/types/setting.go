package types

type Setting struct {
	BaseModel
	Name  string `json:"name"`
	Value string `json:"value"`
}

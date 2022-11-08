package bucket

type File struct {
	Name string      `json:"name"`
	Link string      `json:"url"`
	Info interface{} `json:"info"`
}

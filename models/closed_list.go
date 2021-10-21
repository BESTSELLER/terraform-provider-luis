package models

// ClosedListEntity holds the JSON reponse from luis.ai
type ClosedListEntity struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TypeID       int    `json:"typeId"`
	ReadableType string `json:"readableType"`
	SubLists     []struct {
		ID            int      `json:"id"`
		CanonicalForm string   `json:"canonicalForm"`
		List          []string `json:"list"`
	} `json:"subLists"`
	Roles []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"roles"`
}

// ClosedListSublist holds the JSON request to luis.ai
type ClosedListSublist struct {
	CanonicalForm string   `json:"canonicalForm"`
	List          []string `json:"list"`
}

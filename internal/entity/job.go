package entity

type Job struct {
	ID              string `json:"id"`
	URL             string `json:"url"`
	Title           string `json:"title"`
	Location        string `json:"location"`
	Level           string `json:"level"` // senior, mid-level, entry level
	Description     string `json:"description"`
	Type            string `json:"type"` // fulltime, part-time
	Company         string `json:"company"`
	CompanyIndustry string `json:"company_industry"` // software development, IT service and IT consulting
	Remote          bool   `json:"Remote"`
	Relocation      bool   `json:"Relocation"`
}

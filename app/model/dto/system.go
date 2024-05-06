package dto

type Stat struct {
	Runtime Runtime `json:"runtime"`
	CPU     CPU     `json:"cpu"`
	Disk    Disk    `json:"disk"`
	Memory  Memory  `json:"memory"`
}
type Runtime struct {
	NpmVersion  string `json:"npmVersion"`
	NodeVersion string `json:"nodeVersion"`
	GoVersion   string `json:"goVersion"`
	Os          string `json:"os"`
	Arch        string `json:"arch"`
}
type CoresLoad struct {
	RawLoad     int     `json:"rawLoad"`
	RawLoadIdle int     `json:"rawLoadIdle"`
	CoreLoad    float64 `json:"coreLoad"`
}
type CPU struct {
	Manufacturer       string      `json:"manufacturer"`
	Brand              string      `json:"brand"`
	PhysicalCores      int         `json:"physicalCores"`
	Model              string      `json:"model"`
	Speed              float64     `json:"speed"`
	RawCurrentLoad     int         `json:"rawCurrentLoad"`
	RawCurrentLoadIdle int         `json:"rawCurrentLoadIdle"`
	AllLoad            float64     `json:"allLoad"`
	CoresLoad          []CoresLoad `json:"coresLoad"`
}
type Disk struct {
	Size            int64  `json:"size"`
	Used            int64  `json:"used"`
	Available       int64  `json:"available"`
	ConstructorName string `json:"_constructor-name_"`
}
type Memory struct {
	Total     int64 `json:"total"`
	Available int64 `json:"available"`
}

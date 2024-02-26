package entities

type Covid struct {
	ConfirmDate    string
	No             int
	Age            *int
	Gender         string
	GenderEn       string
	Nation         string
	NationEn       string
	Province       string
	ProvinceId     int
	District       string
	ProvinceEn     string
	StatQuarantine int
}

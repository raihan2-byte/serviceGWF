package input

type InputAddressBuyer struct {
	Province    string `json:"province" binding:"required"`
	City        string `json:"city" binding:"required"`
	SubDistrict string `json:"sub_district" binding:"required"`
	HomeAddress string `json:"home_address" binding:"required"`
	Courier     string `json:"courier" binding:"required"`
}

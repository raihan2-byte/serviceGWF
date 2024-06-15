package entity

type Province struct {
	ID   string `json:"province_id"`
	Name string `json:"province"`
}

type ProvinceData struct {
	Result []Province `json:"results"`
}

type ProvinceResponse struct {
	ProvinceData ProvinceData `json:"rajaongkir"`
}

type City struct {
	ID         string `json:"city_id"`
	Name       string `json:"city_name"`
	PostalCode string `json:"postal_code"`
	ProvinceID string `json:"province_id"`
}

type CityData struct {
	Result []City `json:"results"`
}

type CityResponse struct {
	CityData CityData `json:"rajaongkir"`
}

// input
type ShippingFeeParams struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
	HomeAddress string `json:"home_address"`
}

type OriginDetails struct {
	CityID   string `json:"city_id"`
	CityName string `json:"city_name"`
}

type DestinationDetails struct {
	CityID     string `json:"city_id"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
	Province   string `json:"province"`
}

type OngkirCost struct {
	Service     string       `json:"service"`
	Description string       `json:"Description"`
	Cost        []CostDetail `json:"cost"`
}

type CostDetail struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type OngkirResult struct {
	Code  string       `json:"code"`
	Name  string       `json:"name"`
	Costs []OngkirCost `json:"costs"`
}

type OngkirData struct {
	OriginDetails      OriginDetails      `json:"origin_details"`
	DestinationDetails DestinationDetails `json:"destination_details"`
	Result             []OngkirResult     `json:"results"`
}

type OngkirResponse struct {
	OngkirData OngkirData `json:"rajaongkir"`
}

type ShippingFeeOption struct {
	Service            string             `json:"service"`
	Fee                int                `json:"fee"`
	DestinationDetails DestinationDetails `json:"destination_details"`
}

type ApplyShippingResponse struct {
	ID          int
	ShippingFee int    `json:"shipping_fee"`
	TotalWeight int    `json:"total_weight"`
	CityID      string `json:"city_id"`
	CityName    string `json:"city_name"`
	PostalCode  string `json:"postal_code"`
	Province    string `json:"province"`
	HomeAddress string `json:"home_address" binding:"required"`
	Courier     string `json:"courier"`
	// DestinationDetails DestinationDetails `json:"destination_details"`
	UserID int
	User   User `gorm:"foreignKey:UserID"`
}

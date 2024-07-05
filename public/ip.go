package public

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPInfo struct {
	Address string `json:"address"`
	Content struct {
		AddressDetail struct {
			Province     string `json:"province"`
			City         string `json:"city"`
			District     string `json:"district"`
			StreetNumber string `json:"street_number"`
			Adcode       string `json:"adcode"`
			Street       string `json:"street"`
			CityCode     int    `json:"city_code"`
		} `json:"address_detail"`
		Point struct {
			Y string `json:"y"`
			X string `json:"x"`
		} `json:"point"`
		Address string `json:"address"`
	} `json:"content"`
	Status int `json:"status"`
}

func GetIPInfo(ip string) IPInfo {
	apiKey := "gMbGAcMvKWGSNaHVfwgN1lC2quFVijMF"
	url := fmt.Sprintf("http://api.map.baidu.com/location/ip?ak=%s&ip=%s", apiKey, ip)
	var ipInfo IPInfo
	resp, err := http.Get(url)
	if err != nil {
		return ipInfo
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&ipInfo)
	return ipInfo
}

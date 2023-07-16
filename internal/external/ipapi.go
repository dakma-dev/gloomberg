package external

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/utils"
)

type IPInfo struct {
	IP                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryName        string  `json:"country_name"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
	Hostname           string  `json:"hostname"`
	Error              bool    `json:"error"`
	Reason             string  `json:"reason"`
	Reserved           bool    `json:"reserved"`
}

func GetIPInfo(ctx context.Context, ipAddr net.Addr) (*IPInfo, error) {
	// get ip address without port
	idx := strings.LastIndex(ipAddr.String(), ":")
	if idx == -1 {
		return nil, errors.New("could not find ':' in remote address")
	}

	addrWithoutPort := ipAddr.String()[:idx]

	// build url
	url := fmt.Sprintf("https://ipapi.co/%s/json/", addrWithoutPort)

	// create required header
	header := http.Header{}
	header.Add("User-Agent", "ipapi.co/#go-v1.5")

	gbl.Log.Debugf("🩲 getting IPInfo fron ipapi.co for %+v", addrWithoutPort)

	response, err := utils.HTTP.GetWithTLS12AndHeader(ctx, url, header)
	if err != nil && os.IsTimeout(err) {
		gbl.Log.Debugf("⌛️ timeout while fetching ipapi.co IPInfo: %+v", err.Error())

		return nil, err
	} else if err != nil {
		gbl.Log.Debugf("❌ ipapi.co IPInfo error: %+v", err.Error())

		return nil, err
	}

	defer response.Body.Close()

	gbl.Log.Debugf("ipapi.co IPInfo response status: %s", response.Status)

	// map to EventSignatures
	var ipInfo IPInfo

	// read response body
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		gbl.Log.Debugf("❌ error reading ipapi.co IPInfo response: %+v", err.Error())

		return nil, err
	}

	gbl.Log.Debugf("ipapi.co IPInfo response body: %s", string(bodyBytes))

	// unmarshal body
	err = json.Unmarshal(bodyBytes, &ipInfo)
	if err != nil {
		gbl.Log.Debugf("❌ error decoding ipapi.co IPInfo response: %+v", err.Error())

		return nil, err
	} else if ipInfo.Error {
		gbl.Log.Debugf("❌ ipapi.co IPInfo error: %+v", ipInfo.Reason)

		ipInfo.CountryCode = "00"

		return &ipInfo, errors.New(ipInfo.Reason)
	}

	return &ipInfo, nil
}

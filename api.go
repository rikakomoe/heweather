package heweather

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	APIEndPoint = "https://free-api.heweather.com/s6/%s?%s"
)

func (api WeatherAPI) MakeRequest(endpoint string) (APIResponseData, error) {
	v := url.Values{}
	v.Add("location", api.Location)
	v.Add("username", api.Username)
	if api.Language != "" {
		v.Add("lang", api.Language)
	}
	if api.Unit != "" {
		v.Add("unit", api.Unit)
	}
	t := time.Now()
	v.Add("t", strconv.FormatInt(t.Unix(), 10))
	v.Add("key", api.Key)
	u := fmt.Sprintf(APIEndPoint, endpoint, v.Encode())
	resp, err := http.Get(u)
	if err != nil {
		return APIResponseData{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIResponseData{}, err
	}
	var apiresp APIResponse
	err = json.Unmarshal(body, &apiresp)
	if err != nil {
		return APIResponseData{}, err
	}
	if len(apiresp.HeWeather6) == 0 {
		return APIResponseData{}, errors.New("empty response")
	}
	return apiresp.HeWeather6[0], nil
}

func (api WeatherAPI) Forecast() ([]DailyForecast, error) {
	resp, err := api.MakeRequest("weather/forecast")
	if err != nil {
		return []DailyForecast{}, err
	}
	return resp.DailyForecast, nil
}

func (api WeatherAPI) Now() (RealTimeWeather, error) {
	resp, err := api.MakeRequest("weather/now")
	if err != nil {
		return RealTimeWeather{}, err
	}
	return resp.RealTimeWeather, nil
}

func (api WeatherAPI) Hourly() ([]HourlyWeather, error) {
	resp, err := api.MakeRequest("weather/hourly")
	if err != nil {
		return []HourlyWeather{}, err
	}
	return resp.HourlyWeather, nil
}

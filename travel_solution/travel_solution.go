package travel_solution

import (
	"net/http"
	"os"
	"utils"
)

func SendRequest(client *http.Client) (string, error) {
	endpoint := os.Getenv("TRAVEL_SOLUTION_ENDPOINT")
	method := "POST"
	headers := map[string]string{
		"Content-Type":        "text/xml",
		"X-IBM-Client-Id":     os.Getenv("X_IBM_CLIENT_ID"),
		"X-IBM-Client-Secret": os.Getenv("X_IBM_CLIENT_SECRET"),
	}

	body := CreateRequestBody()

	respBody, err := utils.SendHTTPRequest(client, method, endpoint, body, headers)
	if err != nil {
		// handle error
		return "", err
	}
	return string(respBody), nil
}

func CreateRequestBody() []byte {
	requestBody := "<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:xsd='http://www.w3.org/2001/XMLSchema' xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance' xmlns:ns2='http://PicoServiceModel/Sale/SolutionEngine/TravelSolution/' xmlns:soapenv='http://schemas.xmlsoap.org/soap/envelope/'><soapenv:Body><ns2:search><searchRequest><serviceContext><correlationId>123</correlationId><channelId><id>452353</id></channelId></serviceContext><searchCriteria><departureTimesStart>2024-06-30T03:30:00</departureTimesStart><departureTimesEnd>2024-07-01T03:29:00</departureTimesEnd><departure><id>830008409</id></departure><arrival><id>830006900</id></arrival><direction>ONE_WAY</direction><channelFilter><id>452353</id></channelFilter><parameters><maxNumberOfChanges>2</maxNumberOfChanges></parameters></searchCriteria></searchRequest></ns2:search></soapenv:Body></soapenv:Envelope>"

	return []byte(requestBody)
}

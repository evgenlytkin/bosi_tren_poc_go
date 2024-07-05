package admissible_service

import (
	"net/http"
	"os"
	"utils"
)

func SendRequest(client *http.Client) (string, error) {
	endpoint := os.Getenv("ADMISSIBLE_DATA_ENDPOINT")
	method := "POST"
	headers := map[string]string{
		"Content-Type":        "text/xml",
		"X-IBM-Client-Id":     os.Getenv("X_IBM_CLIENT_ID"),
		"X-IBM-Client-Secret": os.Getenv("X_IBM_CLIENT_SECRET"),
		"SOAPAction":          os.Getenv("SOAP_ACTION_SERVICE"),
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
	requestBody := "<?xml version='1.0' encoding='UTF-8'?><soapenv:Envelope xmlns:xsd='http://www.w3.org/2001/XMLSchema' xmlns:xsi='http://www.w3.org/2001/XMLSchema-instance' xmlns:chan='http://PicoServiceModel/Channels/ChannelEngine/ChannelsCoordinator/' xmlns:soapenv='http://schemas.xmlsoap.org/soap/envelope/'><soapenv:Body><chan:getAdmissibleServices><getAdmissibleServicesRequest><id><id>452353</id></id></getAdmissibleServicesRequest></chan:getAdmissibleServices></soapenv:Body></soapenv:Envelope>"

	return []byte(requestBody)
}

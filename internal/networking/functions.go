package networking

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os/exec"
	"repositories/cryptobros/internal/config"
	"time"
)

var (
	EndpointIfconfig = "https://ifconfig.co/json"
	SleepTime        = 10
	HTTPClient       = &http.Client{Timeout: 30 * time.Second}
)

// TODO : Find another way to retrieve new IPv6 - This is almost impossible to work in server environment
func RenewIPv6() error {
	currentIPv6, err := GetIPv6()
	if err != nil {
		return errors.New("error RenewIPv6 cannot getIPv6:" + err.Error())
	}

	for {
		log.Println("starting the loop to renew IPv6 - current IPv6:", currentIPv6)

		cmd := exec.Command("bash", "-c", config.RestartNetworkManagerCLICommand)
		_, err = cmd.CombinedOutput()
		if err != nil {
			return errors.New("error RenewIPv6 cannot restart NetworkManager: " + err.Error())
		}

		newIPv6 := ""
		for {
			newIPv6, err = GetIPv6()
			if err != nil {
				log.Println("trying to fetch the new ipv6")
				time.Sleep(time.Duration(SleepTime) * time.Second)
				continue
				// TODO : After some retries, return error
				// return errors.New("error RenewIPv6 cannot getIPv6:" + err.Error())
			}
			break
		}

		if newIPv6 != currentIPv6 {
			log.Println("IPv6 changed - new IPv6:", newIPv6)
			break
		}
	}

	return nil
}

// TODO : It would be a good idea to retrieve the IPv6 from internal network manager - not from an external third party
func GetIPv6() (string, error) {
	body, err := GetRequest(EndpointIfconfig)
	if err != nil {
		return "", errors.New("error getIPv6 cannot getRequest: " + err.Error())
	}

	var ifconfig Ifconfig
	err = json.Unmarshal(body, &ifconfig)
	if err != nil {
		return "", errors.New("error getIPv6 cannot Unmarshal: " + err.Error())
	}

	ipv6 := ifconfig.IP
	if ipv6 == "" {
		return "", errors.New("error getIPv6 empty IP")
	}

	return ipv6, nil
}

func GetRequest(url string) ([]byte, error) {
	HTTPClient.CloseIdleConnections()
	resp, err := HTTPClient.Get(url)
	if err != nil {
		return []byte{}, errors.New("error GetRequest on Get:" + err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, errors.New("error GetRequest on ReadAll:" + err.Error())
	}

	return body, nil
}

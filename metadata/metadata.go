package metadata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Metadata struct {
	Client      *http.Client
	Endpoint    string
	ResourceMap map[string]string
}

func New(endpoint string) *Metadata {
	resourceMap := make(map[string]string)
	resourceMap["ami-id"] = "latest/meta-data/ami-id"
	resourceMap["ami-launch-index"] = "latest/meta-data/ami-launch-index"
	resourceMap["availability-zone"] = "latest/meta-data/placement/availability-zone"
	resourceMap["hostname"] = "latest/meta-data/hostname"
	resourceMap["instance-id"] = "latest/meta-data/instance-id"
	resourceMap["instance-type"] = "latest/meta-data/instance-type"
	resourceMap["local-hostname"] = "latest/meta-data/local-hostname"
	resourceMap["local-ipv4"] = "latest/meta-data/local-ipv4"
	resourceMap["profile"] = "latest/meta-data/profile"
	resourceMap["public-hostname"] = "latest/meta-data/public-hostname"
	resourceMap["public-ipv4"] = "latest/meta-data/public-ipv4"

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return &Metadata{
		Client:      client,
		Endpoint:    endpoint,
		ResourceMap: resourceMap,
	}
}

func (m *Metadata) Get(key string) (string, error) {
	resource, ok := m.ResourceMap[key]

	if !ok {
		return "", fmt.Errorf("key %s could not be found", key)
	}

	res, err := m.Client.Get(m.Endpoint + "/" + resource)
	defer res.Body.Close()

	if err != nil {
		return "", fmt.Errorf("key %s could not be found", key)
	}

	if res.StatusCode != 200 {
		return "", fmt.Errorf("key %s could not be found", key)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("key %s could not be found", key)
	}

	return string(body), nil
}

package airthings

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

type Client struct {
	Endpoint string
	http     *http.Client
	conf     *clientcredentials.Config
}

func Authorize(ctx context.Context, clientId, clientSecret string, scopes []string) (*Client, error) {
	if scopes == nil {
		scopes = []string{"read:device:current_values"}
	}
	conf := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		TokenURL:     "https://accounts-api.airthings.com/v1/token",
	}

	httpClient := conf.Client(ctx)
	return &Client{Endpoint: "https://ext-api.airthings.com/", http: httpClient, conf: conf}, nil
}

type Authentication struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`

	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`

	RefreshToken string `json:"refresh_token"`

	Scope []string `json:"scope"`
}

type ListDevicesOptions struct {
	ShowInactive   bool
	OrganizationId string
	UserGroupId    string
}

// ListDevices get devices belonging to the account
func (c *Client) ListDevices(opts ListDevicesOptions) ([]Device, error) {
	var query []string
	if opts.ShowInactive {
		query = append(query, "showInactive=true")
	}
	if opts.OrganizationId != "" {
		query = append(query, fmt.Sprintf("organizationId=%s", url.QueryEscape(opts.OrganizationId)))
	}
	if opts.UserGroupId != "" {
		query = append(query, fmt.Sprintf("userGroupId=%s", url.QueryEscape(opts.UserGroupId)))
	}
	var response struct {
		Devices []Device `json:"devices"`
	}
	err := c.get("/v1/devices?"+strings.Join(query, "&"), &response)
	return response.Devices, err
}

type GetDeviceOptions struct {
	SerialNumber   string
	OrganizationId string
	UserGroupId    string
}

var ErrNoSerialNumber = errors.New("missing serial number")

// GetDevice returns information about a particular device
func (c *Client) GetDevice(opts GetDeviceOptions) (Device, error) {
	if opts.SerialNumber == "" {
		return Device{}, ErrNoSerialNumber
	}
	var query []string
	if opts.OrganizationId != "" {
		query = append(query, fmt.Sprintf("organizationId=%s", url.QueryEscape(opts.OrganizationId)))
	}
	if opts.UserGroupId != "" {
		query = append(query, fmt.Sprintf("userGroupId=%s", url.QueryEscape(opts.UserGroupId)))
	}
	var device Device
	err := c.get(fmt.Sprintf("/v1/devices/%s?%s", opts.SerialNumber, strings.Join(query, "&")), &device)
	return device, err
}

type GetLatestSamplesOptions struct {
	SerialNumber   string // required
	OrganizationId string
	UserGroupId    string
}

func (c *Client) GetLatestSamples(opts GetLatestSamplesOptions) (map[SensorType]interface{}, error) {
	if opts.SerialNumber == "" {
		return nil, ErrNoSerialNumber
	}
	var query []string
	if opts.OrganizationId != "" {
		query = append(query, fmt.Sprintf("organizationId=%s", url.QueryEscape(opts.OrganizationId)))
	}
	if opts.UserGroupId != "" {
		query = append(query, fmt.Sprintf("userGroupId=%s", url.QueryEscape(opts.UserGroupId)))
	}
	var response struct {
		Data map[SensorType]interface{} `json:"data"`
	}
	err := c.get(fmt.Sprintf("/v1/devices/%s/latest-samples?%s", opts.SerialNumber, strings.Join(query, "&")), &response)
	return response.Data, err
}

type Device struct {
	SerialNumber string         `json:"id"` // this is the serial number
	DeviceType   DeviceType     `json:"deviceType"`
	Sensors      []SensorType   `json:"sensors"`
	Segment      DeviceSegment  `json:"segment"`
	Location     DeviceLocation `json:"location"`
}

type DeviceSegment struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StartedAt string `json:"started"`
	Active    bool   `json:"active"`
}

type DeviceLocation struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c *Client) get(path string, output interface{}) error {
	for attempts := 0; attempts < 3; attempts++ {
		res, err := c.http.Get(c.Endpoint + path)
		if err != nil {
			return err
		}
		if res.StatusCode >= 500 {
			res.Body.Close()
			continue
		}
		defer res.Body.Close()
		if res.StatusCode < 200 || res.StatusCode >= 300 {
			return fmt.Errorf("bad response: %d %s", res.StatusCode, res.Status)
		}
		buf, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read: %w", err)
		}
		return json.Unmarshal(buf, output)
	}
	return fmt.Errorf("failed after 3 attempts")
}

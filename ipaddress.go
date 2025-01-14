package cherrygo

import (
	"fmt"
	"strings"
)

// GetIP interface metodas isgauti team'sus
type GetIP interface {
	List(projectID string, ipID string, opts *GetOptions) (IPAddresses, *Response, error)
	Create(projectID string, request *CreateIPAddress) (IPAddresses, *Response, error)
	Remove(projectID string, request *RemoveIPAddress) (IPAddresses, *Response, error)
	Update(projectID string, ipID string, request *UpdateIPAddress) (IPAddresses, *Response, error)
}

// IPClient paveldi client
type IPClient struct {
	client *Client
}

// CreateIPAddress fields for adding addition IP address
type CreateIPAddress struct {
	Type       string `json:"type,omitempty"`
	Region     string `json:"region,omitempty"`
	PtrRecord  string `json:"ptr_record,omitempty"`
	ARecord    string `json:"a_record,omitempty"`
	RoutedTo   string `json:"routed_to,omitempty"`
	AssignedTo string `json:"assigned_to,omitempty"`
	TargetedTo string `json:"targeted_to,omitempty"`
	Tags       *map[string]string `json:"tags,omitempty"`
}

// UpdateIPAddress fields for updating IP address
type UpdateIPAddress struct {
	PtrRecord  string             `json:"ptr_record,omitempty"`
	ARecord    string             `json:"a_record,omitempty"`
	RoutedTo   string             `json:"routed_to,omitempty"`
	AssignedTo string             `json:"assigned_to,omitempty"`
	TargetedTo string             `json:"targeted_to,omitempty"`
	Tags       *map[string]string `json:"tags,omitempty"`
}

// RemoveIPAddress fields for removing IP address
type RemoveIPAddress struct {
	ID string `json:"id,omitempty"`
}

// List func lists teams
func (i *IPClient) List(projectID string, ipID string, opts *GetOptions) (IPAddresses, *Response, error) {
	//root := new(teamRoot)

	path := strings.Join([]string{baseIPSPath, projectID, endIPSPath, ipID}, "/")
	pathQuery := opts.WithQuery(path)

	var trans IPAddresses

	resp, err := i.client.MakeRequest("GET", pathQuery, nil, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}

	return trans, resp, err
}

// Create function orders new floating IP address
func (i *IPClient) Create(projectID string, request *CreateIPAddress) (IPAddresses, *Response, error) {

	var trans IPAddresses

	ipAddressPath := strings.Join([]string{baseIPSPath, projectID, endIPSPath}, "/")

	resp, err := i.client.MakeRequest("POST", ipAddressPath, request, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}
	return trans, resp, err

}

// Update function updates existing floating IP address
func (i *IPClient) Update(projectID string, ipID string, request *UpdateIPAddress) (IPAddresses, *Response, error) {

	var trans IPAddresses

	ipAddressPath := strings.Join([]string{baseIPSPath, projectID, endIPSPath, ipID}, "/")

	resp, err := i.client.MakeRequest("PUT", ipAddressPath, request, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}

	return trans, resp, err
}

// Remove function remove existing floating IP address
func (i *IPClient) Remove(projectID string, request *RemoveIPAddress) (IPAddresses, *Response, error) {

	var trans IPAddresses

	ipAddressPath := strings.Join([]string{baseIPSPath, projectID, endIPSPath, request.ID}, "/")

	resp, err := i.client.MakeRequest("DELETE", ipAddressPath, request, &trans)
	if err != nil {
		err = fmt.Errorf("Error: %v", err)
	}
	return trans, resp, err
}

package maas

import (
	"fmt"
	"net/url"
)

func (c *Client) GetFabrics(filters ...string) (*[]Fabric, error) {
	params, err := FilterToUrlParams(filters...)
	if err != nil {
		return nil, err
	}

	var rsp []byte
	rsp, err = c.TurnResponse(c.Get("fabrics", "", params))
	if err != nil {
		return nil, err
	}

	fabrics := make([]Fabric, 0, 2)
	err = Unmarshal(rsp, &fabrics)
	if err != nil {
		return nil, err
	}
	return &fabrics, nil
}

func (c *Client) CreateFabric(name, description, classType string) (*Fabric, error) {
	params := url.Values{}

	if len(name) != 0 {
		params.Add("name", name)
	}
	if len(description) != 0 {
		params.Add("description", description)
	}
	if len(classType) != 0 {
		params.Add("class_type", classType)
	}

	rsp, err := c.TurnResponse(c.Post("fabrics", "", params, nil))
	if err != nil {
		return nil, err
	}

	fabric := Fabric{}
	err = Unmarshal(rsp, &fabric)
	if err != nil {
		return nil, err
	}
	return &fabric, nil
}

func (c *Client) GetFabric(id int) (*Fabric, error) {
	resourceUrl := fmt.Sprintf("fabrics/%d", id)
	rsp, err := c.TurnResponse(c.Get(resourceUrl, "", nil))

	fabric := Fabric{}
	err = Unmarshal(rsp, &fabric)
	if err != nil {
		return nil, err
	}
	return &fabric, nil
}

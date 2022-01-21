package maas

import (
	"encoding/json"
	"net/url"
	"strconv"
)

func (n *Node) GetBlockDevices() ([]*BlockDevice, error) {
	c := n.getClient()
	if c == nil {
		return nil, ErrEmptyClient
	}

	rsp, err := c.TurnResponse(c.Get(JoinURLs(n.ResourceUri, "blockdevices"), "", nil))
	if err != nil {
		return nil, err
	}
	res := make([]*BlockDevice, 0, 2)
	if err := json.Unmarshal(rsp, &res); err != nil {
		return nil, err
	}

	for _, b := range res {
		b.setClient(n.getClient())
		b.recursiveClient()
	}

	return res, nil
}

func (n *Node) CreateBlockDevices(name, model, serial, idPath string, size, blockSize int) (*BlockDevice, error) {
	c := n.getClient()
	if c == nil {
		return nil, ErrEmptyClient
	}

	if len(name) == 0 {
		return nil, ErrRequireParametersEmpty
	}

	params := url.Values{}
	if len(model) > 0 {
		params.Add("model", model)
	}
	if len(serial) > 0 {
		params.Add("serial", serial)
	}
	if len(model) > 0 {
		params.Add("model", model)
	}
	if size > 0 {
		params.Add("block_size", strconv.Itoa(size))
	}
	if blockSize > 0 {
		params.Add("block_size", strconv.Itoa(blockSize))
	}
	c.Post(JoinURLs(n.ResourceUri, "blockdevices"), "", params, nil)

}

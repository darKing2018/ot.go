package session

import "ot/ot/selection"

type Client struct {
	Name      string              `json:"name"`
	Selection selection.Selection `json:"selection"`
}

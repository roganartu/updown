package main

import (
	"crypto/md5"
	"time"
)

type Message struct {
	Client   *Client
	Body     []byte
	Press    *Press
	Received time.Time
}

type Client struct {
	IP        string
	UserAgent string
}

// Fingerprint returns a hash fingerprint of the client details.
func (c *Client) Fingerprint() string {
	a := md5.Sum([]byte(c.IP + c.UserAgent))
	return string((&a)[:])
}

package config

import "os"

type Host struct {
	IP   string
	Port string
}

func (h *Host) Init() Host {
	h.IP = os.Getenv("IP")
	h.Port = os.Getenv("PORT")
	return *h
}

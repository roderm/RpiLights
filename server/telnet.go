package server

import (
	"playground/light"
)

type TelnetSever struct{
	light light.RpiLight
}

type telnetHandler struct {}
func (t *TelnetSever) Server(port int){



}

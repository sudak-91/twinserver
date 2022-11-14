package main

import (
	"fmt"

	"github.com/sudak-91/engcore_modbus"
)

func main() {
	Registers, err := engcore_modbus.NewModbusRegisters(20, 20, 20, 20)
	if err != nil {
		panic(err)
	}
	Serv := engcore_modbus.NewModbusServer("0.0.0.0", 502, Registers)
	go Serv.StartServer()
	quit := make(chan bool)
	<-quit
	fmt.Println("Exit")

}

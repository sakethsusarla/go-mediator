package main

import "fmt"

type SmartHomeMediator interface {
	SendCommand(sender string, deviceType string, action string)
}

type Device interface {
	HandleCommand(action string)
	GetType() string
}

type User interface {
	ControlDevice(deviceType string, action string)
}

type SmartFloorLamp struct {
}

func (f *SmartFloorLamp) HandleCommand(action string) {
	fmt.Printf("Floor Lamp: Okay, performing the action %s on the floor lamp 🛋️\n", action)
}

func (f *SmartFloorLamp) GetType() string {
	return "lamp"
}

type SmartLock struct {
}

func (l *SmartLock) HandleCommand(action string) {
	fmt.Printf("Smart Lock: The door of Apartment 4A 2311 North Robles Avenue is now %sed\n", action)
}

func (l *SmartLock) GetType() string {
	return "lock"
}

type SmartSpeaker struct {
}

func (s *SmartSpeaker) HandleCommand(action string) {
	fmt.Printf("Smart Speaker: Okay, about to %s 🎵\n", action)
}

func (s *SmartSpeaker) GetType() string {
	return "speaker"
}

type SmartTelevision struct {
}

func (t *SmartTelevision) HandleCommand(action string) {
	fmt.Printf("Smart TV: Okay, about to %s 📺\n", action)
}

func (t *SmartTelevision) GetType() string {
	return "tv"
}

type SmartThermostat struct {
}

func (th *SmartThermostat) HandleCommand(action string) {
	fmt.Printf("Smart Thermostat: Okay, the thermostat is about to %s 🌡️\n", action)
}

func (th *SmartThermostat) GetType() string {
	return "thermostat"
}

type BazingaSmartHomeController struct {
	devices map[string]Device
}

func NewBazingaSmartHomeController() *BazingaSmartHomeController {
	return &BazingaSmartHomeController{
		devices: make(map[string]Device),
	}
}

func (c *BazingaSmartHomeController) RegisterDevice(device Device) {
	c.devices[device.GetType()] = device
}

func (c *BazingaSmartHomeController) SendCommand(sender string, deviceType string, action string) {
	if device, isPresent := c.devices[deviceType]; isPresent {
		fmt.Printf("👍 Sure, %s. Sending your command to the %s now.\n", sender, deviceType)
		device.HandleCommand(action)
	} else {
		fmt.Printf("👎 Sorry, %s. The %s device is not registered in the smart home system. Unable to perform the action: %s\n", sender, deviceType, action)
	}
}

type Character struct {
	name     string
	mediator SmartHomeMediator
}

func NewCharacter(name string, mediator SmartHomeMediator) *Character {
	return &Character{
		name:     name,
		mediator: mediator,
	}
}

func (ch *Character) ControlDevice(deviceType string, action string) {
	ch.mediator.SendCommand(ch.name, deviceType, action)
}

func main() {
	mediator := NewBazingaSmartHomeController()

	smartTv := &SmartTelevision{}
	smartSpeaker := &SmartSpeaker{}
	smartFloorLamp := &SmartFloorLamp{}
	smartThermostat := &SmartThermostat{}
	smartLock := &SmartLock{}

	mediator.RegisterDevice(smartTv)
	mediator.RegisterDevice(smartSpeaker)
	mediator.RegisterDevice(smartFloorLamp)
	mediator.RegisterDevice(smartThermostat)
	mediator.RegisterDevice(smartLock)

	wolowitz := NewCharacter("Howard", mediator)
	cooper := NewCharacter("Sheldon", mediator)
	hofstadter := NewCharacter("Leonard", mediator)
	koothrapalli := NewCharacter("Rajesh", mediator)
	bloom := NewCharacter("Stuart", mediator)
	kripke := NewCharacter("Bawwy", mediator)

	wolowitz.ControlDevice("tv", "play Star Trek")
	fmt.Println("--------------------------")
	kripke.ControlDevice("assistant", "wecommend a westauwant")
	fmt.Println("--------------------------")
	cooper.ControlDevice("thermostat", "start heating")
	fmt.Println("--------------------------")
	hofstadter.ControlDevice("lock", "unlock")
	fmt.Println("--------------------------")
	koothrapalli.ControlDevice("lamp", "dim")
	fmt.Println("--------------------------")
	bloom.ControlDevice("speaker", "play Walking on Sunshine by Katrina and the Waves")
	fmt.Println("--------------------------")
}

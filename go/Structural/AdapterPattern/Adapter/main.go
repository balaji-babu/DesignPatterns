package main

import "fmt"

type transport interface {
	navigateToDestination()
}

type client struct {
}

func (t *client) startNavigation(trans transport) {
	fmt.Println("Client starting the navigation process.")
	trans.navigateToDestination()
}

type boat struct {
}

func (b *boat) navigateToDestination() {
	fmt.Println("Navigating to an Island.")
}

type car struct {
}

func (c *car) driveToDestination() {
	fmt.Println("Going to the destination.")
}

// Car adapter struct to allow cat to go to an island
type carAdapter struct {
	carTransport *car
}

func (c *carAdapter) navigateToDestination() {
	fmt.Println("Car Adapter modified to allow navigation.")
	c.carTransport.driveToDestination()
}

func main() {
	client := &client{}
	boat := &boat{}

	client.startNavigation(boat)

	car := &car{}
	carAdapter := &carAdapter{
		carTransport: car,
	}
	client.startNavigation(carAdapter)
}

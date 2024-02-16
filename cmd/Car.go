package main

import "fmt"

type car map[string]string

func newcar(registration_number string,brand string,color string,places string,yearr string) (car){
	newcar := make(car)
	newcar["registration_number"] = registration_number
    newcar["brand"] = brand
    newcar["color"] = color
	newcar["places"] = places
	newcar["yearr"] = yearr

	return newcar

}

func printCar(myCar map[string]string){
	fmt.Println(myCar)

}
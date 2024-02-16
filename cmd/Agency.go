package main

import "fmt"

type agency map[string]string

func NewAgency() agency {
    // Initialize the map
    myAgency := make(agency)

    // Add key-value pairs to the map
    myAgency["name"] = "Hadia Rental Car agency"
    myAgency["Phone"] = "00213555555"
    myAgency["Address"] = "ALGERIA, Algiers number 5"
    myAgency["turnover"] = "200000000"

    return myAgency
}


func GetInfo(myAgency map[string]string) {

    fmt.Println(myAgency)
}

func (myAgency agency)UpdatePhone(newphone string)(map[string]string) {

    myAgency["Phone"]=newphone
    return myAgency
}
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}

func PrintUser(u *User) {
	fmt.Printf("Name: %s\n", u.Name)
	fmt.Printf("Type: %s\n", u.Type)
	fmt.Printf("Age: %d\n", u.Age)
	fmt.Printf("Social VK: %s abd Facebook: %s\n", u.Social.Vk, u.Social.Facebook)
}

// User Internal User representation
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Vk       string `json:"vk"`
	Facebook string `json:"facebook"`
}

func main() {
	//fmt.Println("Hello world")
	//
	//var (
	//	personName string = "Bob"
	//	personAge         = 45
	//	personUID  int
	//)
	//
	//fmt.Printf("Name %s\nAge %d\nUID %d\n", personName, personAge, personUID)
	//
	//var a, b = 30, " asd"
	//fmt.Println(a, b)
	//a = 300
	//fmt.Println(a, b)
	//
	//var count int
	//fmt.Println(count)
	//count++
	//fmt.Println(count)
	//
	//var (
	//	age  int
	//	name string
	//)
	//fmt.Scan(&age)
	//fmt.Scan(&name)
	//
	//fmt.Println(name, age)
	//
	//c1 := complex(5, 7)
	//c2 := 12 + 32i
	//
	//fmt.Println(c1 + c2)
	//fmt.Println(c1 * c2)

	//totalString, subString := "ASDADS", "asd"
	//
	//fmt.Println(strings.Contains(totalString, subString))

	//:Loops
	//outer:
	//	for i := 0; i <= 2; i++ {
	//		for j := 1; j <= 3; j++ {
	//			fmt.Println(i, " ", j)
	//			if i == j {
	//				break outer
	//			}
	//		}
	//	}

	//	var password string
	//outer2:
	//	for {
	//		fmt.Println("Insert password")
	//		fmt.Scan(&password)
	//		if strings.Contains(password, "1234") {
	//			fmt.Println("Weak password Try again")
	//			continue
	//		} else {
	//			fmt.Println("Password accepted")
	//			break outer2
	//		}
	//	}

	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor successfully created")

	var users Users

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(byteValue, &users)
	//fmt.Println(users)

	for _, u := range users.Users {
		fmt.Println("===========")
		PrintUser(&u)
	}
}

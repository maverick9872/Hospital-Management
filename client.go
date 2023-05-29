package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

type Patient struct {
	ID       int
	Name     string
	Phone    string
	Medicine string
}

func printMenu() {
	fmt.Println("Menu:")
	fmt.Println("1. Get patient details")
	fmt.Println("2. Remove patient")
	fmt.Println("3. Sort and display all patients")
	fmt.Println("0. Exit")
}

func readIntInput(prompt string) (int, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func readStringInput(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var choice int

	for {
		printMenu()
		choice, err = readIntInput("Enter your choice: ")
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch choice {
		case 1:
			id, err := readIntInput("Enter patient ID: ")
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			var patient Patient
			err = client.Call("HospitalServer.GetPatientDetails", id, &patient)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Patient Details:")
			fmt.Println("Name:", patient.Name)
			fmt.Println("Phone:", patient.Phone)
			fmt.Println("Medicine:", patient.Medicine)
		case 2:
			id, err := readIntInput("Enter patient ID: ")
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			var success bool
			err = client.Call("HospitalServer.RemovePatient", id, &success)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if success {
				fmt.Println("Patient removed successfully")
			} else {
				fmt.Println("Patient not found")
			}
		case 3:
			var patients []Patient
			err = client.Call("HospitalServer.GetAllPatients", struct{}{}, &patients)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("All Patients:")
			for _, patient := range patients {
				fmt.Printf("ID: %d, Name: %s, Phone: %s, Medicine: %s\n", patient.ID, patient.Name, patient.Phone, patient.Medicine)
				fmt.Printf("")
			}
		case 0:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println()
	}
}

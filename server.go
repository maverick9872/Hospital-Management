package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"sort"
)

type Patient struct {
	ID       int
	Name     string
	Phone    string
	Medicine string
}

type HospitalServer struct {
	Patients map[int]Patient
}

func (s *HospitalServer) GetPatientDetails(id int, reply *Patient) error {
	patient, exists := s.Patients[id]
	if !exists {
		return errors.New("patient not found")
	}
	*reply = patient
	return nil
}

func (s *HospitalServer) RemovePatient(id int, reply *bool) error {
	_, exists := s.Patients[id]
	if !exists {
		return errors.New("patient not found")
	}
	delete(s.Patients, id)
	*reply = true
	return nil
}

func (s *HospitalServer) GetAllPatients(_ struct{}, reply *[]Patient) error {
	patients := make([]Patient, 0, len(s.Patients))
	for _, patient := range s.Patients {
		patients = append(patients, patient)
	}
	sort.Slice(patients, func(i, j int) bool {
		return patients[i].ID < patients[j].ID
	})
	*reply = patients
	return nil
}

func main() {
	server := new(HospitalServer)
	server.Patients = make(map[int]Patient)

	server.Patients[1] = Patient{
		ID:       3,
		Name:     "Aadhithian Biju",
		Phone:    "1234567890",
		Medicine: "Paracetamol",
	}
	server.Patients[2] = Patient{
		ID:       1,
		Name:     "Parimarjan Mishra",
		Phone:    "9876543210",
		Medicine: "Crocin DS",
	}
	server.Patients[3] = Patient{
		ID:       2,
		Name:     "Mithun Menon",
		Phone:    "4562949204",
		Medicine: "Poviodine Gargle",
	}

	rpc.Register(server)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Starting server...")
	http.Serve(l, nil)
}

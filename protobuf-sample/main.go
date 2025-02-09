package main

import (
	"fmt"
	"log"
	"os"
	"protobuf-sample/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:           1,
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		Occupation:   pb.Occupation_ENGINEER,
		PhoneNumbers: []string{"1234567890", "0987654321"},
		Projects:     map[string]*pb.Company_Project{},
		Profile: &pb.Employee_Text{
			Text: "Hello, World!",
		},
		Birthday: &pb.Date{
			Year:  1990,
			Month: 1,
			Day:   1,
		},
	}

	binData, err := proto.Marshal(employee)
	if err != nil {
		log.Fatal("Failed to marshal employee: ", err)
	}

	if err := os.WriteFile("employee.bin", binData, 0666); err != nil {
		log.Fatal("Failed to write employee to file: ", err)
	}

	in, err := os.ReadFile("employee.bin")
	if err != nil {
		log.Fatal("Failed to read employee from file: ", err)
	}

	employee2 := &pb.Employee{}
	if err := proto.Unmarshal(in, employee2); err != nil {
		log.Fatal("Failed to unmarshal employee: ", err)
	}

	fmt.Println("Unmarshalled employee: ", employee2)
}

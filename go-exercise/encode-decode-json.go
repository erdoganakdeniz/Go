package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID int
	Ad string
	Soyad string
	Unvan string
}

func main() {
	emp:=Employee{ID: 100,Ad: "Erdogan",Soyad: "Akdeniz",Unvan:"Muhendis"}
	data,err:=json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	jsonStr:=string(data)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)
	b:=[]byte(`{"ID":101,"Ad":"Irene","Soyad":"Rose","Unvan":"Developer"}`)
	var emp1 Employee
	err=json.Unmarshal(b,&emp1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Employee value is: ")
	fmt.Printf("ID : %d,Ad : %s %s ,Unvan : %s ",emp1.ID,emp1.Ad,emp1.Soyad,emp1.Unvan)

}

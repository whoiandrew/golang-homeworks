// let's go school homework
// author: @whoiandrew
package main

import (
	"fmt"
	"sync"
)

var (
	mutex sync.RWMutex
)

type selfIntroducer interface {
	tellName()
	tellCompanyName()
	tellPosition()
}

//Human represents basics human's attributes
type Human struct {
	name        string
	secondname  string
	location    string
	citizenship string
	age         uint
	isMale      bool
}

//Employee represents advanced stack of employee's attributes
type Employee struct {
	nickname           string
	salary             float64
	workingHoursPerDay uint
	isUnique           bool
	isRemote           bool
	position           string
	yearsOfExperience  uint
	companyName        string
	isChief            bool
	human              *Human
}

//Driver represents driver's profession basic attributes
type Driver struct {
	DriversCardCategory string `json:"category"`
	HasOwnVehicle       bool   `json:",omitempty"`
	employee            *Employee
}

//Doctor represents doctor's profession basic attributes
type Doctor struct {
	Category string `json:"category"`
	Hospital uint   `json:"hospitalNumber,omitempty"`
	employee *Employee
}

//SoftwareDeveloper represents software dev's profession basic attributes
type SoftwareDeveloper struct {
	Rank           string   `json:",omitempty"`
	Specialization string   `json:",omitempty"`
	Stack          []string `json:"stackOfTechnologies,omitempty"`
	HasDiploma     bool     `json:"diploma ,omitempty"`
	employee       *Employee
}

//Barber represents barber's profession basic attributes
type Barber struct {
	LongHairSkillLevel   uint `json:"long,omitempty"`
	MiddleHairSkillLevel uint `json:"middle,omitempty"`
	ShortHairSkillLevel  uint `json:"short,omitempty"`
	IsShavingAvialable   bool `json:"-"`
	employee             *Employee
}

//Teacher represents teacher's profession basic attributes
type Teacher struct {
	Subject     string `json:",omitempty"`
	School      uint   `json:"numberOfSchool,omitempty"`
	HasOwnClass bool   `json:"-"`
	employee    *Employee
}

func (e Employee) tellName() {
	fmt.Printf("\nMy name is %v", e.human.name)
}

func (e Employee) tellCompanyName() {
	fmt.Printf("\nI work in %v company", e.companyName)
}

func (e Employee) tellPosition() {
	fmt.Printf("\nI work as a %v", e.position)
}

func introduce(s selfIntroducer, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("\n")
	s.tellName()
	s.tellCompanyName()
	s.tellPosition()
	mutex.Unlock()
	wg.Done()
}

func fillCache(arr []Employee) map[string]Employee {
	var m = make(map[string]Employee)
	for _, value := range arr {
		m[value.nickname] = value
	}
	return m
}

func getTypes(m map[string]Employee) map[string]string {
	var types = make(map[string]string)
	for _, value := range m {
		types[value.nickname] = fmt.Sprintf("%T", value)
	}
	return types
}

func employeeToHuman(e Employee) (h Human) {
	return *e.human
}

func chiefsCounter(earr []Employee) int {
	var counter int
	for _, v := range earr {
		if v.isChief {
			counter++
		}
	}
	return counter
}

func main() {
	richard := Driver{
		"D",
		true,
		&Employee{
			nickname:           "rich328",
			salary:             10.5,
			workingHoursPerDay: 12,
			yearsOfExperience:  2,
			companyName:        "DHL",
			isChief:            true,
			position:           "deliveryman",
			human: &Human{
				"Richard",
				"Jackson",
				"Nashville, USA",
				"USA",
				25,
				true,
			},
		},
	}

	mark := SoftwareDeveloper{
		Stack:          []string{"mongo", "express", "vue", "node"},
		Specialization: "FullStack",
		Rank:           "Middle",
		employee: &Employee{
			position:           "Team lead",
			nickname:           "marker007",
			salary:             300000,
			workingHoursPerDay: 8,
			yearsOfExperience:  12,
			companyName:        "Twitter",
			isChief:            true,
			human: &Human{
				"Mark",
				"Meyer",
				"Munich, Germany",
				"Germany",
				30,
				true,
			},
		},
	}

	anastasia := Doctor{
		Category: "dantist",
		Hospital: 453,
		employee: &Employee{
			nickname:           "an8320",
			position:           "main doctor",
			salary:             20.6,
			workingHoursPerDay: 6,
			yearsOfExperience:  10,
			isChief:            false,
			companyName:        "Boris",
			human: &Human{
				"Anastasia",
				"Petrenko",
				"Kyiv, Ukraine",
				"Ukraine",
				43,
				false,
			},
		},
	}

	mike := Barber{
		ShortHairSkillLevel:  5,
		MiddleHairSkillLevel: 4,
		LongHairSkillLevel:   8,
		employee: &Employee{
			position: "junior barber",
			nickname: "",
			salary:   13.,
			isChief:  false,
			human: &Human{
				name:        "Mike",
				citizenship: "Argentina",
				location:    "Lissabon, Portugal",
				age:         29,
			},
		},
	}
	deborah := Teacher{
		Subject: "italian language",
		employee: &Employee{
			position: "ordinary teacher",
			nickname: "deb987",
			isChief:  false,
			human: &Human{
				"Deborah",
				"DeLuca",
				"Roma, Italy",
				"Spain",
				34,
				false,
			},
		},
	}

	employeesArr := []Employee{*richard.employee, *mark.employee, *anastasia.employee, *mike.employee, *deborah.employee}
	employeesCache := fillCache(employeesArr)
	//employeesTypes := getTypes(employeesCache)

	fmt.Printf("%+v\n\n%+v\n\n%+v\n\n%+v\n\n%+v\n", deborah, mike, anastasia, mark, richard)

	var wgChiefs sync.WaitGroup
	var wgNonChiefs sync.WaitGroup

	wgChiefs.Add(chiefsCounter(employeesArr))
	wgNonChiefs.Add(len(employeesArr) - chiefsCounter(employeesArr))

	for _, v := range employeesCache {
		if v.isChief {
			go introduce(v, &wgChiefs)
		}
	}
	wgChiefs.Wait()

	for _, v := range employeesCache {
		if !v.isChief {
			go introduce(v, &wgNonChiefs)
		}
	}
	wgNonChiefs.Wait()

	//deborahHuman := employeeToHuman(*deborah.employee)

	//fmt.Printf("\nTypes: %+v", employeesTypes)

}

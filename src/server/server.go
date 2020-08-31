// let's go school homework
// author: @whoiandrew
package main

import (
	//"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var (
	mutex sync.RWMutex
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

type selfIntroducer interface {
	tellName()
	tellCompanyName()
	tellPosition()
}

//Human represents basics human's attributes
type Human struct {
	Name        string
	Secondname  string
	Location    string
	Citizenship string
	Age         uint
	IsMale      bool
}

//Employee represents advanced stack of employee's attributes
type Employee struct {
	Nickname           string
	Salary             float64
	WorkingHoursPerDay uint64
	IsUnique           bool
	IsRemote           bool
	Position           string
	YearsOfExperience  uint64
	CompanyName        string
	IsChief            bool
	Human              *Human
}

//Driver represents driver's profession basic attributes
type Driver struct {
	DriversCardCategory string `json:"category"`
	HasOwnVehicle       bool   `json:",omitempty"`
	Employee            *Employee
}

//Doctor represents doctor's profession basic attributes
type Doctor struct {
	Category string `json:"category"`
	Hospital uint   `json:"hospitalNumber,omitempty"`
	Employee *Employee
}

//SoftwareDeveloper represents software dev's profession basic attributes
type SoftwareDeveloper struct {
	Rank           string   `json:",omitempty"`
	Specialization string   `json:",omitempty"`
	Stack          []string `json:"stackOfTechnologies,omitempty"`
	HasDiploma     bool     `json:"diploma ,omitempty"`
	Employee       *Employee
}

//Barber represents barber's profession basic attributes
type Barber struct {
	LongHairSkillLevel   uint `json:"long,omitempty"`
	MiddleHairSkillLevel uint `json:"middle,omitempty"`
	ShortHairSkillLevel  uint `json:"short,omitempty"`
	IsShavingAvialable   bool `json:"-"`
	Employee             *Employee
}

//Teacher represents teacher's profession basic attributes
type Teacher struct {
	Subject     string `json:",omitempty"`
	School      uint   `json:"numberOfSchool,omitempty"`
	HasOwnClass bool   `json:"-"`
	Employee    *Employee
}

func (e Employee) tellName() {
	fmt.Printf("\nMy name is %v", e.Human.Name)
}

func (e Employee) tellCompanyName() {
	fmt.Printf("\nI work in %v company", e.CompanyName)
}

func (e Employee) tellPosition() {
	fmt.Printf("\nI work as a %v", e.Position)
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
		m[value.Nickname] = value
	}
	return m
}

func getTypes(m map[string]Employee) map[string]string {
	var types = make(map[string]string)
	for _, value := range m {
		types[value.Nickname] = fmt.Sprintf("%T", value)
	}
	return types
}

func employeeToHuman(e Employee) (h Human) {
	return *e.Human
}

func chiefsCounter(earr []Employee) int {
	var counter int
	for _, v := range earr {
		if v.IsChief {
			counter++
		}
	}
	return counter
}

func ToHash(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func reversedMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func main() {
	richard := Driver{
		"D",
		true,
		&Employee{
			Nickname:           "rich328",
			Salary:             10.5,
			WorkingHoursPerDay: 12,
			YearsOfExperience:  2,
			CompanyName:        "DHL",
			IsChief:            true,
			Position:           "deliveryman",
			Human: &Human{
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
		Employee: &Employee{
			Position:           "Team lead",
			Nickname:           "marker007",
			Salary:             300000,
			WorkingHoursPerDay: 8,
			YearsOfExperience:  12,
			CompanyName:        "Twitter",
			IsChief:            true,
			Human: &Human{
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
		Employee: &Employee{
			Nickname:           "an8320",
			Position:           "main doctor",
			Salary:             20.6,
			WorkingHoursPerDay: 6,
			YearsOfExperience:  10,
			IsChief:            false,
			CompanyName:        "Boris",
			Human: &Human{
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
		Employee: &Employee{
			Position: "junior barber",
			Nickname: "mikee1",
			Salary:   13.,
			IsChief:  false,
			Human: &Human{
				Name:        "Mike",
				Citizenship: "Argentina",
				Location:    "Lissabon, Portugal",
				Age:         29,
			},
		},
	}
	deborah := Teacher{
		Subject: "italian language",
		Employee: &Employee{
			Position: "ordinary teacher",
			Nickname: "deb987",
			IsChief:  false,
			Human: &Human{
				"Deborah",
				"DeLuca",
				"Roma, Italy",
				"Spain",
				34,
				false,
			},
		},
	}

	employeesArr := []Employee{*richard.Employee, *mark.Employee, *anastasia.Employee, *mike.Employee, *deborah.Employee}
	employeesCache := fillCache(employeesArr)
	pwdCache := make(map[string]string)

	var tokensHashes = struct {
		sync.RWMutex
		m map[string]string
	}{m: make(map[string]string)}

	for _, e := range employeesArr {
		pwdCache[e.Nickname] = ToHash(fmt.Sprintf("%v0000", e.Nickname))
	}

	var wgChiefs sync.WaitGroup
	var wgNonChiefs sync.WaitGroup

	wgChiefs.Add(chiefsCounter(employeesArr))
	wgNonChiefs.Add(len(employeesArr) - chiefsCounter(employeesArr))

	for _, v := range employeesCache {
		if v.IsChief {
			go introduce(v, &wgChiefs)
		}
	}
	wgChiefs.Wait()

	for _, v := range employeesCache {
		if !v.IsChief {
			go introduce(v, &wgNonChiefs)
		}
	}
	wgNonChiefs.Wait()

	http.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == MethodPost {
			err := req.ParseForm()
			if err != nil {
				panic(err)
			}

			nickname := req.PostFormValue("nickname")
			password := ToHash(req.PostFormValue("pwd1"))

			if _, ok := pwdCache[nickname]; ok {
				fmt.Fprintf(w, "User %v already exists, \n Want to login?", nickname)
			} else {
				pwdCache[nickname] = password
				fmt.Fprintf(w, "%v, your account has created succesfully, welcome! ", nickname)
			}

		}

	})

	http.HandleFunc("/login", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == MethodPost {
			err := req.ParseForm()
			if err != nil {
				panic(err)
			}

			nickname := req.PostFormValue("nickname")
			password := req.PostFormValue("pwd")

			if pwd, ok := pwdCache[nickname]; ok && CheckPasswordHash(password, pwd) {
				fmt.Fprintf(w, "Welcome to the club, %v", nickname)
			} else {
				fmt.Fprint(w, "Wrong login or password")
			}

		}

	})

	http.HandleFunc("/loginFromServer", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == MethodPost {
			err := req.ParseForm()
			if err != nil {
				fmt.Println(err)
				return
			}

			nickname := req.PostFormValue("nickname")
			password := req.PostFormValue("pwd")

			if pwd, ok := pwdCache[nickname]; ok && CheckPasswordHash(password, pwd) {
				fmt.Fprintf(w, "Welcome to the club, %v\n", nickname)
				tokensHashes.RLock()
				_, loginExists := tokensHashes.m[nickname]
				tokensHashes.RUnlock()
				if loginExists {
					tokensHashes.RLock()
					w.Write([]byte(fmt.Sprintf(`{"token": %v}`, tokensHashes.m[nickname])))
					tokensHashes.RUnlock()

				} else {
					tokensHashes.Lock()
					tokensHashes.m[nickname] = ToHash(nickname)
					fmt.Fprintf(w, "token - %v", tokensHashes.m[nickname])
					tokensHashes.Unlock()
				}
			} else {
				fmt.Fprint(w, "Wrong login or password")
			}
		}
	})

	http.HandleFunc("/loginWithTokenFromServer", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == MethodPost {
			err := req.ParseForm()
			if err != nil {
				fmt.Println(err)
				return
			}

			token := req.PostFormValue("token")

			tokensHashes.Lock()
			nick, ok := reversedMap(tokensHashes.m)[token]
			tokensHashes.Unlock()

			if ok {
				fmt.Fprintf(w, "Logged in by token, %v", nick)
			} else {
				fmt.Fprintln(w, "Wrong token")
			}

		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method == MethodGet {
			queries := req.URL.Query()
			fmt.Fprintf(w, req.Method)
			for k, v := range queries {
				if k == "nick" {
					elem, ok := employeesCache[v[0]]
					if !ok {
						fmt.Fprintf(w, "\nEmployee %v does not exist", v)
					} else {
						fmt.Fprintf(w, "\n%v", elem)
					}
				} else {
					fmt.Fprintf(w, "\nPls, input nick=<nickname>")
				}
			}
		} else if req.Method == MethodPost {
			fmt.Fprintf(w, req.Method)
			err := req.ParseForm()
			if err != nil {
				panic(err)
			}

			nickname := req.PostFormValue("nickname")
			salary, _ := strconv.ParseFloat(req.PostFormValue("salary"), 64)
			workingHoursPerDay, _ := strconv.ParseUint(req.PostFormValue("workingHoursPerDay"), 10, 64)

			if _, ok := employeesCache[nickname]; ok {
				fmt.Fprintf(w, "\nUser %v has already created", nickname)
			} else {
				employeesCache[nickname] = Employee{
					Nickname:           nickname,
					Salary:             salary,
					WorkingHoursPerDay: workingHoursPerDay,
				}
			}

			fmt.Fprintf(w, "\n%+v", employeesCache[nickname])

		}

	})

	port := ":8082"

	fmt.Printf("\n\nListening on port %v\n", port)
	http.ListenAndServe(port, nil)

}

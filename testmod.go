package testmod
import (
"errors"
"fmt"
"reflect"
"strings"
"sync"
)

func WorkersToPerson(w IWorker) Person{
	var person Person
	switch value := w.(type) {
	case Teacher:
		person = *value.Person
	case Developer:
		person = *value.Person
	case Accountant:
		person = *value.Person
	case Manager:
		person = *value.Person
	case Doctor:
		person = *value.Person
	default:
		fmt.Printf("Don't know about %T  %v\n",value,value)
		person = Person{}
	}
	return person
}
//struct Cash
type Cache struct{
	sync.RWMutex
	Items  map[string]IWorker
}

//New Cash
func NewCash() *Cache{
	items := make(map[string]IWorker)
	return &Cache{
		Items: items,
	}
}

//Delete Item
func (c *Cache) Delete(key string) error {
	c.RLock()
	defer c.RUnlock()
	if _, found := c.Items[key]; !found {
		return errors.New("Key not found")
	}
	delete(c.Items, key)
	return nil
}
//Delete all
func (c *Cache) DeleteAll() error {
	if c == nil {
		return errors.New("Cache is 'nil'")
	}
	c.Lock()
	defer c.Unlock()
	for k ,_ := range c.Items {
		delete(c.Items, k)
	}
	return nil
}

//All Workers Talk
func  AllWorkersTalk(c *Cache,quit chan int){
	for {
		select {
		case <-quit:
			c.RLock()
			defer c.RUnlock()
			for _, worker := range c.Items {
				fmt.Println(worker.Talk())
			}
			break
		}
	}
}

func AllChiefsTalk(c *Cache,quit chan int){
	c.RLock()
	defer c.RUnlock()
	for _, chief := range c.Items {
		fmt.Println(chief.Talk())
	}
	quit <- 0
}

//Add Worker to cash map
func (c *Cache) Add(key string, item IWorker){
	c.Lock()
	defer c.Unlock()
	c.Items[key] = item
}

// interface for workers
type IWorker interface {
	Talk()  string
}

type Person struct {
	Name        string
	Surname     string
	Patronymic  string
	PhoneNumber uint64
	Email       string
	HomeAddress string
	Age         int
	Education   string
}

func (p Person) String() string{
	return fmt.Sprintf("\nName: %v\nSurname: %v\nAge: %v\nEducation: %v\nEmail: %v\nAddress: %v\nPatronimic: %v\nPhoneNumber: +%v",p.Name,p.Surname,p.Age,p.Education,p.Email,p.Email,p.Patronymic,p.PhoneNumber)
}

func (p *Person) HappyBirthday(){
	fmt.Println("CONGRADULATIONS TO ",p.Name," ",p.Surname," Happy Birthday")
}

func (p *Person) ChangeAge() {
	p.Age = p.Age + 1
}

type Developer struct {
	Position             string	  `json:"position,omitempty"`
	ProgrammingLanguages []string `json:"programmingLanguages,omitempty",xml:",omitempty"`
	SkillLevel           string   `json:"skillLevel,omitempty",xml:"SkillLev"`
	ExperienceInYears    int      `json:"experience,omitempty",xml:",omitempty"`
	Person               *Person
}
// method talk for Developer to implement IWorker
func (w Developer) Talk() string{
	s := fmt.Sprint(reflect.TypeOf(w))
	str := strings.Split(s,".")
	return fmt.Sprint("Hello I am ",w.Person.Name,"! My profession is ",str[1]," and I am working as ",w.Position,".")
}

type Teacher struct {
	Position string				  `json:"position,omitempty"`
	Subject                string `json:"subject,omitempty",xml:"Subj,omitempty"`
	LessonsInWeek          int    `json:"lessonsInWeek,omitempty",xml:"LessInWeek,omitempty"`
	NumberOfTeachingGroups int    `json:"numberOfGroups,omitempty",xml:"GroupsNum,omitempty"`
	Person                 *Person
}
// method talk for Teacher to implement IWorker
func (w Teacher) Talk() string{
	s := fmt.Sprint(reflect.TypeOf(w))
	str := strings.Split(s,".")
	return fmt.Sprint("Hello I am ",w.Person.Name,"! My profession is ",str[1]," and I am working as ",w.Position,".")
}


type Accountant struct {
	Position string				      `json:"position,omitempty"`
	ProgramToWorkWith         string  `json:"programToWorkWith,omitempty",xml:"Program,omitempty"`
	MaxWorkingHoursInDay      int     `json:"maxWorkingHoursInDay,omitempty",xml:"MaxHoursInDay,omitempty"`
	NeededExperienceInYears   int     `json:"neededExperienceInYears,omitempty",xml:"ExperienseYeas,omitempty"`
	TypeOfDocumentsToWorkWith string  `json:"typeOfDocumentsToWorkWith,omitempty",xml:"DocumentsType,omitempty"`
	Person                    *Person
}
// method talk for Accountant to implement IWorker
func (w Accountant) Talk() string{
	s := fmt.Sprint(reflect.TypeOf(w))
	str := strings.Split(s,".")
	return fmt.Sprint("Hello I am ",w.Person.Name,"! My profession is ",str[1]," and I am working as ",w.Position,".")
}

type Doctor struct {
	Position string				      `json:"position,omitempty"`
	Specialization          string    `json:"specialization,omitempty",xml:"Secesilization,omitempty"`
	DoctorGraduation        string    `json:"doctorGraduation,omitempty",xml:"Gruduation,omitempty"`
	Certificates            []string  `json:"certificates,omitempty",xml:"-"`
	PreviousWorkingPosition string    `json:"previousWorkingPosition,omitempty",xml:"-"`
	Person                  *Person
}
// method talk for Doctor to implement IWorker
func (w Doctor) Talk() string{
	s := fmt.Sprint(reflect.TypeOf(w))
	str := strings.Split(s,".")
	return fmt.Sprint("Hello I am ",w.Person.Name,"! My profession is ",str[1]," and I am working as ",w.Position,".")
}

type Manager struct {
	Position string				     `json:"position,omitempty"`
	ManagingSphere       string      `json:"managingSphere,omitempty",xml:"-"`
	NeededSkills         []string    `json:"neededSkills,omitempty"xml:"Skills,omitempty"`
	MinWorkingHoursInDay int         `json:"-",xml:"-"`
	WorkingDaysInWeek    int         `json:"workingDaysInWeek,omitempty",xml:"-"`
	Person               *Person
}
// method talk for Manager to implement IWorker
func (w Manager) Talk() string{
	s := fmt.Sprint(reflect.TypeOf(w))
	str := strings.Split(s,".")
	return fmt.Sprint("Hello I am ",w.Person.Name,"! My profession is ",str[1]," and I am working as ",w.Position,".")
}

func println(){
	fmt.Println("----------------------------------------------------------------------------------")
}

func StandartWorkersMap() *Cache{
	developer := &Developer{
		Position: "Junior Developer",
		ProgrammingLanguages: []string{"C#", "Go"},
		SkillLevel:           "Junior",
		ExperienceInYears:    1,
		Person: &Person{
			Name:        "Dmytro",
			Surname:     "Ptushkovich",
			Patronymic:  "Dmytrovich",
			PhoneNumber: 380678848393,
			Email:       "dmitrorezn@gmail.com",
			HomeAddress: "st. Myloslavska",
			Age:         20,
			Education:   "student",
		},
	}
	teacher := &Teacher{
		Position: "High school teacher",
		Subject:                "Math",
		LessonsInWeek:          12,
		NumberOfTeachingGroups: 4,
		Person: &Person{
			Name:        "Dmytro",
			Surname:     "Ivanenko",
			Patronymic:  "Volodymirovich",
			PhoneNumber: 380672444466,
			Email:       "dmitrorezn@gmail.com",
			HomeAddress: "st. Vatutina",
			Age:         19,
			Education:   "student",
		},
	}
	accountant := &Accountant{
		Position: "Junior Accountant",
		ProgramToWorkWith:         "C1",
		MaxWorkingHoursInDay:      10,
		NeededExperienceInYears:   1,
		TypeOfDocumentsToWorkWith: "finances",
		Person: &Person{
			Name:        "Ivan",
			Surname:     "Sokolenko",
			Patronymic:  "Igorovich",
			PhoneNumber: 380672350466,
			Email:       "dmitrorezn@gmail.com",
			HomeAddress: "st. Vatutina",
			Age:         20,
			Education:   "student",
		},
	}
	doctor := &Doctor{
		Position: "Nurse",
		Specialization:          "Neiroherurg",
		DoctorGraduation:        "Doctor",
		Certificates:            []string{"Practical Courses of Bogomoltsa University 2019"},
		PreviousWorkingPosition: "Doctor in city hospital",
		Person: &Person{
			Name:        "Igor",
			Surname:     "Mosenko",
			Patronymic:  "Igorovich",
			PhoneNumber: 380672350466,
			Email:       "dmitrorezn@gmail.com",
			HomeAddress: "st. Vatutina",
			Age:         18,
			Education:   "student",
		},
	}
	manager := &Manager{
		Position: "Product manager",
		ManagingSphere:       "IT",
		NeededSkills:         nil,
		MinWorkingHoursInDay: 6,
		WorkingDaysInWeek:    4,
		Person: &Person{
			Name:        "Dmytro",
			Surname:     "Adamenko",
			Patronymic:  "Victorovich",
			PhoneNumber: 380672350466,
			Email:       "dmitrorezn@gmail.com",
			HomeAddress: "st. Vatutina",
			Age:         21,
			Education:   "student",
		},
	}
	
	//Cash testing
	cache1 := NewCash()
	cache1.Add("developer",developer)
	cache1.Add("doctor",doctor)
	cache1.Add("manager",manager)
	cache1.Add("accountant",accountant)
	cache1.Add("teacher",teacher)
	
	developerC := *developer
	developerC.Position = "Chief"
	doctorC := *doctor
	doctorC.Position = "Chief"
	managerC := *manager
	managerC.Position = "Chief"
	accountantC := *accountant
	accountantC.Position = "Chief"
	teacherC := *teacher
	teacherC.Position = "Chief"

	cache1.Add("developerCh",&developerC)
	cache1.Add("doctorCh",&doctorC)
	cache1.Add("managerCh",&managerC)
	cache1.Add("accountantCh",&accountantC)
	cache1.Add("teacherCh",&teacherC)
    
	return cache1
}

func WorkersTalk(cache *Cache,waitgrWo,waitgrCh *sync.WaitGroup){
	waitgrCh.Wait()
	cache.RLock()
	defer cache.RUnlock()
	for _,worker := range cache.Items{
		fmt.Println(worker.Talk())
	}
	waitgrWo.Done()
}
func ChiefsTalk(cache *Cache,waitgrCh *sync.WaitGroup){
	cache.RLock()
	defer cache.RUnlock()
	for _,chief := range cache.Items{
		fmt.Println(chief.Talk())
	}
	waitgrCh.Done()
}







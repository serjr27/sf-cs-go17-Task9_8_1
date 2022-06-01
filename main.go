package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := map[string]Man{
		"Иван":     {Name: "Иван", LastName: "Иванов", Age: 25, Gender: "Мужской", Crimes: 2},
		"Петр":     {Name: "Петр", LastName: "Петров", Age: 32, Gender: "Мужской", Crimes: 3},
		"Федор":    {Name: "Федор", LastName: "Федоров", Age: 60, Gender: "Мужской", Crimes: 5},
		"Василий":  {Name: "Василий", LastName: "Васильев", Age: 45, Gender: "Мужской", Crimes: 4},
		"Антонина": {Name: "Антонина", LastName: "Антонова", Age: 28, Gender: "Женский", Crimes: 0},
	}

	suspects := []string{"Иван", "Петр", "Федор", "Василий", "Антонина"}

	crimes := 0
	name := ""

	for _, v := range suspects {
		if c := people[v].Crimes; c > crimes {
			crimes = c
			name = v
		}
	}

	if crimes == 0 {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым.")
		return
	}

	fmt.Println("В базе данных найден подозреваемый с наибольшим числом преступлений:")
	fmt.Println("Имя:\t\t", people[name].Name)
	fmt.Println("Фамилия:\t", people[name].LastName)
	fmt.Println("Возраст:\t", people[name].Age)
	fmt.Println("Пол:\t\t", people[name].Gender)
	fmt.Println("Преступлений:\t", people[name].Crimes)
}

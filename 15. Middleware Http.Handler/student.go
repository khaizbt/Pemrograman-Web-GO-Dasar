package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() { //Fungsi Init akan otomatis terpanggil
	students = append(students, &Student{Id: "s001", Name: "tamam", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "khaiz", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "khafif", Grade: 2})
}

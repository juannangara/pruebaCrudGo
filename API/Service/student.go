package service

import (
	models "purebaCrud/API/Models"
	repository "purebaCrud/API/Repository"
)

type StudentService interface {
	GetStudents() ([]models.Student, error)
	CreateStudent(student models.NewStudent) error
	GetById(id int) (models.Student, error)
	UpdateStudent(student models.Student) error
	DeleteStudent(student models.Student) error
}

type studentService struct {
	studentRepo repository.StudentRepo
}

func CreateStudentService(studentRepo repository.StudentRepo) StudentService {
	return &studentService{
		studentRepo: studentRepo,
	}
}

func (studentService studentService) CreateStudent(student models.NewStudent) error {

	return studentService.studentRepo.CreateStudent(student)
}

func (studentService studentService) GetById(id int) (models.Student, error) {
	return studentService.studentRepo.GetById(id)
}

func (studentService studentService) GetStudents() ([]models.Student, error) {
	return studentService.studentRepo.GetStudents()
}

func (studentService studentService) UpdateStudent(student models.Student) error {
	return studentService.studentRepo.UpdateStudent(student)
}

func (studentService studentService) DeleteStudent(student models.Student) error {
	return studentService.studentRepo.DeleteStudent(student)
}

package repository

import (
	"database/sql"
	models "purebaCrud/API/Models"
	"time"
)

type StudentRepo interface {
	GetStudents() ([]models.Student, error)
	CreateStudent(u models.NewStudent) error
	GetById(id int) (models.Student, error)
	UpdateStudent(u models.Student) error
	DeleteStudent(u models.Student) error
}

type studentRepo struct {
	db *sql.DB
}

func CreateSqlGestor(db *sql.DB) StudentRepo {
	return &studentRepo{
		db: db,
	}
}

func (sql studentRepo) CreateStudent(u models.NewStudent) error {
	_, err := sql.db.Exec("INSERT INTO student(name, crationDate) VALUES (?, ?)", u.Name, time.Now())
	return err
}

func (sql studentRepo) GetById(id int) (models.Student, error) {
	var student models.Student
	err := sql.db.QueryRow("SELECT id, name, nickname FROM usuarios WHERE studentID = ?", id).Scan(&student.Student_id, &student.Name, &student.Creation_date)
	return student, err
}

func (sql studentRepo) GetStudents() ([]models.Student, error) {

	rows, err := sql.db.Query("SELECT * FROM student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var s models.Student
		err := rows.Scan(&s.Student_id, &s.Name, &s.Creation_date)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (sql studentRepo) UpdateStudent(s models.Student) error {
	_, err := sql.db.Exec("UPDATE usuarios SET name = ? WHERE nickname = ?", s.Name)
	return err
}

func (sql studentRepo) DeleteStudent(u models.Student) error {
	_, err := sql.db.Exec("DELETE FROM student WHERE studentID  = ? ", u.Student_id)
	return err
}

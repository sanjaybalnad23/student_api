package database

import (
	"database/sql"

	"github.com/Drstrange373/student-api/internal/config"
	"github.com/Drstrange373/student-api/internal/types"
	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Database, error) {
	db, err := sql.Open("postgres", cfg.PostgresConnectionString)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		name TEXT,
		age INTEGER,
		email TEXT UNIQUE
	)`)
	if err != nil {
		return nil, err
	}
	return &Database{Db: db}, nil
}

func (d Database) CreateStudent(name string, email string, age int) (id int64, err error) {
	// TODO : Implement this func
	err = d.Db.QueryRow("INSERT INTO students(name, email, age) VALUES($1, $2, $3) RETURNING id", name, email, age).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// returns nill if  student found
func (d Database) GetStudentById(id int64) (*types.Student, error) {
	result, err := d.Db.Query("SELECT id, name, age, email FROM students WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var student types.Student

	if result.Next() {
		result.Scan(&student.Id, &student.Name, &student.Age, &student.Email)
		return &student, nil
	}

	return nil, nil
}

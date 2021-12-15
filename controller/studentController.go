package controller

import (
	"example/mvc/model"
	"fmt"
)

func GetStudent() ([]model.Student, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`SELECT * FROM students`)

	if stmtErr == nil {
		data, dataErr := stmt.Query()

		if dataErr == nil {
			var result []model.Student

			for data.Next() {
				var student model.Student

				if scanErr := data.Scan(
					&student.Id,
					&student.Name,
					&student.Email,
					&student.Phone,
					&student.Birthday,
					&student.Password,
				); scanErr == nil {
					result = append(result, student)
				} else {
					return []model.Student{}, scanErr
				}
			}

			return result, nil
		}

		return []model.Student{}, dataErr
	}

	return []model.Student{}, stmtErr
}

func GetStudentById(id int) ([]model.Student, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`SELECT * FROM students WHERE id=?`)

	if stmtErr == nil {
		data, dataErr := stmt.Query(id)

		if dataErr == nil {
			var result []model.Student

			for data.Next() {
				var student model.Student

				if scanErr := data.Scan(
					&student.Id,
					&student.Name,
					&student.Email,
					&student.Phone,
					&student.Birthday,
					&student.Password,
				); scanErr == nil {
					result = append(result, student)
				} else {
					return []model.Student{}, scanErr
				}
			}

			return result, nil
		}

		return []model.Student{}, dataErr
	}

	return []model.Student{}, nil
}

func GetStudentByName(name string) ([]model.Student, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`SELECT * FROM students WHERE name LIKE ?`)

	if stmtErr == nil {
		data, dataErr := stmt.Query(fmt.Sprintf("%%%v%%", name))

		if dataErr == nil {
			var result []model.Student

			for data.Next() {
				var student model.Student

				if scanErr := data.Scan(
					&student.Id,
					&student.Name,
					&student.Email,
					&student.Phone,
					&student.Birthday,
					&student.Password,
				); scanErr == nil {
					result = append(result, student)
				} else {
					return []model.Student{}, scanErr
				}
			}

			return result, nil
		}

		return []model.Student{}, dataErr
	}

	return []model.Student{}, stmtErr
}

func AddStudent(student *model.Student) (*model.Student, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`INSERT INTO students (name,email,phone,birthday,password) VALUES (?,?,?,?,?)`)

	if stmtErr == nil {
		exe, exeErr := stmt.Exec(student.Name, student.Email, student.Phone, student.Birthday, student.Password)

		if exeErr == nil {
			id, idErr := exe.LastInsertId()

			if idErr == nil {
				student.Id = id
				return student, nil
			}

			return student, idErr
		}

		return student, exeErr
	}

	return student, stmtErr
}

func EditStudent(student *model.Student) (*model.Student, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`UPDATE students SET name=?,email=?,phone=?,birthday=?,password=? WHERE id=?`)

	if stmtErr == nil {
		exe, exeErr := stmt.Exec(student.Name, student.Email, student.Phone, student.Birthday, student.Password, student.Id)

		if exeErr == nil {
			aff, affErr := exe.RowsAffected()

			if aff > 0 {
				return student, nil
			}

			return nil, affErr
		}

		return nil, exeErr
	}

	return nil, stmtErr
}

func DeleteStudentById(id int64) error {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`DELETE FROM students WHERE id=?`)

	if stmtErr == nil {
		exe, exeErr := stmt.Exec(id)

		if exeErr == nil {
			aff, affErr := exe.RowsAffected()

			if aff > 0 {
				return nil
			}

			return affErr
		}

		return exeErr
	}

	return stmtErr
}

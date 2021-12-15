package controller

import (
	"example/mvc/model"
	"fmt"
)

func GetSubject() ([]model.Subject, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`SELECT * FROM subjects`)

	if stmtErr == nil {
		data, dataErr := stmt.Query()

		if dataErr == nil {
			var result []model.Subject

			for data.Next() {
				var subject model.Subject

				if scanErr := data.Scan(
					&subject.Id,
					&subject.Name,
					&subject.Sec,
					&subject.Unit,
				); scanErr == nil {
					result = append(result, subject)
				} else {
					return []model.Subject{}, scanErr
				}
			}

			return result, nil
		}

		return []model.Subject{}, dataErr
	}

	return []model.Subject{}, stmtErr
}

func GetSubjectById(id int64) ([]model.Subject, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`SELECT * FROM subjects WHERE id=?`)

	if stmtErr == nil {
		data, dataErr := stmt.Query(id)

		if dataErr == nil {
			var result []model.Subject

			for data.Next() {
				var subject model.Subject

				if scanErr := data.Scan(
					&subject.Id,
					&subject.Name,
					&subject.Sec,
					&subject.Unit,
				); scanErr == nil {
					result = append(result, subject)
				} else {
					return []model.Subject{}, scanErr
				}
			}

			return result, nil
		}

		return []model.Subject{}, dataErr
	}

	return []model.Subject{}, stmtErr
}

func GetSubjectByName(name string) ([]model.Subject, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`SELECT * FROM subjects WHERE name LIKE ?`)

	if stmtErr == nil {
		data, dataErr := stmt.Query(fmt.Sprintf(`%%%v%%`, name))

		if dataErr == nil {
			var result []model.Subject

			for data.Next() {
				var subject model.Subject

				if scanErr := data.Scan(
					&subject.Id,
					&subject.Name,
					&subject.Sec,
					&subject.Unit,
				); scanErr == nil {
					result = append(result, subject)
				} else {
					return []model.Subject{}, scanErr
				}
			}

			return result, nil
		}

		return []model.Subject{}, dataErr
	}

	return []model.Subject{}, stmtErr
}

func AddSubject(subject *model.Subject) (*model.Subject, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`INSERT INTO subjects (name,sec,unit) VALUES (?,?,?)`)

	if stmtErr == nil {
		exe, exeErr := stmt.Exec(subject.Name, subject.Sec, subject.Unit)

		if exeErr == nil {
			id, idErr := exe.LastInsertId()

			if idErr == nil {
				subject.Id = id
				return subject, nil
			}

			return nil, idErr
		}

		return nil, exeErr
	}

	return nil, stmtErr
}

func EditSubject(subject *model.Subject) (*model.Subject, error) {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`UPDATE subjects SET name=?, sec=?, unit=? WHERE id=?`)

	if stmtErr == nil {
		exe, exeErr := stmt.Exec(subject.Name, subject.Sec, subject.Unit, subject.Id)

		if exeErr == nil {
			aff, affErr := exe.RowsAffected()

			if aff > 0 {
				return subject, nil
			}

			return nil, affErr
		}

		return nil, exeErr
	}

	return nil, stmtErr
}

func DeleteSubjectById(id int64) error {
	stmt, stmtErr := DB_CONFIG.Conn.Prepare(`DELETE FROM subjects WHERE id=?`)

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

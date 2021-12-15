package service

import (
	"encoding/json"
	"example/mvc/controller"
	"example/mvc/model"
	"fmt"
	"net/http"
	"strconv"
)

func GetStudentServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var data []model.Student
	var err error

	if values := r.URL.Query(); values[`id`] != nil {
		id, _ := strconv.Atoi(values[`id`][0])
		data, err = controller.GetStudentById(id)
	} else if values[`name`] != nil {
		name := values[`name`][0]
		data, err = controller.GetStudentByName(name)
	} else {
		data, err = controller.GetStudent()
	}

	if err == nil {
		status = 200
		result = ResultSet{fmt.Sprintf("Found %v rows.", len(data)), data}
	} else {
		status = 400
		result = ResultSet{err.Error(), nil}
	}

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

func AddStudentServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var input model.Student

	if bodyErr := json.NewDecoder(r.Body).Decode(&input); bodyErr == nil {
		if _, addErr := controller.AddStudent(&input); addErr == nil {
			status = 201
			result = ResultSet{fmt.Sprintf("Student %v is added.", input.Id), input}
		} else {
			status = 400
			result = ResultSet{addErr.Error(), nil}
		}
	} else {
		status = 400
		result = ResultSet{bodyErr.Error(), nil}
	}

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

func EditStudentServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var input model.Student

	if bodyErr := json.NewDecoder(r.Body).Decode(&input); bodyErr == nil {
		if _, editErr := controller.EditStudent(&input); editErr == nil {
			status = 200
			result = ResultSet{fmt.Sprintf("Student %v is edited.", input.Id), input}
		} else {
			status = 400
			result = ResultSet{editErr.Error(), nil}
		}
	} else {
		status = 400
		result = ResultSet{bodyErr.Error(), nil}
	}

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

func DeleteStudentServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var input model.Student

	if bodyErr := json.NewDecoder(r.Body).Decode(&input); bodyErr == nil {
		if deleteErr := controller.DeleteStudentById(input.Id); deleteErr == nil {
			status = 204
			result = ResultSet{fmt.Sprintf("Student %v is deleted.", input.Id), nil}
		} else {
			status = 400
			result = ResultSet{deleteErr.Error(), nil}
		}
	} else {
		status = 400
		result = ResultSet{bodyErr.Error(), nil}
	}

	res, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(res)
}

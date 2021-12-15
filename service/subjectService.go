package service

import (
	"encoding/json"
	"example/mvc/controller"
	"example/mvc/model"
	"fmt"
	"net/http"
	"strconv"
)

func GetSubjectServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var data []model.Subject
	var err error

	if values := r.URL.Query(); values[`id`] != nil {
		id, _ := strconv.Atoi(values[`id`][0])
		data, err = controller.GetSubjectById(int64(id))
	} else if values[`name`] != nil {
		name := values[`name`][0]
		data, err = controller.GetSubjectByName(name)
	} else {
		data, err = controller.GetSubject()
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

func AddSubjectServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var input model.Subject

	if bodyErr := json.NewDecoder(r.Body).Decode(&input); bodyErr == nil {
		if _, addErr := controller.AddSubject(&input); addErr == nil {
			status = 200
			result = ResultSet{fmt.Sprintf(`Subject %v is added.`, input.Id), input}
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

func EditSubjectServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var input model.Subject

	if bodyErr := json.NewDecoder(r.Body).Decode(&input); bodyErr == nil {
		if _, editErr := controller.EditSubject(&input); editErr == nil {
			status = 200
			result = ResultSet{fmt.Sprintf(`Subject %v is edited.`, input.Id), input}
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

func DeleteSubjectServ(w http.ResponseWriter, r *http.Request) {
	var result ResultSet
	var status int
	var input model.Subject

	if bodyErr := json.NewDecoder(r.Body).Decode(&input); bodyErr == nil {
		if deleteErr := controller.DeleteSubjectById(input.Id); deleteErr == nil {
			status = 200
			result = ResultSet{fmt.Sprintf(`Subject %v is deleted.`, input.Id), nil}
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

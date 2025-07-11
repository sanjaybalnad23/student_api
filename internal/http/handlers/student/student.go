package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Drstrange373/student-api/internal/storage"
	"github.com/Drstrange373/student-api/internal/types"
	"github.com/Drstrange373/student-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		// validation
		if err := validator.New().Struct(student); err != nil {
			validationErr, ok := err.(validator.ValidationErrors)
			if !ok {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
				return
			}
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validationErr))
			return
		}

		id, err := storage.CreateStudent(student.Name, student.Email, student.Age)

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}
		slog.Info("Student created successfully", slog.String("id", fmt.Sprint(id)))

		response.WriteJson(w, http.StatusCreated, response.GeneralResponse(map[string]int64{"id": id}))
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = r.PathValue("id")
		slog.Info("Getting a student ", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, err := storage.GetStudentById(intId)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if student == nil {
			response.WriteJson(w, http.StatusNotFound, response.GeneralError(fmt.Errorf("student not found")))
			return
		}

		response.WriteJson(w, http.StatusOK, response.GeneralResponse(student))
	}
}

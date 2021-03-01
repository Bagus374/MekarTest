package handler

import (
	"TestMekar/common"
	"TestMekar/model"
	"TestMekar/users"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService users.UserService
}

func (h UserHandler) signup(writer http.ResponseWriter, request *http.Request) {
	user := new(model.User)
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		common.HandleError(writer, http.StatusInternalServerError, "Oppss, something error")
		fmt.Printf("[UserHandler.signup] Error when decoder data from body with error : %v\n", err)
		return
	}

	response, err := h.userService.SignUp(user)
	if err != nil {
		common.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[UserHandler.InsertData] Error when request data to usecase with error: %v\n", err)
		return
	}

	common.HandleSuccess(writer, http.StatusOK, response)
}

func (h UserHandler) findAll(writter http.ResponseWriter, request *http.Request) {
	result, err := h.userService.FindAll()
	if err != nil {
		common.HandleError(writter, http.StatusNoContent, "Result Empty!")
		return
	}
	common.HandleSuccess(writter, http.StatusOK, result)
}

func (h UserHandler) GetUserById(writter http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		common.HandleError(writter, http.StatusBadRequest, "ID not valid !!")
		fmt.Printf("[UserHandler] Error when convert pathvar with error: %v\n", err)
	}

	data, err := h.userService.GetUserByID(id)
	if err != nil {
		common.HandleError(writter, http.StatusInternalServerError, "Opps, something error")
		fmt.Printf("[UserHandler.getByID]Error when request data with error : %v \n", err)
		return
	}

	common.HandleSuccess(writter, http.StatusOK, data)
}

func CreateUserHandler(r *mux.Router, userService users.UserService)  {
	userHandler := UserHandler{userService}

	r.HandleFunc("/api/v1/signup", userHandler.signup).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/findAll", userHandler.findAll).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user/{id}", userHandler.GetUserById).Methods(http.MethodGet)

}



package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/api-user-auth/app/models"
	"github.com/lailiseptiandi/api-user-auth/app/services"
	"github.com/lailiseptiandi/api-user-auth/app/utils"
)

type UserController struct {
	userSevice services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService}
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var userInput models.UserRegiserInput
	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		resp := utils.ResponseError(nil, err.Error())
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	if userInput.Password != userInput.PasswordConfirm {
		resp := utils.ResponseError(nil, "Password do not match")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	newUser, err := uc.userSevice.RegisterUser(userInput)
	if err != nil {
		resp := utils.ResponseError(nil, "Failed Register User")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	token, _ := utils.GenerateToken(newUser)

	formatter := models.FomatterUserRegister(newUser, token)
	resp := utils.ResponseSuccess(formatter, "Successfully register user")
	ctx.JSON(http.StatusCreated, resp)
	return

}

func (uc *UserController) LoginUser(ctx *gin.Context) {
	var userInput models.UserLoginInput
	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		resp := utils.ResponseError(nil, err.Error())
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	user, err := uc.userSevice.LoginUser(userInput)
	if err != nil {
		resp := utils.ResponseError(nil, err.Error())
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	token, _ := utils.GenerateToken(user)
	formatter := models.FomatterUserRegister(user, token)
	resp := utils.ResponseSuccess(formatter, "Successfully login user")
	ctx.JSON(http.StatusCreated, resp)
	return

}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var userInput models.UserUpdateInput
	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}
		resp := utils.ResponseError(errorMessage, "error")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	idUser := ctx.Param("id")

	userID, _ := strconv.Atoi(idUser)

	err = uc.userSevice.UpdateUser(uint(userID), userInput)
	if err != nil {
		resp := utils.ResponseError(nil, "Failed Updated User")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	resp := utils.ResponseSuccess(nil, "Successfully Updated User")
	ctx.JSON(http.StatusOK, resp)
	return
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	idUser := ctx.Param("id")

	userID, _ := strconv.Atoi(idUser)

	err := uc.userSevice.DeleteUser(uint(userID))
	if err != nil {
		resp := utils.ResponseError(nil, "Failed Deleted User")
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}

	resp := utils.ResponseSuccess(nil, "Successfully Deleted User")
	ctx.JSON(http.StatusOK, resp)
	return
}

func (uc *UserController) GetUser(ctx *gin.Context) {

	user, err := uc.userSevice.GetUser()
	if err != nil {
		resp := utils.ResponseError(nil, err.Error())
		ctx.JSON(http.StatusOK, resp)
		return
	}

	formatter := models.FormatterGetUser(user)
	resp := utils.ResponseSuccess(formatter, "Success get data user")
	ctx.JSON(http.StatusOK, resp)
	return
}

func (uc *UserController) DetailUser(ctx *gin.Context) {
	idUser := ctx.Param("id")
	userID, _ := strconv.Atoi(idUser)
	user, err := uc.userSevice.FindUserById(uint(userID))
	if err != nil {
		resp := utils.ResponseError(nil, err.Error())
		ctx.JSON(http.StatusOK, resp)
		return
	}

	formatter := models.FormatterDetailUser(user)
	resp := utils.ResponseSuccess(formatter, "Success get detail user")
	ctx.JSON(http.StatusOK, resp)
	return
}

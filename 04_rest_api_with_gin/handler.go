package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertHandler(context *gin.Context) {
	defer Log("Insert Handler")

	if !IsNotNil(db) {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Database connection is nil"})
		return
	}

	user := User{}

	context.BindJSON(&user)

	Log("ReqBody:", user)
	id, err := Insert(db, user)

	if IsNotNil(err) {
		context.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	message := fmt.Sprintf("Row created successfully, with id %v", id)
	context.JSON(http.StatusOK, gin.H{"message": message})
}

func SelectAllHandler(context *gin.Context) {
	defer Log("Select controller")

	if !IsNotNil(db) {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Database connection is nil"})
		return
	}

	users, err := SelectAll(db)

	if IsNotNil(err) {
		context.JSON(http.StatusOK,
			gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, users)
}

func SelectOneByIdHandler(context *gin.Context) {
	defer Log("Select by id controller")

	if !IsNotNil(db) {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Database connection is nil"})
		return
	}

	id, idCastErr := ToInt(context.Param("id"))

	if IsNil(id) || IsNotNil(idCastErr) {
		context.JSON(
			http.StatusOK, gin.H{"message": "Id is required"})
		return
	}

	user, err := SelectById(db, id)

	if IsNotNil(err) {
		context.JSON(http.StatusOK,
			gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func UpdateByIdHandler(context *gin.Context) {
	defer Log("Select by id controller")

	if !IsNotNil(db) {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Database connection is nil"})
		return
	}

	id, idCastErr := ToInt(context.Param("id"))

	if IsNil(id) || IsNotNil(idCastErr) {
		context.JSON(
			http.StatusOK, gin.H{"message": "Id is required"})
		return
	}

	// user to update
	user := User{Id: id}
	Log(user)
	context.BindJSON(&user)
	Log(user)

	err := UpdateById(db, user)
	Log(err)

	if IsNotNil(err) {
		context.JSON(http.StatusOK,
			gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK,
		gin.H{"message": "User update successfully"})
}

func DeleteByIdHandler(context *gin.Context) {
	defer Log("Select by id controller")

	if !IsNotNil(db) {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Database connection is nil"})
		return
	}

	id, idCastErr := ToInt(context.Param("id"))

	if IsNil(id) || IsNotNil(idCastErr) {
		context.JSON(
			http.StatusOK, gin.H{"message": "Id is required"})
		return
	}

	err := DeleteById(db, id)

	if IsNotNil(err) {
		context.JSON(http.StatusOK,
			gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK,
		gin.H{"message": "User deleted successfully"})
}

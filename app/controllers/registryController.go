package controllers

import (
	"errors"
	"fmt"
	"net/http"

	logs "github.com/ashishbhatt01/registeryApp/app/logging"
	"github.com/ashishbhatt01/registeryApp/app/models"
	"github.com/ashishbhatt01/registeryApp/app/services"

	"github.com/gin-gonic/gin"
)

type RegistryController struct {
	registryService services.IRegistery
}

//create new controller
func NewRegistryController(registryService services.IRegistery) *RegistryController {
	return &RegistryController{
		registryService: registryService,
	}
}

//it adds to the global value registry
func (rc *RegistryController) Add(ctx *gin.Context) {
	data, err := rc.BindRequestData(ctx)

	if err != nil {
		rc.HandleResponse(ctx, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = rc.registryService.AddData(*data)
	if err != nil {
		rc.HandleResponse(ctx, http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	} else {
		rc.HandleResponse(ctx, http.StatusOK, gin.H{"isAdded": true})
	}
}

//it substract from the global value registry
func (rc *RegistryController) Subs(ctx *gin.Context) {
	data, err := rc.BindRequestData(ctx)

	if err != nil {
		rc.HandleResponse(ctx, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = rc.registryService.SubstractData(*data)

	if err != nil {
		rc.HandleResponse(ctx, http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	} else {
		rc.HandleResponse(ctx, http.StatusOK, gin.H{"isSubstracted": true})
	}
}

//it returns the current value in registry
func (rc *RegistryController) Get(ctx *gin.Context) {
	data := rc.registryService.GetData()
	rc.HandleResponse(ctx, http.StatusOK, data)
}

//BindRequestData() => validate the incoming request
func (rc *RegistryController) BindRequestData(ctx *gin.Context) (*models.Register, error) {
	logs.LogInfo("Unmarshaling incoming request")
	var data *models.Register
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		logs.LogError(fmt.Sprintf("cannot unmarshal body, error: %s", err))
		return nil, errors.New("invalid data received")
	}
	return data, nil
}

//HandleResponse() => handle the Json response
func (rc *RegistryController) HandleResponse(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}

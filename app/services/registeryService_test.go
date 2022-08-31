package services

import (
	"log"
	"os"
	"testing"

	logs "github.com/ashishbhatt01/registeryApp/app/logging"
	"github.com/ashishbhatt01/registeryApp/app/models"
	"github.com/stretchr/testify/assert"
)

func Test_AddData_Success(t *testing.T) {
	intializeTestLogger()
	service := NewRegistryService()
	TotalValue = models.Register{
		Value: 0,
	}
	sampleData := models.Register{
		Value: 10,
	}
	err := service.AddData(sampleData)
	data := service.GetData()

	assert.Nil(t, err, "it should not return err")
	assert.Equal(t, data.Value, 10, "return value should be 10")
}

func Test_AddData_Failure(t *testing.T) {
	intializeTestLogger()
	service := NewRegistryService()
	TotalValue = models.Register{
		Value: 0,
	}
	sampleData := models.Register{
		Value: -10,
	}
	err := service.AddData(sampleData)
	data := service.GetData()

	assert.NotNil(t, err, "it should return error")
	assert.Equal(t, data.Value, 0, "returned value should be zero")
}

func Test_SubstractData_Success(t *testing.T) {
	intializeTestLogger()
	service := NewRegistryService()
	TotalValue = models.Register{
		Value: 100,
	}
	sampleData := models.Register{
		Value: 40,
	}
	err := service.SubstractData(sampleData)
	data := service.GetData()

	assert.Nil(t, err, "it should not return error")
	assert.Equal(t, data.Value, 60, "returned value should be 60")
}

func Test_SubstractData_Failure_1(t *testing.T) {
	intializeTestLogger()
	service := NewRegistryService()
	TotalValue = models.Register{
		Value: 10,
	}
	sampleData := models.Register{
		Value: 30,
	}
	err := service.SubstractData(sampleData)
	data := service.GetData()

	assert.NotNil(t, err, "it should return error")
	assert.Equal(t, data.Value, 10, "returned value should be 10")
}

func Test_SubstractData_Failure_2(t *testing.T) {
	intializeTestLogger()
	service := NewRegistryService()
	TotalValue = models.Register{
		Value: 10,
	}
	sampleData := models.Register{
		Value: -20,
	}
	err := service.SubstractData(sampleData)
	data := service.GetData()

	assert.NotNil(t, err, "it should return error")
	assert.Equal(t, data.Value, 10, "returned value should be 10")
}

func intializeTestLogger() {
	logs.InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logs.WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logs.ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

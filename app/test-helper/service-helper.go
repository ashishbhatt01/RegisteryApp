package testhelper

import (
	"log"
	"os"

	logs "github.com/ashishbhatt01/registeryApp/app/logging"
	"github.com/ashishbhatt01/registeryApp/app/models"
	"github.com/ashishbhatt01/registeryApp/app/services"
)

type MockRegisterService struct {
}

//mock AddData func from service
func (m MockRegisterService) AddData(data models.Register) error {
	return nil
}

//mock SubstractData func from service
func (m MockRegisterService) SubstractData(data models.Register) error {
	return nil
}

//mock GetData func from service
func (m MockRegisterService) GetData() models.Register {
	return services.TotalValue
}

//mock the TotalValue in services
func (m MockRegisterService) SetData(data models.Register) {
	services.TotalValue = data
}

//IntializeTestLogger() => initialize the test logger for testing
func IntializeTestLogger() {
	logs.InfoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logs.WarningLogger = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	logs.ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

package services

import (
	"errors"
	"fmt"
	"log"

	logs "github.com/ashishbhatt01/registeryApp/app/logging"
	"github.com/ashishbhatt01/registeryApp/app/models"
)

const (
	add string = "addition"
	sub string = "substraction"
)

var TotalValue models.Register

//registry interface
type IRegistery interface {
	AddData(data models.Register) error
	SubstractData(data models.Register) error
	GetData() models.Register
}

type Registery struct {
}

//create new Registry
func NewRegistryService() IRegistery {
	return &Registery{}
}

//initialize the TotalValue with 0
func init() {
	log.Println("Initializing TotalValue variable...")
	TotalValue = models.Register{
		Value: 0,
	}
	log.Println("TotalValue initialize with value 0.")
}

//AddData() => add data to global value registry
func (r Registery) AddData(data models.Register) error {
	err := r.Validatedata(add, data)
	if err != nil {
		logs.LogWarning(err.Error())
		return err
	}
	logs.LogInfo(fmt.Sprintf("Adding value %d", data.Value))
	TotalValue.Value = TotalValue.Value + data.Value
	return nil
}

//substractData() => substract the value from global registry
func (r Registery) SubstractData(data models.Register) error {
	err := r.Validatedata(sub, data)
	if err != nil {
		logs.LogWarning(err.Error())
		return err
	}
	logs.LogInfo(fmt.Sprintf("Substracting value %d", data.Value))
	TotalValue.Value = TotalValue.Value - data.Value
	return nil
}

//GetData() => get totalvalue in registry
func (r Registery) GetData() models.Register {
	logs.LogInfo(fmt.Sprintf("Current Value %d", TotalValue.Value))
	return TotalValue
}

//Validatedata() => it will validate the incoming data against the operation
func (r Registery) Validatedata(operation string, data models.Register) error {
	logs.LogInfo(fmt.Sprintf("Validating data received for operation %s", operation))
	switch operation {
	case add:
		if data.Value <= 0 {
			return fmt.Errorf("cannot add negative or zero value %d", data.Value)
		} else {
			return nil
		}
	case sub:
		currentData := r.GetData()
		if data.Value <= 0 {
			return fmt.Errorf("cannot substract negative or zero value %d", data.Value)
		} else if currentData.Value < data.Value {
			return fmt.Errorf("cannot substract, value received %d  is greater than existing value %d", data.Value, currentData.Value)
		} else if currentData.Value <= 0 {
			return fmt.Errorf("cannot substract %d, existing value is %d", data.Value, currentData.Value)
		} else {
			return nil
		}
	default:
		return errors.New("no matching operation found")
	}
}

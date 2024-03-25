package gormvalidator

import (
	"errors"
	"gorm.io/gorm"
	"reflect"
)

type SqlNullFieldValidator struct {
	DB        *gorm.DB     // gorm.DB Instance
	Table     reflect.Type //
	TableName string       // Table Name
}

// TableExists method to check if a table exists in the database
func (v *SqlNullFieldValidator) TableExists() bool {
	return v.DB.Migrator().HasTable(v.TableName)
}

func (v *SqlNullFieldValidator) ValidateNullFields() ([]interface{}, error) {
	//	Check if the table exists
	if !v.TableExists() {
		return nil, errors.New("the table does not exist in the database")
	}

	// Create a new instance of the table structure
	tablePtr := reflect.New(reflect.SliceOf(v.Table)).Elem().Interface()

	// Realizar la consulta y mapear los resultados a un slice de la estructura de tabla
	result := v.DB.Find(&tablePtr)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convertir tablePtr a un slice del tipo correcto
	registros := reflect.ValueOf(tablePtr)
	nullRecords := make([]interface{}, 0)
	for i := 0; i < registros.Len(); i++ {
		// Obtener el elemento actual del slice
		registro := registros.Index(i).Interface()

		//	Check if the record has null fields
		if hasNullFields(registro) {
			nullRecords = append(nullRecords, registro)
		}
	}
	return nullRecords, nil
}

// Helper function to check is a strcut has null fields
func hasNullFields(record interface{}) bool {
	val := reflect.ValueOf(record)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		//	Check if field is null
		if field.Kind() == reflect.Ptr && field.IsNil() {
			return true
		}
	}
	return false
}

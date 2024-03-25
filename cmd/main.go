package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"reflect"
	"time"
	"validator-slq/internal/domain"
	gormvalidator "validator-slq/pkg/validator/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3307)/C_14A?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//log.Fatalf("Error al conectar a la base de datos: %v", err)
		panic("Error al conectar a la base de datos")
	}

	// Migrar los esquemas (crear la tabla si aun no existe)
	if err := db.AutoMigrate(&domain.Producto{}); err != nil {
		log.Fatalf("Error al migrar los esquemas: %v", err)
	}

	//	Crear una instancia del validador para la tabla productos
	productValidator := &gormvalidator.SqlNullFieldValidator{
		DB:        db,
		Table:     reflect.TypeOf(domain.Producto{}),
		TableName: "productos",
	}

	//TEST:::
	//	Verificar si la tabla existe
	if !productValidator.TableExists() {
		log.Printf("La tabla %s no existe en la base de datos", productValidator.TableName)
	}

	//	Ejecuta la validacion de campos nulos
	nullRecords, err := productValidator.ValidateNullFields()
	if err != nil {
		log.Fatalf("Error al validar campos nulos: %v", err)
	}

	//	Imprimir registros con campos nulos
	fmt.Println("Registros con campos nulos:")
	for _, record := range nullRecords {
		// Convertir cada registro a tipo domain.Producto
		producto := record.(domain.Producto)
		fmt.Printf("%#v\n", producto)
		err := setDefaultFields(&producto, db)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func setDefaultFields(nullRecords *domain.Producto, db *gorm.DB) error {

	if nullRecords.Expiration == nil {
		//log.Printf("%+v", product)
		//nullExpirationProducts = append(nullExpirationProducts, product)

		// Establecer un valor predeterminado para Expiration
		defaultExpiration := time.Now().AddDate(1, 0, 0)
		defaultExpirationStr := defaultExpiration.Format("2006-01-02") // Convertir a string
		nullRecords.Expiration = &defaultExpirationStr

		// Guardar los cambios en la base de datos
		if err := db.Save(nullRecords).Error; err != nil {
			return err
		}
		log.Printf("%+v", nullRecords)
	}
	return nil
}

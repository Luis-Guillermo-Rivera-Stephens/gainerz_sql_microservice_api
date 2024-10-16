package data

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Get_db() (*gorm.DB, error) {
	var err error
	var url string
	if DB == nil { //Si la base no esta inicializada
		url, err = getDataBaseUrl() //funcion que retorna el DNS del .env
		if err != nil {
			fmt.Println("No se pudo obtener el link de la base de datos")
			return nil, err
		}

		DB, err := gorm.Open(sqlserver.Open(url), &gorm.Config{}) //inicia la conexión

		if err != nil {
			fmt.Println("Error connecting to database")
			return nil, err
		}

		return DB, nil
	}

	sqls, err := DB.DB()
	if err != nil {
		fmt.Println("Error al comprobar la conexion de la base de datos")
		return nil, err
	}

	if err = sqls.Ping(); err != nil { //si el ping no fue exitoso
		fmt.Println("DB ping unsuccessful")
		return nil, err
	}

	fmt.Println("DB is conected")
	return DB, nil
}

func getDataBaseUrl() (string, error) {
	err := godotenv.Load("./DB.env") //carga el .env, lo hace accesible
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	databaseURL := os.Getenv("DATABASE_URL") //busca en el .env esa variable
	if databaseURL == "" {
		fmt.Println("No se encontro el url")
		return "", err
	}

	return databaseURL, nil
}
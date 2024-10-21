package data

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Get_db(flag bool, test bool) (*gorm.DB, error) {
	var err error
	var url string
	if DB == nil { //Si la base no esta inicializada
		url, err = getDataBaseUrl(test) //funcion que retorna el DNS del .env
		if err != nil {
			fmt.Println("No se pudo obtener el link de la base de datos")
			return nil, err
		}

		DB, err := gorm.Open(sqlserver.Open(url), &gorm.Config{}) //inicia la conexión

		if err != nil {
			fmt.Println("Error connecting to database")
			return nil, err
		}

		if flag {
			DB.AutoMigrate(&GYMCHAIN{}, &GYMIMAGE{}, &USER{}, &VIP{}, &COACH{}, &GYM{})
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

func getDataBaseUrl(test bool) (string, error) {
	var server, port, user, password, database string
	err := godotenv.Load("./app/data/DB.env") //carga el .env, lo hace accesible
	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	if test {
		server = os.Getenv("DB_SERVER_TEST")
		port = os.Getenv("DB_PORT_TEST")
		user = os.Getenv("DB_USER_TEST")
		password = os.Getenv("DB_PASSWORD_TEST")
		database = os.Getenv("DB_NAME_TEST")

	} else {
		server = os.Getenv("DB_SERVER")
		port = os.Getenv("DB_PORT")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		database = os.Getenv("DB_NAME")
	}

	databaseURL := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable",
		user, password, server, port, database)

	return databaseURL, nil
}

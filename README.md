## Jobsity

- This interview challenge is designed to assess a candidate's ability to build a full-stack web application using a combination of front-end
  and back-end technologies.

## Folder structure

- Go: containt the code of the backend app
- Angular: containt the code of the frontend

## Run the backend Go 1.22 or latest

Move to the `Go` folder and configure the values for the database inside. In the database folder, you will find a file named `database.go`.

```go
func Connection() {
	host := getEnv("DB_HOST", "localhost")         // host
	port := getEnv("DB_PORT", "3306")              // port of you database
	user := getEnv("DB_USER", "root")              // your database user
	password := getEnv("DB_PASSWORD", "123456789") // your database password
	database := getEnv("DB_NAME", "jobsity")       // name of you database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, database)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	d.AutoMigrate(&models.Task{})
	DB = d
}
```

move to `cmd` folder and run the following command

- `go run .`

## Run the fron-end

#### First be sure you have NodeJS 20 LST or latest and run teh following commndas in angular folder

- npm install
- ng serve

be sure that the apiUrl is setup in the file `environment.development.ts` you find this file in `src > environments`

```TS
export const environment = {
  apiUrl: 'http://localhost:8080/',
};
```

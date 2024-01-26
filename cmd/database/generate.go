package main

import (
	"fmt"
	"os/exec"

	"github.com/kasfulk/orders-backend/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func GenerateStruct() {
	config, _ := configs.LoadConfig("./configs")
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable TimeZone=Asia/Makassar", config.Database.DBHost, config.Database.Username, config.Database.Password, config.Database.DBName, config.Database.DBPort)
	g := gen.NewGenerator(gen.Config{
		OutPath:       "domain",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		fmt.Println(err)
	}

	GenerateTable(g, db, "public")

	exec.Command("cp", "--recursive", "model/*", "domain/.").Run()
	exec.Command("rm", "-rf", "model/*").Run()
}

func GenerateTable(g *gen.Generator, db *gorm.DB, path string) {
	db.Exec("set search_path='" + path + "'")
	g.UseDB(db)
	g.GenerateAllTable()
	g.Execute()
}

func main() {
	GenerateStruct()
}

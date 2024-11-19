package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"victorbetoni/trabalho-web/internal/domain/crypt"
	"victorbetoni/trabalho-web/internal/infra/db"
	"victorbetoni/trabalho-web/internal/infra/http/jwt"
	"victorbetoni/trabalho-web/internal/infra/http/router"
	"victorbetoni/trabalho-web/internal/infra/repository"
	"victorbetoni/trabalho-web/pkg/uow"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	p, _ := crypt.EncryptPassword("1234")
	fmt.Println(string(p))
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/trabalhoweb?parseTime=true")

	if err != nil {
		panic(err)
	}

	if err := dtb.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Database connection stablished.")

	defer dtb.Close()

	uow, err := uow.NewUow(ctx, dtb)
	if err != nil {
		panic(err)
	}

	registerRepositories(uow)

	pwd, err := os.Getwd()

	if err != nil {
		log.Fatal("Couldnt get current working directory\n")
		return
	}

	pubKeyFile, err := os.ReadFile(fmt.Sprintf("%s/keys/publicKey.pem", pwd))

	if err != nil {
		log.Fatalf("Couldnt retrieve RSA keys: %s\n", err.Error())
		return
	}

	privateKeyFile, err := os.ReadFile(fmt.Sprintf("%s/keys/privateKey.pem", pwd))

	if err != nil {
		log.Fatalf("Couldnt retrieve RSA keys: %s\n", err.Error())
		return
	}

	jwt.DefineKeyPair(privateKeyFile, pubKeyFile)

	rout := router.Build()
	err = rout.Run(fmt.Sprintf(":%s", "8085"))
	if err != nil {
		log.Fatalf("Couldn't start HTTP server: %s\n", err.Error())
		return
	}
	fmt.Println("Listening on port 8085")
}

func registerRepositories(uow *uow.Uow) {
	uow.Register("ProfessorRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewProfessorRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("LoginRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewLoginRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}

package models

import (
	"encoding/json"
	"errors"
	db2 "github.com/glaubergoncalves/go-api-jwt-token/api/db"
	"github.com/glaubergoncalves/go-api-jwt-token/api/db/redis"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nome      string    `gorm:"size:255;not null;unique" json:"nome"`
	Nickname  string    `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Senha     string    `gorm:"size:100;not null;" json:"senha"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerificaSenha(hashedSenha, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(senha))
}

func (u *Usuario) AntesSalvar() error {
	hashedSenha, err := Hash(u.Senha)
	if err != nil {
		return err
	}
	u.Senha = string(hashedSenha)
	return nil
}

func (u *Usuario) Prepara() {
	u.ID = 0
	u.Nome = html.EscapeString(strings.TrimSpace(u.Nome))
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *Usuario) Valida(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Nickname invalido")
		}
		if u.Senha == "" {
			return errors.New("Senha invalida")
		}
		if u.Email == "" {
			return errors.New("Email invalido")
		}

		return nil
	case "login":
		if u.Senha == "" {
			return errors.New("Senha invalida")
		}
		if u.Email == "" {
			return errors.New("Email invalido")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Nickname invalido")
		}
		if u.Senha == "" {
			return errors.New("Senha invalido")
		}
		if u.Email == "" {
			return errors.New("Email invalido")
		}
		return nil
	}
}

// SaveUsuario - cria um usuario
func (u *Usuario) SaveUsuario(db *gorm.DB) (*Usuario, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	return u, nil
}

// FindAllUsuarios - busca todos os usuarios
func (u *Usuario) FindAllUsuarios() (*[]Usuario, error) {

	usuarios := []Usuario{}
	reply, err := redis.Get("usuarios")

	if err != nil {
		log.Println("Buscando no mysql")

		db := db2.Connect()
		defer db.Close()

		var err error

		err = db.Debug().Model(&Usuario{}).Limit(100).Find(&usuarios).Error
		if err != nil {
			return &[]Usuario{}, err
		}

		u, _ := json.Marshal(usuarios)
		redis.Set("usuarios", []byte(u))

		return &usuarios, err
	}

	log.Println("Buscando no redis")
	json.Unmarshal(reply, &usuarios)

	return &usuarios, nil
}

// FindUsuarioByID - busca um usuario no cambo por id
func (u *Usuario) FindUsuarioByID(db *gorm.DB, uid uint32) (*Usuario, error) {
	var err error
	err = db.Debug().Model(Usuario{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Usuario{}, errors.New("Usuario não encontrado")
	}
	return u, err
}

func (u *Usuario) UpdateUsuario(db *gorm.DB, uid uint32) (*Usuario, error) {

	// To hash the password
	err := u.AntesSalvar()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&Usuario{}).UpdateColumns(
		map[string]interface{}{
			"senha":     u.Senha,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Usuario{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Usuario{}, err
	}
	return u, nil
}

func (u *Usuario) DeleteUsuario(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Usuario{}).Where("id = ?", uid).Take(&Usuario{}).Delete(&Usuario{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

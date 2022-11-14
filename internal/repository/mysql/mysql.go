package repository

import (
	"fmt"
	"log"

	"github.com/AdiKhoironHasan/privy-cake-store/internal/models"
	"github.com/AdiKhoironHasan/privy-cake-store/internal/repository"
	servErrors "github.com/AdiKhoironHasan/privy-cake-store/pkg/errors"
	"github.com/jmoiron/sqlx"
)

const (
	AddNewCake   = `INSERT INTO cake_store (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, now(), now())`
	ShowAllCake  = `SELECT * FROM cake_store WHERE %s ORDER BY rating DESC`
	ShowCakeByID = `SELECT * FROM cake_store WHERE id = ? LIMIT 1`
	UpdateCake   = `UPDATE cake_store SET title = ?, description = ?, rating = ?, image = ? WHERE id = ?`
	DeleteCake   = `DELETE FROM cake_store WHERE id = ?`
)

var statement PreparedStatement

type PreparedStatement struct {
	addNewCake   *sqlx.Stmt
	showCakeByID *sqlx.Stmt
	updateCake   *sqlx.Stmt
	deleteCake   *sqlx.Stmt
}

type MySQLSQLRepo struct {
	Conn *sqlx.DB
}

func NewRepo(Conn *sqlx.DB) repository.SqlRepository {
	repo := &MySQLSQLRepo{Conn}
	InitPreparedStatement(repo)
	return repo
}

func (m *MySQLSQLRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := m.Conn.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

func InitPreparedStatement(m *MySQLSQLRepo) {
	statement = PreparedStatement{
		addNewCake:   m.Preparex(AddNewCake),
		showCakeByID: m.Preparex(ShowCakeByID),
		updateCake:   m.Preparex(UpdateCake),
		deleteCake:   m.Preparex(DeleteCake),
	}
}

func (m *MySQLSQLRepo) AddNewCake(dataCake *models.CakeModels) error {
	result, err := statement.addNewCake.Exec(dataCake.Title, dataCake.Description, dataCake.Rating, dataCake.Image)
	if err != nil {
		log.Println("Failed Query AddNewCake: ", err.Error())
		return fmt.Errorf(servErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()

	if err != nil {
		log.Println("Failed RowAffectd AddNewCake: ", err.Error())
		return fmt.Errorf(servErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("AddNewCake: No Data Changed")
		return fmt.Errorf(servErrors.ErrorNoDataChange)
	}

	return nil
}

func (m *MySQLSQLRepo) ShowAllCake(where string) ([]*models.CakeModels, error) {
	var dataCakes []*models.CakeModels

	var query string

	if where != "" && where != "%s" {
		query = fmt.Sprintf(ShowAllCake, where)
	} else {
		query = fmt.Sprintf(ShowAllCake, "1=1")
	}

	err := m.Conn.Select(&dataCakes, query)

	if err != nil {
		log.Println("Failed Query ShowAllCake : ", err.Error())
		return nil, fmt.Errorf(servErrors.ErrorDB)
	}

	if len(dataCakes) == 0 {
		log.Println("Data Not Found ShowAllCake")
		return nil, nil
	}

	return dataCakes, nil
}

func (m *MySQLSQLRepo) ShowCakeByID(id int) ([]*models.CakeModels, error) {
	var dataCake []*models.CakeModels

	err := m.Conn.Select(&dataCake, ShowCakeByID, id)

	if err != nil {
		log.Println("Failed Query ShowCakeByID : ", err.Error())
		return nil, fmt.Errorf(servErrors.ErrorDB)
	}

	if len(dataCake) == 0 {
		log.Println("Data Not Found ShowCakeByID")
		return nil, servErrors.ErrNotFound
	}

	return dataCake, nil
}

func (m *MySQLSQLRepo) UpdateCake(dataCake *models.CakeModels) error {
	result, err := statement.updateCake.Exec(dataCake.Title, dataCake.Description, dataCake.Rating, dataCake.Image, dataCake.ID)
	if err != nil {
		log.Println("Failed Query UpdateCake : ", err.Error())
		return fmt.Errorf(servErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("Failed RowAffectd UpdateCake : ", err.Error())
		return fmt.Errorf(servErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("UpdateCake: No Data Changed")
		return servErrors.ErrInvalidRequest
	}

	return nil
}

func (m *MySQLSQLRepo) DeleteCake(id int) error {
	result, err := statement.deleteCake.Exec(id)
	if err != nil {
		log.Println("Failed Query DeleteCake : ", err.Error())
		return fmt.Errorf(servErrors.ErrorDB)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Println("Failed RowAffectd DeleteCake : ", err.Error())
		return fmt.Errorf(servErrors.ErrorDB)
	}

	if rows < 1 {
		log.Println("DeleteCake: No Data Deleted")
		return servErrors.ErrNotFound
	}

	return nil
}

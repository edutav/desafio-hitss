package repositories

import (
	"database/sql"

	"github.com/edutav/golang-api-clients/internal/models"
)

type ClientsRepository struct {
	db *sql.DB
}

type ProductController interface {
	GetAll() (*[]models.Clients, error)
	GetByID(id int) (*models.Clients, error)
	Update(id int, c models.Clients) error
	Delete(id int) error
}

func NewClientsRepository(db *sql.DB) *ClientsRepository {
	return &ClientsRepository{db: db}
}

func (r *ClientsRepository) GetAll() (*[]models.Clients, error) {
	clients := []models.Clients{}
	rows, err := r.db.Query("SELECT id, nome, sobrenome, contato, endereco, data_nascimento, cpf FROM clientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c models.Clients

		if err := rows.Scan(&c.ID, &c.Nome, &c.Sobrenome, &c.Contato, &c.Endereco, &c.DataNascimento, &c.CPF); err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	return &clients, nil
}

func (r *ClientsRepository) GetByID(id int) (*models.Clients, error) {
	var c models.Clients

	row := r.db.QueryRow("SELECT id, nome, sobrenome, contato, endereco, data_nascimento, cpf FROM clientes WHERE id=$1", id)

	if err := row.Scan(&c.ID, &c.Nome, &c.Sobrenome, &c.Contato, &c.Endereco, &c.DataNascimento, &c.CPF); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &c, nil
}

func (r *ClientsRepository) Update(id int, c *models.Clients) error {
	_, err := r.db.Exec("UPDATE clientes SET nome=$1, sobrenome=$2, contato=$3, endereco=$4, data_nascimento=$5, cpf=$6 WHERE id=$7",
		c.Nome, c.Sobrenome, c.Contato, c.Endereco, c.DataNascimento, c.CPF, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *ClientsRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM clientes WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}

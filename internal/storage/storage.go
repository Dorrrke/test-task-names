package storage

import "github.com/jackc/pgx/v5/pgxpool"

type Storage struct {
	db *pgxpool.Pool
}

func New(dbConn *pgxpool.Pool) *Storage {
	return &Storage{
		db: dbConn,
	}
}

func (stor *Storage) GetNameById() {

}
func (stor *Storage) GetAllNames() {

}
func (stor *Storage) GetNameBySurname() {

}
func (stor *Storage) GetNamesByPatronymic() {

}
func (stor *Storage) GetNamesByAge() {

}
func (stor *Storage) GetNamesBySex() {

}
func (stor *Storage) GetNamesByNational() {

}
func (stor *Storage) SaveName() {

}
func (stor *Storage) DeleteName() {

}
func (stor *Storage) UpdateName() {

}

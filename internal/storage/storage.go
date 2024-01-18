package storage

import (
	"context"
	"strings"

	"github.com/Dorrrke/test-task-names/internal/domain/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Storage struct {
	db *pgxpool.Pool
}

func New(dbConn *pgxpool.Pool) *Storage {
	return &Storage{
		db: dbConn,
	}
}

func (s *Storage) GetNameById(ctx context.Context, id int) (models.NameData, error) {
	rows := s.db.QueryRow(ctx, `SELECT name, surname, patronymic, age, gender, "national" FROM names where "nid" = $1`, id)

	var name models.NameData
	if err := rows.Scan(&name.Name, &name.Surname, &name.Patronymic, &name.Age, &name.Gender, &name.National); err != nil {
		return models.NameData{}, errors.Wrap(err, "Error parsing db info")
	}

	name.Name = strings.TrimSpace(name.Name)
	name.Surname = strings.TrimSpace(name.Surname)
	name.Patronymic = strings.TrimSpace(name.Patronymic)
	name.Gender = strings.TrimSpace(name.Gender)
	name.National = strings.TrimSpace(name.National)

	return name, nil
}

func (s *Storage) GetAllNames(ctx context.Context) ([]models.NameData, error) {
	rows, err := s.db.Query(ctx, `SELECT name, surname, patronymic, age, gender, "national" FROM names`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var names []models.NameData

	for rows.Next() {
		var name models.NameData
		err = rows.Scan(&name.Name, &name.Surname, &name.Patronymic, &name.Age, &name.Gender, &name.National)
		if err != nil {
			return nil, err
		}
		name.Name = strings.TrimSpace(name.Name)
		name.Surname = strings.TrimSpace(name.Surname)
		name.Patronymic = strings.TrimSpace(name.Patronymic)
		name.Gender = strings.TrimSpace(name.Gender)
		name.National = strings.TrimSpace(name.National)
		names = append(names, name)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return names, nil
}

func (s *Storage) GetNameBySurname(ctx context.Context, surname string) (models.NameData, error) {
	rows := s.db.QueryRow(ctx, `SELECT name, surname, patronymic, age, gender, "national" FROM names where surname = $1`, surname)

	var name models.NameData

	if err := rows.Scan(&name.Name, &name.Surname, &name.Patronymic, &name.Age, &name.Gender, &name.National); err != nil {
		return models.NameData{}, errors.Wrap(err, "Error parsing db info")
	}

	name.Name = strings.TrimSpace(name.Name)
	name.Surname = strings.TrimSpace(name.Surname)
	name.Patronymic = strings.TrimSpace(name.Patronymic)
	name.Gender = strings.TrimSpace(name.Gender)
	name.National = strings.TrimSpace(name.National)

	return name, nil
}
func (s *Storage) GetNamesByPatronymic(ctx context.Context, patronymic string) (models.NameData, error) {
	rows := s.db.QueryRow(ctx, `SELECT name, surname, patronymic, age, gender, "national" FROM names where patronymic = $1`, patronymic)

	var name models.NameData

	if err := rows.Scan(&name.Name, &name.Surname, &name.Patronymic, &name.Age, &name.Gender, &name.National); err != nil {
		return models.NameData{}, errors.Wrap(err, "Error parsing db info")
	}

	name.Name = strings.TrimSpace(name.Name)
	name.Surname = strings.TrimSpace(name.Surname)
	name.Patronymic = strings.TrimSpace(name.Patronymic)
	name.Gender = strings.TrimSpace(name.Gender)
	name.National = strings.TrimSpace(name.National)

	return name, nil

}
func (s *Storage) GetNamesByAge(ctx context.Context, age int) (models.NameData, error) {
	rows := s.db.QueryRow(ctx, `SELECT name, surname, patronymic, age, gender, "national" FROM names where age = $1`, age)

	var name models.NameData

	if err := rows.Scan(&name.Name, &name.Surname, &name.Patronymic, &name.Age, &name.Gender, &name.National); err != nil {
		return models.NameData{}, errors.Wrap(err, "Error parsing db info")
	}

	name.Name = strings.TrimSpace(name.Name)
	name.Surname = strings.TrimSpace(name.Surname)
	name.Patronymic = strings.TrimSpace(name.Patronymic)
	name.Gender = strings.TrimSpace(name.Gender)
	name.National = strings.TrimSpace(name.National)

	return name, nil

}
func (s *Storage) GetNamesByGender(ctx context.Context, gender string) (models.NameData, error) {
	rows := s.db.QueryRow(ctx, `SELECT name, surname, patronymic, age, gender, "national" FROM names where gender = $1`, gender)

	var name models.NameData

	if err := rows.Scan(&name.Name, &name.Surname, &name.Patronymic, &name.Age, &name.Gender, &name.National); err != nil {
		return models.NameData{}, errors.Wrap(err, "Error parsing db info")
	}

	name.Name = strings.TrimSpace(name.Name)
	name.Surname = strings.TrimSpace(name.Surname)
	name.Patronymic = strings.TrimSpace(name.Patronymic)
	name.Gender = strings.TrimSpace(name.Gender)
	name.National = strings.TrimSpace(name.National)

	return name, nil

}
func (s *Storage) GetNamesByNational(ctx context.Context, national string) (models.NameData, error) {
	rows := s.db.QueryRow(ctx, `SELECT name, surname, patronymic, age, gender, "national" FROM names where national = $1`, national)

	var name models.NameData

	if err := rows.Scan(&name.Name, &name.Surname, &name.Patronymic, &name.Age, &name.Gender, &name.National); err != nil {
		return models.NameData{}, errors.Wrap(err, "Error parsing db info")
	}

	name.Name = strings.TrimSpace(name.Name)
	name.Surname = strings.TrimSpace(name.Surname)
	name.Patronymic = strings.TrimSpace(name.Patronymic)
	name.Gender = strings.TrimSpace(name.Gender)
	name.National = strings.TrimSpace(name.National)

	return name, nil

}
func (s *Storage) SaveName(ctx context.Context, name models.NameData) error {
	_, err := s.db.Exec(ctx, `INSERT INTO names (name, surname, patronymic, age, gender, "national") values ($1, $2, $3, $4, $5, $6)`,
		name.Name,
		name.Surname,
		name.Patronymic,
		name.Age,
		name.Gender,
		name.National,
	)
	if err != nil {
		return errors.Wrap(err, "Error while inserting row in db")
	}
	return nil

}
func (s *Storage) DeleteName(ctx context.Context, id int) error {
	_, err := s.db.Exec(ctx, `DELETE FROM names where "nid" = $1`, id)

	return err
}
func (s *Storage) UpdateName(ctx context.Context, name models.NameData, id int) error {
	_, err := s.db.Exec(ctx, `UPDATE names SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, "national" = $6, WHERE "nid" = $7`,
		name.Name,
		name.Surname,
		name.Patronymic,
		name.Age,
		name.Gender,
		name.National,
		id)

	return err
}

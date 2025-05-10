package bioskop

import (
	"database/sql"

	"formative-14/modules/bioskop/dto/req"
	"formative-14/modules/bioskop/dto/res"
)

// Contract 
type Repository interface {
	CreateBioskop(bioskop req.CreateBioskopDto) (res.GetBioskopDto, error)
	GetAllBioskop() ([]res.GetBioskopDto, error)
	GetBioskopById(id int) (res.GetBioskopDto, error)
	HardDeleteBioskop(id int) error
	UpdateBioskop(id int, bioskop req.UpdateBioskopDto) (res.GetBioskopDto, error)
}

// Repository struct holds the DB connection.
type bioskopRepository struct {
	DB *sql.DB
}

// NewRepository creates a new Bioskop repository instance.
func NewRepository(db *sql.DB) *bioskopRepository {
	return &bioskopRepository{
		DB: db,
	}
}

func (r *bioskopRepository) CreateBioskop(bioskop req.CreateBioskopDto) (res.GetBioskopDto, error) {
	bioskopEntity := Bioskop{
		Nama:   bioskop.Nama,
		Lokasi: bioskop.Lokasi,
		Rating: bioskop.Rating,
	}
	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskopEntity.ID)
	if err != nil {
		return res.GetBioskopDto{}, err
	}
	return res.GetBioskopDto{
		ID:     bioskopEntity.ID,
		Nama:   bioskopEntity.Nama,
		Lokasi: bioskopEntity.Lokasi,
		Rating: bioskopEntity.Rating,
	}, nil
}

func (r *bioskopRepository) GetAllBioskop() ([]res.GetBioskopDto, error){
	bioskops := []res.GetBioskopDto{}
	query := `SELECT id, nama, lokasi, rating FROM bioskop`
	rows, err := r.DB.Query(query)
	if err != nil {
		return []res.GetBioskopDto{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var bioskop res.GetBioskopDto
		if err := rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating); err != nil {
			return []res.GetBioskopDto{}, err
		}
		bioskops = append(bioskops, bioskop)
	}
	return bioskops, nil
}

func (r *bioskopRepository) GetBioskopById(id int) (res.GetBioskopDto, error) {
	bioskop := res.GetBioskopDto{}
	query := `SELECT id, nama, lokasi, rating FROM bioskop WHERE id = $1`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return res.GetBioskopDto{}, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating); err != nil {
			return res.GetBioskopDto{}, err
		}
		return bioskop, nil
	}
	return res.GetBioskopDto{}, sql.ErrNoRows
}

func (r *bioskopRepository) HardDeleteBioskop(id int) error {
	query := `DELETE FROM bioskop WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *bioskopRepository) UpdateBioskop(id int, bioskop req.UpdateBioskopDto) (res.GetBioskopDto, error) {
	bioskopEntity := Bioskop{
		Nama:   bioskop.Nama,
		Lokasi: bioskop.Lokasi,
		Rating: bioskop.Rating,
	}
	query := `UPDATE bioskop SET nama = $1, lokasi = $2, rating = $3 WHERE id = $4 RETURNING id`
	err := r.DB.QueryRow(query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating, id).Scan(&bioskopEntity.ID)
	if err != nil {
		return res.GetBioskopDto{}, err
	}
	return res.GetBioskopDto{
		ID:     bioskopEntity.ID,
		Nama:   bioskopEntity.Nama,
		Lokasi: bioskopEntity.Lokasi,
		Rating: bioskopEntity.Rating,
	}, nil
}
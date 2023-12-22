package user

import "gorm.io/gorm"

// Kontrak kerja yang mendefinisikan apa yang harus dilakukan, tetapi tidak mendefinisikan bagaimana melakukannya. Misalnya, "Anda harus bisa bahasa Go", tetapi tidak mendefinisikan bagaimana Anda belajar Go.
type Repository interface {
	// Tipe data lain yang mengimplementasikan metode Save dengan parameter dan tipe pengembalian yang sama dapat dianggap sebagai tipe Repository
	Save(user User) (User, error)
}

// Pekerja yang menandatangani kontrak dan menunjukkan bahwa mereka bisa bahasa Go dengan menyediakan implementasi untuk metode Save.
type repository struct {
	// pointer ke instance gorm.DB
	db *gorm.DB
}

// Inputan berupa db
// Return berupa struct repository
func NewRepository(db *gorm.DB) *repository {
	// Membuat instance repository
	return &repository{db}
}

// Function untuk struct repository
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

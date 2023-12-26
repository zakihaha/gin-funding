package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
}

type repositry struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositry {
	return &repositry{db}
}

func (r *repositry) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("campaign_id = ?", campaignID).Preload("User").Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

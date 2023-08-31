package services

import (
	"github.com/google/uuid"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/database"
	"www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/models"
	reviewSchema "www.github.com/BalkanID-University/vit-2025-summer-engineering-internship-task-souvik150/internal/schemas/review"
)

func CreateReview(userID uuid.UUID, bookID uuid.UUID, userName string, payload reviewSchema.CreateReviewSchema) (models.Review, error) {
	review := models.Review{
		UserID:   userID,
		BookID:   bookID,
		Username: userName,
		Rating:   float64(payload.Rating),
		Comment:  payload.Comment,
		//Approved: false,
	}

	err := database.DB.Create(&review).Error
	if err != nil {
		return models.Review{}, err
	}

	return review, nil
}

func GetReviewByUserIDAndBookID(userID uuid.UUID, bookID uuid.UUID) (models.Review, error) {
	var review models.Review
	result := database.DB.First(&review, "user_id = ? AND book_id = ?", userID, bookID)
	if result.Error != nil {
		return models.Review{}, result.Error
	}

	return review, nil
}

func UpdateReview(userID uuid.UUID, bookID uuid.UUID, userName string, payload reviewSchema.CreateReviewSchema) (models.Review, error) {
	review := models.Review{
		UserID:   userID,
		BookID:   bookID,
		Username: userName,
		Rating:   float64(payload.Rating),
		Comment:  payload.Comment,
		//Approved: false,
	}

	result := database.DB.Model(&review).Where("user_id = ? AND book_id = ?", userID, bookID).Updates(review)
	if result.Error != nil {
		return models.Review{}, result.Error
	}

	return review, nil
}

func GetReviewByID(reviewID uuid.UUID) (models.Review, error) {
	var review models.Review
	result := database.DB.First(&review, "id = ?", reviewID)
	if result.Error != nil {
		return models.Review{}, result.Error
	}

	return review, nil
}

func GetReviewByBookID(bookID uuid.UUID) ([]models.Review, error) {
	var reviews []models.Review
	result := database.DB.Find(&reviews, "book_id = ?", bookID)
	if result.Error != nil {
		return nil, result.Error
	}

	return reviews, nil
}

func DeleteReview(reviewID uuid.UUID) error {
	review, err := GetReviewByID(reviewID)
	if err != nil {
		return err
	}

	result := database.DB.Delete(&review)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

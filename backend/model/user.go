package model

func GetUser(googleUserId string) (User, error) {
	var user User

	tx := db.Where("google_user_id = ?", googleUserId).First(&user)

	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}

func CreateUser(user User) (User, error) {
	tx := db.Create(&user)
	return user, tx.Error
}
package database

import "fmt"

func GetUserFavoriteKey(userID string) string {
	return fmt.Sprintf("%s:favorite", userID)
}

func GetUserHistoryKey(userID string) string {
	return fmt.Sprintf("%s:history", userID)
}

package controllers

import (
	"fmt"
	"time"

	"[[.project]]/app/models"
)

func GenerateSystemToken(account int64, remote string, times int64, duration time.Duration) (*models.SystemToken, error) {
	var token models.SystemToken
	tx := db.Begin()
	tx.Delete(token, "system_account_id = ?", account)
	token.SystemAccountID = account
	token.Remote = remote
	token.Times = times
	token.ExpiredAt = time.Now().Add(duration)
	tx.Create(models.SystemTokenCipher(&token))
	if err := tx.Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &token, nil
}

func HitSystemToken(secret string, remote string) (*models.SystemToken, error) {
	var token models.SystemToken
	if err := db.Where("secret = ?", secret).First(&token).Error; err != nil {
		return nil, err
	}

	if token.Remote != remote {
		return nil, fmt.Errorf("ip (%s) changed.", remote)
	}

	token.Hits++
	if token.Times > 0 && token.Hits > token.Times {
		return nil, fmt.Errorf("hit times exceed.")
	}

	if token.ExpiredAt.Before(time.Now()) {
		return nil, fmt.Errorf("token expired.")
	}
	tx := db.Begin()
	if err := tx.Model(&token).Update("hits", token.Hits).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &token, nil
}

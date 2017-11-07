// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package syncronizer

import (
	"time"

	"github.com/fatih/color"
	"github.com/zhuharev/qiwi-admin/models"
	"github.com/zhuharev/qiwi-admin/pkg/qiwi"
)

// NewContext init synchronizer adn run it in background
func NewContext() (err error) {
	go func() {
		for {
			wallets, _ := models.GetAllWallets()
			for _, wallet := range wallets {
				err = Sync(wallet.ID)
				if err != nil {
					color.Red("%s", err)
				}
			}
			time.Sleep(5 * time.Minute)
		}
	}()
	return
}

// Sync specified wallet
func Sync(walletID uint) (err error) {
	wallet, err := models.GetWallet(walletID)
	if err != nil {
		return
	}

	lastTxnID, _ := models.GetLastTxn(wallet.ID)

	txns, err := qiwi.GetLastTxns(wallet.Token, wallet.WalletID)
	if err != nil {
		return
	}

	var (
		insertTxns []models.Txn
	)

	for _, txn := range txns {
		if txn.ID > lastTxnID {
			insertTxns = append(insertTxns, txn)
		}
	}

	err = models.CreateMultipleTxns(walletID, insertTxns)
	if err != nil {
		return
	}

	return
}
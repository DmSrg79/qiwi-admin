// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/zhuharev/qiwi-admin/pkg/setting"

	"github.com/jinzhu/gorm"

	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB

	dbName = "db.sqlite"
)

// DB returns db for global usage
func DB() *gorm.DB {
	return db
}

// NewContext init db instance
func NewContext() (err error) {
	dbPath := filepath.Join(setting.App.DataDir, dbName)
	log.Println("database path", dbPath)
	db, err = gorm.Open("sqlite3", dbPath)
	if err != nil {
		return
	}
	db.LogMode(true)

	err = db.AutoMigrate(&User{},
		&Txn{},
		&Wallet{},
		&Group{},
		&App{},
		&Autotransfer{}).Error
	if err != nil {
		return fmt.Errorf("automigrate err: %s", err)
	}

	err = NewStormContext()
	return
}

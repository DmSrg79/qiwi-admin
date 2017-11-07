// Copyright 2017 Kirill Zhuharev. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package routers

import (
	"github.com/zhuharev/qiwi-admin/models"
	"github.com/zhuharev/qiwi-admin/pkg/context"
	"github.com/zhuharev/qiwi-admin/pkg/syncronizer"
)

// Index shows index page
func Index(ctx *context.Context) {
	ctx.HTML(200, "index")
}

// GlobalInit inits all packaegs
func GlobalInit() (err error) {
	err = models.NewContext()
	if err != nil {
		return
	}
	err = syncronizer.NewContext()
	if err != nil {
		return
	}
	return
}
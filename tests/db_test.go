package tests

import (
	"fmt"
	"testing"

	"github.com/wangxso/wangxsoblog/models"
)

func TestDbConnect(t *testing.T) {
	db := models.GetDB()
	fmt.Print(db)
}

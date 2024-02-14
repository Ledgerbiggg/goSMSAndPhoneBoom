package test

import (
	"goSMSBoom/utils"
	"testing"
)

func TestGetRandomString(t *testing.T) {

	randomString, err := utils.GetRandomString(10)
	if err != nil {
		t.Error(err)
	}
	t.Log(randomString)

}

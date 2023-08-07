package sdk_test

import (
	"github.com/al-online/mzapi-client-sdk/services"
	"github.com/al-online/mzapi-client-sdk/services/common/models"
	"log"
	"testing"
)

func TestSDK(t *testing.T) {
	client, err := services.NewClientWithSecret("chery", "abcdefg")
	if err != nil {
		log.Fatal(err)
	}
	resp1, err := client.GetNameByGETMethod("君莫惜1")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp1)
	resp2, err := client.GetNameByPOSTMethod("君莫惜2")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp2)
	resp3, err := client.GetNameByJSON(models.User{Username: "君莫惜3"})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(resp3)
}

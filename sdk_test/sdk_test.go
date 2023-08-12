package sdk_test

import (
	"github.com/al-online/mzapi-client-sdk/services"
	"github.com/al-online/mzapi-client-sdk/services/common/models"
	"log"
	"testing"
)

func TestSDK(t *testing.T) {
	client, err := services.NewClientWithSecret("0i0973inmxchozyp8nazmz3i9udev08bwdr8tlt1e9k05b1xdi6vpzz2", "uqcbi5l4r0r4fff43hfpo34clk88i5n5l6ib5x1mddezts4hn4l36ua6m8bz892u")
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

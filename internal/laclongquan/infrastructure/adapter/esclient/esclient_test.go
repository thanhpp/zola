package esclient_test

import (
	"testing"

	"github.com/thanhpp/zola/internal/laclongquan/domain/entity"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/accountcipher"
	"github.com/thanhpp/zola/internal/laclongquan/infrastructure/adapter/esclient"
)

var (
	accCipher, _ = accountcipher.New("12345678901234567890123456789012")
	userFac      = entity.NewUserFactory(accCipher)
	validUser, _ = userFac.NewUser("0965340948", "ThisIsAPwd", "Test user", "")
)

func newESClient() *esclient.EsClient {
	var (
		host = "https://zola-search.herokuapp.com"
	)
	return esclient.NewEsClient(host)
}

func TestCreateOrUpdateUser(t *testing.T) {
	var (
		es = newESClient()
	)
	err := es.CreateOrUpdateUser(validUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSearchUser(t *testing.T) {
	var (
		es = newESClient()
	)
	ids, err := es.SearchUser("09", 0, 20)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ids)
}

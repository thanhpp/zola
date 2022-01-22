package llqclient_test

import (
	"testing"

	"github.com/thanhpp/zola/internal/auco/infra/adapter/llqclient"
)

var (
	remoteHost = "https://zola.thanhpp.ninja"
)

func newLLQClient() *llqclient.LacLongQuanClient {
	return llqclient.NewLLQClient(remoteHost)
}

func TestValidateToken(t *testing.T) {
	var (
		c         = newLLQClient()
		testToken = `eyJhbGciOiJSUzI1NiIsImtpZCI6IkxhY0xvbmdRdWFuS2V5IiwidHlwIjoiSldUIn0.eyJleHAiOjE2NDI4ODU3MzcsImlhdCI6MTY0Mjg0MjUzNywiaXNzIjoiTGFjTG9uZ1F1YW4iLCJqdGkiOiJlZTBlM2ExZC03YjYyLTExZWMtYWY0OC0wMjQyYWMxMzAwMDMiLCJ1c2VyIjp7ImlkIjoiMGZjOWVmNzEtNzA4ZS0xMWVjLWJkMDEtMDI0MmMwYTgzMDAzIiwicm9sZSI6InVzZXIifX0.FyRQXSS-9JcI1pKg60zFn2ptWBdWKKrOYTEklD5IL2Oa9FuDtEfp4HGtyt4pQ1ok6UdqiJBojMz47DIeay76LuokATIVpUDaPUiZMGbFdw9pCxXEze_HOzjZHhRdSz6C-XKsclUu406Gz8sGNrsUisI23wRyytJNjhnzMDPdEoQgPSFBhrwyd9ZnFaBsRThyGTZd39uVTG18NkPKFQznAGv7p7bQ3rClHcPj44MuaSCKAarF4iDhah6gqi0dielosssz1ATPkqQDgZQn4qUGQ8P9Tq_JmwiNbvAz8H12Pr6XiyvcsgFQNpaf_eej_Kdc1nVteJnCYYvSc_wQfNpngjfsQHWR5mIYNTttuaADlUV13uu6AFLe-UcbLZ_f5QF6vHRb_mHCS7ckwJm00jpYmPtMW8zVtiXVOnoYzHwzNwXao4oODQ6SJkX4vK4L_bTvawCIDKeBt5BFT00j5DI_FlNgEcM_NV_xZS-_421ctWv0ALakfiMSeGJo8-kQLZ97`
	)

	resp, err := c.ValidateToken(testToken)
	if err != nil {
		t.Errorf("validate token error %v", err)
		return
	}
	t.Logf("validate token response %v", resp)
}

func TestGetUserInfo(t *testing.T) {
	var (
		c          = newLLQClient()
		testUserID = `ea033d22-708d-11ec-bd01-0242c0a83003`
	)

	resp, err := c.GetUserInfo(testUserID)
	if err != nil {
		t.Errorf("get user info error %v", err)
		return
	}
	t.Logf("get user info response %v", resp)
}

func TestIsBlockUser(t *testing.T) {
	var (
		c          = newLLQClient()
		testUserID = `ea033d22-708d-11ec-bd01-0242c0a83003`
	)

	resp, err := c.IsBlock(testUserID, testUserID)
	if err != nil {
		t.Errorf("is block user error %v", err)
		return
	}
	t.Logf("is block user response %v", resp)
}

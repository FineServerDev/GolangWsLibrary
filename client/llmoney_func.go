package client

import "GolangWsLibrary/protocol/economy"

import "C"

//export SetMoney
func getMoney(userid string) int {
	Response, err := SendRequest[economy.GetUserCreditResponse](economy.CreateGetUserCreditRequest(userid))
	if err != nil {
		return -1
	}
	return Response.Credit
}

//export SetMoney
func setMoney(userid string, money int) bool {
	_, err := SendRequest[economy.SetUserCreditResponse](economy.CreateSetUserCreditRequest(userid, money))
	if err != nil {
		return false
	}
	return true
}

//export AddMoney
func addMoney(userid string, money int) bool {
	_, err := SendRequest[economy.AlterUserCreditResponse](economy.CreateAlterUserCreditRequest(userid, money))
	if err != nil {
		return false
	}
	return true
}

//export removeMoney
func removeMoney(userid string, money int) bool {
	_, err := SendRequest[economy.AlterUserCreditResponse](economy.CreateAlterUserCreditRequest(userid, -money))
	if err != nil {
		return false
	}
	return true
}

//export transferMoney
func transferMoney(from string, to string, money int) bool {
	_, err := SendRequest[economy.TransferUserCreditResponse](economy.CreateTransferUserCreditRequest(from, to, money))
	if err != nil {
		return false
	}
	return true
}

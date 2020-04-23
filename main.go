package main

import (
	"github.com/Ungigdu/pirate_pool_query_example/contracts/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"math/big"
)

var logger, _ = logging.GetLogger("pirate_example")

//ethereum ropsten network access point
var AccessPoint = "https://ropsten.infura.io/v3/f3245cef90ed440897e43efc6b3dd0f7"
//Micropayment system contract address
var MicroPayAddr = "0x4291d9Ff189D90Ba875E0fc1Da4D602406DD7D6e"
var TokenAddr = "0xAd44c8493dE3FE2B070f33927A315b50Da9a0e25"

//pool info struct
type pool struct {
	MainAddress common.Address
	PayAddress common.Address
	Name string
	Email string
	Url string
}

func GetOpts(blockNumber uint64) *bind.CallOpts {
	var opts = new(bind.CallOpts)
	if blockNumber == 0 {
		opts = nil
	}else{
		opts.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}
	return opts
}

func RecoverPaySysContract() *generated.MicroPaySystem {
	conn,err := ethclient.Dial(AccessPoint)
	if err != nil {
		logger.Fatal("can't connect to ethereum, error : ",err)
		return nil
	}
	sys, err := generated.NewMicroPaySystem(common.HexToAddress(MicroPayAddr),conn)
	if err!=nil{
		logger.Fatal("can't recover contract, error : ",err)
		return nil
	}
	return sys
}

func RecoverToken() *generated.Token{
	conn,err := ethclient.Dial(AccessPoint)
	if err != nil {
		logger.Fatal("can't connect to ethereum, error : ",err)
		return nil
	}
	token, err := generated.NewToken(common.HexToAddress(TokenAddr),conn)
	if err!=nil{
		logger.Fatal("can't recover contract, error : ", err)
		return nil
	}
	return token
}

//query pool data and update to Pools
func QueryPoolAddresses() []common.Address{
	sys := RecoverPaySysContract()
	poolLists,err := sys.Pools(GetOpts(0))
	if err!=nil{
		logger.Error("can't get pool address list, error : ",err)
		return nil
	}
	//filter address(0)
	ret:=[]common.Address{}
	for _,d := range poolLists {
		if d.String() != "0x0000000000000000000000000000000000000000" {
			ret = append(ret, d)
		}
	}
	return ret
}

//query pool data detail
func QueryPoolDetail(poolAddr common.Address) *pool {
	sys := RecoverPaySysContract()
	d,err := sys.PoolData(GetOpts(0),poolAddr)
	if err!=nil{
		logger.Error("query pool error : ", err)
		return &pool{}
	}else{
		return &pool{
			MainAddress: d.MainAddr,
			PayAddress:  d.PayerAddr,
			Name:        string(d.ShortName),
			Email:       string(d.Email),
			Url:         string(d.Url),
		}
	}
}

//query pool users
func QueryUsersOfPool(poolAddr common.Address) []common.Address {
	sys := RecoverPaySysContract()
	users,err := sys.UsersOfPool(GetOpts(0),poolAddr)
	if err!=nil {
		logger.Error("query users error, the pool is : ", poolAddr.String(), "error : ", err)
		return nil
	}else{
		return users
	}
}

//query user balance
func GetUserBalance(userAddr common.Address) *big.Int{
	token := RecoverToken()
	n,err := token.BalanceOf(GetOpts(0),userAddr)
	if err!=nil{
		logger.Error("query user balance error : ", err)
		return nil
	}else{
		return n
	}
}

func main(){


	showBalance()
}

func showBalance(){
	pools := QueryPoolAddresses()
	println("there are those pools : ")
	for _,d := range pools {
		pd := QueryPoolDetail(d)
		println("name : ", pd.Name, "addr : ", pd.MainAddress.String(), "pool balance : ", GetUserBalance(d).String())
		println("  pool claim history : ")
		sys := RecoverPaySysContract()
		filter := &bind.FilterOpts{
			Start:0,
			End: nil,
			Context:context.Background(),
		}
		users := QueryUsersOfPool(d)
		it,err := sys.FilterPoolClaim(filter,[]common.Address{d},users)
		if err!=nil{
			logger.Error("loop over claim event error : ", err)
		}else{
			for it.Next(){
				println("  ++ claimed from user ", it.Event.User.String(), "with amount : ",it.Event.Tonken.String())
			}
		}
		println("  ====user list====")
		for _, u := range users {
			balance := GetUserBalance(u)
			println("  |_user : ", u.String(), "balance : ", balance.String())
		}
	}
}


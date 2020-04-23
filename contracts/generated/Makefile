SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	abigen --abi MicroPaySystem.abi --pkg generated --type MicroPaySystem --out MicroPaySystem.go
	abigen --abi hyperorchidprotocol.abi --pkg generated --type Token --out Token.go
	abigen --abi ExtendToken.abi --pkg generated --type ExtendToken --out ExtendToken.go
	abigen --abi ApplyToken.abi --pkg generated --type ApplyToken --out ApplyToken.go
	abigen --abi HopShop.abi --pkg generated --type HopShop --out HopShop.go
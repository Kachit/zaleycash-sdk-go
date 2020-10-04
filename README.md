# ZaleyCash API SDK GO (Unofficial)
[![Build Status](https://travis-ci.org/Kachit/zaleycash-sdk-go.svg?branch=main)](https://travis-ci.org/Kachit/zaleycash-sdk-go)
[![codecov](https://codecov.io/gh/Kachit/zaleycash-sdk-go/branch/main/graph/badge.svg?token=81py9uxbmw)](https://codecov.io/gh/Kachit/zaleycash-sdk-go)
[![Release](https://img.shields.io/github/v/release/Kachit/zaleycash-sdk-go.svg)](https://github.com/Kachit/zaleycash-sdk-go/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/zaleycash-sdk-go/blob/master/LICENSE)

## Description
Unofficial ZaleyCash API Client for Go

## API documentation
https://docs.google.com/document/d/1-B6mVykI6vh17HnjkXam3fUdJochBwkNe0oeuauqgeY/edit

## Download
```shell
go get -u github.com/kachit/zaleycash-sdk-go
```

## Usage

```go
package main

import (
	zaleycash_sdk "github.com/kachit/zaleycash-sdk-go"
	"net/http"
	"fmt"
)

func main(){
	httpClient :=&http.Client{}
    	config := zaleycash_sdk.NewConfig("secret-key", "public-key")
    	auth := zaleycash_sdk.NewAuthFromConfig(config, httpClient)
    	response, err := auth.GetToken()
    	if err != nil {
    		fmt.Println(err)
    	}
    	if !response.IsSuccess() {
    		fmt.Println(response.GetError())
    	}
    
    	var token zaleycash_sdk.Token
    	err = response.Unmarshal(&token)
    
    	client := zaleycash_sdk.NewClientFromConfig(config, &token, httpClient)
    	
    	//get myTarget token
    	response, err = client.MyTarget().GetToken("myTarget-account-id")
    	if err != nil {
    		fmt.Println(err)
    	}
    	if !response.IsSuccess() {
    		fmt.Println(response.GetError())
    	}
}
```

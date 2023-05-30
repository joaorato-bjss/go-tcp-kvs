package store

func DoStorePut(key string, value string) StorePutResponse {
	responseChannel := make(chan StorePutResponse)
	request := StorePutRequest{
		Key:         key,
		Data:        value,
		RespChannel: responseChannel,
	}

	requestChannel <- request
	return <-responseChannel
}

func DoStoreGet(key string) StoreGetResponse {
	responseChannel := make(chan StoreGetResponse)
	request := StoreGetRequest{
		Key:         key,
		RespChannel: responseChannel,
	}

	requestChannel <- request
	return <-responseChannel
}

func DoStoreDelete(key string) StoreDeleteResponse {
	responseChannel := make(chan StoreDeleteResponse)
	request := StoreDeleteRequest{
		Key:         key,
		RespChannel: responseChannel,
	}

	requestChannel <- request
	return <-responseChannel
}

func DoListGet(key string) ListGetResponse {
	responseChannel := make(chan ListGetResponse)
	request := ListGetRequest{
		Key:         key,
		RespChannel: responseChannel,
	}

	requestChannel <- request
	return <-responseChannel
}

func DoListGetAll() ListGetAllResponse {
	responseChannel := make(chan ListGetAllResponse)
	request := ListGetAllRequest{
		RespChannel: responseChannel,
	}

	requestChannel <- request
	return <-responseChannel
}

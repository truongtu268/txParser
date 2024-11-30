package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func convertHexToInt(hex string) (int64, error) {
	hexaStr := strings.Replace(hex, "0x", "", -1)
	dec, err := strconv.ParseInt(hexaStr, 16, 64)
	if err != nil {
		return 0, err
	}
	return dec, nil
}

func convertIntToHex(dec int64) string {
	return fmt.Sprintf("0x%x", dec)
}

func callToEth(methodName MethodName, params []any) ([]byte, error) {
	client := &http.Client{}
	reqData := Request{
		Method:  string(methodName),
		Params:  params,
		Jsonrpc: "2.0",
		ID:      1,
	}
	bodyData, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ethereumUrl, bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

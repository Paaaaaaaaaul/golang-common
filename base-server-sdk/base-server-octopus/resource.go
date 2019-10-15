package base_server_octopus

import (
	"encoding/json"
	"github.com/becent/golang-common/base-server-sdk"
	"github.com/becent/golang-common/gin-handler"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

func Upload(orgId int, formFile map[string]string) (map[string]string, *base_server_sdk.Error) {
	if orgId <= 0 || len(formFile) == 0 {
		return nil, base_server_sdk.ErrInvalidParams
	}

	params := make(map[string]string)
	params["orgId"] = strconv.Itoa(orgId)

	files := make(map[string]string, len(formFile))
	for fieldName, file := range formFile {
		fileName := filepath.Base(file)
		key := fieldName + gin_handler.FILES_SEPARATOR + fileName
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, base_server_sdk.ErrOpenFile
		}
		files[key] = string(data)
	}

	client := base_server_sdk.Instance
	data, err := client.DoRequest(client.Hosts.OctopusServerHost, "resource", "upload", params, files)
	if err != nil {
		return nil, err
	}
	var resp map[string]string
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, base_server_sdk.ErrServiceBusy
	}

	return resp, nil
}

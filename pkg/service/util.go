/*
 *  Copyright 2020 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package service

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/akraino-edge-stack/ealt-edge/mep/mepagent/pkg/model"
)

// get yaml and parse to struct
func GetConf(path string) (model.AppInstanceInfo, error) {
	yamlFile, err := ioutil.ReadFile(path)
	var info model.AppInstanceInfo
	if err != nil {
		return info, err
	}

	err = yaml.UnmarshalStrict(yamlFile, &info)

	if err != nil {
		return info, err
	}

	return info, nil
}

func GetAppConf(FilePath string) (model.appconf, error) {
	var Appinfo model.appconf
	yamlFile, err := ioutil.ReadFile(FilePath)
	if err != nil {
		return nil, err
	}

	err1 := yaml.UnmarshalStrict(yamlFile, Appinfo)
	if err1 != nil {
		return nil, err
	}

	return Appinfo, nil

}

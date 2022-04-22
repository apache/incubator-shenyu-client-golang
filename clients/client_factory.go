/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package clients

import (
	"github.com/incubator-shenyu-client-golang/clients/admin_client"
	"github.com/incubator-shenyu-client-golang/common/constants"
	"github.com/incubator-shenyu-client-golang/model"
	"reflect"
)

/**
 * Get ShenYuAdminClient
 **/
func NewShenYuAdminClient(client *model.ShenYuAdminClient) (adminToken model.AdminToken, err error) {
	headers := map[string][]string{}
	headers[constants.DEFAULT_CONNECTION] = []string{constants.DEFAULT_CONNECTION_VALUE}
	headers[constants.DEFAULT_CONTENT_TYPE] = []string{constants.DEFAULT_CONTENT_TYPE_VALUE}

	params := map[string]string{}
	if reflect.DeepEqual(client, model.ShenYuAdminClient{}) || client.UserName == "" || client.Password == "" {
		params[constants.ADMIN_USERNAME] = constants.DEFAULT_ADMIN_ACCOUNT
		params[constants.ADMIN_PASSWORD] = constants.DEFAULT_ADMIN_PASSWORD
	} else {
		params[constants.ADMIN_USERNAME] = client.UserName
		params[constants.ADMIN_PASSWORD] = client.Password
	}

	tokenRequest := &model.ShenYuCommonRequest{
		Url:       constants.DEFAULT_SHENYU_ADMIN_URL + constants.DEFAULT_SHENYU_TOKEN,
		Header:    headers,
		Params:    params,
		TimeoutMs: constants.DEFAULT_REQUEST_TIME,
	}

	adminToken, err = admin_client.GetShenYuAdminUser(tokenRequest)
	if err == nil {
		return adminToken, nil
	} else {
		return model.AdminToken{}, err
	}
}

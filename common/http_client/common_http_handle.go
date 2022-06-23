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

package http_client

import (
	"encoding/json"
	"net/url"
	"strings"
)

/**
 * Common http_client util
 **/
func handleCommonUrl(url string, params map[string]string) string {
	if !strings.HasSuffix(url, "?") {
		url = url + "?"
	}
	for key, value := range params {
		url = url + key + "=" + value + "&"
	}
	url = strings.TrimSuffix(url, "&")
	return url
}

func ToJsonString(object interface{}) string {
	js, _ := json.Marshal(object)
	return string(js)
}

func GetUrlFormedMap(source map[string]string) (urlEncoded string) {
	urlEncoder := url.Values{}
	for key, value := range source {
		urlEncoder.Add(key, value)
	}
	urlEncoded = urlEncoder.Encode()
	return
}

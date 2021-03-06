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

package main

import (
	"encoding/json"
	"github.com/apache/incubator-shenyu-client-golang/clients/consul_client"
	"github.com/apache/incubator-shenyu-client-golang/common/constants"
	"github.com/apache/incubator-shenyu-client-golang/common/shenyu_sdk_client"
	"github.com/apache/incubator-shenyu-client-golang/model"
	"github.com/hashicorp/go-uuid"
	"github.com/wonderivan/logger"
	"time"
)

func main() {

	//Create ShenYuConsulClient  start
	ccp := &consul_client.ConsulClientParam{
		Host:  "127.0.0.1",
		Port:  8500,
		Token: "",
	}

	sdkClient := shenyu_sdk_client.GetFactoryClient(constants.CONSUL_CLIENT)
	client, createResult, err := sdkClient.NewClient(ccp)

	if !createResult && err != nil {
		logger.Fatal("Create ShenYuConsulClient error : %+V", err)
	}

	scc := client.(*consul_client.ShenYuConsulClient)
	//Create ShenYuConsulClient end
	uuid1, _ := uuid.GenerateUUID()
	uuid2, _ := uuid.GenerateUUID()
	uuid3, _ := uuid.GenerateUUID()

	//RegisterServiceInstance start
	//init MetaDataRegister
	metaData1 := &model.MetaDataRegister{
		ServiceId: uuid1,
		AppName:   "testMetaDataRegister1", //require user provide
		Path:      "/your/path1",           //require user provide
		Enabled:   true,                    //require user provide
		Host:      "127.0.0.1",             //require user provide
		Port:      "8080",                  //require user provide
		RPCType:   "http",                  //require user provide
	}

	metaData2 := &model.MetaDataRegister{
		ServiceId: uuid2,
		AppName:   "testMetaDataRegister2", //require user provide
		Path:      "/your/path2",           //require user provide
		Enabled:   true,                    //require user provide
		Host:      "127.0.0.1",             //require user provide
		Port:      "8181",                  //require user provide
		RPCType:   "http",                  //require user provide
	}

	metaData3 := &model.MetaDataRegister{
		ServiceId: uuid3,
		AppName:   "testMetaDataRegister3", //require user provide
		Path:      "/your/path3",           //require user provide
		Enabled:   true,                    //require user provide
		Host:      "127.0.0.1",             //require user provide
		Port:      "8282",                  //require user provide
		RPCType:   "http",                  //require user provide
	}

	//register multiple metaData
	registerResult1, err := scc.RegisterServiceInstance(metaData1)
	if !registerResult1 && err != nil {
		logger.Fatal("Register consul Instance error : %+V", err)
	}

	registerResult2, err := scc.RegisterServiceInstance(metaData2)
	if !registerResult2 && err != nil {
		logger.Fatal("Register consul Instance error : %+V", err)
	}

	registerResult3, err := scc.RegisterServiceInstance(metaData3)
	if !registerResult3 && err != nil {
		logger.Fatal("Register consul Instance error : %+V", err)
	}
	//RegisterServiceInstance end

	time.Sleep(time.Second)

	//GetServiceInstanceInfo start
	instanceDetail, err := scc.GetServiceInstanceInfo(metaData1)
	nodes1, ok := instanceDetail.([]*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get consul client metaData error %+v:", err)
	}

	//range nodes
	for index, node := range nodes1 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
		}
	}

	instanceDetail2, err := scc.GetServiceInstanceInfo(metaData2)
	nodes2, ok := instanceDetail2.([]*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get consul client metaData error %+v:", err)
	}

	//range nodes1
	for index, node := range nodes2 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
		}
	}

	instanceDetail3, err := scc.GetServiceInstanceInfo(metaData3)
	nodes3, ok := instanceDetail3.([]*model.MetaDataRegister)
	if !ok {
		logger.Fatal("get consul client metaData error %+v:", err)
	}

	for index, node := range nodes3 {
		nodeJson, err := json.Marshal(node)
		if err == nil {
			logger.Info("GetNodesInfo ,success Index", index, string(nodeJson))
		}
	}
	//GetServiceInstanceInfo end

	//DeregisterServiceInstance start
	//your can chose to invoke,not require
	logger.Info("> DeregisterServiceInstance start")
	deRegisterResult1, err := scc.DeregisterServiceInstance(metaData1)
	if err != nil {
		panic(err)
	}

	deRegisterResult2, err := scc.DeregisterServiceInstance(metaData2)
	if err != nil {
		panic(err)
	}

	deRegisterResult3, err := scc.DeregisterServiceInstance(metaData3)
	if err != nil {
		panic(err)
	}

	if deRegisterResult1 && deRegisterResult2 && deRegisterResult3 {
		logger.Info("DeregisterServiceInstance success !")
	}
	//DeregisterServiceInstance end

	//do your logic

}

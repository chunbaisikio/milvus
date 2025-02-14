// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package datanode

import (
	"log"
	"testing"
)

func TestParamTable_DataNode(t *testing.T) {

	Params.Init()

	t.Run("Test NodeID", func(t *testing.T) {
		id := Params.NodeID
		log.Println("NodeID:", id)
	})

	t.Run("Test flowGraphMaxQueueLength", func(t *testing.T) {
		length := Params.FlowGraphMaxQueueLength
		log.Println("flowGraphMaxQueueLength:", length)
	})

	t.Run("Test flowGraphMaxParallelism", func(t *testing.T) {
		maxParallelism := Params.FlowGraphMaxParallelism
		log.Println("flowGraphMaxParallelism:", maxParallelism)
	})

	t.Run("Test FlushInsertBufSize", func(t *testing.T) {
		size := Params.FlushInsertBufferSize
		log.Println("FlushInsertBufferSize:", size)
	})

	t.Run("Test InsertBinlogRootPath", func(t *testing.T) {
		path := Params.InsertBinlogRootPath
		log.Println("InsertBinlogRootPath:", path)
	})

	t.Run("Test PulsarAddress", func(t *testing.T) {
		address := Params.PulsarAddress
		log.Println("PulsarAddress:", address)
	})

	t.Run("Test SegmentStatisticsChannelName", func(t *testing.T) {
		path := Params.SegmentStatisticsChannelName
		log.Println("SegmentStatisticsChannelName:", path)
	})

	t.Run("Test TimeTickChannelName", func(t *testing.T) {
		name := Params.TimeTickChannelName
		log.Println("TimeTickChannelName:", name)
	})

	t.Run("Test msgChannelSubName", func(t *testing.T) {
		name := Params.MsgChannelSubName
		log.Println("MsgChannelSubName:", name)
	})

	t.Run("Test EtcdEndpoints", func(t *testing.T) {
		endpoints := Params.EtcdEndpoints
		log.Println("EtcdEndpoints:", endpoints)
	})

	t.Run("Test MetaRootPath", func(t *testing.T) {
		path := Params.MetaRootPath
		log.Println("MetaRootPath:", path)
	})

	t.Run("Test minioAccessKeyID", func(t *testing.T) {
		id := Params.MinioAccessKeyID
		log.Println("MinioAccessKeyID:", id)
	})

	t.Run("Test minioSecretAccessKey", func(t *testing.T) {
		key := Params.MinioSecretAccessKey
		log.Println("MinioSecretAccessKey:", key)
	})

	t.Run("Test MinioUseSSL", func(t *testing.T) {
		useSSL := Params.MinioUseSSL
		log.Println("MinioUseSSL:", useSSL)
	})

	t.Run("Test MinioBucketName", func(t *testing.T) {
		name := Params.MinioBucketName
		log.Println("MinioBucketName:", name)
	})
}

/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mainline

import (
	"context"
	"fmt"
	"net/http"

	"configcenter/src/common/metadata"
)

func (m *mainline) CreateManyModelClassification(ctx context.Context, h http.Header) (resp *metadata.TopoModelNode, err error) {
	resp = new(metadata.TopoModelNode)
	subPath := "/read/mainline/model"

	err = m.client.Post().
		WithContext(ctx).
		SubResource(subPath).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

func (m *mainline) CreateModelClassification(ctx context.Context, h http.Header, bkBizID string) (resp *metadata.TopoInstanceNode, err error) {
	resp = new(metadata.TopoInstanceNode)
	subPath := fmt.Sprintf("/read/mainline/instance/%s", bkBizID)

	err = m.client.Post().
		WithContext(ctx).
		SubResource(subPath).
		WithHeaders(h).
		Do().
		Into(resp)
	return
}

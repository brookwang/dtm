/*
 * Copyright (c) 2021 yedf. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package test

import (
	"fmt"
	"github.com/dtm-labs/dtm/dtmsvr"
	"testing"
)

func TestUtils(t *testing.T) {
	fmt.Println("TestUtils")
	dtmsvr.CronExpiredTrans(1)
}

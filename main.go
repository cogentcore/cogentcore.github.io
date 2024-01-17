// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "cogentcore.org/core/gi"

func main() {
	b := gi.NewAppBody("Cogent Core")
	gi.NewLabel(b).SetType(gi.LabelHeadlineLarge).SetText("Cogent Core")
	b.NewWindow().Run().Wait()
}

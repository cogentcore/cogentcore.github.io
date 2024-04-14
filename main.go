// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "cogentcore.org/core/core"

func main() {
	b := core.NewBody("Cogent Core")
	core.NewLabel(b).SetType(core.LabelHeadlineLarge).SetText("Cogent Core")
	core.NewLabel(b).SetText("A free and open source software ecosystem for all platforms, built around a powerful, fast, and cogent core framework allowing you to Code Once, Run Everywhere.")
	core.NewLabel(b).SetText(`See <a href="https://cogentcore.org/core">core</a> for more information.`)
	b.RunMainWindow()
}

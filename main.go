// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "cogentcore.org/core/core"

func main() {
	b := core.NewBody("Cogent Core")
	core.NewText(b).SetType(core.TextHeadlineLarge).SetText("Cogent Core")
	core.NewText(b).SetText("A free and open source software ecosystem for all platforms, built around a powerful, fast, and elegant framework allowing you to Code Once, Run Everywhere.")
	core.NewText(b).SetText(`See <a href="https://cogentcore.org/core">core</a> for more information.`)
	b.RunMainWindow()
}

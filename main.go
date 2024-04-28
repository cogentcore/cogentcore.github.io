// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

//go:embed icon.svg name.png
var resources embed.FS

func main() {
	b := core.NewBody("Cogent Core")

	frame := core.NewFrame(b).Style(func(s *styles.Style) {
		s.Direction = styles.Column
		s.CenterAll()
	})
	errors.Log(core.NewSVG(frame).OpenFS(resources, "icon.svg"))
	img := core.NewImage(frame)
	errors.Log(img.OpenFS(resources, "name.png"))
	img.Style(func(s *styles.Style) {
		x := func(uc *units.Context) float32 {
			return min(uc.Dp(612), uc.Vw(90))
		}
		s.Min.Set(units.Custom(x), units.Custom(func(uc *units.Context) float32 {
			return x(uc) * (128.0 / 612.0)
		}))
	})
	core.NewText(frame).SetType(core.TextHeadlineMedium).SetText("A free and open source software ecosystem for all platforms, built around a powerful, fast, and elegant framework")
	core.NewButton(frame).SetText("Learn about the Cogent Core framework").OnClick(func(e events.Event) {
		core.TheApp.OpenURL("https://cogentcore.org/core")
	})

	b.RunMainWindow()
}

// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/htmlcore"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/pages"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

//go:embed name.png
var resources embed.FS

//go:embed content
var content embed.FS

func main() {
	b := core.NewBody("Cogent Core")
	pg := pages.NewPage(b).SetContent(content)
	b.AddAppBar(pg.MakeToolbar)
	b.AddAppBar(func(p *tree.Plan) {
		tree.Add(p, func(w *core.Button) {
			w.SetText("Blog").SetIcon(icons.RssFeed)
			w.OnClick(func(e events.Event) {
				pg.Context.OpenURL("/blog")
			})
		})
		tree.Add(p, func(w *core.Button) {
			w.SetText("GitHub").SetIcon(icons.GitHub)
			w.OnClick(func(e events.Event) {
				pg.Context.OpenURL("https://github.com/cogentcore")
			})
		})
	})

	htmlcore.ElementHandlers["home-page"] = func(ctx *htmlcore.Context) bool {
		frame := core.NewFrame(ctx.BlockParent)
		frame.Styler(func(s *styles.Style) {
			s.Direction = styles.Column
			s.Grow.Set(1, 1)
			s.CenterAll()
		})
		errors.Log(core.NewSVG(frame).ReadString(core.AppIcon))
		img := core.NewImage(frame)
		errors.Log(img.OpenFS(resources, "name.png"))
		img.Styler(func(s *styles.Style) {
			x := func(uc *units.Context) float32 {
				return min(uc.Dp(612), uc.Vw(90))
			}
			s.Min.Set(units.Custom(x), units.Custom(func(uc *units.Context) float32 {
				return x(uc) * (128.0 / 612.0)
			}))
		})
		core.NewText(frame).SetType(core.TextHeadlineMedium).SetText(core.AppAbout)
		core.NewButton(frame).SetText("Learn about the Cogent Core framework").OnClick(func(e events.Event) {
			core.TheApp.OpenURL("https://cogentcore.org/core")
		})
		return true
	}

	b.RunMainWindow()
}

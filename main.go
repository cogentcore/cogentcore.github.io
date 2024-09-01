// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"image/color"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/htmlcore"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/pages"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
	_ "cogentcore.org/core/yaegicore"
)

//go:embed name.png
var resources embed.FS

//go:embed content
var content embed.FS

func main() {
	b := core.NewBody("Cogent Core")
	pg := pages.NewPage(b).SetContent(content)
	b.AddTopBar(func(bar *core.Frame) {
		tb := core.NewToolbar(bar)
		tb.Maker(pg.MakeToolbar)
		tb.Maker(func(p *tree.Plan) {
			tree.Add(p, func(w *core.Button) {
				w.SetText("Blog").SetIcon(icons.RssFeed)
				w.OnClick(func(e events.Event) {
					pg.Context.OpenURL("/blog")
				})
			})
			tree.Add(p, func(w *core.Button) {
				w.SetText("Videos").SetIcon(icons.VideoLibrary)
				w.OnClick(func(e events.Event) {
					pg.Context.OpenURL("https://youtube.com/@CogentCore")
				})
			})
			tree.Add(p, func(w *core.Button) {
				w.SetText("GitHub").SetIcon(icons.GitHub)
				w.OnClick(func(e events.Event) {
					pg.Context.OpenURL("https://github.com/cogentcore")
				})
			})
			tree.Add(p, func(w *core.Button) {
				w.SetText("Community").SetIcon(icons.Forum)
				w.OnClick(func(e events.Event) {
					pg.Context.OpenURL("/community")
				})
			})
			tree.Add(p, func(w *core.Button) {
				w.SetText("Sponsor").SetIcon(icons.Favorite)
				w.OnClick(func(e events.Event) {
					pg.Context.OpenURL("https://github.com/sponsors/cogentcore")
				})
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
			s.Min.X.SetCustom(func(uc *units.Context) float32 {
				return min(uc.Dp(612), uc.Vw(80))
			})
		})
		core.NewText(frame).SetType(core.TextHeadlineMedium).SetText(core.AppAbout)
		core.NewButton(frame).SetText("Learn about the Cogent Core framework").OnClick(func(e events.Event) {
			core.TheApp.OpenURL("https://cogentcore.org/core")
		})
		return true
	}
	htmlcore.ElementHandlers["color-scheme-control"] = func(ctx *htmlcore.Context) bool {
		type theme struct {
			Theme core.Themes `default:"Auto"`
			Color color.RGBA  `default:"#4285f4"`
		}
		th := &theme{core.AppearanceSettings.Theme, core.AppearanceSettings.Color}
		fm := core.NewForm(ctx.BlockParent).SetStruct(th)
		fm.OnChange(func(e events.Event) {
			core.AppearanceSettings.Theme = th.Theme
			core.AppearanceSettings.Color = th.Color
			core.UpdateSettings(ctx.BlockParent, core.AppearanceSettings)
		})
		return true
	}

	b.OnShow(func(e events.Event) {
		b.Update() // TODO: needed for image sizing on initial load (core/#1037)
	})
	b.RunMainWindow()
}

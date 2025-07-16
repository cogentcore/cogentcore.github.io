// Copyright (c) 2023, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"image/color"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/content"
	"cogentcore.org/core/core"
	"cogentcore.org/core/events"
	"cogentcore.org/core/htmlcore"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

//go:embed name.png
var resources embed.FS

//go:embed content
var econtent embed.FS

func main() {
	b := core.NewBody("Cogent Core")
	ct := content.NewContent(b).SetContent(econtent)
	ctx := ct.Context
	b.AddTopBar(func(bar *core.Frame) {
		tb := core.NewToolbar(bar)
		tb.Maker(ct.MakeToolbar)
		tb.Maker(func(p *tree.Plan) {
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://cogentcore.org/blog")
				w.SetText("Blog").SetIcon(icons.RssFeed)
			})
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://youtube.com/@CogentCore")
				w.SetText("Videos").SetIcon(icons.VideoLibrary)
			})
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://github.com/cogentcore")
				w.SetText("GitHub").SetIcon(icons.GitHub)
			})
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "/community")
				w.SetText("Community").SetIcon(icons.Forum)
			})
			tree.Add(p, func(w *core.Button) {
				ctx.LinkButton(w, "https://github.com/sponsors/cogentcore")
				w.SetText("Sponsor").SetIcon(icons.Favorite)
			})
		})
	})

	ctx.ElementHandlers["home-page"] = func(ctx *htmlcore.Context) bool {
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

		buttons := core.NewFrame(frame)
		cc := core.NewButton(buttons).SetText("Cogent Core")
		ctx.LinkButton(cc, "https://cogentcore.org/core")
		cl := core.NewButton(buttons).SetText("Cogent Lab").SetType(core.ButtonTonal)
		ctx.LinkButton(cl, "https://cogentcore.org/lab")
		blog := core.NewButton(buttons).SetText("Blog").SetType(core.ButtonTonal)
		ctx.LinkButton(blog, "https://cogentcore.org/blog")
		return true
	}
	ctx.ElementHandlers["color-scheme-control"] = func(ctx *htmlcore.Context) bool {
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

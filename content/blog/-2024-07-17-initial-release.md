Today we are announcing the initial public release of the Cogent Core GUI framework. Cogent Core prioritizes the ability to Code Once, Run Everywhere (Core). You are reading this blog in a Cogent Core app running on the web via wasm (web assembly). The same code can run on macOS, Windows, Linux, iOS, Android, and the web, through a `core` command line tool that manages all the details for running and building apps for each platform.

Cogent Core is written in Go (Golang), and inherits many of the best features from this language and its associated ecosystem. Go code is simple, easy to read and write, and emphasizes a minimalist approach without the extra syntax and boilerplate that clutters other languages.

Here's a simple hello world app in Cogent Core:

```Go
package main

import "cogentcore.org/core/core"

func main() {
    b := core.NewBody()
    core.NewButton(b).SetText("Hello, World!")
    b.RunMainWindow()
}
```

See that "Hello, World!" button? That is a live rendering of the code shown in the text editor above. You can change the message, press `Ctrl+Enter` or just click off, and you'll see it update!

Cogent Core supports all the usual types of GUI widgets, along with some fairly advanced ones not found in other frameworks. Interactive, editable examples of all major widgets are available on the main [docs](https://www.cogentcore.org/core) page. (which is also a Cogent Core app running via wasm).

Here's a small sample of widgets, and a few things you can do with them:

```Go
core.NewButton(b).SetText("Hello, World!").SetIcon(icons.Send).OnClick(func(e events.Event) {
    core.MessageSnackbar(b, "Message sent to the world!")
})
core.NewText(b).SetText("Name:").SetTooltip("Enter your name in the text field")
core.NewTextField(b).SetPlaceholder("Jane Doe")
value := float32(0.5)
spinner := core.Bind(&value, core.NewSpinner(b))
slider := core.Bind(&value, core.NewSlider(b))
spinner.OnChange(func(e events.Event) {
    slider.Update()
})
slider.OnChange(func(e events.Event) {
    spinner.Update()
})
core.NewColorButton(b).SetColor(colors.Orange)

type language struct {
    Name   string
    Rating int
}
sl := []language{{"Go", 10}, {"Python", 5}}
core.NewTable(b).SetSlice(&sl).OnChange(func(e events.Event) {
    core.MessageSnackbar(b, fmt.Sprintf("Languages: %v", sl))
})
```

Again, you can modify any of the code above and immediately see the effects!

You can even make interactive plots of data:

```Go
type Data struct {
	Time       float32
	Population float32
}
data := []Data{
    {0, 500},
    {1, 800},
    {2, 1600},
    {3, 1400},
}
dt := errors.Log1(table.NewSliceTable(data))
pe := plotcore.NewPlotEditor(b).SetTable(dt)
pe.Options.XAxisColumn = "Time"
pe.ColumnOptions("Population").On = true
```

## Key Features

* Extensive styling properties allow everything to be customized, including a powerful automatic layout system that solves all the hard layout problems for you. Anyone coming from the CSS world should be able to quickly adapt, and others {insert something to appeal to people not coming from the CSS world}

* Vulkan for high performance 2D and 3D rendering

* Full-featured SVG for 2D rendering and icons

* Powerful HCT color space integrated throughout: allows instant light / dark and alternate color scheme customization.

* Efficient mechanism for dynamically updating content that captures the best of imperative and declarative mode programming.

* Focus on efficient keyboard navigation and customizable mappings, with full support for emacs mode.

*add more...*

## Our Story

We are committed to supporting Cogent Core and growing a full software ecosystem around it, building on where Cogent Core came from and where we want it to go.

The initial version of this software was named [GoKi](https://github.com/goki/gi). It was written in 2018 by Professor Randy O'Reilly to enable him to develop advanced [neural network models](https://emersim.org) of the brain using Go, instead of C++. He had grown frustrated with the long build times and tiresome boilerplate involved in coding in C++. At that time, Python was becoming increasingly popular, but it is really just a wrapper around C++, resulting in a complex and unpleasant combination of two languages. Go, by contrast, compiles nearly instantly, and runs nearly as fast as C++. The small difference in compute time (less than 5-10%) was more than made up for by the massive increase in coding efficiency and overall happiness from using Go.

The only thing missing from the Go ecosystem at the time was a full-featured native GUI framework, so Randy built on his experience with [Qt](https://en.wikipedia.org/wiki/Qt_(software)) to write one in Go. GoKi provided a powerful 2D and 3D interface that enabled experts, as well as undergraduate and graduate students in classes taught by professors around the world, to better understand and develop new ideas about how the brain works. However, as a first effort in Go, GoKi retained too much of the C++ style.

Meanwhile, Randy's son Kai was busy experimenting with different frameworks and languages for various coding projects. He eventually came to the same conclusion that Go is truly the best language around. After exploring existing GUI frameworks in Go, Kai decided that a major overhaul of GoKi might produce a better framework than any of the other options.

So the father and son team (more son than father, to be clear) spent the next year rewriting this codebase many times over, peeling away layers of complexity and finding the most robust and elegant solutions to the many problems such a framework must solve. The [principles](https://cogentcore.org/core/architecture/principles) capture some of our hard-won lessons. We hope that the experience of using this framework demonstrates the resulting simplicity and power of the approach.

Throughout this process, Randy maintained what is now Cogent Code as his primary everyday code editor, and the neural network models based on this framework were used by {convey the breadth of users in a non-redundant way}. Therefore, the core of the framework is well-tested and ready to use, even as we continue to build out more features and welcome input from the broader community for how to make it even better.

We each have a long-term commitment to the future of the Cogent Core framework. Randy and other scientists and instructors rely on it for research and teaching. Kai plans to build his career around the framework. He is programming apps in it and is available for consulting projects to develop Cogent Core solutions.

We think Go is such a special language that it deserves to be used for everything and anything, far beyond its traditional role as a server-side and cli-based workhorse. We hope that the excitement of building toward a world-class GUI framework in our beloved Go language will grow a vibrant community of users and developers.

## Future directions

One important future direction, evident in the interactive editing ability shown above, is to use the [yaegi](https://github.com/traefik/yaegi) Go interpreter as a replacement for the traditional role that Python has played relative to backend languages such as C++, so that you can transparently have a dynamic, interpreted experience as well as the lightning-fast compilation of Go. We think this can provide an ideal combination of rapid prototyping and hot-reloading (as in the Flutter framework), within a strongly typed and robust language that scales to large-scale applications (unlike Python and Javascript).

Furthermore, we have written a shell language variant of Go, called `cosh` (Cogent Shell), which allows direct intermixing of shell-like execution of command-line tools, with standard Go control constructs, using the yaegi interpreter. Everything can be transpiled into standard Go and built the usual way as a fully compiled executable as well. Next, we plan to extend this general approach to the numerical computing and data science domain, in the Cogent Numbers framework, to provide a viable competitor in this Python-dominated domain.

In addition, we will be completing the Cogent Canvas app, which provides editing of SVG-based vector graphics, and the Cogent Mail client. We also plan to make a video editing app and are working on a web browser.

## Comparisons with other frameworks

In the remainder of this blog, we provide some commentary about how we think Cogent Core compares with various other widely-used GUI frameworks and languages. As you can tell if you've read this far, we think the Go language is the best, so if you strongly prefer another language, Cogent Core may not be for you. But if you've been slogging along in the Javascript + HTML + CSS world, or come from a Qt/C++ or Flutter background, you might find something to like here.

### Javascript + HTML + CSS frameworks

Most of the world's GUI software is written using some web framework.

* Everything is built on some version of JS, HTML, and (S)CSS, all of which are poorly designed and extremely difficult to use effectively in conjunction with each other for any significant use case.

* basic widgets significantly lacking: takes huge amount of time and effort to code a basic slider or spinner. so much CSS garbage.. Kai, take it away here..

* send to link comparing tic-tac-toe implementations?

### Platform-specific frameworks

Most mobile apps are written using tools optimized for each of the two major mobile platforms (iOS and Android), requiring significant duplication of effort, and mastery of vastly different software ecosystems (Swift/Objective-C vs. Java). By contrast, Cogent Core allows you to truly write one app in one language, and have it work well on both platforms (Core = "Code Once, Run Everywhere").

### Flutter

* lots of complex boilerplate, relatively obscure language (Dart), apparently losing some support from Google.

* send to link comparing ours vs. theirs?

### Go GUIs: Fyne and Gio

* Fyne is "fine" overall, but it lacks the support for customization and complexity that is necessary to build powerful real-world apps.

* Gio has a very powerful set of platform-specific backend code, but the low-level nature of the immediate-mode design makes it very difficult to develop complex, large-scale applications, since there are too many steps to accomplish things.

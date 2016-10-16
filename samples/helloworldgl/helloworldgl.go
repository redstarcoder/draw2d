// Open an OpenGl window and display a rectangle and "Hello World" using a OpenGl GraphicContext
package main

import (
	"image/color"
	"log"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/redstarcoder/draw2d"
	"github.com/redstarcoder/draw2d/draw2dgl"
	"github.com/redstarcoder/draw2d/draw2dkit"
)

var (
	// global rotation
	rotate        int
	width, height int
	mx, my	      int
	redraw        = true
	font          draw2d.FontData
)

func reshape(window *glfw.Window, w, h int) {
	gl.ClearColor(1, 1, 1, 1)
	/* Establish viewing area to cover entire window. */
	gl.Viewport(0, 0, int32(w), int32(h))
	/* PROJECTION Matrix mode. */
	gl.MatrixMode(gl.PROJECTION)
	/* Reset project matrix. */
	gl.LoadIdentity()
	/* Map abstract coords directly to window coords. */
	gl.Ortho(0, float64(w), 0, float64(h), -1, 1)
	/* Invert Y axis so increasing Y goes down. */
	gl.Scalef(1, -1, 1)
	/* Shift origin up to upper-left corner. */
	gl.Translatef(0, float32(-h), 0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Disable(gl.DEPTH_TEST)
	width, height = w, h
	redraw = true
}

// Ask to refresh
func invalidate() {
	redraw = true
}

func display() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.LineWidth(1)
	gc := draw2dgl.NewGraphicContext(width, height)
	gc.SetFontData(draw2d.FontData{
		Name:   "luxi",
		Family: draw2d.FontFamilyMono,
		Style:  draw2d.FontStyleBold | draw2d.FontStyleItalic})
	gc.SetFontSize(14)
	// Draw Rectangle
	gc.BeginPath()
	draw2dkit.RoundedRectangle(gc, 200, 200, 600, 600, 100, 100)
	gc.SetFillColor(color.RGBA{0, 0, 0, 0xff})
	gc.Fill()
	// Display Hello World
	gc.BeginPath()
	gc.FillStringAt("Hello World", 8, 52)

	gl.Flush() /* Single buffered, so needs a flush. */
}

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	width, height = 800, 800
	window, err := glfw.CreateWindow(width, height, "Show RoundedRect", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	window.SetSizeCallback(reshape)
	window.SetKeyCallback(onKey)
	window.SetCharCallback(onChar)
	window.SetCursorPosCallback(onMMove)
	window.SetMouseButtonCallback(onMClick)

	glfw.SwapInterval(1)

	err = gl.Init()
	if err != nil {
		panic(err)
	}

	reshape(window, width, height)
	for !window.ShouldClose() {
		if redraw {
			display()
			window.SwapBuffers()
			redraw = false
		}
		glfw.PollEvents()
		//		time.Sleep(2 * time.Second)
	}
}

func onChar(w *glfw.Window, char rune) {
	log.Println(char)
}

func onMMove(w *glfw.Window, xpos, ypos float64) {
	mx, my = int(xpos), int(ypos)
}

func onMClick(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	log.Printf("Pos=(%d, %d) Btn=%d Pressed=%t", mx, my, button, action==glfw.Press)
}

func onKey(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch {
	case key == glfw.KeyEscape && action == glfw.Press,
		key == glfw.KeyQ && action == glfw.Press:
		w.SetShouldClose(true)
	}
}

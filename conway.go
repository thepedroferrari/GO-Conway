package main
import (
    "github.com/hajimehoshi/ebiten"
    "image/color"
    "math/rand"
    "log"
)
const scale=2
const WIDTH=160
const HEIGHT=120
var black color.RGBA=color.RGBA{75,139,190,255}//95,95,95
var white color.RGBA=color.RGBA{255,232,115,255}//233,233,233
var grid [WIDTH][HEIGHT]uint8=[WIDTH][HEIGHT]uint8{}
var buffer [WIDTH][HEIGHT]uint8=[WIDTH][HEIGHT]uint8{}
var counter int=0

// Logic
func update() error{
	for x:=1;x<WIDTH-1;x++ {
		for y:=1;y<HEIGHT-1;y++ {
			buffer[x][y]=0
			n:=grid[x-1][y-1]+grid[x-1][y+0]+grid[x-1][y+1]+grid[x+0][y-1]+grid[x+0][y+1]+grid[x+1][y-1]+grid[x+1][y+0]+grid[x+1][y+1]
			if grid[x][y]==0 && n==3 {
				buffer[x][y]=1
			}else if n>3 || n<2 {
				buffer[x][y]=0
			}else{
				buffer[x][y]=grid[x][y]
			}
		}
	}
	temp:=buffer
	buffer=grid
	grid=temp
	return nil
}

// Main
func render(screen *ebiten.Image){
	screen.Fill(white)
	for x:=0;x<WIDTH;x++ {
		for y:=0;y<HEIGHT;y++ {
			if grid[x][y]>0{
				for x1:=0;x1<scale;x1++ {
					for y1:=0;y1<scale;y1++ {
						screen.Set((x*scale)+x1,(y*scale)+y1,black)
					}
				}
			}
		}
	}
}
func frame(screen *ebiten.Image) error{
	counter++
	var err error=nil
	if counter==20 {
		err=update()
		counter=0
	}
	if !ebiten.IsDrawingSkipped(){
    render(screen)
  }
  return err
}
func main() {
	for x:=1;x<WIDTH-1;x++ {
		for y:=1;y<HEIGHT-1;y++ {
			if(rand.Float32()<0.5){
				grid[x][y]=1
			}
		}
	}
	if err:=ebiten.Run(frame,WIDTH*scale,HEIGHT*scale,2,"Conway's Game of Go");err!=nil{
    log.Fatal(err)
	}
}





// Ebiten (v2)

// PkgGoDev Build Status Build Status Go Report Card

// A dead simple 2D game library for Go

// Ebiten is an open source game library for the Go programming language. Ebiten's simple API allows you to quickly and easily develop 2D games that can be deployed across multiple platforms.

//     Website (ebiten.org)
//     API Reference
//     Cheat Sheet

// Overview
// Platforms

//     Windows (No Cgo!)
//     macOS
//     Linux
//     FreeBSD
//     Android
//     iOS
//     WebAssembly

// Note: Gamepad and keyboard are not available on Android/iOS.

// For installation on desktops, see the installation instruction.
// Features

//     2D Graphics (Geometry/Color matrix transformation, Various composition modes, Offscreen rendering, Fullscreen, Text rendering, Automatic batches, Automatic texture atlas)
//     Input (Mouse, Keyboard, Gamepads, Touches)
//     Audio (Ogg/Vorbis, MP3, WAV, PCM)

// Packages

//     ebiten
//         audio
//             mp3
//             vorbis
//             wav
//         ebitenutil
//         inpututil
//         mobile
//         text

// Community
// Slack

// #ebiten channel in Gophers Slack
// License

// Ebiten is licensed under Apache license version 2.0. See LICENSE file.
// Collapse ▴
// Documentation
// Rendered for
// Overview ¶

//     Environment variables
//     Build tags 

// Package ebiten provides graphics and input API to develop a 2D game.

// You can start the game by calling the function RunGame.

// // Game implements ebiten.Game interface.
// type Game struct{}

// // Update proceeds the game state.
// // Update is called every tick (1/60 [s] by default).
// func (g *Game) Update() error {
//     // Write your game's logical update.
//     return nil
// }

// // Draw draws the game screen.
// // Draw is called every frame (typically 1/60[s] for 60Hz display).
// func (g *Game) Draw(screen *ebiten.Image) {
//     // Write your game's rendering.
// }

// // Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// // If you don't have to adjust the screen size with the outside size, just return a fixed size.
// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
//     return 320, 240
// }

// func main() {
//     game := &Game{}
//     // Sepcify the window size as you like. Here, a doulbed size is specified.
//     ebiten.SetWindowSize(640, 480)
//     ebiten.SetWindowTitle("Your game's title")
//     // Call ebiten.RunGame to start your game loop.
//     if err := ebiten.RunGame(game); err != nil {
//         log.Fatal(err)
//     }
// }

// In the API document, 'the main thread' means the goroutine in init(), main() and their callees without 'go' statement. It is assured that 'the main thread' runs on the OS main thread. There are some Ebiten functions that must be called on the main thread under some conditions (typically, before ebiten.RunGame is called).
// Environment variables ¶

// `EBITEN_SCREENSHOT_KEY` environment variable specifies the key to take a screenshot. For example, if you run your game with `EBITEN_SCREENSHOT_KEY=q`, you can take a game screen's screenshot by pressing Q key. This works only on desktops.

// `EBITEN_INTERNAL_IMAGES_KEY` environment variable specifies the key to dump all the internal images. This is valid only when the build tag 'ebitendebug' is specified. This works only on desktops.
// Build tags ¶

// `ebitendebug` outputs a log of graphics commands. This is useful to know what happens in Ebiten. In general, the number of graphics commands affects the performance of your game.

// `ebitengl` forces to use OpenGL in any environments.
// Index ¶

//     Constants
//     func CurrentFPS() float64
//     func CurrentTPS() float64
//     func CursorPosition() (x, y int)
//     func DeviceScaleFactor() float64
//     func GamepadAxis(id GamepadID, axis int) float64
//     func GamepadAxisNum(id GamepadID) int
//     func GamepadButtonNum(id GamepadID) int
//     func GamepadName(id GamepadID) string
//     func GamepadSDLID(id GamepadID) string
//     func InputChars() []rune
//     func IsFocused() bool
//     func IsFullscreen() bool
//     func IsGamepadButtonPressed(id GamepadID, button GamepadButton) bool
//     func IsKeyPressed(key Key) bool
//     func IsMouseButtonPressed(mouseButton MouseButton) bool
//     func IsRunnableOnUnfocused() bool
//     func IsScreenClearedEveryFrame() bool
//     func IsScreenTransparent() bool
//     func IsVsyncEnabled() bool
//     func IsWindowDecorated() bool
//     func IsWindowFloating() bool
//     func IsWindowMaximized() bool
//     func IsWindowMinimized() bool
//     func IsWindowResizable() bool
//     func MaxTPS() int
//     func MaximizeWindow()
//     func MinimizeWindow()
//     func RestoreWindow()
//     func RunGame(game Game) error
//     func RunGameWithoutMainLoop(game Game)
//     func ScreenSizeInFullscreen() (int, int)
//     func SetCursorMode(mode CursorModeType)
//     func SetFullscreen(fullscreen bool)
//     func SetInitFocused(focused bool)
//     func SetMaxTPS(tps int)
//     func SetRunnableOnUnfocused(runnableOnUnfocused bool)
//     func SetScreenClearedEveryFrame(cleared bool)
//     func SetScreenTransparent(transparent bool)
//     func SetVsyncEnabled(enabled bool)
//     func SetWindowDecorated(decorated bool)
//     func SetWindowFloating(float bool)
//     func SetWindowIcon(iconImages []image.Image)
//     func SetWindowPosition(x, y int)
//     func SetWindowResizable(resizable bool)
//     func SetWindowSize(width, height int)
//     func SetWindowTitle(title string)
//     func TouchPosition(id TouchID) (int, int)
//     func Wheel() (xoff, yoff float64)
//     func WindowPosition() (x, y int)
//     func WindowSize() (int, int)
//     type Address
//     type ColorM
//         func (c *ColorM) Apply(clr color.Color) color.Color
//         func (c *ColorM) ChangeHSV(hueTheta float64, saturationScale float64, valueScale float64)
//         func (c *ColorM) Concat(other ColorM)
//         func (c *ColorM) Element(i, j int) float64
//         func (c *ColorM) Invert()
//         func (c *ColorM) IsInvertible() bool
//         func (c *ColorM) Reset()
//         func (c *ColorM) RotateHue(theta float64)
//         func (c *ColorM) Scale(r, g, b, a float64)
//         func (c *ColorM) SetElement(i, j int, element float64)
//         func (c *ColorM) String() string
//         func (c *ColorM) Translate(r, g, b, a float64)
//     type CompositeMode
//     type CursorModeType
//         func CursorMode() CursorModeType
//     type DrawImageOptions
//     type DrawRectShaderOptions
//     type DrawTrianglesOptions
//     type DrawTrianglesShaderOptions
//     type Filter
//     type Game
//     type GamepadButton
//     type GamepadID
//         func GamepadIDs() []GamepadID
//     type GeoM
//         func (g *GeoM) Apply(x, y float64) (float64, float64)
//         func (g *GeoM) Concat(other GeoM)
//         func (g *GeoM) Element(i, j int) float64
//         func (g *GeoM) Invert()
//         func (g *GeoM) IsInvertible() bool
//         func (g *GeoM) Reset()
//         func (g *GeoM) Rotate(theta float64)
//         func (g *GeoM) Scale(x, y float64)
//         func (g *GeoM) SetElement(i, j int, element float64)
//         func (g *GeoM) Skew(skewX, skewY float64)
//         func (g *GeoM) String() string
//         func (g *GeoM) Translate(tx, ty float64)
//     type Image
//         func NewImage(width, height int) *Image
//         func NewImageFromImage(source image.Image) *Image
//         func (i *Image) At(x, y int) color.Color
//         func (i *Image) Bounds() image.Rectangle
//         func (i *Image) Clear()
//         func (i *Image) ColorModel() color.Model
//         func (i *Image) Dispose()
//         func (i *Image) DrawImage(img *Image, options *DrawImageOptions)
//         func (i *Image) DrawRectShader(width, height int, shader *Shader, options *DrawRectShaderOptions)
//         func (i *Image) DrawTriangles(vertices []Vertex, indices []uint16, img *Image, options *DrawTrianglesOptions)
//         func (i *Image) DrawTrianglesShader(vertices []Vertex, indices []uint16, shader *Shader, ...)
//         func (i *Image) Fill(clr color.Color)
//         func (i *Image) ReplacePixels(pixels []byte)
//         func (i *Image) Set(x, y int, clr color.Color)
//         func (i *Image) Size() (width, height int)
//         func (i *Image) SubImage(r image.Rectangle) image.Image
//     type Key
//         func (k Key) String() string
//     type MouseButton
//     type Shader
//         func NewShader(src []byte) (*Shader, error)
//         func (s *Shader) Dispose()
//     type TouchID
//         func TouchIDs() []TouchID
//     type Vertex

// Constants ¶
// View Source

// const ColorMDim = affine.ColorMDim

// ColorMDim is a dimension of a ColorM.
// View Source

// const DefaultTPS = 60

// DefaultTPS represents a default ticks per second, that represents how many times game updating happens in a second.
// View Source

// const GeoMDim = 3

// GeoMDim is a dimension of a GeoM.
// View Source

// const MaxIndicesNum = graphics.IndicesNum

// MaxIndicesNum is the maximum number of indices for DrawTriangles.
// View Source

// const UncappedTPS = clock.UncappedTPS

// UncappedTPS is a special TPS value that means the game doesn't have limitation on TPS.
// Variables ¶
// This section is empty.
// Functions ¶
// func CurrentFPS ¶

// func CurrentFPS() float64

// CurrentFPS returns the current number of FPS (frames per second), that represents how many swapping buffer happens per second.

// On some environments, CurrentFPS doesn't return a reliable value since vsync doesn't work well there. If you want to measure the application's speed, Use CurrentTPS.

// CurrentFPS is concurrent-safe.
// func CurrentTPS ¶

// func CurrentTPS() float64

// CurrentTPS returns the current TPS (ticks per second), that represents how many update function is called in a second.

// CurrentTPS is concurrent-safe.
// func CursorPosition ¶

// func CursorPosition() (x, y int)

// CursorPosition returns a position of a mouse cursor relative to the game screen (window). The cursor position is 'logical' position and this considers the scale of the screen.

// CursorPosition is concurrent-safe.
// func DeviceScaleFactor ¶

// func DeviceScaleFactor() float64

// DeviceScaleFactor returns a device scale factor value of the current monitor which the window belongs to.

// DeviceScaleFactor returns a meaningful value on high-DPI display environment, otherwise DeviceScaleFactor returns 1.

// DeviceScaleFactor might panic on init function on some devices like Android. Then, it is not recommended to call DeviceScaleFactor from init functions.

// DeviceScaleFactor must be called on the main thread before the main loop, and is concurrent-safe after the main loop.
// func GamepadAxis ¶

// func GamepadAxis(id GamepadID, axis int) float64

// GamepadAxis returns the float value [-1.0 - 1.0] of the given gamepad (id)'s axis (axis).

// GamepadAxis is concurrent-safe.

// GamepadAxis always returns 0 on mobiles.
// func GamepadAxisNum ¶

// func GamepadAxisNum(id GamepadID) int

// GamepadAxisNum returns the number of axes of the gamepad (id).

// GamepadAxisNum is concurrent-safe.

// GamepadAxisNum always returns 0 on mobiles.
// func GamepadButtonNum ¶

// func GamepadButtonNum(id GamepadID) int

// GamepadButtonNum returns the number of the buttons of the given gamepad (id).

// GamepadButtonNum is concurrent-safe.

// GamepadButtonNum always returns 0 on mobiles.
// func GamepadName ¶

// func GamepadName(id GamepadID) string

// GamepadName returns a string with the name. This function may vary in how it returns descriptions for the same device across platforms. for example the following drivers/platforms see a Xbox One controller as the following:

// - Windows: "Xbox Controller"
// - Chrome: "Xbox 360 Controller (XInput STANDARD GAMEPAD)"
// - Firefox: "xinput"

// GamepadName always returns an empty string on mobiles.

// GamepadName is concurrent-safe.
// func GamepadSDLID ¶

// func GamepadSDLID(id GamepadID) string

// GamepadSDLID returns a string with the GUID generated in the same way as SDL. To detect devices, see also the community project of gamepad devices database: https://github.com/gabomdq/SDL_GameControllerDB

// GamepadSDLID always returns an empty string on browsers and mobiles.

// GamepadSDLID is concurrent-safe.
// func InputChars ¶

// func InputChars() []rune

// InputChars return "printable" runes read from the keyboard at the time update is called.

// InputChars represents the environment's locale-dependent translation of keyboard input to Unicode characters.

// IsKeyPressed is based on a mapping of device (US keyboard) codes to input device keys. "Control" and modifier keys should be handled with IsKeyPressed.

// InputChars is concurrent-safe.

// On Android (ebitenmobile), EbitenView must be focusable to enable to handle keyboard keys.

// Keyboards don't work on iOS yet (#1090).
// func IsFocused ¶

// func IsFocused() bool

// IsFocused returns a boolean value indicating whether the game is in focus or in the foreground.

// IsFocused will only return true if IsRunnableOnUnfocused is false.

// IsFocused is concurrent-safe.
// func IsFullscreen ¶

// func IsFullscreen() bool

// IsFullscreen reports whether the current mode is fullscreen or not.

// IsFullscreen always returns false on browsers or mobiles.

// IsFullscreen is concurrent-safe.
// func IsGamepadButtonPressed ¶

// func IsGamepadButtonPressed(id GamepadID, button GamepadButton) bool

// IsGamepadButtonPressed returns the boolean indicating the given button of the gamepad (id) is pressed or not.

// If you want to know whether the given button of gamepad (id) started being pressed in the current frame, use inpututil.IsGamepadButtonJustPressed

// IsGamepadButtonPressed is concurrent-safe.

// The relationships between physical buttons and buttion IDs depend on environments. There can be differences even between Chrome and Firefox.

// IsGamepadButtonPressed always returns false on mobiles.
// func IsKeyPressed ¶

// func IsKeyPressed(key Key) bool

// IsKeyPressed returns a boolean indicating whether key is pressed.

// If you want to know whether the key started being pressed in the current frame, use inpututil.IsKeyJustPressed

// Known issue: On Edge browser, some keys don't work well:

// - KeyKPEnter and KeyKPEqual are recognized as KeyEnter and KeyEqual.
// - KeyPrintScreen is only treated at keyup event.

// IsKeyPressed is concurrent-safe.

// On Android (ebitenmobile), EbitenView must be focusable to enable to handle keyboard keys.

// Keyboards don't work on iOS yet (#1090).
// func IsMouseButtonPressed ¶

// func IsMouseButtonPressed(mouseButton MouseButton) bool

// IsMouseButtonPressed returns a boolean indicating whether mouseButton is pressed.

// If you want to know whether the mouseButton started being pressed in the current frame, use inpututil.IsMouseButtonJustPressed

// IsMouseButtonPressed is concurrent-safe.
// func IsRunnableOnUnfocused ¶

// func IsRunnableOnUnfocused() bool

// IsRunnableOnUnfocused returns a boolean value indicating whether the game runs even in background.

// IsRunnableOnUnfocused is concurrent-safe.
// func IsScreenClearedEveryFrame ¶

// func IsScreenClearedEveryFrame() bool

// IsScreenClearedEveryFrame returns true if the frame isn't cleared at the beginning.

// IsScreenClearedEveryFrame is concurrent-safe.
// func IsScreenTransparent ¶

// func IsScreenTransparent() bool

// IsScreenTransparent reports whether the window is transparent.

// IsScreenTransparent is concurrent-safe.
// func IsVsyncEnabled ¶

// func IsVsyncEnabled() bool

// IsVsyncEnabled returns a boolean value indicating whether the game uses the display's vsync.

// IsVsyncEnabled is concurrent-safe.
// func IsWindowDecorated ¶

// func IsWindowDecorated() bool

// IsWindowDecorated reports whether the window is decorated.

// IsWindowDecorated is concurrent-safe.
// func IsWindowFloating ¶

// func IsWindowFloating() bool

// IsWindowFloating reports whether the window is always shown above all the other windows.

// IsWindowFloating returns false on browsers and mobiles.

// IsWindowFloating is concurrent-safe.
// func IsWindowMaximized ¶

// func IsWindowMaximized() bool

// IsWindowMaximized reports whether the window is maximized or not.

// IsWindowMaximized returns false when the window is not resizable.

// IsWindowMaximized always returns false on browsers and mobiles.

// IsWindowMaximized is concurrent-safe.
// func IsWindowMinimized ¶

// func IsWindowMinimized() bool

// IsWindowMinimized reports whether the window is minimized or not.

// IsWindowMinimized always returns false on browsers and mobiles.

// IsWindowMinimized is concurrent-safe.
// func IsWindowResizable ¶

// func IsWindowResizable() bool

// IsWindowResizable reports whether the window is resizable by the user's dragging on desktops. On the other environments, IsWindowResizable always returns false.

// IsWindowResizable is concurrent-safe.
// func MaxTPS ¶

// func MaxTPS() int

// MaxTPS returns the current maximum TPS.

// MaxTPS is concurrent-safe.
// func MaximizeWindow ¶

// func MaximizeWindow()

// MaximizeWindow maximizes the window.

// MaximizeWindow panics when the window is not resizable.

// MaximizeWindow does nothing on browsers or mobiles.

// MaximizeWindow is concurrent-safe.
// func MinimizeWindow ¶

// func MinimizeWindow()

// MinimizeWindow minimizes the window.

// If the main loop does not start yet, MinimizeWindow does nothing.

// MinimizeWindow does nothing on browsers or mobiles.

// MinimizeWindow is concurrent-safe.
// func RestoreWindow ¶

// func RestoreWindow()

// RestoreWindow restores the window from its maximized or minimized state.

// RestoreWindow panics when the window is not maximized nor minimized.

// RestoreWindow is concurrent-safe.
// func RunGame ¶

// func RunGame(game Game) error

// RunGame starts the main loop and runs the game. game's Update function is called every tick to update the game logic. game's Draw function is, if it exists, called every frame to draw the screen. game's Layout function is called when necessary, and you can specify the logical screen size by the function.

// game must implement Game interface. Game's Draw function is optional, but it is recommended to implement Draw to seperate updating the logic and rendering.

// RunGame is a more flexibile form of Run due to game's Layout function. You can make a resizable window if you use RunGame, while you cannot if you use Run. RunGame is more sophisticated way than Run and hides the notion of 'scale'.

// While Run specifies the window size, RunGame does not. You need to call SetWindowSize before RunGame if you want. Otherwise, a default window size is adopted.

// Some functions (ScreenScale, SetScreenScale, SetScreenSize) are not available with RunGame.

// On browsers, it is strongly recommended to use iframe if you embed an Ebiten application in your website.

// RunGame must be called on the main thread. Note that Ebiten bounds the main goroutine to the main OS thread by runtime.LockOSThread.

// Ebiten tries to call game's Update function 60 times a second by default. In other words, TPS (ticks per second) is 60 by default. This is not related to framerate (display's refresh rate).

// RunGame returns error when 1) OpenGL error happens, 2) audio error happens or 3) f returns error. In the case of 3), RunGame returns the same error.

// The size unit is device-independent pixel.

// Don't call RunGame twice or more in one process.
// func RunGameWithoutMainLoop ¶

// func RunGameWithoutMainLoop(game Game)

// RunGameWithoutMainLoop runs the game, but don't call the loop on the main (UI) thread. Different from Run, RunGameWithoutMainLoop returns immediately.

// Ebiten users should NOT call RunGameWithoutMainLoop. Instead, functions in github.com/hajimehoshi/ebiten/v2/mobile package calls this.
// func ScreenSizeInFullscreen ¶

// func ScreenSizeInFullscreen() (int, int)

// ScreenSizeInFullscreen returns the size in device-independent pixels when the game is fullscreen. The adopted monitor is the 'current' monitor which the window belongs to. The returned value can be given to Run or SetSize function if the perfectly fit fullscreen is needed.

// On browsers, ScreenSizeInFullscreen returns the 'window' (global object) size, not 'screen' size since an Ebiten game should not know the outside of the window object. For more details, see SetFullscreen API comment.

// On mobiles, ScreenSizeInFullscreen returns (0, 0) so far.

// ScreenSizeInFullscreen's use cases are limited. If you are making a fullscreen application, you can use RunGame and the Game interface's Layout function instead. If you are making a not-fullscreen application but the application's behavior depends on the monitor size, ScreenSizeInFullscreen is useful.

// ScreenSizeInFullscreen must be called on the main thread before ebiten.Run, and is concurrent-safe after ebiten.Run.
// func SetCursorMode ¶

// func SetCursorMode(mode CursorModeType)

// SetCursorMode sets the render and capture mode of the mouse cursor. CursorModeVisible sets the cursor to always be visible. CursorModeHidden hides the system cursor when over the window. CursorModeCaptured hides the system cursor and locks it to the window.

// On browsers, only CursorModeVisible and CursorModeHidden are supported.

// SetCursorMode does nothing on mobiles.

// SetCursorMode is concurrent-safe.
// func SetFullscreen ¶

// func SetFullscreen(fullscreen bool)

// SetFullscreen changes the current mode to fullscreen or not on desktops.

// On fullscreen mode, the game screen is automatically enlarged to fit with the monitor. The current scale value is ignored.

// On desktops, Ebiten uses 'windowed' fullscreen mode, which doesn't change your monitor's resolution.

// SetFullscreen does nothing on browsers or mobiles.

// SetFullscreen is concurrent-safe.
// func SetInitFocused ¶

// func SetInitFocused(focused bool)

// SetInitFocused sets whether the application is focused on show. The default value is true, i.e., the application is focused. Note that the application does not proceed if this is not focused by default. This behavior can be changed by SetRunnableInBackground.

// SetInitFocused does nothing on mobile.

// SetInitFocused panics if this is called after the main loop.

// SetInitFocused is cuncurrent-safe.
// func SetMaxTPS ¶

// func SetMaxTPS(tps int)

// SetMaxTPS sets the maximum TPS (ticks per second), that represents how many updating function is called per second. The initial value is 60.

// If tps is UncappedTPS, TPS is uncapped and the game is updated per frame. If tps is negative but not UncappedTPS, SetMaxTPS panics.

// SetMaxTPS is concurrent-safe.
// func SetRunnableOnUnfocused ¶

// func SetRunnableOnUnfocused(runnableOnUnfocused bool)

// SetRunnableOnUnfocused sets the state if the game runs even in background.

// If the given value is true, the game runs even in background e.g. when losing focus. The initial state is true.

// Known issue: On browsers, even if the state is on, the game doesn't run in background tabs. This is because browsers throttles background tabs not to often update.

// SetRunnableOnUnfocused does nothing on mobiles so far.

// SetRunnableOnUnfocused is concurrent-safe.
// func SetScreenClearedEveryFrame ¶

// func SetScreenClearedEveryFrame(cleared bool)

// SetScreenClearedEveryFrame enables or disables the clearing of the screen at the beginning of each frame. The default value is false and the screen is cleared each frame by default.

// SetScreenClearedEveryFrame is concurrent-safe.
// func SetScreenTransparent ¶

// func SetScreenTransparent(transparent bool)

// SetScreenTransparent sets the state if the window is transparent.

// SetScreenTransparent panics if SetScreenTransparent is called after the main loop.

// SetScreenTransparent does nothing on mobiles.

// SetScreenTransparent is concurrent-safe.
// func SetVsyncEnabled ¶

// func SetVsyncEnabled(enabled bool)

// SetVsyncEnabled sets a boolean value indicating whether the game uses the display's vsync.

// If the given value is true, the game tries to sync the display's refresh rate. If false, the game ignores the display's refresh rate. The initial value is true. By disabling vsync, the game works more efficiently but consumes more CPU.

// Note that the state doesn't affect TPS (ticks per second, i.e. how many the run function is updated per second).

// SetVsyncEnabled does nothing on mobiles so far.

// SetVsyncEnabled is concurrent-safe.
// func SetWindowDecorated ¶

// func SetWindowDecorated(decorated bool)

// SetWindowDecorated sets the state if the window is decorated.

// The window is decorated by default.

// SetWindowDecorated works only on desktops. SetWindowDecorated does nothing on other platforms.

// SetWindowDecorated is concurrent-safe.
// func SetWindowFloating ¶

// func SetWindowFloating(float bool)

// SetWindowFloating sets the state whether the window is always shown above all the other windows.

// SetWindowFloating does nothing on browsers or mobiles.

// SetWindowFloating is concurrent-safe.
// func SetWindowIcon ¶

// func SetWindowIcon(iconImages []image.Image)

// SetWindowIcon sets the icon of the game window.

// If len(iconImages) is 0, SetWindowIcon reverts the icon to the default one.

// For desktops, see the document of glfwSetWindowIcon of GLFW 3.2:

// This function sets the icon of the specified window.
// If passed an array of candidate images, those of or closest to the sizes
// desired by the system are selected.
// If no images are specified, the window reverts to its default icon.

// The desired image sizes varies depending on platform and system settings.
// The selected images will be rescaled as needed.
// Good sizes include 16x16, 32x32 and 48x48.

// As macOS windows don't have icons, SetWindowIcon doesn't work on macOS.

// SetWindowIcon doesn't work on browsers or mobiles.

// SetWindowIcon is concurrent-safe.
// func SetWindowPosition ¶

// func SetWindowPosition(x, y int)

// SetWindowPosition sets the window position. The origin position is the left-upper corner of the current monitor. The unit is device-independent pixels.

// SetWindowPosition does nothing on fullscreen mode.

// SetWindowPosition does nothing on browsers and mobiles.

// SetWindowPosition is concurrent-safe.
// func SetWindowResizable ¶

// func SetWindowResizable(resizable bool)

// SetWindowResizable sets whether the window is resizable by the user's dragging on desktops. On the other environments, SetWindowResizable does nothing.

// The window is not resizable by default.

// If SetWindowResizable is called with true and Run is used, SetWindowResizable panics. Use RunGame instead.

// SetWindowResizable is concurrent-safe.
// func SetWindowSize ¶

// func SetWindowSize(width, height int)

// SetWindowSize sets the window size on desktops. SetWindowSize does nothing on other environments.

// On fullscreen mode, SetWindowSize sets the original window size.

// SetWindowSize panics if width or height is not a positive number.

// SetWindowSize is concurrent-safe.
// func SetWindowTitle ¶

// func SetWindowTitle(title string)

// SetWindowTitle sets the title of the window.

// SetWindowTitle does nothing on browsers or mobiles.

// SetWindowTitle is concurrent-safe.
// func TouchPosition ¶

// func TouchPosition(id TouchID) (int, int)

// TouchPosition returns the position for the touch of the specified ID.

// If the touch of the specified ID is not present, TouchPosition returns (0, 0).

// TouchPosition is cuncurrent-safe.
// func Wheel ¶

// func Wheel() (xoff, yoff float64)

// Wheel returns the x and y offset of the mouse wheel or touchpad scroll. It returns 0 if the wheel isn't being rolled.

// Wheel is concurrent-safe.
// func WindowPosition ¶

// func WindowPosition() (x, y int)

// WindowPosition returns the window position. The origin position is the left-upper corner of the current monitor. The unit is device-independent pixels.

// WindowPosition panics if the main loop does not start yet.

// WindowPosition returns the last window position on fullscreen mode.

// WindowPosition returns (0, 0) on browsers and mobiles.

// WindowPosition is concurrent-safe.
// func WindowSize ¶

// func WindowSize() (int, int)

// WindowSize returns the window size on desktops. WindowSize returns (0, 0) on other environments.

// On fullscreen mode, WindowSize returns the original window size.

// WindowSize is concurrent-safe.
// Types ¶
// type Address ¶

// type Address int

// Address represents a sampler address mode.

// const (
// 	// AddressUnsafe means there is no guarantee when the texture coodinates are out of range.
// 	AddressUnsafe Address = Address(driver.AddressUnsafe)

// 	// AddressClampToZero means that out-of-range texture coordinates return 0 (transparent).
// 	AddressClampToZero Address = Address(driver.AddressClampToZero)

// 	// AddressRepeat means that texture coordinates wrap to the other side of the texture.
// 	AddressRepeat Address = Address(driver.AddressRepeat)
// )

// type ColorM ¶

// type ColorM struct {
// 	// contains filtered or unexported fields
// }

// A ColorM represents a matrix to transform coloring when rendering an image.

// A ColorM is applied to the straight alpha color while an Image's pixels' format is alpha premultiplied. Before applying a matrix, a color is un-multiplied, and after applying the matrix, the color is multiplied again.

// The initial value is identity.
// func (*ColorM) Apply ¶

// func (c *ColorM) Apply(clr color.Color) color.Color

// Apply pre-multiplies a vector (r, g, b, a, 1) by the matrix where r, g, b, and a are clr's values in straight-alpha format. In other words, Apply calculates ColorM * (r, g, b, a, 1)^T.
// func (*ColorM) ChangeHSV ¶

// func (c *ColorM) ChangeHSV(hueTheta float64, saturationScale float64, valueScale float64)

// ChangeHSV changes HSV (Hue-Saturation-Value) values. hueTheta is a radian value to rotate hue. saturationScale is a value to scale saturation. valueScale is a value to scale value (a.k.a. brightness).

// This conversion uses RGB to/from YCrCb conversion.
// func (*ColorM) Concat ¶

// func (c *ColorM) Concat(other ColorM)

// Concat multiplies a color matrix with the other color matrix. This is same as muptiplying the matrix other and the matrix c in this order.
// func (*ColorM) Element ¶

// func (c *ColorM) Element(i, j int) float64

// Element returns a value of a matrix at (i, j).
// func (*ColorM) Invert ¶

// func (c *ColorM) Invert()

// Invert inverts the matrix. If c is not invertible, Invert panics.
// func (*ColorM) IsInvertible ¶

// func (c *ColorM) IsInvertible() bool

// IsInvertible returns a boolean value indicating whether the matrix c is invertible or not.
// func (*ColorM) Reset ¶

// func (c *ColorM) Reset()

// Reset resets the ColorM as identity.
// func (*ColorM) RotateHue ¶

// func (c *ColorM) RotateHue(theta float64)

// RotateHue rotates the hue. theta represents rotating angle in radian.
// func (*ColorM) Scale ¶

// func (c *ColorM) Scale(r, g, b, a float64)

// Scale scales the matrix by (r, g, b, a).
// func (*ColorM) SetElement ¶

// func (c *ColorM) SetElement(i, j int, element float64)

// SetElement sets an element at (i, j).
// func (*ColorM) String ¶

// func (c *ColorM) String() string

// String returns a string representation of ColorM.
// func (*ColorM) Translate ¶

// func (c *ColorM) Translate(r, g, b, a float64)

// Translate translates the matrix by (r, g, b, a).
// type CompositeMode ¶

// type CompositeMode int

// CompositeMode represents Porter-Duff composition mode.

// const (
// 	// Regular alpha blending
// 	// c_out = c_src + c_dst × (1 - α_src)
// 	CompositeModeSourceOver CompositeMode = CompositeMode(driver.CompositeModeSourceOver)

// 	// c_out = 0
// 	CompositeModeClear CompositeMode = CompositeMode(driver.CompositeModeClear)

// 	// c_out = c_src
// 	CompositeModeCopy CompositeMode = CompositeMode(driver.CompositeModeCopy)

// 	// c_out = c_dst
// 	CompositeModeDestination CompositeMode = CompositeMode(driver.CompositeModeDestination)

// 	// c_out = c_src × (1 - α_dst) + c_dst
// 	CompositeModeDestinationOver CompositeMode = CompositeMode(driver.CompositeModeDestinationOver)

// 	// c_out = c_src × α_dst
// 	CompositeModeSourceIn CompositeMode = CompositeMode(driver.CompositeModeSourceIn)

// 	// c_out = c_dst × α_src
// 	CompositeModeDestinationIn CompositeMode = CompositeMode(driver.CompositeModeDestinationIn)

// 	// c_out = c_src × (1 - α_dst)
// 	CompositeModeSourceOut CompositeMode = CompositeMode(driver.CompositeModeSourceOut)

// 	// c_out = c_dst × (1 - α_src)
// 	CompositeModeDestinationOut CompositeMode = CompositeMode(driver.CompositeModeDestinationOut)

// 	// c_out = c_src × α_dst + c_dst × (1 - α_src)
// 	CompositeModeSourceAtop CompositeMode = CompositeMode(driver.CompositeModeSourceAtop)

// 	// c_out = c_src × (1 - α_dst) + c_dst × α_src
// 	CompositeModeDestinationAtop CompositeMode = CompositeMode(driver.CompositeModeDestinationAtop)

// 	// c_out = c_src × (1 - α_dst) + c_dst × (1 - α_src)
// 	CompositeModeXor CompositeMode = CompositeMode(driver.CompositeModeXor)

// 	// Sum of source and destination (a.k.a. 'plus' or 'additive')
// 	// c_out = c_src + c_dst
// 	CompositeModeLighter CompositeMode = CompositeMode(driver.CompositeModeLighter)

// 	// The product of source and destination (a.k.a 'multiply blend mode')
// 	// c_out = c_src * c_dst
// 	CompositeModeMultiply CompositeMode = CompositeMode(driver.CompositeModeMultiply)
// )

// This name convention follows CSS compositing: https://drafts.fxtf.org/compositing-2/.

// In the comments, c_src, c_dst and c_out represent alpha-premultiplied RGB values of source, destination and output respectively. α_src and α_dst represent alpha values of source and destination respectively.
// type CursorModeType ¶

// type CursorModeType int

// CursorModeType represents a render and coordinate mode of a mouse cursor.

// const (
// 	CursorModeVisible  CursorModeType = CursorModeType(driver.CursorModeVisible)
// 	CursorModeHidden   CursorModeType = CursorModeType(driver.CursorModeHidden)
// 	CursorModeCaptured CursorModeType = CursorModeType(driver.CursorModeCaptured)
// )

// func CursorMode ¶

// func CursorMode() CursorModeType

// CursorMode returns the current cursor mode.

// On browsers, only CursorModeVisible and CursorModeHidden are supported.

// CursorMode returns CursorModeHidden on mobiles.

// CursorMode is concurrent-safe.
// type DrawImageOptions ¶

// type DrawImageOptions struct {
// 	// GeoM is a geometry matrix to draw.
// 	// The default (zero) value is identity, which draws the image at (0, 0).
// 	GeoM GeoM

// 	// ColorM is a color matrix to draw.
// 	// The default (zero) value is identity, which doesn't change any color.
// 	ColorM ColorM

// 	// CompositeMode is a composite mode to draw.
// 	// The default (zero) value is regular alpha blending.
// 	CompositeMode CompositeMode

// 	// Filter is a type of texture filter.
// 	// The default (zero) value is FilterNearest.
// 	Filter Filter
// }

// DrawImageOptions represents options for DrawImage.
// type DrawRectShaderOptions ¶

// type DrawRectShaderOptions struct {
// 	// GeoM is a geometry matrix to draw.
// 	// The default (zero) value is identity, which draws the rectangle at (0, 0).
// 	GeoM GeoM

// 	// CompositeMode is a composite mode to draw.
// 	// The default (zero) value is regular alpha blending.
// 	CompositeMode CompositeMode

// 	// Uniforms is a set of uniform variables for the shader.
// 	// The keys are the names of the uniform variables.
// 	// The values must be float or []float.
// 	// If the uniform variable type is an array, a vector or a matrix,
// 	// you have to specify linearly flattened values as a slice.
// 	// For example, if the uniform variable type is [4]vec4, the number of the slice values will be 16.
// 	Uniforms map[string]interface{}

// 	// Images is a set of the source images.
// 	// All the image must be the same size with the rectangle.
// 	Images [4]*Image
// }

// DrawRectShaderOptions represents options for DrawRectShader.

// This API is experimental.
// type DrawTrianglesOptions ¶

// type DrawTrianglesOptions struct {
// 	// ColorM is a color matrix to draw.
// 	// The default (zero) value is identity, which doesn't change any color.
// 	// ColorM is applied before vertex color scale is applied.
// 	//
// 	// If Shader is not nil, ColorM is ignored.
// 	ColorM ColorM

// 	// CompositeMode is a composite mode to draw.
// 	// The default (zero) value is regular alpha blending.
// 	CompositeMode CompositeMode

// 	// Filter is a type of texture filter.
// 	// The default (zero) value is FilterNearest.
// 	Filter Filter

// 	// Address is a sampler address mode.
// 	// The default (zero) value is AddressUnsafe.
// 	Address Address
// }

// DrawTrianglesOptions represents options for DrawTriangles.
// type DrawTrianglesShaderOptions ¶

// type DrawTrianglesShaderOptions struct {
// 	// CompositeMode is a composite mode to draw.
// 	// The default (zero) value is regular alpha blending.
// 	CompositeMode CompositeMode

// 	// Uniforms is a set of uniform variables for the shader.
// 	// The keys are the names of the uniform variables.
// 	// The values must be float or []float.
// 	// If the uniform variable type is an array, a vector or a matrix,
// 	// you have to specify linearly flattened values as a slice.
// 	// For example, if the uniform variable type is [4]vec4, the number of the slice values will be 16.
// 	Uniforms map[string]interface{}

// 	// Images is a set of the source images.
// 	// All the image must be the same size.
// 	Images [4]*Image
// }

// DrawTrianglesShaderOptions represents options for DrawTrianglesShader.

// This API is experimental.
// type Filter ¶

// type Filter int

// Filter represents the type of texture filter to be used when an image is maginified or minified.

// const (
// 	// FilterNearest represents nearest (crisp-edged) filter
// 	FilterNearest Filter = Filter(driver.FilterNearest)

// 	// FilterLinear represents linear filter
// 	FilterLinear Filter = Filter(driver.FilterLinear)
// )

// type Game ¶

// type Game interface {
// 	// Update updates a game by one tick. The given argument represents a screen image.
// 	//
// 	// Update updates only the game logic and Draw draws the screen.
// 	//
// 	// In the first frame, it is ensured that Update is called at least once before Draw. You can use Update
// 	// to initialize the game state.
// 	//
// 	// After the first frame, Update might not be called or might be called once
// 	// or more for one frame. The frequency is determined by the current TPS (tick-per-second).
// 	Update() error

// 	// Draw draws the game screen by one frame.
// 	//
// 	// The give argument represents a screen image. The updated content is adopted as the game screen.
// 	Draw(screen *Image)

// 	// Layout accepts a native outside size in device-independent pixels and returns the game's logical screen
// 	// size.
// 	//
// 	// On desktops, the outside is a window or a monitor (fullscreen mode). On browsers, the outside is a body
// 	// element. On mobiles, the outside is the view's size.
// 	//
// 	// Even though the outside size and the screen size differ, the rendering scale is automatically adjusted to
// 	// fit with the outside.
// 	//
// 	// Layout is called almost every frame.
// 	//
// 	// If Layout returns non-positive numbers, the caller can panic.
// 	//
// 	// You can return a fixed screen size if you don't care, or you can also return a calculated screen size
// 	// adjusted with the given outside size.
// 	Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int)
// }

// Game defines necessary functions for a game.
// type GamepadButton ¶

// type GamepadButton int

// A GamepadButton represents a gamepad button.

// const (
// 	GamepadButton0   GamepadButton = GamepadButton(driver.GamepadButton0)
// 	GamepadButton1   GamepadButton = GamepadButton(driver.GamepadButton1)
// 	GamepadButton2   GamepadButton = GamepadButton(driver.GamepadButton2)
// 	GamepadButton3   GamepadButton = GamepadButton(driver.GamepadButton3)
// 	GamepadButton4   GamepadButton = GamepadButton(driver.GamepadButton4)
// 	GamepadButton5   GamepadButton = GamepadButton(driver.GamepadButton5)
// 	GamepadButton6   GamepadButton = GamepadButton(driver.GamepadButton6)
// 	GamepadButton7   GamepadButton = GamepadButton(driver.GamepadButton7)
// 	GamepadButton8   GamepadButton = GamepadButton(driver.GamepadButton8)
// 	GamepadButton9   GamepadButton = GamepadButton(driver.GamepadButton9)
// 	GamepadButton10  GamepadButton = GamepadButton(driver.GamepadButton10)
// 	GamepadButton11  GamepadButton = GamepadButton(driver.GamepadButton11)
// 	GamepadButton12  GamepadButton = GamepadButton(driver.GamepadButton12)
// 	GamepadButton13  GamepadButton = GamepadButton(driver.GamepadButton13)
// 	GamepadButton14  GamepadButton = GamepadButton(driver.GamepadButton14)
// 	GamepadButton15  GamepadButton = GamepadButton(driver.GamepadButton15)
// 	GamepadButton16  GamepadButton = GamepadButton(driver.GamepadButton16)
// 	GamepadButton17  GamepadButton = GamepadButton(driver.GamepadButton17)
// 	GamepadButton18  GamepadButton = GamepadButton(driver.GamepadButton18)
// 	GamepadButton19  GamepadButton = GamepadButton(driver.GamepadButton19)
// 	GamepadButton20  GamepadButton = GamepadButton(driver.GamepadButton20)
// 	GamepadButton21  GamepadButton = GamepadButton(driver.GamepadButton21)
// 	GamepadButton22  GamepadButton = GamepadButton(driver.GamepadButton22)
// 	GamepadButton23  GamepadButton = GamepadButton(driver.GamepadButton23)
// 	GamepadButton24  GamepadButton = GamepadButton(driver.GamepadButton24)
// 	GamepadButton25  GamepadButton = GamepadButton(driver.GamepadButton25)
// 	GamepadButton26  GamepadButton = GamepadButton(driver.GamepadButton26)
// 	GamepadButton27  GamepadButton = GamepadButton(driver.GamepadButton27)
// 	GamepadButton28  GamepadButton = GamepadButton(driver.GamepadButton28)
// 	GamepadButton29  GamepadButton = GamepadButton(driver.GamepadButton29)
// 	GamepadButton30  GamepadButton = GamepadButton(driver.GamepadButton30)
// 	GamepadButton31  GamepadButton = GamepadButton(driver.GamepadButton31)
// 	GamepadButtonMax GamepadButton = GamepadButton31
// )

// GamepadButtons
// type GamepadID ¶

// type GamepadID = driver.GamepadID

// GamepadID represents a gamepad's identifier.
// func GamepadIDs ¶

// func GamepadIDs() []GamepadID

// GamepadIDs returns a slice indicating available gamepad IDs.

// GamepadIDs is concurrent-safe.

// GamepadIDs always returns an empty slice on mobiles.
// type GeoM ¶

// type GeoM struct {
// 	// contains filtered or unexported fields
// }

// A GeoM represents a matrix to transform geometry when rendering an image.

// The initial value is identity.
// func (*GeoM) Apply ¶

// func (g *GeoM) Apply(x, y float64) (float64, float64)

// Apply pre-multiplies a vector (x, y, 1) by the matrix. In other words, Apply calculates GeoM * (x, y, 1)^T. The return value is x and y values of the result vector.
// func (*GeoM) Concat ¶

// func (g *GeoM) Concat(other GeoM)

// Concat multiplies a geometry matrix with the other geometry matrix. This is same as muptiplying the matrix other and the matrix g in this order.
// func (*GeoM) Element ¶

// func (g *GeoM) Element(i, j int) float64

// Element returns a value of a matrix at (i, j).
// func (*GeoM) Invert ¶

// func (g *GeoM) Invert()

// Invert inverts the matrix. If g is not invertible, Invert panics.
// func (*GeoM) IsInvertible ¶

// func (g *GeoM) IsInvertible() bool

// IsInvertible returns a boolean value indicating whether the matrix g is invertible or not.
// func (*GeoM) Reset ¶

// func (g *GeoM) Reset()

// Reset resets the GeoM as identity.
// func (*GeoM) Rotate ¶

// func (g *GeoM) Rotate(theta float64)

// Rotate rotates the matrix by theta. The unit is radian.
// func (*GeoM) Scale ¶

// func (g *GeoM) Scale(x, y float64)

// Scale scales the matrix by (x, y).
// func (*GeoM) SetElement ¶

// func (g *GeoM) SetElement(i, j int, element float64)

// SetElement sets an element at (i, j).
// func (*GeoM) Skew ¶

// func (g *GeoM) Skew(skewX, skewY float64)

// Skew skews the matrix by (skewX, skewY). The unit is radian.
// func (*GeoM) String ¶

// func (g *GeoM) String() string

// String returns a string representation of GeoM.
// func (*GeoM) Translate ¶

// func (g *GeoM) Translate(tx, ty float64)

// Translate translates the matrix by (tx, ty).
// type Image ¶

// type Image struct {
// 	// contains filtered or unexported fields
// }

// Image represents a rectangle set of pixels. The pixel format is alpha-premultiplied RGBA. Image implements image.Image and draw.Image.
// func NewImage ¶

// func NewImage(width, height int) *Image

// NewImage returns an empty image.

// If width or height is less than 1 or more than device-dependent maximum size, NewImage panics.
// func NewImageFromImage ¶

// func NewImageFromImage(source image.Image) *Image

// NewImageFromImage creates a new image with the given image (source).

// If source's width or height is less than 1 or more than device-dependent maximum size, NewImageFromImage panics.
// func (*Image) At ¶

// func (i *Image) At(x, y int) color.Color

// At returns the color of the image at (x, y).

// At loads pixels from GPU to system memory if necessary, which means that At can be slow.

// At always returns a transparent color if the image is disposed.

// Note that an important logic should not rely on values returned by At, since the returned values can include very slight differences between some machines.

// At can't be called outside the main loop (ebiten.Run's updating function) starts.
// func (*Image) Bounds ¶

// func (i *Image) Bounds() image.Rectangle

// Bounds returns the bounds of the image.
// func (*Image) Clear ¶

// func (i *Image) Clear()

// Clear resets the pixels of the image into 0.

// When the image is disposed, Clear does nothing.
// func (*Image) ColorModel ¶

// func (i *Image) ColorModel() color.Model

// ColorModel returns the color model of the image.
// func (*Image) Dispose ¶

// func (i *Image) Dispose()

// Dispose disposes the image data. After disposing, most of image functions do nothing and returns meaningless values.

// Calling Dispose is not mandatory. GC automatically collects internal resources that no objects refer to. However, calling Dispose explicitly is helpful if memory usage matters.

// When the image is disposed, Dipose does nothing.
// func (*Image) DrawImage ¶

// func (i *Image) DrawImage(img *Image, options *DrawImageOptions)

// DrawImage draws the given image on the image i.

// DrawImage accepts the options. For details, see the document of DrawImageOptions.

// For drawing, the pixels of the argument image at the time of this call is adopted. Even if the argument image is mutated after this call, the drawing result is never affected.

// When the image i is disposed, DrawImage does nothing. When the given image img is disposed, DrawImage panics.

// When the given image is as same as i, DrawImage panics.

// DrawImage works more efficiently as batches when the successive calls of DrawImages satisfy the below conditions:

// * All render targets are same (A in A.DrawImage(B, op))
// * Either all ColorM element values are same or all the ColorM have only
//    diagonal ('scale') elements
//   * If only (*ColorM).Scale is applied to a ColorM, the ColorM has only
//     diagonal elements. The other ColorM functions might modify the other
//     elements.
// * All CompositeMode values are same
// * All Filter values are same

// Even when all the above conditions are satisfied, multiple draw commands can be used in really rare cases. Ebiten images usually share an internal automatic texture atlas, but when you consume the atlas, or you create a huge image, those images cannot be on the same texture atlas. In this case, draw commands are separated. The texture atlas size is 4096x4096 so far. Another case is when you use an offscreen as a render source. An offscreen doesn't share the texture atlas with high probability.

// For more performance tips, see https://ebiten.org/documents/performancetips.html
// func (*Image) DrawRectShader ¶

// func (i *Image) DrawRectShader(width, height int, shader *Shader, options *DrawRectShaderOptions)

// DrawRectShader draws a rectangle with the specified width and height with the specified shader.

// For the details about the shader, see https://ebiten.org/documents/shader.html.

// When one of the specified image is non-nil and is disposed, DrawRectShader panics.

// When the image i is disposed, DrawRectShader does nothing.

// This API is experimental.
// func (*Image) DrawTriangles ¶

// func (i *Image) DrawTriangles(vertices []Vertex, indices []uint16, img *Image, options *DrawTrianglesOptions)

// DrawTriangles draws triangles with the specified vertices and their indices.

// If len(indices) is not multiple of 3, DrawTriangles panics.

// If len(indices) is more than MaxIndicesNum, DrawTriangles panics.

// The rule in which DrawTriangles works effectively is same as DrawImage's.

// When the given image is disposed, DrawTriangles panics.

// When the image i is disposed, DrawTriangles does nothing.
// func (*Image) DrawTrianglesShader ¶

// func (i *Image) DrawTrianglesShader(vertices []Vertex, indices []uint16, shader *Shader, options *DrawTrianglesShaderOptions)

// DrawTrianglesShader draws triangles with the specified vertices and their indices with the specified shader.

// For the details about the shader, see https://ebiten.org/documents/shader.html.

// If len(indices) is not multiple of 3, DrawTrianglesShader panics.

// If len(indices) is more than MaxIndicesNum, DrawTrianglesShader panics.

// When a specified image is non-nil and is disposed, DrawTrianglesShader panics.

// When the image i is disposed, DrawTrianglesShader does nothing.

// This API is experimental.
// func (*Image) Fill ¶

// func (i *Image) Fill(clr color.Color)

// Fill fills the image with a solid color.

// When the image is disposed, Fill does nothing.
// func (*Image) ReplacePixels ¶

// func (i *Image) ReplacePixels(pixels []byte)

// ReplacePixels replaces the pixels of the image with p.

// The given p must represent RGBA pre-multiplied alpha values. len(pix) must equal to 4 * (bounds width) * (bounds height).

// ReplacePixels works on a sub-image.

// When len(pix) is not appropriate, ReplacePixels panics.

// When the image is disposed, ReplacePixels does nothing.
// func (*Image) Set ¶

// func (i *Image) Set(x, y int, clr color.Color)

// Set sets the color at (x, y).

// Set loads pixels from GPU to system memory if necessary, which means that Set can be slow.

// In the current implementation, successive calls of Set invokes loading pixels at most once, so this is efficient.

// If the image is disposed, Set does nothing.
// func (*Image) Size ¶

// func (i *Image) Size() (width, height int)

// Size returns the size of the image.
// func (*Image) SubImage ¶

// func (i *Image) SubImage(r image.Rectangle) image.Image

// SubImage returns an image representing the portion of the image p visible through r. The returned value shares pixels with the original image.

// The returned value is always *ebiten.Image.

// If the image is disposed, SubImage returns nil.

// In the current Ebiten implementation, SubImage is available only as a rendering source.
// type Key ¶

// type Key int

// A Key represents a keyboard key. These keys represent pysical keys of US keyboard. For example, KeyQ represents Q key on US keyboards and ' (quote) key on Dvorak keyboards.

// const (
// 	Key0            Key = Key(driver.Key0)
// 	Key1            Key = Key(driver.Key1)
// 	Key2            Key = Key(driver.Key2)
// 	Key3            Key = Key(driver.Key3)
// 	Key4            Key = Key(driver.Key4)
// 	Key5            Key = Key(driver.Key5)
// 	Key6            Key = Key(driver.Key6)
// 	Key7            Key = Key(driver.Key7)
// 	Key8            Key = Key(driver.Key8)
// 	Key9            Key = Key(driver.Key9)
// 	KeyA            Key = Key(driver.KeyA)
// 	KeyB            Key = Key(driver.KeyB)
// 	KeyC            Key = Key(driver.KeyC)
// 	KeyD            Key = Key(driver.KeyD)
// 	KeyE            Key = Key(driver.KeyE)
// 	KeyF            Key = Key(driver.KeyF)
// 	KeyG            Key = Key(driver.KeyG)
// 	KeyH            Key = Key(driver.KeyH)
// 	KeyI            Key = Key(driver.KeyI)
// 	KeyJ            Key = Key(driver.KeyJ)
// 	KeyK            Key = Key(driver.KeyK)
// 	KeyL            Key = Key(driver.KeyL)
// 	KeyM            Key = Key(driver.KeyM)
// 	KeyN            Key = Key(driver.KeyN)
// 	KeyO            Key = Key(driver.KeyO)
// 	KeyP            Key = Key(driver.KeyP)
// 	KeyQ            Key = Key(driver.KeyQ)
// 	KeyR            Key = Key(driver.KeyR)
// 	KeyS            Key = Key(driver.KeyS)
// 	KeyT            Key = Key(driver.KeyT)
// 	KeyU            Key = Key(driver.KeyU)
// 	KeyV            Key = Key(driver.KeyV)
// 	KeyW            Key = Key(driver.KeyW)
// 	KeyX            Key = Key(driver.KeyX)
// 	KeyY            Key = Key(driver.KeyY)
// 	KeyZ            Key = Key(driver.KeyZ)
// 	KeyApostrophe   Key = Key(driver.KeyApostrophe)
// 	KeyBackslash    Key = Key(driver.KeyBackslash)
// 	KeyBackspace    Key = Key(driver.KeyBackspace)
// 	KeyCapsLock     Key = Key(driver.KeyCapsLock)
// 	KeyComma        Key = Key(driver.KeyComma)
// 	KeyDelete       Key = Key(driver.KeyDelete)
// 	KeyDown         Key = Key(driver.KeyDown)
// 	KeyEnd          Key = Key(driver.KeyEnd)
// 	KeyEnter        Key = Key(driver.KeyEnter)
// 	KeyEqual        Key = Key(driver.KeyEqual)
// 	KeyEscape       Key = Key(driver.KeyEscape)
// 	KeyF1           Key = Key(driver.KeyF1)
// 	KeyF2           Key = Key(driver.KeyF2)
// 	KeyF3           Key = Key(driver.KeyF3)
// 	KeyF4           Key = Key(driver.KeyF4)
// 	KeyF5           Key = Key(driver.KeyF5)
// 	KeyF6           Key = Key(driver.KeyF6)
// 	KeyF7           Key = Key(driver.KeyF7)
// 	KeyF8           Key = Key(driver.KeyF8)
// 	KeyF9           Key = Key(driver.KeyF9)
// 	KeyF10          Key = Key(driver.KeyF10)
// 	KeyF11          Key = Key(driver.KeyF11)
// 	KeyF12          Key = Key(driver.KeyF12)
// 	KeyGraveAccent  Key = Key(driver.KeyGraveAccent)
// 	KeyHome         Key = Key(driver.KeyHome)
// 	KeyInsert       Key = Key(driver.KeyInsert)
// 	KeyKP0          Key = Key(driver.KeyKP0)
// 	KeyKP1          Key = Key(driver.KeyKP1)
// 	KeyKP2          Key = Key(driver.KeyKP2)
// 	KeyKP3          Key = Key(driver.KeyKP3)
// 	KeyKP4          Key = Key(driver.KeyKP4)
// 	KeyKP5          Key = Key(driver.KeyKP5)
// 	KeyKP6          Key = Key(driver.KeyKP6)
// 	KeyKP7          Key = Key(driver.KeyKP7)
// 	KeyKP8          Key = Key(driver.KeyKP8)
// 	KeyKP9          Key = Key(driver.KeyKP9)
// 	KeyKPAdd        Key = Key(driver.KeyKPAdd)
// 	KeyKPDecimal    Key = Key(driver.KeyKPDecimal)
// 	KeyKPDivide     Key = Key(driver.KeyKPDivide)
// 	KeyKPEnter      Key = Key(driver.KeyKPEnter)
// 	KeyKPEqual      Key = Key(driver.KeyKPEqual)
// 	KeyKPMultiply   Key = Key(driver.KeyKPMultiply)
// 	KeyKPSubtract   Key = Key(driver.KeyKPSubtract)
// 	KeyLeft         Key = Key(driver.KeyLeft)
// 	KeyLeftBracket  Key = Key(driver.KeyLeftBracket)
// 	KeyMenu         Key = Key(driver.KeyMenu)
// 	KeyMinus        Key = Key(driver.KeyMinus)
// 	KeyNumLock      Key = Key(driver.KeyNumLock)
// 	KeyPageDown     Key = Key(driver.KeyPageDown)
// 	KeyPageUp       Key = Key(driver.KeyPageUp)
// 	KeyPause        Key = Key(driver.KeyPause)
// 	KeyPeriod       Key = Key(driver.KeyPeriod)
// 	KeyPrintScreen  Key = Key(driver.KeyPrintScreen)
// 	KeyRight        Key = Key(driver.KeyRight)
// 	KeyRightBracket Key = Key(driver.KeyRightBracket)
// 	KeyScrollLock   Key = Key(driver.KeyScrollLock)
// 	KeySemicolon    Key = Key(driver.KeySemicolon)
// 	KeySlash        Key = Key(driver.KeySlash)
// 	KeySpace        Key = Key(driver.KeySpace)
// 	KeyTab          Key = Key(driver.KeyTab)
// 	KeyUp           Key = Key(driver.KeyUp)
// 	KeyAlt          Key = Key(driver.KeyReserved0)
// 	KeyControl      Key = Key(driver.KeyReserved1)
// 	KeyShift        Key = Key(driver.KeyReserved2)
// 	KeyMax          Key = KeyShift
// )

// Keys.
// func (Key) String ¶

// func (k Key) String() string

// String returns a string representing the key.

// If k is an undefined key, String returns an empty string.
// type MouseButton ¶

// type MouseButton int

// A MouseButton represents a mouse button.

// const (
// 	MouseButtonLeft   MouseButton = MouseButton(driver.MouseButtonLeft)
// 	MouseButtonRight  MouseButton = MouseButton(driver.MouseButtonRight)
// 	MouseButtonMiddle MouseButton = MouseButton(driver.MouseButtonMiddle)
// )

// MouseButtons
// type Shader ¶

// type Shader struct {
// 	// contains filtered or unexported fields
// }

// Shader represents a compiled shader program.

// For the details about the shader, see https://ebiten.org/documents/shader.html.
// func NewShader ¶

// func NewShader(src []byte) (*Shader, error)

// NewShader compiles a shader program in the shading language Kage, and retruns the result.

// If the compilation fails, NewShader returns an error.

// For the details about the shader, see https://ebiten.org/documents/shader.html.
// func (*Shader) Dispose ¶

// func (s *Shader) Dispose()

// Dispose disposes the shader program. After disposing, the shader is no longer available.
// type TouchID ¶

// type TouchID = driver.TouchID

// TouchID represents a touch's identifier.
// func TouchIDs ¶

// func TouchIDs() []TouchID

// TouchIDs returns the current touch states.

// If you want to know whether a touch started being pressed in the current frame, use inpututil.JustPressedTouchIDs

// TouchIDs returns nil when there are no touches. TouchIDs always returns nil on desktops.

// TouchIDs is concurrent-safe.
// type Vertex ¶

// type Vertex struct {
// 	// DstX and DstY represents a point on a destination image.
// 	DstX float32
// 	DstY float32

// 	// SrcX and SrcY represents a point on a source image.
// 	// Be careful that SrcX/SrcY coordinates are on the image's bounds.
// 	// This means that a left-upper point of a sub-image might not be (0, 0).
// 	SrcX float32
// 	SrcY float32

// 	// ColorR/ColorG/ColorB/ColorA represents color scaling values.
// 	// 1 means the original source image color is used.
// 	// 0 means a transparent color is used.
// 	ColorR float32
// 	ColorG float32
// 	ColorB float32
// 	ColorA float32
// }

// Vertex represents a vertex passed to DrawTriangles.

// +build !wasm

package audio

import (
	"github.com/g3n/engine/audio"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["audio.direction"] = &AudioDirection{}
}

type AudioDirection struct {
	pc1 *PlayerCone
	pc2 *PlayerCone
	pc3 *PlayerCone
	pc4 *PlayerCone
	pc5 *PlayerCone
	pc6 *PlayerCone
}

// Start is called once at the start of the demo.
func (t *AudioDirection) Start(a *app.App) {

	// Create axes helper
	axes := helper.NewAxes(1)
	a.Scene().Add(axes)

	// Show grid helper
	grid := helper.NewGrid(100, 1, &math32.Color{0.4, 0.4, 0.4})
	a.Scene().Add(grid)

	// Add directional white light
	l1 := light.NewDirectional(&math32.Color{1, 1, 1}, 1.0)
	l1.SetPosition(10, 10, 10)
	a.Scene().Add(l1)

	// Sets camera position
	a.Camera().SetPosition(0, 0, 10)
	a.Camera().UpdateSize(10)
	a.Camera().LookAt(&math32.Vector3{0, 0, 0}, &math32.Vector3{0, 1, 0})

	// Creates listener and adds it to the current camera
	// The listener must have the same initial direction as the camera
	listener := audio.NewListener()
	cam := a.Camera()
	cdir := cam.Direction()
	listener.SetDirectionVec(&cdir)
	cam.Add(listener)

	// Creates player cones
	t.pc1 = NewPlayerCone(a, "Vivaldi1.wav", &math32.Color{1, 0, 0})
	t.pc1.SetPosition(0, 0, 3)
	t.pc1.SetDirection(&math32.Vector3{0, 0, 1})
	t.pc1.player.SetLooping(true)
	t.pc1.player.Play()
	a.Scene().Add(t.pc1)

	t.pc2 = NewPlayerCone(a, "Bach1.ogg", &math32.Color{0, 1, 0})
	t.pc2.SetPosition(3, 0, 0)
	t.pc2.SetDirection(&math32.Vector3{1, 0, 0})
	t.pc2.player.SetLooping(true)
	t.pc2.player.Play()
	a.Scene().Add(t.pc2)

	t.pc3 = NewPlayerCone(a, "engine.ogg", &math32.Color{0, 0, 1})
	t.pc3.SetPosition(0, 0, -3)
	t.pc3.SetDirection(&math32.Vector3{0, 0, -1})
	t.pc3.player.SetLooping(true)
	t.pc3.player.Play()
	a.Scene().Add(t.pc3)

	t.pc4 = NewPlayerCone(a, "bomb2.ogg", &math32.Color{0, 1, 1})
	t.pc4.SetPosition(-3, 0, 0)
	t.pc4.SetDirection(&math32.Vector3{-1, 0, 0})
	t.pc4.player.SetLooping(true)
	t.pc4.player.Play()
	a.Scene().Add(t.pc4)

	t.pc5 = NewPlayerCone(a, "tone_440hz.wav", &math32.Color{1, 1, 1})
	t.pc5.SetPosition(0, 3, 0)
	t.pc5.SetDirection(&math32.Vector3{0, 1, 0})
	t.pc5.player.SetLooping(true)
	t.pc5.player.Play()
	a.Scene().Add(t.pc5)

	t.pc6 = NewPlayerCone(a, "tone_1khz.wav", &math32.Color{1, 0, 1})
	t.pc6.SetPosition(0, -3, 0)
	t.pc6.SetDirection(&math32.Vector3{0, -1, 0})
	t.pc6.player.SetLooping(true)
	t.pc6.player.Play()
	a.Scene().Add(t.pc6)

	// Add controls
	if a.ControlFolder() == nil {
		return
	}
	g1 := a.ControlFolder().AddGroup("Play sources")
	cb1 := g1.AddCheckBox("Vivaldi1").SetValue(true)
	cb1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc1.Toggle()
	})
	cb2 := g1.AddCheckBox("Bach1").SetValue(true)
	cb2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc2.Toggle()
	})
	cb3 := g1.AddCheckBox("engine").SetValue(true)
	cb3.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc3.Toggle()
	})
	cb4 := g1.AddCheckBox("bomb2").SetValue(true)
	cb4.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc4.Toggle()
	})
	cb5 := g1.AddCheckBox("tone_440hz").SetValue(true)
	cb5.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc5.Toggle()
	})
	cb6 := g1.AddCheckBox("tone_1khz").SetValue(true)
	cb6.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc6.Toggle()
	})

	g2 := a.ControlFolder().AddGroup("Sound Cone")
	g1s1 := g2.AddSlider("Outer Cone:", 360, t.pc1.player.OuterCone())
	g1s1.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc1.player.SetOuterCone(g1s1.Value())
		t.pc2.player.SetOuterCone(g1s1.Value())
		t.pc3.player.SetOuterCone(g1s1.Value())
		t.pc4.player.SetOuterCone(g1s1.Value())
		t.pc5.player.SetOuterCone(g1s1.Value())
		t.pc6.player.SetOuterCone(g1s1.Value())
	})
	g1s2 := g2.AddSlider("Inner Cone:", 360, t.pc1.player.InnerCone())
	g1s2.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		t.pc1.player.SetInnerCone(g1s2.Value())
		t.pc2.player.SetInnerCone(g1s2.Value())
		t.pc3.player.SetInnerCone(g1s2.Value())
		t.pc4.player.SetInnerCone(g1s2.Value())
		t.pc5.player.SetInnerCone(g1s2.Value())
		t.pc6.player.SetInnerCone(g1s2.Value())
	})
}

// Update is called every frame.
func (t *AudioDirection) Update(a *app.App, deltaTime time.Duration) {}

// Cleanup is called once at the end of the demo.
func (t *AudioDirection) Cleanup(a *app.App) {}

type PlayerCone struct {
	graphic.Mesh
	player    *audio.Player
	start     time.Time
	label     *graphic.Mesh
	direction math32.Vector3
}

func NewPlayerCone(app *app.App, filename string, color *math32.Color) *PlayerCone {

	pc := new(PlayerCone)

	// Creates audio source
	player, err := audio.NewPlayer(app.DirData() + "/audio/" + filename)
	if err != nil {
		app.Log().Fatal("error:%s", err)
	}

	player.SetRolloffFactor(1)
	pc.player = player

	// Creates cone geometry and orients it so the cone base is pointed
	// in the +Z direction
	geom := geometry.NewCone(1, 1, 32, 1, true)
	geom.ApplyMatrix(math32.NewMatrix4().MakeRotationX(-math32.Pi / 2))
	pc.direction.Set(0, 0, 1)

	// Creates materials and initialize mesh
	mat1 := material.NewStandard(color)
	mat2 := material.NewStandard(color)
	mat2.SetEmissiveColor(color)
	pc.Mesh.Init(geom, nil)
	pc.Mesh.AddGroupMaterial(mat1, 0)
	pc.Mesh.AddGroupMaterial(mat2, 1)

	// Adds plane with file name
	font := gui.StyleDefault().Font
	font.SetPointSize(48)
	font.SetColor(&math32.Color4{0, 0, 0, 1})
	width, height := font.MeasureText(filename)
	aspect := float32(width) / float32(height)
	img := font.DrawText(filename)
	tex := texture.NewTexture2DFromRGBA(img)
	plane_geom := geometry.NewPlane(2.0, 2.0/aspect)
	plane_mat := material.NewStandard(math32.NewColor("white"))
	plane_mat.AddTexture(tex)
	plane_mat.SetTransparent(true)
	pc.label = graphic.NewMesh(plane_geom, plane_mat)
	pc.label.SetPosition(0, 0, 0.51)
	pc.Add(pc.label)

	// Set up player and adds it to the cone mesh
	pc.player.SetLooping(true)
	pc.player.SetOuterCone(90)
	pc.player.SetInnerCone(0)
	pc.start = time.Now()
	pc.Add(pc.player)
	return pc
}

func (pc *PlayerCone) SetDirection(vdir *math32.Vector3) {

	newdir := *vdir
	newdir.Normalize()
	var q math32.Quaternion
	q.SetFromUnitVectors(&pc.direction, &newdir)
	pc.Mesh.SetQuaternionQuat(&q)
}

func (pc *PlayerCone) Direction() math32.Vector3 {

	var wdir math32.Vector3
	pc.Mesh.WorldDirection(&wdir)
	return wdir
}

func (pc *PlayerCone) Toggle() {

	if pc.Visible() {
		pc.player.Pause()
		pc.SetVisible(false)
	} else {
		pc.player.Play()
		pc.SetVisible(true)
	}
}

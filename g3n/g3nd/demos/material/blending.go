package material

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/texture"
	"github.com/g3n/g3nd/app"
	"time"
)

func init() {
	app.DemoMap["material.blending"] = &Blending{}
}

type Blending struct {
	texbg *texture.Texture2D
}

// Start is called once at the start of the demo.
func (t *Blending) Start(a *app.App) {

	a.Camera().SetPositionZ(600)
	a.Camera().UpdateSize(600)
	a.AmbLight().SetIntensity(2)

	// Creates checker board textures for background
	c1 := &math32.Color{0.7, 0.7, 0.7}
	c2 := &math32.Color{0.3, 0.3, 0.3}
	t.texbg = texture.NewBoard(16, 16, c1, c2, c2, c1, 1)
	t.texbg.SetWrapS(gls.REPEAT)
	t.texbg.SetWrapT(gls.REPEAT)
	t.texbg.SetRepeat(64, 64)

	// Creates background plane
	matbg := material.NewStandard(&math32.Color{1, 1, 1})
	matbg.SetPolygonOffset(1, 1)
	matbg.AddTexture(t.texbg)
	geombg := geometry.NewPlane(4000, 3000)
	meshbg := graphic.NewMesh(geombg, matbg)
	meshbg.SetPosition(0, 0, -1)
	a.Scene().Add(meshbg)

	// Builds list of textures
	texnames := []string{
		"uvgrid.jpg", "sprite0.jpg",
		"sprite0.png", "lensflare0.png",
		"lensflare0_alpha.png",
	}
	texlist := []*texture.Texture2D{}
	for _, tname := range texnames {
		tex, err := texture.NewTexture2DFromImage(a.DirData() + "/images/" + tname)
		if err != nil {
			a.Log().Fatal("Error loading texture: %s", err)
		}
		texlist = append(texlist, tex)
	}

	blendings := []struct {
		blending string
		value    material.Blending
	}{
		{"NoBlending", material.BlendNone},
		{"NormalBlending", material.BlendNormal},
		{"AdditiveBlending", material.BlendAdditive},
		{"SubtractiveBlending", material.BlendSubtractive},
		{"MultiplyBlending", material.BlendMultiply},
	}

	// This geometry will be shared by several meshes
	// For each mesh which uses this geometry we need to increment its refcount
	geo1 := geometry.NewPlane(100, 100)
	defer geo1.Dispose()

	// Internal function go generate a row of images
	var addImageRow = func(tex *texture.Texture2D, y int) {
		for i := 0; i < len(blendings); i++ {
			material := material.NewStandard(&math32.Color{1, 1, 1})
			material.SetOpacity(1)
			material.SetTransparent(true)
			material.AddTexture(tex)
			material.SetBlending(blendings[i].value)
			x := (float32(i) - float32(len(blendings))/2) * 110
			mesh := graphic.NewMesh(geo1.Incref(), material)
			mesh.SetPosition(x, float32(y), 0)
			a.Scene().Add(mesh)
		}
	}

	addImageRow(texlist[0], 300)
	addImageRow(texlist[1], 150)
	addImageRow(texlist[2], 0)
	addImageRow(texlist[3], -150)
	addImageRow(texlist[4], -300)
}

// Update is called every frame.
func (t *Blending) Update(a *app.App, deltaTime time.Duration) {

	tt := float32(a.RunTime().Seconds()) * 0.003
	rx, ry := t.texbg.Repeat()
	ox := math32.Mod(tt*rx, 1.0)
	oy := math32.Mod(tt*ry, 1.0)
	t.texbg.SetOffset(-ox, oy)
}

// Cleanup is called once at the end of the demo.
func (t *Blending) Cleanup(a *app.App) {}

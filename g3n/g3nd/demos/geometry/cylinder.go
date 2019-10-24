package geometry

import (
	"time"

	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/helper"
	"github.com/g3n/g3nd/app"
)

func init() {
	app.DemoMap["geometry.cone-cylinder"] = &ConeCylinder{}
}

type ConeCylinder struct {
	mesh    *graphic.Mesh
	normals *helper.Normals
}

// Start is called once at the start of the demo.
func (t *ConeCylinder) Start(a *app.App) {

	// Add directional red light from right
	l1 := light.NewDirectional(&math32.Color{1, 0, 0}, 1.0)
	l1.SetPosition(1, 0, 0)
	a.Scene().Add(l1)

	// Add directional green light from top
	l2 := light.NewDirectional(&math32.Color{0, 1, 0}, 1.0)
	l2.SetPosition(0, 1, 0)
	a.Scene().Add(l2)

	// Add directional blue light from front
	l3 := light.NewDirectional(&math32.Color{0, 0, 1}, 1.0)
	l3.SetPosition(0, 0, 1)
	a.Scene().Add(l3)

	// Left cylinder
	geom1 := geometry.NewCylinder(0.8, 2, 16, 2, true, true)
	mat1 := material.NewStandard(&math32.Color{0, 1, 0})
	mat1.SetWireframe(true)
	mat1.SetSide(material.SideDouble)
	mat1.SetUseLights(material.UseLightNone)
	t.mesh = graphic.NewMesh(geom1, mat1)
	t.mesh.SetPosition(-2, 0, 0)
	a.Scene().Add(t.mesh)

	// Middle cylinder
	geom2 := geometry.NewCylinder(0.8, 2, 32, 16, false, true)
	mat2 := material.NewStandard(&math32.Color{1, 1, 1})
	mat2.SetSide(material.SideDouble)
	mesh := graphic.NewMesh(geom2, mat2)
	mesh.SetPosition(0, 0, 0)
	a.Scene().Add(mesh)

	// Right cylinder
	geom3 := geometry.NewTruncatedCone(0.4, 0.8, 2, 32, 1, false, true)
	mat3 := material.NewStandard(&math32.Color{1, 1, 1})
	mat3.SetSide(material.SideDouble)
	mesh3 := graphic.NewMesh(geom3, mat3)
	mesh3.SetPosition(2, 0, 0)
	a.Scene().Add(mesh3)

	// Create axes helper
	axes := helper.NewAxes(2)
	a.Scene().Add(axes)

	// Adds normals helper
	t.normals = helper.NewNormals(t.mesh, 0.5, &math32.Color{0, 1, 0}, 1)
	a.Scene().Add(t.normals)
}

// Update is called every frame.
func (t *ConeCylinder) Update(a *app.App, deltaTime time.Duration) {

	t.mesh.RotateY(0.005)
	t.normals.Update()
}

// Cleanup is called once at the end of the demo.
func (t *ConeCylinder) Cleanup(a *app.App) {}

package camera

import "math"

type Camera struct {
	X, Y                      float64
	screenWidth, screenHeight float64
}

func NewCamera(x, y, screenwidth, screenheight float64) *Camera {
	return &Camera{
		X:            x,
		Y:            y,
		screenWidth:  screenwidth,
		screenHeight: screenheight,
	}
}

func (c *Camera) FollowTarget(targetX, targetY float64) {
	c.X = -targetX + c.screenWidth/2.0
	c.Y = -targetY + c.screenHeight/2.0
}

func (c *Camera) Constrain(tilemapWidthPixels, tilemapHeightPixels float64) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	c.X = math.Max(c.X, c.screenWidth-tilemapWidthPixels)
	c.Y = math.Max(c.Y, c.screenHeight-tilemapHeightPixels)
}

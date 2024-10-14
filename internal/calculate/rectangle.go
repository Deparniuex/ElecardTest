package calculate

import (
	"ElecardTest/internal/entity"
)

func CalcRectangles(tasks *entity.Tasks) []entity.Rectangle {
	var minX, maxX, minY, maxY float64
	var rectangles []entity.Rectangle
	for _, task := range tasks.Result {
		minX, maxX, minY, maxY = task[0].X-task[0].R, task[0].X+task[0].R, task[0].Y-task[0].R, task[0].Y+task[0].R
		for _, circle := range task {
			if circle.X-circle.R < minX {
				minX = circle.X - circle.R
			}
			if circle.X+circle.R > maxX {
				maxX = circle.X + circle.R
			}
			if circle.Y-circle.R < minY {
				minY = circle.Y - circle.R
			}
			if circle.Y+circle.R > maxY {
				maxY = circle.Y + circle.R
			}
		}
		rectangles = append(rectangles, entity.Rectangle{
			LeftBottom: entity.Coordinate{
				X: minX,
				Y: minY,
			},
			RightTop: entity.Coordinate{
				X: maxX,
				Y: maxY,
			},
		})
	}
	return rectangles
}

package models

type Card struct {
	Id            int
	Name          string
	Svg           string
	TopologyHoles int
	TopologyParts int
}

var deck []Card
var field []Card

func init() {
	deck = []Card{
		{
			Id:   1,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   2,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   3,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   4,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   5,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   6,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   7,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   8,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   9,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
		{
			Id:   10,
			Name: "ao",
			TopologyHoles: 2,
			TopologyParts: 2,
			Svg: `
<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="50pt" height="50pt" viewBox="0 0 50 50" version="1.1">
<g id="surface1">
<path style=" stroke:none;fill-rule:nonzero;fill:rgb(0%,0%,0%);fill-opacity:1;" d="M 22.984375 7.386719 L 22.984375 2.429688 L 26.644531 2.429688 L 26.644531 7.386719 L 44.734375 7.386719 L 44.734375 10.242188 L 26.644531 10.242188 L 26.644531 13.074219 L 42.488281 13.074219 L 42.488281 15.835938 L 26.644531 15.835938 L 26.644531 18.886719 L 47.25 18.886719 L 47.25 21.792969 L 2.695312 21.792969 L 2.695312 18.886719 L 22.984375 18.886719 L 22.984375 15.835938 L 7.429688 15.835938 L 7.429688 13.074219 L 22.984375 13.074219 L 22.984375 10.242188 L 5.183594 10.242188 L 5.183594 7.386719 Z M 39.949219 24.527344 L 39.949219 43.226562 C 39.949219 44.496094 39.683594 45.398438 39.144531 45.9375 C 38.476562 46.605469 37.160156 46.9375 35.191406 46.9375 C 32.960938 46.9375 30.625 46.839844 28.183594 46.644531 L 27.597656 43.03125 C 30.152344 43.472656 32.429688 43.691406 34.433594 43.691406 C 35.734375 43.691406 36.386719 43.054688 36.386719 41.785156 L 36.386719 39.078125 L 13.828125 39.078125 L 13.828125 47.425781 L 10.289062 47.425781 L 10.289062 24.527344 Z M 13.828125 27.382812 L 13.828125 30.386719 L 36.386719 30.386719 L 36.386719 27.382812 Z M 13.828125 33.144531 L 13.828125 36.316406 L 36.386719 36.316406 L 36.386719 33.144531 Z M 13.828125 33.144531 "/>
</g>
</svg>
			`,
		},
	}
}

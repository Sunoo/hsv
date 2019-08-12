// image/color.Color implementation for HSV representation
package hsv

// HSVColor defines a color in the Hue-Saturation-Value scheme.
// Hue is a value [0 - 360] specifying the color
// Saturation is a value [0 - 100] specifying the strength of the color
// Value is a value [0 - 100] specifying the brightness of the color (currently does nothing)
type HSVColor struct {
	H, S float64
	V int
}

func (h HSVColor) RGBA() (r, g, b, a uint32) {
	// Direct implementation of the graph in this image:
	// https://en.wikipedia.org/wiki/HSL_and_HSV#/media/File:HSV-RGB-comparison.svg
	var max uint32 = 255
	var min uint32 = uint32(255 - (h.S / 100 * 255))

	var segment uint32 = uint32(h.H / 60)
	var offset uint32 = uint32(h.H) % 60
	var mid uint32 = ((max - min) * offset) / 60

	switch segment {
		case 0:
			return max, min + mid, min, 255
		case 1:
			return max - mid, max, min, 255
		case 2:
			return min, max, min + mid, 255
		case 3:
			return min, max - mid, max, 255
		case 4:
			return min + mid, min, max, 255
		case 5:
			return max, min, max - mid, 255
	}

	return 0, 0, 0, 255
}

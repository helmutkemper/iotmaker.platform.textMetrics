package iotmaker_platform_textMetrics

type TextMetrics struct {
	// en: Is a double giving the calculated width of a segment of inline text in CSS
	// pixels. It takes into account the current font of the context.
	Width float64

	// en: Is a double giving the distance parallel to the baseline from the alignment
	// point given by the CanvasRenderingContext2D.textAlign property to the left side
	// of the bounding rectangle of the given text, in CSS pixels.
	ActualBoundingBoxLeft float64

	// en: Is a double giving the distance parallel to the baseline from the alignment
	// point given by the CanvasRenderingContext2D.textAlign property to the right side
	// of the bounding rectangle of the given text, in CSS pixels.
	ActualBoundingBoxRight float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline attribute to the top of the highest
	// bounding rectangle of all the fonts used to render the text, in CSS pixels.
	FontBoundingBoxAscent float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline attribute to the bottom of the bounding
	// rectangle of all the fonts used to render the text, in CSS pixels.
	FontBoundingBoxDescent float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline attribute to the top of the bounding
	// rectangle used to render the text, in CSS pixels.
	ActualBoundingBoxAscent float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline attribute to the bottom of the bounding
	// rectangle used to render the text, in CSS pixels.
	ActualBoundingBoxDescent float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline property to the top of the em square in
	// the line box, in CSS pixels.
	EmHeightAscent float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline property to the bottom of the em square in
	// the line box, in CSS pixels.
	EmHeightDescent float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline property to the hanging baseline of the
	// line box, in CSS pixels.
	HangingBaseline float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline property to the alphabetic baseline of the
	// line box, in CSS pixels.
	AlphabeticBaseline float64

	// en: Is a double giving the distance from the horizontal line indicated by the
	// CanvasRenderingContext2D.textBaseline property to the ideographic baseline of
	// the line box, in CSS pixels.
	IdeographicBaseline float64
}

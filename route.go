package route

// Point is route point
type Point struct {
	Routes map[*Point]*Route
	Name   string
}

// Route has left/right ways to point
type Route struct {
	Left  *Point
	Right *Point
}

// NewPoint returns new point structure
func NewPoint(name string) *Point {
	return &Point{Name: name, Routes: make(map[*Point]*Route)}
}

// NewRoute returns Question's Route structure
func NewRoute() (*Point, *Point) {
	a := NewPoint("A")
	b := NewPoint("B")
	c := NewPoint("C")
	d := NewPoint("D")
	e := NewPoint("E")
	f := NewPoint("F")

	a.Routes[d] = &Route{Left: c, Right: b}
	a.Routes[b] = &Route{Left: d, Right: c}
	a.Routes[c] = &Route{Left: b, Right: d}

	b.Routes[a] = &Route{Left: c, Right: e}
	b.Routes[c] = &Route{Left: e, Right: a}
	b.Routes[e] = &Route{Left: a, Right: c}

	c.Routes[a] = &Route{Left: f, Right: b}
	c.Routes[b] = &Route{Left: a, Right: f}
	c.Routes[f] = &Route{Left: b, Right: a}

	d.Routes[a] = &Route{Left: e, Right: f}
	d.Routes[e] = &Route{Left: f, Right: a}
	d.Routes[f] = &Route{Left: a, Right: e}

	e.Routes[b] = &Route{Left: f, Right: d}
	e.Routes[d] = &Route{Left: b, Right: f}
	e.Routes[f] = &Route{Left: d, Right: b}

	f.Routes[c] = &Route{Left: d, Right: e}
	f.Routes[d] = &Route{Left: e, Right: c}
	f.Routes[e] = &Route{Left: c, Right: d}

	return b, a
}

// Stranger walks Point structured route.
type Stranger struct {
	Current  *Point
	Previous *Point
}

// NewStranger returns new stranger struct
func NewStranger() *Stranger {
	previous, current := NewRoute()
	return &Stranger{Current: current, Previous: previous}
}

// Walk is direction that stranger walks
func (s *Stranger) Walk(to string) string {
	if to != "l" && to != "r" && to != "b" {
		panic("invalid input")
	}

	if to == "l" {
		current := s.Current
		s.Current = s.Current.Routes[s.Previous].Left
		s.Previous = current
	} else if to == "r" {
		current := s.Current
		s.Current = s.Current.Routes[s.Previous].Right
		s.Previous = current
	} else {
		current := s.Current
		s.Current = s.Previous
		s.Previous = current
	}

	return s.Current.Name
}

func (s *Stranger) solve(input string) string {
	result := s.Current.Name
	for _, r := range input {
		result = result + s.Walk(string(r))
	}
	return result
}

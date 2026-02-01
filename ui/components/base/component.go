package base

type Component struct{
	Width int
}

func New(width int) *Component {
	return &Component{Width: width}
}

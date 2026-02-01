package endpoints

import (
	"reflect"
	"sort"
	"stat_by_sites/domain"
	"stat_by_sites/ui/components/base"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type Table struct {
	base.Component
	Data       []Endpoint
}

func NewTable(width int) *Table{
	return &Table{
		base.Component{Width: width},
		make([]Endpoint, 0),
	}
}

func (t *Table) Update(data []domain.Endpoint) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].URL < data[j].URL
	})
	
	ep := EndpointPresenter{}
	t.Data = ep.Present(data)
}

func (e Table) buildHeaders() string {
	t := reflect.TypeOf(Endpoint{})
	colNum := t.NumField()
	headerCells := make([]string, colNum)

	for i := range colNum{
		field := t.Field(i)
		fieldName := field.Tag.Get("header")
		fieldWidth, _ := strconv.Atoi(field.Tag.Get("width"))
		headerCells[i] = cellStyle.Width(fieldWidth).Render(headerStyle.Render(fieldName))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, headerCells...)
}

func (e Table) buildRow(endpoint Endpoint) string {
	values := endpoint.formateValues()
	cells := make([]string, len(values))
	t := reflect.TypeOf(endpoint)
	colNum := t.NumField()

	for i := range colNum {
		field := t.Field(i)
		fieldWidth, _ := strconv.Atoi(field.Tag.Get("width"))
		cells[i] = cellStyle.Width(fieldWidth).Render(values[i])
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, cells...)
}
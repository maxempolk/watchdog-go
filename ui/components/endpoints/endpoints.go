package endpoints

import (
	"reflect"
	"slices"
	"stat_by_sites/domain/endpoint"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")).Underline(true)
	cellStyle   = lipgloss.NewStyle()
)

type Endpoint struct {
	URL      	string `header:"ENDPOINT"   width:"30"`
	Status    int    `header:"STATUS"     width:"12"`
	Latency   string `header:"LATENCY"    width:"12"`
	LastCheck string `header:"LAST CHECK" width:"15"`
	Trend     []bool `header:"TREND"      width:"11"`
}

func (e Endpoint) generateTrendRepresentation() string{
	// TODO: нормально ли разворачивать тут
	reversed := make([]bool, endpoint.TrendSize)
	copy(reversed, e.Trend)
	slices.Reverse(reversed)

	if len(reversed) < 4 {
		return "○ ○ ○ ○ "
	}
	
	var trend strings.Builder
	trend.Grow(8)
	
	for i := 0; i < 4; i++ {
		if reversed[i] {
			trend.WriteString("●")
		} else {
			trend.WriteString("○")
		}
		trend.WriteString(" ")
	}
	
	return trend.String()
}

func (e Endpoint) generateStatusRepresentation() string{
	return "[ " + strconv.Itoa(e.Status) + " ]"
}

func (e Endpoint) generateEndpointNameRepresentation() string{
	t := reflect.TypeOf(e)
	field, _ := t.FieldByName("URL")
	width, _ := strconv.Atoi( field.Tag.Get("width") )

	formatedName := e.URL;
	if(len(e.URL) > width && width-4 > 0){
		formatedName = e.URL[:width-4] + "..."
	}

	return formatedName
}

func (e Endpoint) formateValues() []string{
	return []string{
		e.generateEndpointNameRepresentation(),
		e.generateStatusRepresentation(),
		e.Latency,
		e.LastCheck,
		e.generateTrendRepresentation(),
	}
}
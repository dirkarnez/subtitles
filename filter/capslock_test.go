package filter

import (
	"fmt"
	"testing"

	"github.com/martinlindhe/go-subber/caption"
	"github.com/martinlindhe/go-subber/testExtras"
	"github.com/stretchr/testify/assert"
)

func TestCreateMovieHashFromMovieFile(t *testing.T) {

	var in []caption.Caption
	in = append(in, caption.Caption{
		Seq:   1,
		Text:  []string{"GO NINJA!", "NINJA GO!"},
		Start: testExtras.MakeTime(0, 0, 4, 630),
		End:   testExtras.MakeTime(0, 0, 6, 18)})

	var expected []caption.Caption
	expected = append(expected, caption.Caption{
		Seq:   1,
		Text:  []string{"Go ninja!", "Ninja go!"},
		Start: testExtras.MakeTime(0, 0, 4, 630),
		End:   testExtras.MakeTime(0, 0, 6, 18)})

	fmt.Println(CapslockStripper(in))

	assert.Equal(t, expected, CapslockStripper(in))
}

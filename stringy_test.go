package listy_test

import (
	"testing"

	"github.com/xdg/listy"
	"github.com/xdg/testy"
)

func TestSJoin(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	cases := []struct {
		Label  string
		Input  []string
		Sep    string
		Output string
	}{
		{Label: "empty list", Input: []string{}, Sep: " ", Output: ""},
		{Label: "one item", Input: []string{"one"}, Sep: " ", Output: "one"},
		{Label: "two items", Input: []string{"one", "two"}, Sep: " ", Output: "one two"},
		{Label: "three items", Input: []string{"one", "two", "three"}, Sep: " ", Output: "one two three"},
	}

	for _, v := range cases {
		xs := listy.S(v.Input)
		got := xs.Join(v.Sep)
		is.Label(v.Label).Equal(got, v.Output)
	}

}

func TestBSJoin(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	cases := []struct {
		Label  string
		Input  []string
		Sep    string
		Output string
	}{
		{Label: "empty list", Input: []string{}, Sep: " ", Output: ""},
		{Label: "one item", Input: []string{"one"}, Sep: " ", Output: "one"},
		{Label: "two items", Input: []string{"one", "two"}, Sep: " ", Output: "one two"},
		{Label: "three items", Input: []string{"one", "two", "three"}, Sep: " ", Output: "one two three"},
	}

	for _, v := range cases {
		bs := make([][]byte, 0)
		for _, v := range v.Input {
			bs = append(bs, []byte(v))
		}
		xs := listy.BS(bs)
		got := xs.Join([]byte(v.Sep))
		is.Label(v.Label).Equal(got, []byte(v.Output))
	}

}

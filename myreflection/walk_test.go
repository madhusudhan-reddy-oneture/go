package myreflection

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	testCases := []struct {
		desc          string
		input         interface{}
		expectedCalls []string
	}{
		{
			desc: "struct with one string field",
			input: struct {
				Name string
			}{"Chandler"},
			expectedCalls: []string{"Chandler"},
		},
		{
			desc: "struct with two string fields",
			input: struct {
				Name string
				City string
			}{"Trump", "New York City"},
			expectedCalls: []string{"Trump", "New York City"},
		},
		{
			desc: "struct with non string field",
			input: struct {
				Name string
				Age  int
			}{"Trump", 79},
			expectedCalls: []string{"Trump"},
		},
		{
			desc: "nested fields",
			input: Person{
				"Trump",
				Profile{Age: 79, City: "New York City"},
			},
			expectedCalls: []string{"Trump", "New York City"},
		},
		{
			desc: "pointers to things",
			input: &Person{
				"Trump",
				Profile{Age: 79, City: "New York City"},
			},
			expectedCalls: []string{"Trump", "New York City"},
		},
		{
			desc: "slices",
			input: []Profile{
				{33, "London"},
				{28, "Mumbai"},
			},
			expectedCalls: []string{"London", "Mumbai"},
		},
		{
			desc: "arrays",
			input: [2]Profile{
				{33, "London"},
				{28, "Mumbai"},
			},
			expectedCalls: []string{"London", "Mumbai"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var got []string
			walk(tC.input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, tC.expectedCalls) {
				t.Errorf("got %v, want %v", got, tC.expectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Batman":    "Gotham",
			"SpiderMan": "New York City",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Gotham")
		assertContains(t, got, "New York City")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{28, "Lisbon"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Lisbon"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Queens"}, Profile{26, "Bronx"}
		}

		var got []string
		want := []string{"Queens", "Bronx"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, neddle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == neddle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, neddle)
	}
}

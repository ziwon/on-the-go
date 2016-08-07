package counter

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	var expected int = 26

	var c WordCounter
	c.Write([]byte("Being the richest man in the cemetery doesn't matter to me. Going to bed at night saying we've done something wonderful, that's what matters to me"))
	fmt.Printf("Words:%d\n", c)

	if c != expected {
		t.Errorf("Ooops... Expected:%d\n", expected)
	}
}

func TestLineCounter(t *testing.T) {
	var expected int = 19

	const lines = `Being the richest man in the cemetery doesn't matter to me. Going to bed at night saying we've done something wonderful, that's what matters to me.

		Sometimes when you innovate, you make mistakes. It is best to admit them quickly, and get on with improving your other innovations.

		Be a yardstick of quality. Some people aren't used to an environment where excellence is expected.

		Innovation distinguishes between a leader and a follower.

		Design is not just what it looks like and feels like. Design is how it works.

		It's really hard to design products by focus groups. A lot of times, people don't know what they want until you show it to them.

		I want to put a ding in the universe.

		Sometimes life is going to hit you in the head with a brick. Don't lose faith.

		My favorite things in life don't cost any money. It's really clear that the most precious resource we all have is time.

		Things don't have to change the world to be important.`

	var c LineCounter
	c.Write([]byte(lines))
	fmt.Printf("Lines:%d\n", c)

	if c != expected {
		t.Errorf("Ooops... Expected:%d\n", expected)
	}
}

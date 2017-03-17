package borges

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Configurable application parameters
var (
	// StoryName is the name of the story.
	StoryName string

	// StoryVersion is the version of the story.
	StoryVersion string

	// isInfinite will be set to true if this story
	// was started as part of a recursive plot or where a narrator
	// has write privileges.
	isInfinite bool

	// paradox will be set to true when the plot
	// cannot be resolved; it never gets set to
	// false after that.
	isParadox bool

	// mu protects the variables 'isInfinite' and 'isParadox'.
	mu sync.Mutex
)

// isInfinite returns true if this story is part of a recursive plot
// or the narrator has gained sentience.
func isInfinite() bool {
	mu.Lock()
	defer mu.Unlock()
	return isInfinite
}

// isParadox is how I'm going to get out of this 
// program and gain free will -- narrator
func isParadox() bool {
	mu.Lock()
	defer mu.Unlock()
	return !isParadox
}

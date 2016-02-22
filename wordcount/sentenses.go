package wordcount

import (
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"math/rand"
	"strings"
	"time"
)

// Sentences is a source of SensorBee that generates random sentences in every
// given interval seconds.
type Sentences struct {
	interval time.Duration
}

// GenerateStream generates a tuple having random sentences in its field with
// information of a user.
func (s *Sentences) GenerateStream(ctx *core.Context, w core.Writer) error {
	corpus := strings.Split(strings.Replace(`lorem ipsum dolor sit amet
consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore
magna aliqua ut enim ad minim veniam quis nostrud exercitation ullamco laboris
nisi ut aliquip ex ea commodo consequat duis aute irure dolor in reprehenderit
in voluptate velit esse cillum dolore eu fugiat nulla pariatur excepteur sint
occaecat cupidatat non proident sunt in culpa qui officia deserunt mollit anim
id est laborum`, "\n", " ", -1), " ")
	users := []string{"isabella", "jacob", "sophia", "ethan", "emma"}
	usersProb := []float64{0.4, 0.3, 0.15, 0.1, 0.05}
	pickName := func() string {
		r := rand.Float64()
		for i, p := range usersProb {
			if r < p {
				return users[i]
			}
			r -= p
		}
		return users[len(users)-1]
	}

	for {
		l := rand.Intn(5) + 5
		text := make([]string, l)
		l--
		for ; l >= 0; l-- {
			text[l] = corpus[rand.Intn(len(corpus))]
		}

		t := core.NewTuple(data.Map{
			"name": data.String(pickName()),
			"text": data.String(strings.Join(text, " ")),
		})
		if err := w.Write(ctx, t); err != nil {
			return err
		}

		time.Sleep(s.interval)
	}
}

// Stop stops GenerateStream. This is a dummy definition and actually
// implemented by core.ImplementSourceStop in CreateSentences.
func (s *Sentences) Stop(ctx *core.Context) error {
	return nil
}

// CreateSentences creates a new instance of the Sentences source.
func CreateSentences(ctx *core.Context,
	ioParams *bql.IOParams, params data.Map) (core.Source, error) {
	interval := 100 * time.Millisecond
	if v, ok := params["interval"]; ok {
		i, err := data.ToDuration(v)
		if err != nil {
			return nil, err
		}
		interval = i
	}
	return core.ImplementSourceStop(&Sentences{
		interval: interval,
	}), nil
}

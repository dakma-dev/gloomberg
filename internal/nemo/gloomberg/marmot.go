package gloomberg

import (
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

// marmot is the job runner and scheduler in gloomberg.
// it is used to run several async tasks like filling the degendb, updating the token ranks, etc.
// jobs will be put into marmots queue and will be executed in a separate goroutine when it's their turn.
//
// rate limits/throttling is handled by the marmot too,
// so we don't hit the api limits of opensea, etherscan, whatever.

type marmot struct {
	// gb *Gloomberg

	intervals map[string]time.Duration
	lastRuns  map[string]time.Time
}

var Jobs = make(chan *marmotJob, 1024)

func defaultIntervals() map[string]time.Duration {
	defaultIntervals := make(map[string]time.Duration)

	log.Printf("defaultIntervals: %+v", viper.GetStringMap("marmot.defaults.intervals"))

	for key, interval := range viper.GetStringMap("marmot.defaults.intervals") {
		var ok bool

		defaultIntervals[key], ok = interval.(time.Duration)
		if !ok {
			log.Errorf("could not cast %v to time.Duration", interval)
		}
	}

	log.Printf("defaultIntervals p: %+v", defaultIntervals)

	return defaultIntervals
}

func AddJob(key string, run func(), params ...any) {
	job := &marmotJob{
		key:    key,
		job:    run,
		params: params,
	}

	Jobs <- job

	gbl.Log.Infof("created scheduled task %+v", job)
}

func (m *marmot) AddJob(key string, run func(), params ...any) {
	job := &marmotJob{
		key:    key,
		job:    run,
		params: params,
	}

	Jobs <- job

	gbl.Log.Infof("created scheduled task %+v", job)
}

func newMarmot() *marmot {
	marmot := &marmot{
		intervals: defaultIntervals(),
		lastRuns:  make(map[string]time.Time),
	}

	return marmot
}

// // run starts the marmot task runner/scheduler.
// func (m *marmot) run() {
// 	for {
// 		job := *<-Jobs

// 		job.runJob(m.gb, job)

// 		time.Sleep(time.Millisecond * 100)
// 	}
// }

// marmotJob implementations.
type marmotJob struct {
	key    string
	job    func()
	params []any
}

// func (j *marmotJob) runJob(_ *Gloomberg, params ...any) {
// 	log.Printf("running job in queue %s | params: %#v | runAt: %s", j.key, params, time.Now().Format("15:04:05"))
// }

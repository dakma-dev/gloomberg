package jobs

import (
	"fmt"
	"sync"
	"time"

	"github.com/benleb/gloomberg/internal"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/log"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/xid"
	"github.com/spf13/viper"
)

// jobs/runner is the job runner and scheduler in gloomberg.

var jobsCounter = promauto.NewCounter(prometheus.CounterOpts{
	Name: "gloomberg_jobs_count_total",
	Help: "No of jobs run by gloomberg.",
},
)

// Job implementations.
type Job struct {
	ID     xid.ID
	Name   string
	key    string
	job    func(...any)
	params []any
}

func (j *Job) runJob(params ...any) { j.job(params...) }
func (j *Job) String() string {
	return fmt.Sprintf(
		"%s | %s | %s",
		style.BoldAlmostWhite(j.Name),
		j.key,
		style.GrayStyle.Render(j.ID.String()),
	)
}

type Runner struct {
	intervals map[string]time.Duration
	lastRuns  map[string]time.Time

	jobs *chan *Job

	*sync.RWMutex
}

var jobQueue = make(chan *Job, 1024)

func NewJobRunner() *Runner {
	runner := &Runner{
		intervals: defaultIntervals(),
		lastRuns:  make(map[string]time.Time),

		jobs: &jobQueue,

		RWMutex: &sync.RWMutex{},
	}

	for i := 0; i < viper.GetInt("jobs.numRunner"); i++ {
		go runner.run()
		time.Sleep(time.Millisecond * 137)
	}

	return runner
}

// run starts the job runner.
func (m *Runner) run() {
	for {
		job := <-*m.jobs

		// get last run and interval for this jobs key
		m.RLock()
		keyLastRun := m.lastRuns[job.key]
		keyInterval := m.intervals[job.key]
		m.RUnlock()

		keyWaitUntil := keyLastRun.Add(keyInterval)

		// check if job can be run or if it's too early
		if keyWaitUntil.Before(time.Now()) {
			m.Lock()
			m.lastRuns[job.key] = time.Now()
			m.Unlock()

			job.runJob(job.params...)

			jobsCounter.Inc()

			gbl.Log.Debugf(
				"✔️ %s | %v %s",
				style.TrendGreenStyle.Render("done"),
				job,
				style.DarkGrayStyle.Render(fmt.Sprintf("| queue: %s", style.GrayStyle.Render(fmt.Sprint(len(*m.jobs))))),
			)

			if jobsProcessed := int64(utils.GetMetricValue(jobsCounter)); jobsProcessed%viper.GetInt64("jobs.status_every") == 0 {
				msg := fmt.Sprintf(
					"✔️ %s jobs %s | %s jobs/min %s",
					style.BoldAlmostWhite(fmt.Sprint(jobsProcessed)),
					style.TrendGreenStyle.Render("processed"),
					style.BoldAlmostWhite(fmt.Sprintf("%.2f", float64(jobsProcessed)/time.Since(internal.RunningSince).Minutes())),
					style.DarkGrayStyle.Render(fmt.Sprintf("|🔜  queue: %s", style.GrayStyle.Render(fmt.Sprint(len(*m.jobs))))),
				)

				gbl.Log.Info(msg)
				log.Print(msg)
			}
		} else {
			*m.jobs <- job

			gbl.Log.Debugf(
				"🔙 re-enqued | %s | next run in %ss %s",
				style.BoldAlmostWhite(job.key),
				style.BoldAlmostWhite(fmt.Sprint(time.Until(keyWaitUntil).Seconds())),
				style.DarkGrayStyle.Render(fmt.Sprintf("| queue: %s", style.GrayStyle.Render(fmt.Sprint(len(*m.jobs))))),
			)

			time.Sleep(time.Millisecond * 137)
		}
	}
}

func (m *Runner) AddJob(name string, key string, run func(...any), params ...any) {
	AddJob(name, key, run, params...)
}

func defaultIntervals() map[string]time.Duration {
	defaultIntervals := make(map[string]time.Duration)

	intervals, ok := viper.Get("jobs.defaults.intervals").(map[string]time.Duration)
	if !ok {
		gbl.Log.Error("🤦‍♀️ invalid jobs.defaults.intervals")

		return defaultIntervals
	}

	for key, interval := range intervals {
		defaultIntervals[key] = interval
	}

	return defaultIntervals
}

func AddJob(name string, key string, run func(...any), params ...any) {
	job := &Job{
		ID:     xid.New(),
		Name:   name,
		key:    key,
		job:    run,
		params: params,
	}

	jobQueue <- job

	gbl.Log.Debugf("🔜 %s | %+v %s", style.LightGrayStyle.Render("new"), job, style.DarkGrayStyle.Render(fmt.Sprintf("| queue: %s", style.GrayStyle.Render(fmt.Sprint(len(jobQueue))))))
}

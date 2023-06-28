package gloomberg

import (
	"time"

	"github.com/benleb/gloomberg/internal/gbl"
)

// marmot is a simple task runner/scheduler.
//
// example usage
//
//	gb.CreatePeriodicTask("testing", 5*time.Second, func(gb *gloomberg.Gloomberg) {
//	    log.Printf("testing tasks lol! %+v", len(gb.Ranks))
//	})
//
//	gb.CreateScheduledTask("testing", time.Now().Add(17*time.Second), func(gb *gloomberg.Gloomberg) {
//	    log.Printf("testing scheduled tasks lol! %+v", len(gb.Ranks))
//	})

type marmot struct {
	gb *Gloomberg

	running bool
	Tasks   []marmotTask
}

func (m *marmot) CreatePeriodicTask(name string, runEvery time.Duration, run func(*Gloomberg)) {
	task := &marmotPeriodicTask{
		Name:     name,
		RunEvery: runEvery,
		Run:      run,
	}

	m.addTask(task)

	gbl.Log.Infof("created periodic task %s | runEvery: %s", task.Name, task.RunEvery.String())
}

func (m *marmot) CreateScheduledTask(name string, runAt time.Time, run func(*Gloomberg)) {
	task := &marmotScheduledTask{
		Name:  name,
		RunAt: runAt,
		Run:   run,
	}

	m.addTask(task)

	gbl.Log.Infof("created scheduled task %s | runAt: %s", task.Name, task.RunAt.Format("15:04:05"))
}

// addTask adds a task to the marmot task list.
func (m *marmot) addTask(task marmotTask) {
	m.Tasks = append(m.Tasks, task)

	// run tasks added after marmot has started
	if m.running {
		go task.runTask(m.gb)
	}
}

// run starts the marmot task runner/scheduler.
func (m *marmot) run() {
	for _, task := range m.Tasks {
		go task.runTask(m.gb)
	}
	m.running = true
}

// marmotTask is an interface that all tasks must implement.
type marmotTask interface {
	runTask(gb *Gloomberg)
}

// Periodic Tasks (run every X duration).
type marmotPeriodicTask struct {
	Name     string
	RunEvery time.Duration
	Run      func(*Gloomberg)
}

func (mt *marmotPeriodicTask) runTask(gb *Gloomberg) {
	t := time.NewTicker(mt.RunEvery)

	for {
		<-t.C

		gbl.Log.Infof("Running periodic task %s | runEvery: %s", mt.Name, mt.RunEvery.String())

		mt.Run(gb)
	}
}

// Scheduled Tasks (run at a specific time).
type marmotScheduledTask struct {
	Name  string
	RunAt time.Time
	Run   func(*Gloomberg)
}

func (mt *marmotScheduledTask) runTask(gb *Gloomberg) {
	t := time.NewTimer(time.Until(mt.RunAt))

	for {
		<-t.C

		gbl.Log.Infof("Running scheduled task %s | runAt: %s", mt.Name, mt.RunAt.Format("15:04:05"))

		mt.Run(gb)
	}
}

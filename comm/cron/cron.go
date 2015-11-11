// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cron

import (
	"sort"
	"time"
)

// Cron keeps track of any number of entries, invoking the associated func as
// specified by the schedule. It may be started, stopped, and the entries may
// be inspected while running.
type Cron struct {
	entries  []*Entry
	stop     chan struct{}
	add      chan *Entry
	snapshot chan []*Entry
	running  bool
}

// Job is an interface for submitted cron jobs.
type Job interface {
	Run()
}

// The Schedule describes a job's duty cycle.
type Schedule interface {
	// Return the next activation time, later than the given time.
	// Next is invoked initially, and then each time the job is run.
	Next(time.Time) time.Time
}

// Entry consists of a schedule and the func to execute on that schedule.
type Entry struct {
	// The schedule on which this job should be run.
	Schedule Schedule

	// The next time the job will run. This is the zero time if Cron has not been
	// started or this entry's schedule is unsatisfiable
	Next time.Time

	// The last time this job was run. This is the zero time if the job has never
	// been run.
	Prev time.Time

	// The Job to run.
	Job Job
}

// byTime is a wrapper for sorting the entry array by time
// (with zero time at the end).
type byTime []*Entry

func (me byTime) Len() int      { return len(me) }
func (me byTime) Swap(i, j int) { me[i], me[j] = me[j], me[i] }
func (me byTime) Less(i, j int) bool {
	// Two zero times should return false.
	// Otherwise, zero is "greater" than any other time.
	// (To sort it at the end of the list.)
	if me[i].Next.IsZero() {
		return false
	}
	if me[j].Next.IsZero() {
		return true
	}
	return me[i].Next.Before(me[j].Next)
}

// New returns a new Cron job runner.
func New() *Cron {
	return &Cron{
		entries:  nil,
		add:      make(chan *Entry),
		stop:     make(chan struct{}),
		snapshot: make(chan []*Entry),
		running:  false,
	}
}

// A wrapper that turns a func() into a cron.Job
type FuncJob func()

func (me FuncJob) Run() {
	defer func() {
		if err := recover(); err != nil {
			logger.Printf("Panic recovery FuncJob.Run -> %s\n", err)
		}
	}()
	me()
}

// AddFunc adds a func to the Cron to be run on the given schedule.
func (me *Cron) AddFunc(spec string, cmd func()) error {
	return me.AddJob(spec, FuncJob(cmd))
}

// AddFunc adds a Job to the Cron to be run on the given schedule.
func (me *Cron) AddJob(spec string, cmd Job) error {
	schedule, err := Parse(spec)
	if err != nil {
		return err
	}
	me.Schedule(schedule, cmd)
	return nil
}

// Schedule adds a Job to the Cron to be run on the given schedule.
func (me *Cron) Schedule(schedule Schedule, cmd Job) {
	entry := &Entry{
		Schedule: schedule,
		Job:      cmd,
	}
	if !me.running {
		me.entries = append(me.entries, entry)
		return
	}

	me.add <- entry
}

// Entries returns a snapshot of the cron entries.
func (me *Cron) Entries() []*Entry {
	if me.running {
		me.snapshot <- nil
		x := <-me.snapshot
		return x
	}
	return me.entrySnapshot()
}

// Start the cron scheduler in its own go-routine.
func (me *Cron) Start() {
	me.running = true
	go me.run()
}

// Run the scheduler.. this is private just due to the need to synchronize
// access to the 'running' state variable.
func (me *Cron) run() {
	defer func() {
		if err := recover(); err != nil {
			logger.Printf("Panic recovery Cron.run -> %s\n", err)
		}
	}()
	// Figure out the next activation times for each entry.
	now := time.Now().Local()
	for _, entry := range me.entries {
		entry.Next = entry.Schedule.Next(now)
	}

	for {
		// Determine the next entry to run.
		sort.Sort(byTime(me.entries))

		var effective time.Time
		if len(me.entries) == 0 || me.entries[0].Next.IsZero() {
			// If there are no entries yet, just sleep - it still handles new entries
			// and stop requests.
			effective = now.AddDate(10, 0, 0)
		} else {
			effective = me.entries[0].Next
		}

		select {
		case now = <-time.After(effective.Sub(now)):
			// Run every entry whose next time was this effective time.
			for _, e := range me.entries {
				if e.Next != effective {
					break
				}
				go e.Job.Run()
				e.Prev = e.Next
				e.Next = e.Schedule.Next(effective)
			}
			continue

		case newEntry := <-me.add:
			me.entries = append(me.entries, newEntry)
			newEntry.Next = newEntry.Schedule.Next(now)

		case <-me.snapshot:
			me.snapshot <- me.entrySnapshot()

		case <-me.stop:
			return
		}

		// 'now' should be updated after newEntry and snapshot cases.
		now = time.Now().Local()
	}
}

// Stop the cron scheduler.
func (me *Cron) Stop() {
	me.stop <- struct{}{}
	me.running = false
}

// entrySnapshot returns a copy of the current cron entry list.
func (me *Cron) entrySnapshot() []*Entry {
	entries := []*Entry{}
	for _, e := range me.entries {
		entries = append(entries, &Entry{
			Schedule: e.Schedule,
			Next:     e.Next,
			Prev:     e.Prev,
			Job:      e.Job,
		})
	}
	return entries
}

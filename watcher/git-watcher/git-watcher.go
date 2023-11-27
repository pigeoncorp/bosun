package gitwatcher

import (
	"sync"
	"time"

	"github.com/hashibuto/oof"
	"github.com/pigeoncorp/bosun/watcher/config"
	"github.com/pigeoncorp/bosun/watcher/git"
	log "github.com/sirupsen/logrus"
)

var inst *GitWatcher
var lock = &sync.Mutex{}

type GitWatcher struct {
	stopChannel chan struct{}
	waitGroup   *sync.WaitGroup
}

func NewGitWatcher() *GitWatcher {
	gitWatcher := &GitWatcher{
		stopChannel: make(chan struct{}, 1),
		waitGroup:   &sync.WaitGroup{},
	}

	gitWatcher.waitGroup.Add(1)
	go gitWatcher.tickWatcher()
	return gitWatcher
}

func GetInstance() *GitWatcher {
	lock.Lock()
	defer lock.Unlock()

	if inst == nil {
		inst = NewGitWatcher()
	}

	return inst
}

func (gw *GitWatcher) Stop() {
	gw.stopChannel <- struct{}{}
	gw.waitGroup.Wait()
}

func (gw *GitWatcher) tickWatcher() {
	defer gw.waitGroup.Done()

	tickInterval := time.Duration(config.Config.GitCheckInterval) * time.Second
	timer := time.NewTimer(tickInterval)
	for {
		select {
		case <-gw.stopChannel:
			if !timer.Stop() {
				<-timer.C
			}
			return
		case <-timer.C:
			err := gw.checkForChanges()
			if err != nil {
				log.Error(oof.Tracef("error checking for git changes: %w", err))
			}
		}

	}
}

func (gw *GitWatcher) checkForChanges() error {
	err := git.CloneRepoIfNotExists()
	if err != nil {
		return err
	}

	err = git.StoreCredentials()
	if err != nil {
		return err
	}

	err = git.PullRepoChanges()
	if err != nil {
		return err
	}

	return nil
}

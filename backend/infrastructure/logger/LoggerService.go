package commonLogger

import (
	"fmt"
	loggerInterface "github.com/revazcus/task-tracker/backend/infrastructure/logger/interface"
	logModel "github.com/revazcus/task-tracker/backend/infrastructure/logger/model"
	"sync"
	"time"
)

const (
	defaultInputBufferSize = 100
	defaultJobBufferSize   = 1000
	defaultNumWorkers      = 4
	sendTimeout            = 100 * time.Millisecond
)

type LoggerService struct {
	inputChan  chan *logModel.LogData
	jobChan    chan sendJob
	stopChan   chan struct{}
	numWorkers int
	mutex      sync.RWMutex
	loggers    map[string]loggerInterface.LogPublisher
}

func NewLoggerService(stopChan chan struct{}) *LoggerService {
	return &LoggerService{
		inputChan:  make(chan *logModel.LogData, defaultInputBufferSize),
		jobChan:    make(chan sendJob, defaultJobBufferSize),
		loggers:    make(map[string]loggerInterface.LogPublisher),
		stopChan:   stopChan,
		numWorkers: defaultNumWorkers,
	}
}

func (ls *LoggerService) AddLogger(loggerId string, logger loggerInterface.LogPublisher) {
	ls.mutex.Lock()
	defer ls.mutex.Unlock()
	ls.loggers[loggerId] = logger
}

func (ls *LoggerService) RemoveLogger(loggerId string) {
	ls.mutex.Lock()
	defer ls.mutex.Unlock()
	delete(ls.loggers, loggerId)
}

func (ls *LoggerService) GetInputChan() chan<- *logModel.LogData {
	return ls.inputChan
}

func (ls *LoggerService) Start() {
	go ls.runMainWorker()

	for i := 0; i < ls.numWorkers; i++ {
		go ls.runWorker()
	}
}

func (ls *LoggerService) runMainWorker() {
	defer close(ls.jobChan)
	for {
		select {
		case <-ls.stopChan:
			return
		case logData := <-ls.inputChan:
			if logData == nil {
				continue
			}
			ls.mutex.RLock()
			if len(ls.loggers) == 0 {
				fmt.Println("No loggers configured. Skipping log message")
				continue
			}
			for id, logger := range ls.loggers {
				if logger == nil {
					fmt.Printf("Logger with Id %q is nil. Skipping.\n", id)
					continue
				}
				job := sendJob{loggerId: id, logger: logger, logData: logData}
				ls.jobChan <- job
			}
			ls.mutex.RUnlock()
		}
	}
}

func (ls *LoggerService) runWorker() {
	for job := range ls.jobChan {
		ls.processJob(job)
	}
}

func (ls *LoggerService) processJob(job sendJob) {
	doneChan := make(chan struct{})
	go func() {
		job.logger.SendMessage(job.logData)
		close(doneChan)
	}()

	select {
	case <-doneChan:
	case <-time.After(sendTimeout):
		fmt.Printf("Failed to send log message to logger %q within %v. Original message: %q\n", job.loggerId, sendTimeout, job.logData.Msg)
	}
}

type sendJob struct {
	loggerId string
	logger   loggerInterface.LogPublisher
	logData  *logModel.LogData
}

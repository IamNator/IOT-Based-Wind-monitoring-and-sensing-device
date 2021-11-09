package storage

import "github.com/IamNator/iot-wind/model"

//LogDatabase defines methods for operations involving the entity Log
type LogDatabase interface {
	//FindAllLogs returns the logs stored in database
	FindAllLogs(page, pageSize int) ([]*model.Log, error)

	//CreateLog adds a new log to database
	CreateLog(model.Log) error

	//FindRecent returns the last created log
	FindRecent() (*model.Log, error)
}

//Log implements LogDatabase interface
type Log struct {
	storage *Storage
}

//NewLog creates a new log object for database operation involving logs
func NewLog(s *Storage) LogDatabase {
	return &Log{
		storage: s,
	}
}

//FindAllLogs returns the logs stored in database
func (l *Log) FindAllLogs(page, pageSize int) ([]*model.Log, error) {
	logs := make([]*model.Log, 0)
	result := l.storage.db.Scopes(Paginate(page, pageSize)).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

//CreateLog adds a new log to database
func (l *Log) CreateLog(log model.Log) error {

	result := l.storage.db.Save(&log)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

//FindRecent returns the last created log
func (l *Log) FindRecent() (*model.Log, error) {
	log := new(model.Log)
	result := l.storage.db.Last(&log)
	if result.Error != nil {
		return nil, result.Error
	}
	return log, nil
}

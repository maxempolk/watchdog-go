package repository

import (
	"errors"
	log "stat_by_sites/domain/log"
	"sync"
)

type MemoryLogRepository struct{
	mu        sync.RWMutex
	logs 			[]log.Log
}

func NewMemoryLogRepository() *MemoryLogRepository{
	return &MemoryLogRepository{
		logs: make([]log.Log, 0),
	}
}

func (r *MemoryLogRepository) Add(logs ...log.Log){
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logs = append(r.logs, logs...)
}

func (r *MemoryLogRepository) Length() int{
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.logs)
}

func (r *MemoryLogRepository) List(offset, limit int) ([]log.Log, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	total := len(r.logs)

	if offset < 0 || offset > total {
		return nil, errors.New("offset is invalid")
	}

	if limit < 0 {
		return nil, errors.New("limit is invalid")
	}

	if offset == total {
		return []log.Log{}, nil
	}

	end := offset + limit
	if end > total {
		end = total
	}

	result := make([]log.Log, end-offset)
	copy(result, r.logs[offset:end])

	return result, nil
}
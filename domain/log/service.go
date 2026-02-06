package userlog

type LogService struct {
	repo LogRepository
}

func NewLogService(r LogRepository) *LogService {
	return &LogService{repo: r}
}

func (ls *LogService) Add(logs ...Log) {
	if len(logs) == 0 {
		return
	}

	ls.repo.Add(logs...)
}

func (ls *LogService) List(offset, limit int) ([]Log, error) {
	return ls.repo.List(offset, limit)
}

func (ls *LogService) Count() int {
	return ls.repo.Length()
}

// ListRecent returns logs from newest to oldest.
func (ls *LogService) ListRecent(limit int) ([]Log, error) {
	if limit <= 0 {
		return []Log{}, nil
	}

	total := ls.repo.Length()
	if total == 0 {
		return []Log{}, nil
	}

	if limit > total {
		limit = total
	}

	offset := total - limit
	logs, err := ls.repo.List(offset, limit)
	if err != nil {
		return nil, err
	}

	for i, j := 0, len(logs)-1; i < j; i, j = i+1, j-1 {
		logs[i], logs[j] = logs[j], logs[i]
	}

	return logs, nil
}

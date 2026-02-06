package userlog

type LogRepository interface {
	Add(logs ...Log)
	Length() int
	List(offset, limit int) ([]Log, error)
}

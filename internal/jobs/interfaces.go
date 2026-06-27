package jobs

type Job interface {
	Execute()
	ID()
}

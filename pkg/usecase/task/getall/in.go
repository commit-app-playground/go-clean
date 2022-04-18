package getall

type GetAllTasksIn interface {
}

type getAllTasksIn struct {
}

func NewGetAllTasksIn() (GetAllTasksIn, error) {
	return &getAllTasksIn{}, nil
}

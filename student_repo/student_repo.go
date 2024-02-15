package student_repo

type Student struct {
	Name  string
	Age   int
	Grade int
}

type StudentRepository interface {
	Put(name string, student *Student)
	Get(name string) *Student
	GetStudents() []*Student
}

type MapStudentRepository struct {
	storage map[string]*Student
}

func NewMapStudentRepository() *MapStudentRepository {
	return &MapStudentRepository{
		storage: make(map[string]*Student),
	}
}

func (msr *MapStudentRepository) Put(name string, student *Student) {
	msr.storage[name] = student
}

func (msr *MapStudentRepository) Get(name string) *Student {
	return msr.storage[name]
}

func (msr *MapStudentRepository) GetStudents() []*Student {
	students := make([]*Student, 0, len(msr.storage))
	for _, student := range msr.storage {
		students = append(students, student)
	}
	return students
}

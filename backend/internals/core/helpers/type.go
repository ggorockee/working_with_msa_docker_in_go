package helpers

type Refer struct {
	userId int
}

func NewRefer() *Refer {

	return &Refer{}
}

func (r *Refer) SetUserId(id int) {
	r.userId = id
}

func (r *Refer) UserId() int {
	return r.userId
}

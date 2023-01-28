package productSQL

type Service interface{

}

type service struct{
	 r Repository
}

func NewSqlService(r Repository) Service {
	return &service{
		r,
	}
}
package riders


type RiderSv interface {
}


type RiderServ struct {
	store RiderStore
}

func NewRiderService(store RiderStore) *RiderServ {
	return &RiderServ{store: store}
}

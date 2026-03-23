type GameId string

type CreateGameRequest struct {
	Name    string
	Creator string
}

type ConnectGameRequest struct {
	Id GameId
}

type Creator struct {
	Name string
}
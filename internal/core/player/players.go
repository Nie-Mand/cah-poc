package player

type DirtyPlayers []*DirtyPlayer

func (p *DirtyPlayers) Add(player *DirtyPlayer) {
	*p = append(*p, player)
}

func (p *DirtyPlayers) NextCzar() *DirtyPlayer {
	for _, player := range *p {
		if player.IsCzar() {
			player.UnsetCzar()
			continue
		}
		return player
	}
	return nil
}

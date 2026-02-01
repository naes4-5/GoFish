package game

type Player struct {
	Name  string
	Hand  []Card
	Books int8
}

func NewPlayer(d *Deck, name string, handSize int8) (*Player, error) {
	p := Player{
		Name:  name,
		Hand:  make([]Card, 0, handSize),
		Books: 0,
	}
	for range handSize {
		c, err := d.DrawCard()
		if err != nil {
			return nil, err
		}
		p.Hand = append(p.Hand, c)
	}
	return &p, nil
}

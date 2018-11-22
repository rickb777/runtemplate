package support

type Triple struct {
	Key, Val, Alt string
}

type Triples []Triple

func (triples Triples) TValues() []string {
	var list []string
	for _, p := range triples {
		switch p.Val {
		case "true", "false": // drop
		default:
			v := p.Val
			if len(v) > 0 && v[0] == '*' {
				v = v[1:]
			}
			list = append(list, RichString(v).NoDots().String())
		}
	}
	return list
}

func (triples Triples) PValues() []string {
	var list []string
	for _, p := range triples {
		switch p.Val {
		case "true", "false": // drop
		default:
			list = append(list, p.Val)
		}
	}
	return list
}

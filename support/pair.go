package support

type Pair struct {
	Key, Val string
}

type Pairs []Pair

func (pairs Pairs) Keys() []string {
	var list []string
	for _, p := range pairs {
		list = append(list, p.Key)
	}
	return list
}

func (pairs Pairs) TValues() []string {
	var list []string
	for _, p := range pairs {
		switch p.Val {
		case "true", "false": // drop
		default:
			v := p.Val
			if len(v) > 0 && v[0] == '*' {
				v = v[1:]
			}
			list = append(list, string(RichString(v).NoDots()))
		}
	}
	return list
}

func (pairs Pairs) PValues() []string {
	var list []string
	for _, p := range pairs {
		switch p.Val {
		case "true", "false": // drop
		default:
			list = append(list, p.Val)
		}
	}
	return list
}

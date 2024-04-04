package module

type Category int

const (
	CategoryExploit Category = iota
	CategoryAuxiliary
	CategoryPost
	CategoryPayload
)

func (c Category) String() string {
	switch c {
	case CategoryExploit:
		return "exploit"
	case CategoryAuxiliary:
		return "auxiliary"
	case CategoryPost:
		return "post"
	case CategoryPayload:
		return "payload"
	}
	return "unknown"
}

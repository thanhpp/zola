package dto

func boolTranslate(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

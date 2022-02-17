package api

func metadata(scope string) string {
	if scope != "" {
		return scope
	}
	return "No Metadata"
}

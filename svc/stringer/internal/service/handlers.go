package service

func CountHandler() CountEndpoint {
	svc, _ := New()
	return svc.Count
}

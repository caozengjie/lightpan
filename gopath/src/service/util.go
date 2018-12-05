package service

func checkupdate(updater map[string]interface{}, allowupdate map[string]interface{}) (disableupdatefields []string) {
	disableupdatefields = make([]string, 0)
	for k, _ := range updater {
		if _, have := allowupdate[k]; !have {
			delete(updater, k)
			disableupdatefields = append(disableupdatefields, k)
		}
	}
	return
}

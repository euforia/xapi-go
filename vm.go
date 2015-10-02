package xapi

func (client *Client) GetVMs() (machines []string, err error) {
	var resp StringsResponse
	err = client.SessionCall(&resp, "VM.get_all")
	return resp.Value, err
}

//func (client *Client) GetVMRecords() (machines []VM, err error) {
func (client *Client) GetVMRecords() (machines map[string]VM, err error) {
	var resp VMsResponse
	err = client.SessionCall(&resp, "VM.get_all_records")
	return resp.Value, err
}

func (client *Client) GetVM(opref string) (vm VM, err error) {
	var resp VMResponse
	err = client.SessionCall(&resp, "VM.get_record", opref)
	return resp.Value, err
}

func (client *Client) GetVMByName(name string) (refs []string, err error) {
	var resp StringsResponse
	err = client.SessionCall(&resp, "VM.get_by_name_label", name)
	return resp.Value, err
}

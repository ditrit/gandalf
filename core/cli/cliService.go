package cli

// GandalfAuthenticationService :
type CliService struct {
	client *Client
}

// Login :
func (cs *CliService) Cli() (string, error) {

	req, err := cs.client.newRequest("GET", "/gandalf/cli/", "", nil)
	if err != nil {
		return "", err
	}
	var result string
	err = cs.client.do(req, &result)

	return result, err

}

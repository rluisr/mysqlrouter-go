package mysqlrouter

var (
	client *Client
)

func setup() {
	var err error
	client, err = New("https://mysqlrouter-test.xzy.pw", "luis", "luis", false)
	if err != nil {
		panic(err)
	}
}

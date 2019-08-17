package mysqlrouter

var (
	client *Client
)

func setup() {
	client, _ = New("https://mysqlrouter-test.xzy.pw", "luis", "luis")
}

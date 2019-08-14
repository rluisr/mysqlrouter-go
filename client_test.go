package mysqlrouter

var (
	client *Client
)

func setup() {
	client, _ = New("http://db-router.luis.local:8080", "luis", "luis")
}

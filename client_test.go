package mysqlrouter

var (
	client *Client
)

func setup() {
	client, _ = New("http://localhost:5901", "luis", "luis")
}

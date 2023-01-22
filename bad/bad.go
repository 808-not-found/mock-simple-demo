package bad

import "mock-simple-demo/good"

func Bad() string {
	return good.Good("Yes")
}

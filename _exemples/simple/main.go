package main

import (
	"fmt"

	"github.com/gandarfh/god"
)

var BodySchema = god.Object(god.Map{
	"name": god.String(god.Number()),
	"bool": god.String(god.Boolean()),
	"test": god.Int64(
		god.Required("Test is required"),
		god.Max(40),
		god.Min(10),
		god.Gte(12),
	),
	"email": god.String(
		god.Required("Email is required"),
		god.Email("formato do email incorreto."),
	),
})

var stringSlice = god.Slice(
	god.String(god.Required(), god.Email("formato do email incorreto.")),
	god.Required("slice must have at least one element"),
)

type Body struct {
	Name  interface{} `json:"name"`
	Test  *int64      `json:"test"`
	Email string      `json:"email"`
}

func main() {
	value := int64(11)
	body := map[string]any{
		"name":  "1234",
		"bool":  "false",
		"email": "joao@email.com",
		"test":  &value,
	}

	// if err := god.Validate(body, BodySchema); err != nil {
	// 	fmt.Println(err)
	// }

	// value := int64(1)
	// body2 := Body{
	// 	Test: &value,
	// }

	if err := god.Validate(body, BodySchema); err != nil {
		fmt.Println(err)
	}

	// if err := god.Validate([]string{}, stringSlice); err != nil {
	// 	fmt.Println(err)
	// }

}

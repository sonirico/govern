package main

import (
	"log"

	"github.com/sonirico/govern"
)

func main() {

	type Struct struct {
		Name            string
		Fruit           string
		MaybeOtherFruit *string
	}

	structSchema := govern.New[Struct]().
		Enum(
			govern.FieldEnum[Struct, string](
				"fruits",
				govern.String(
					govern.StringOpts{
						MinLen:   3,
						Required: true,
					},
				),
				func(x Struct) string { return x.Fruit },
				"pera",
				"manzana",
			),
		).
		String(
			"name",
			func(x Struct) string { return x.Name },
			govern.StringOpts{
				MinLen:   3,
				Required: true,
			}).
		Maybe(
			govern.FieldMaybe[Struct, string](
				"maybeOtherFruit",
				govern.String(
					govern.StringOpts{
						MinLen:   3,
						Required: true,
					},
				),
				true,
				func(x Struct) *string { return x.MaybeOtherFruit },
			),
		).
		Maybe(
			govern.FieldMaybe[Struct, string](
				"maybeOtherFruit2",
				govern.Enum[string](
					govern.String(
						govern.StringOpts{
							MinLen:   3,
							Required: true,
						},
					),
					"pera",
					"manzana",
				),
				true,
				func(x Struct) *string { return x.MaybeOtherFruit },
			),
		)

	banana := "pera"

	res := structSchema.Check(Struct{
		Name:            "123",
		Fruit:           "pera",
		MaybeOtherFruit: &banana,
	})

	if res.IsErr() {
		for _, err := range res.Errors {
			log.Println("an error occurred: ", err.Error())
		}
	}
	log.Println("todo ok? ", res.IsOk())
}

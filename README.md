<div align="center">
  <img
    src="assets/god.png"
    alt="httui"
    style="width: 100%"
  />
  <p>God é uma alternativa que roda encima do go-playground/validator</p>
</div>

## Introdução

God é uma biblioteca de validação de dados leve e eficaz escrita em Go. Oferece uma maneira limpa e concisa para validar estruturas complexas e tipos de dados em Go.

## Instalação

Para adicionar o God ao seu projeto Go, execute o seguinte comando:

```bash
go get github.com/gandarfh/god
```

## Uso Básico

God fornece uma série de funções de validação prontas para uso que você pode usar para validar seus dados. Cada função de validação retorna um erro que pode ser usado para determinar se a validação foi bem-sucedida ou não.

```go

import (
	"github.com/gandarfh/god"
)

// Defining the validation schema
var schema = god.Object(god.Map{
    "stringField": god.String(god.Required(), god.Uppercase()),
    "intField":    god.Int64(god.Required()),
    "boolField":   god.Bool(),
    "floatField":  god.Float64(),
    "sliceField":  god.Slice(god.Int64(god.Required())),
    "nestedField": god.Object(god.Map{
        "name":  god.String(god.Required()),
        "email": god.String(god.Required(), god.Contains("@"), god.Email()),
    }),
})

func main () {
	exemple := map[string]interface{}{
		"stringField": "TEST",
		"intField":    int64(10),
		"boolField":   true,
		"floatField":  3.14,
		"sliceField":  []int64{1, 2, 3},
		"nestedField": map[string]interface{}{
			"name":  "John",
			"email": "john@example.com",
		},
	}

	if err := god.Validate(exemple, schema); err != nil {
		fmt.Println(err)
	}
}
```

| Tipo   | Função  | Validações Suportadas                                                              |
| ------ | ------- | ---------------------------------------------------------------------------------- |
| string | String  | Required, Min, Max, Email, URL, Lowercase, Uppercase, Contains, Eq                 |
| int64  | Int64   | Required, Min, Max, Eq, Ne, Gt, Gte, Lt, Lte                                       |
| bool   | Boolean | Required, Eq                                                                       |
| slice  | Slice   | Required, Min, Max                                                                 |
| map    | Object  | N/A (usa um Map para mapear campos de estrutura para seus schemas correspondentes) |

## Criando um Schema Customizado

Você pode facilmente criar um schema personalizado para validar seus próprios dados.

Para ilustrar melhor, vamos criar um Schema personalizado que valida se um valor é uma string e contém exatamente 3 caracteres.

```go
package god

import (
	"fmt"
	"strings"
)

func CustomLengthString(v ...Validation) Schema {
	return func(value interface{}) error {
		return CommonValidation(v, value, "string", func(val interface{}) (interface{}, bool) {
			strVal, ok := val.(string)
			return strVal, ok && len(strings.TrimSpace(strVal)) == 3
		})
	}
}
```

Nesse exemplo, `CustomLengthString` é uma função que cria um schema personalizado que valida se um valor é uma `string` e contém exatamente 3 caracteres.

Agora, você pode usar este schema personalizado em seu código, como qualquer outra validação fornecida pela biblioteca:

```go
func main() {
    userSchema := god.Object( god.Map{
        "username": CustomLengthString(Required()),
        // ...
    })

    user := map[string]interface{}{
        "username": "abc", // esse valor vai passar na validação
        // ...
    }

    if err := god.Validate(user, userSchema); err != nil {
        fmt.Println(err)
    }
}
```

Aqui, 'username' deve ser uma string de exatamente 3 caracteres. Se for diferente, a validação retornará um erro. Dessa forma, você pode criar suas próprias validações para se adequar às necessidades específicas de seu projeto.

No caso acima, a função `CustomLengthString` é uma função de validação personalizada que usa a função `CommonValidation` para fazer o trabalho pesado. Ela simplesmente fornece uma função que sabe como validar o tipo de dado necessário e passa essa função, juntamente com os outros parâmetros necessários, para `CommonValidation`.

## Criando uma validação Customizado

Para criar uma validação personalizada, você precisará fornecer sua própria função de validação. Por exemplo, você pode criar uma função de validação `MyCustomValidation` que exige que uma string seja "hello".

```go
func MyCustomValidation(message ...string) Validation {
	return Validation{
		Tag: "custom",
		Func: func(v interface{}) error {
			str, ok := v.(string)
			if !ok {
				return fmt.Errorf("value is not a string")
			}

			if str != "hello" {
				return fmt.Errorf(god.GetMessage(message, "value is not 'hello'"))
			}

			return nil
		},
		Message: god.GetMessage(message, "Failed on custom validation!"),
	}
}
```

Agora você pode usar MyCustomValidation como qualquer outra validação.

```go
func main() {
	schema := god.String(
		god.Required(),
		god.MyCustomValidation(),
	)

	err := god.Validate("hello", schema)
	if err != nil {
		fmt.Println(err)
	}

	err = god.Validate("world", schema)
	if err != nil {
		fmt.Println(err)  // irá imprimir "Failed on custom validation!"
	}
}
```

No exemplo acima, a validação `MyCustomValidation` verifica se o valor fornecido é a string "hello". Se não for, ele retorna um erro com a mensagem especificada. Isso permite que você crie validações personalizadas que se adequem às necessidades específicas do seu projeto.

## Contato

Twitter - [@gandarfh](https://twitter.com/gandarfh)

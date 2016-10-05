package basic

type serror string

/* Translate Function */
func Translate(input string) (string, string) {

  //Lex Input
  toks, err := Lex(input)

  if err != "" {
      return "", string(err)
  }

  //Translate Input
  trans := Parse(toks)

  //Return input
  return trans, ""
}

/* Classes and Helpers */

type token struct {
  name string
  value string
}

func newToken(name string, value string) *token {
  tok := new(token)
  tok.name = name
  tok.value = value
  return tok
}

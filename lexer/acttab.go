// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"fmt"

	"github.com/sirkon/ldetool/token"
)

type ActionTable [NumStates]ActionRow

type ActionRow struct {
	Accept token.Type
	Ignore string
}

func (a ActionRow) String() string {
	return fmt.Sprintf("Accept=%d, Ignore=%s", a.Accept, a.Ignore)
}

var ActTab = ActionTable{
	ActionRow{ // S0
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S1
		Accept: -1,
		Ignore: "!whitespace",
	},
	ActionRow{ // S2
		Accept: 7,
		Ignore: "",
	},
	ActionRow{ // S3
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S4
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S5
		Accept: 1,
		Ignore: "",
	},
	ActionRow{ // S6
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S7
		Accept: 5,
		Ignore: "",
	},
	ActionRow{ // S8
		Accept: 6,
		Ignore: "",
	},
	ActionRow{ // S9
		Accept: 14,
		Ignore: "",
	},
	ActionRow{ // S10
		Accept: 15,
		Ignore: "",
	},
	ActionRow{ // S11
		Accept: 4,
		Ignore: "",
	},
	ActionRow{ // S12
		Accept: 3,
		Ignore: "",
	},
	ActionRow{ // S13
		Accept: 17,
		Ignore: "",
	},
	ActionRow{ // S14
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S15
		Accept: 13,
		Ignore: "",
	},
	ActionRow{ // S16
		Accept: 16,
		Ignore: "",
	},
	ActionRow{ // S17
		Accept: 8,
		Ignore: "",
	},
	ActionRow{ // S18
		Accept: 12,
		Ignore: "",
	},
	ActionRow{ // S19
		Accept: 9,
		Ignore: "",
	},
	ActionRow{ // S20
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S21
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S22
		Accept: -1,
		Ignore: "!line_comment",
	},
	ActionRow{ // S23
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S24
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S25
		Accept: 11,
		Ignore: "",
	},
	ActionRow{ // S26
		Accept: 2,
		Ignore: "",
	},
	ActionRow{ // S27
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S28
		Accept: 0,
		Ignore: "",
	},
	ActionRow{ // S29
		Accept: 10,
		Ignore: "",
	},
}

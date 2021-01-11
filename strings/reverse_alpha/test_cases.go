package reverse

var testcases = []struct {
	name   string
	input  string
	output string
}{
	{
		name:   "simple",
		input:  "abc",
		output: "cba",
	},
	{
		name:   "with_punctuation",
		input:  "a,b$c",
		output: "c,b$a",
	},
	{
		name:   "with_capital",
		input:  "Ab,c,de!$",
		output: "ed,c,bA!$",
	},
	{
		name:   "with_trailing_space",
		input:  "a,b$c  ",
		output: "c,b$a  ",
	},
}

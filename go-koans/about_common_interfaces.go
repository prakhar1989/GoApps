package go_koans

import "bytes"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		out.WriteString(in.String())

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
		x, _ := in.ReadString('o')
		out.WriteString(x)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}

package marsh

import (
	"encoding"
	"encoding/json"
	"io"
)

// FullWriter is a utility type to couple with Marshal funcs.
type FullWriter struct {
	W io.Writer
	e error
}

func (fw *FullWriter) setErr(e error) error {
	if e != nil {
		fw.e = e
	}
	return fw.e
}

// WriteBytes takes a buffer and an error, and writes it to
// the wrapped Writer inside of FullWriter. It's meant as a
// convenience, so you may call fw.WriteBytes(t.MarshalText())
func (fw *FullWriter) WriteBytes(b []byte, merr error) error {
	if fw.e != nil {
		return fw.e
	}
	if merr != nil {
		return fw.setErr(merr)
	}
	try := 3
	for len(b) > 0 {
		if try == 0 {
			return fw.setErr(io.ErrNoProgress)
		}
		n, werr := fw.W.Write(b)
		if werr != nil {
			return fw.setErr(werr)
		}
		b = b[n:]
		if n == 0 {
			try--
		}
	}
	return nil
}

// BinWrite calls m.MarshalBinary() and writes the result to w.
func BinWrite(w io.Writer, m encoding.BinaryMarshaler) error {
	fw := &FullWriter{W: w}
	return fw.WriteBytes(m.MarshalBinary())
}

// TextWrite calls m.MarshalText() and writes the result to w.
func TextWrite(w io.Writer, m encoding.TextMarshaler) error {
	fw := &FullWriter{W: w}
	return fw.WriteBytes(m.MarshalText())
}

// JSONWrite calls m.MarshalJSON() and writes the result to w.
func JSONWrite(w io.Writer, m json.Marshaler) error {
	fw := &FullWriter{W: w}
	return fw.WriteBytes(m.MarshalJSON())
}

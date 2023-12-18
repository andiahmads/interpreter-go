package ast

import (
	"bytes"
	"go-intepreter/token"
)

// AST adalah singkatan dari "Abstract Syntax Tree" atau "Pohon Sintaksis Abstrak" dalam bahasa Indonesia.
// AST adalah struktur data pohon yang merepresentasikan struktur sintaksis dari suatu kode sumber setelah proses parsing.
// Setiap node dalam AST mewakili elemen-elemen sintaksis dari kode sumber, seperti pernyataan, ekspresi, dan deklarasi.

type Node interface {
	// membuat tiga bidang: satu untuk pengenal, satu untuk ekspresi yang menghasilkan nilai dalam pernyataan let dan satu untuk token.
	TokenLiteral() string
	String() string // mencetak node AST untuk debugging dan membandingkannya dengan node AST lainnya
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// Node Program ini akan menjadi node akar dari setiap AST yang dihasilkan oleh parser kita.
// Setiap program COKLnag yang valid adalah serangkaian pernyataan.
// Pernyataan-pernyataan ini terkandung dalam Program.Statements,
// yang hanya merupakan potongan dari node AST yang mengimplementasikan antarmuka Statement.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement memiliki bidang-bidang yang kita butuhkan: Nama untuk menyimpan pengenal pengikatan dan Nilai untuk yang menghasilkan nilai.
// Dua metode statementNode dan TokenLiteral memenuhi antarmuka Statement dan Node masing-masing.
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

// Metode statementNode dan TokenLiteral ada untuk memenuhi antarmuka Node dan Statement dan terlihat identik dengan metode yang didefinisikan pada *ast.LetStatement.
func (ls *LetStatement) statementNode() {
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Untuk menjaga jumlah tipe simpul yang berbeda tetap kecil,
// kita akan menggunakan Identifier di sini untuk merepresentasikan nama dalam pengikatan variabel dan kemudian menggunakannya kembali,
// untuk merepresentasikan sebuah pengenal sebagai bagian dari atau sebagai lengkap dari sebuah ekspresi.
type Identifier struct {
	Token token.Token
	Value string
}

// Untuk menyimpan pengenal dari pengikatan, x dalam let x = 5;, kita memiliki tipe pengenal struct, yang mengimplementasikan antarmuka Expression.
// Tetapi pengenal dalam pernyataan let tidak menghasilkan nilai, bukan? Jadi mengapa ini adalah sebuah Ekspresi? Ini untuk membuat segalanya tetap sederhana. Pengenal di bagian lain
// lain dari program Monkey memang menghasilkan nilai, contoh: let x = valueProducingIdentifier;.
func (i *Identifier) expressionNode() {
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// Dengan Program, LetStatement dan Identifier mendefinisikan bagian kode sumber COKLang ini;
// let x = 5;

// Parsing Return Statements
// return 5;
// return 10;
// return add(15);
// Pernyataan return hanya terdiri dari kata kunci return dan sebuah ekspresi. Hal ini membuat definisi ast.ReturnStatement menjadi sangat sederhana:
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// Metode statementNode dan TokenLiteral ada untuk memenuhi antarmuka Node dan Statement dan terlihat identik dengan metode yang didefinisikan pada *ast.LetStatement.
func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// let x = 5;
// x + 10;
// Baris pertama adalah pernyataan let, baris kedua adalah pernyataan ekspresi
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (rs *ExpressionStatement) statementNode() {}

func (rs *ExpressionStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// Metode ini hanya membuat sebuah buffer dan menulis nilai kembalian dari setiap
// pernyataan metode String() ke dalamnya. Dan kemudian mengembalikan buffer tersebut sebagai sebuah string. Ini mendelegasikan sebagian besar
// dari pekerjaannya ke Pernyataan *ast.Program.
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// implementasi debugger untuk let statment
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.TokenLiteral())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// implementasi debugger untuk return statement
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// implementasi debugger untuk Expression statement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

func (i *Identifier) String() string {
	return i.Value
}

package main

import "fmt"

// Interface Segregation Principle

type Document struct {
	
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {

}

// forced to implement all three functions
func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {
	
}

func (m MultiFunctionPrinter) Scan(d Document) {
	
}

type OldFashionedPrinter struct {

}

func (m OldFashionedPrinter) Print(d Document) {
	// ok
}

// Deprecated: ...
func (m OldFashionedPrinter) Fax(d Document) {
	fmt.Println("Fax operation not supported on OldFashionedPrinter")
}

func (m OldFashionedPrinter) Scan(d Document) {
	fmt.Println("Scan operation not supported on OldFashionedPrinter")
}

// Fix
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {

}

// Implement interfaces that it only needs
func(m MyPrinter) Print() {

}

type Photocopier struct {

}

// implements both
func(m Photocopier) Print() {

}

func(m Photocopier) Scan() {

}

// combine interfaces
type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	doc := Document{}
	ofp := OldFashionedPrinter{}
	ofp.Scan(doc) // throws error
}

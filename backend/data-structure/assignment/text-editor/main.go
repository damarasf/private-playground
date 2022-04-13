package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	// TODO: answer here
	te.UndoStack.Data = append(te.UndoStack.Data, ch)

	for i := 0; i < len(te.RedoStack.Data); i++ {
		te.RedoStack.Data = te.RedoStack.Data[:i]
	}

	fmt.Println(te.UndoStack.Data)
	fmt.Println(te.RedoStack.Data)
}

func (te *TextEditor) Read() []rune {
	// TODO: answer here
	if te.UndoStack.IsEmpty() {
		return nil
	}

	return te.UndoStack.Data
}

func (te *TextEditor) Undo() {
	// TODO: answer here
	if te.UndoStack.IsEmpty() {
		return
	}

	te.RedoStack.Data = append(te.RedoStack.Data, te.UndoStack.Data[len(te.UndoStack.Data)-1])
	te.UndoStack.Data = te.UndoStack.Data[:len(te.UndoStack.Data)-1]

	fmt.Println(te.UndoStack.Data)
	fmt.Println(te.RedoStack.Data)
}

func (te *TextEditor) Redo() {
	// TODO: answer here
	if te.RedoStack.IsEmpty() {
		return
	}

	te.UndoStack.Data = append(te.UndoStack.Data, te.RedoStack.Data[len(te.RedoStack.Data)-1])
	te.RedoStack.Data = te.RedoStack.Data[:len(te.RedoStack.Data)-1]

	fmt.Println(te.UndoStack.Data)
	fmt.Println(te.RedoStack.Data)
}

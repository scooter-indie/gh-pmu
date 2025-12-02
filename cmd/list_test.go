package cmd

import (
	"bytes"
	"testing"
)

func TestListCommand_Exists(t *testing.T) {
	cmd := NewRootCommand()
	cmd.SetArgs([]string{"list", "--help"})

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("list command should exist: %v", err)
	}

	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("list")) {
		t.Error("Expected help output to mention 'list'")
	}
}

func TestListCommand_HasStatusFlag(t *testing.T) {
	cmd := NewRootCommand()
	listCmd, _, err := cmd.Find([]string{"list"})
	if err != nil {
		t.Fatalf("list command not found: %v", err)
	}

	flag := listCmd.Flags().Lookup("status")
	if flag == nil {
		t.Error("Expected --status flag to exist")
	}
}

func TestListCommand_HasPriorityFlag(t *testing.T) {
	cmd := NewRootCommand()
	listCmd, _, err := cmd.Find([]string{"list"})
	if err != nil {
		t.Fatalf("list command not found: %v", err)
	}

	flag := listCmd.Flags().Lookup("priority")
	if flag == nil {
		t.Error("Expected --priority flag to exist")
	}
}

func TestListCommand_HasJSONFlag(t *testing.T) {
	cmd := NewRootCommand()
	listCmd, _, err := cmd.Find([]string{"list"})
	if err != nil {
		t.Fatalf("list command not found: %v", err)
	}

	flag := listCmd.Flags().Lookup("json")
	if flag == nil {
		t.Error("Expected --json flag to exist")
	}
}

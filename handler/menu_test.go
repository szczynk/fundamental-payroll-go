package handler

import (
	"bytes"
	"fundamental-payroll-go/helper/input"
	"fundamental-payroll-go/mocks"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureStdout() (func(), chan string) {
	// Redirect output using an io.Pipe
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Capture output in a separate goroutine
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// Restore the original stdout
	return func() {
		w.Close()
		os.Stdout = oldStdout
	}, outC
}

func restoreStdout(restore func(), outC chan string) string {
	// Read the captured output
	restore()
	return <-outC
}

func TestMenu_ShowMenu(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		method  string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success add employee",
			method:  "Add",
			input:   "1\n5",
			want:    "Add a new employee",
			wantErr: false,
		},
		{
			name:    "success list employee",
			method:  "List",
			input:   "2\n5",
			want:    "Employee list",
			wantErr: false,
		},
		{
			name:    "success add payroll",
			method:  "Add",
			input:   "3\n5",
			want:    "Add a new payroll",
			wantErr: false,
		},
		{
			name:    "success list salary",
			method:  "List",
			input:   "4\n5",
			want:    "Salary Matrix",
			wantErr: false,
		},
		{
			name:    "success list payroll",
			method:  "List",
			input:   "6\n5",
			want:    "Payroll list",
			wantErr: false,
		},
		{
			name:    "success detail payroll",
			method:  "Detail",
			input:   "7\n5",
			want:    "Payroll detail",
			wantErr: false,
		},
		{
			name:    "back to menu",
			method:  "List",
			input:   "\n5",
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			inputReader := input.NewInputReader(reader)

			mockEmployeeHandler := mocks.NewEmployeeHandler(t)
			if strings.Contains(tt.name, "employee") && !strings.Contains(tt.name, "menu") {
				mockEmployeeHandler.On(tt.method).Return()
			}

			mockPayrollHandler := mocks.NewPayrollHandler(t)
			if strings.Contains(tt.name, "payroll") && !strings.Contains(tt.name, "menu") {
				mockPayrollHandler.On(tt.method).Return()
			}

			mockSalaryHandler := mocks.NewSalaryHandler(t)
			if strings.Contains(tt.name, "salary") && !strings.Contains(tt.name, "menu") {
				mockSalaryHandler.On(tt.method).Return()
			}

			clear := func() error { return nil }
			showMenuList := func() {}

			m := NewMenu(mockEmployeeHandler, mockPayrollHandler, mockSalaryHandler, inputReader, clear, showMenuList)

			restore, outC := captureStdout()
			err := m.ShowMenu()

			got := restoreStdout(restore, outC)

			log.Println("tt.name:", tt.name)
			log.Println("err:", err)
			log.Println("got:", got)
			assert.Equal(t, tt.wantErr, err != nil, "ContactUsecase.Add error = %v, wantErr %v", err, tt.wantErr)
			assert.Contains(t, got, tt.want, "Expected got to contain '%s', but got '%s'", tt.want, got)
		})
	}
}

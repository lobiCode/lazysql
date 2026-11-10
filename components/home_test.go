package components

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/jorgerojas26/lazysql/models"
)

func TestClampTreeWidth(t *testing.T) {
	tests := []struct {
		name           string
		width          int
		containerWidth int
		want           int
	}{
		{
			name:           "width within bounds is unchanged",
			width:          30,
			containerWidth: 120,
			want:           30,
		},
		{
			name:           "does not exceed half the container",
			width:          80,
			containerWidth: 100,
			want:           50,
		},
		{
			name:           "floors at min width",
			width:          10,
			containerWidth: 120,
			want:           minTreeWidth,
		},
		{
			name:           "unknown container keeps requested width above min",
			width:          30,
			containerWidth: 0,
			want:           30,
		},
		{
			name:           "unknown container still floors at min",
			width:          5,
			containerWidth: 0,
			want:           minTreeWidth,
		},
		{
			name:           "tiny container does not overflow",
			width:          30,
			containerWidth: 20,
			want:           20,
		},
		{
			name:           "min width wins over half in small containers",
			width:          30,
			containerWidth: 40,
			want:           minTreeWidth,
		},
		{
			name:           "exactly half is allowed",
			width:          50,
			containerWidth: 100,
			want:           50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := clampTreeWidth(tt.width, tt.containerWidth)
			if got != tt.want {
				t.Errorf("clampTreeWidth(%d, %d) = %d, want %d", tt.width, tt.containerWidth, got, tt.want)
			}
		})
	}
}

// Regression: typing "=" (or "+", "-", "[", "]", ...) into the WHERE filter
// resized the tree or switched tabs when the field had been focused with the
// mouse, because only the "/" shortcut sets the tab's filtering flag.
func TestPrintableKeysReachFocusedTextInput(t *testing.T) {
	changes := []models.DBDMLChange{}
	table := &ResultsTable{
		Table:      tview.NewTable(),
		state:      &ResultsTableState{records: [][]string{}, listOfDBChanges: &changes},
		Page:       tview.NewPages(),
		Wrapper:    tview.NewFlex(),
		Pagination: NewPagination(),
		Filter:     NewResultsFilter(),
	}
	home := &Home{
		TabbedPane:         NewTabbedPane(),
		MainContent:        tview.NewFlex(),
		LeftWrapper:        tview.NewFlex(),
		leftWrapperVisible: true,
		treeWidth:          30,
	}
	home.TabbedPane.AppendTab("orders", table, "db.orders")
	home.TabbedPane.AppendTab("users", table, "db.users")

	previousFocus := App.GetFocus()
	t.Cleanup(func() { App.SetFocus(previousFocus) })
	App.SetFocus(table.Filter.Input)
	if table.GetIsFiltering() {
		t.Fatal("mouse-style focus must not set the filtering flag; the test would not exercise the regression")
	}

	for _, key := range []rune{'=', '+', '-', 'T', 'H', 'q'} {
		event := tcell.NewEventKey(tcell.KeyRune, key, tcell.ModNone)
		if got := home.homeInputCapture(event); got != event {
			t.Errorf("home capture swallowed %q while the filter input was focused", key)
		}
	}
	if home.treeWidth != 30 {
		t.Fatalf("tree width changed to %d while typing into the filter", home.treeWidth)
	}

	for _, key := range []rune{'[', ']', '{', '}', 'X', '<', '>'} {
		event := tcell.NewEventKey(tcell.KeyRune, key, tcell.ModNone)
		if got := home.rightWrapperInputCapture(event); got != event {
			t.Errorf("right wrapper capture swallowed %q while the filter input was focused", key)
		}
	}
	if got := home.TabbedPane.GetCurrentTab().Name; got != "users" {
		t.Fatalf("current tab = %q, want users (tabs must not switch while typing)", got)
	}

	// Without a text input focused the same keys remain shortcuts.
	App.SetFocus(table)
	if got := home.homeInputCapture(tcell.NewEventKey(tcell.KeyRune, '=', tcell.ModNone)); got != nil {
		t.Fatal("= should widen the tree when the table itself is focused")
	}
	// The container has no size in tests, so the clamp decides the exact
	// value; what matters is that the shortcut was applied.
	if home.treeWidth == 30 {
		t.Fatal("= should resize the tree when the table itself is focused")
	}
}

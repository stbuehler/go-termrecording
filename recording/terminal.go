package recording

import (
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	// IsTerminal, MakeRaw, Restore
	terminal "code.google.com/p/go.crypto/ssh/terminal"

	// Setsize
	"github.com/asciinema/asciinema-cli/ptyx"
	// Getsize, Start
	"github.com/kr/pty"
)

type Pty struct {
	Stdin          *os.File
	Stdout         *os.File
	ResizeCallback func(columns uint32, rows uint32)
}

func NewPty() *Pty {
	return &Pty{Stdin: os.Stdin, Stdout: os.Stdout}
}

func (p *Pty) Size() (int, int, error) {
	return pty.Getsize(p.Stdout)
}

func (p *Pty) Record(stdoutCopy io.Writer, command string, args ...string) error {
	// put stdin into raw mode (if it's a tty)
	fd := int(p.Stdin.Fd())
	if terminal.IsTerminal(fd) {
		oldState, err := terminal.MakeRaw(fd)
		if err != nil {
			return err
		}
		defer terminal.Restore(fd, oldState)
	}

	// start command in pty
	master, err := pty.Start(exec.Command(command, args...))
	if err != nil {
		return err
	}
	defer master.Close()

	// install WINCH signal handler
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGWINCH)
	go func() {
		for _ = range signals {
			p.resize(master)
		}
	}()
	defer signal.Stop(signals)
	defer close(signals)

	// do initial resize
	p.resize(master)

	// start stdin -> master copying
	stdinForward, err := Forward(master, p.Stdin)
	if err != nil {
		return err
	}
	defer stdinForward.Stop()

	// copy pty master -> p.stdout & stdoutCopy
	stdout := io.MultiWriter(p.Stdout, stdoutCopy)
	io.Copy(stdout, master)

	return nil
}

func (p *Pty) resize(f *os.File) {
	var rows, cols int

	if terminal.IsTerminal(int(p.Stdout.Fd())) {
		rows, cols, _ = p.Size()
	} else {
		rows = 24
		cols = 80
	}

	ptyx.Setsize(f, rows, cols)
	if nil != p.ResizeCallback {
		p.ResizeCallback(uint32(cols), uint32(rows))
	}
}

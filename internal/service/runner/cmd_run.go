package runner

import (
	"bufio"
	"bytes"
	"os/exec"
	"syscall"
)

// commandRun represents the execution of a command.
type commandRun struct {
	cmd           *exec.Cmd     // Command to be executed
	stopSignal    chan struct{} // Signal channel to stop command execution
	outputBuff    *bytes.Buffer // Buffer to store command output
	output        chan string   // Channel to send command output
	maxBufferSize int           // Maximum buffer size for command output
}

// newCommandRun creates a new instance of commandRun.
func newCommandRun(cmd string) *commandRun {
	return &commandRun{
		cmd:           exec.Command("bash", "-c", cmd),
		stopSignal:    make(chan struct{}),
		output:        make(chan string, 1),
		outputBuff:    new(bytes.Buffer),
		maxBufferSize: 1024 * 1024, // 1 MB maximum buffer size
	}
}

// run executes the command represented by commandRun.
func (cr *commandRun) run() error {
	stdout, err := cr.cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cr.cmd.Start(); err != nil {
		return err
	}

	// Channel to track command completion
	done := make(chan error, 1)
	go func() {
		done <- cr.cmd.Wait()
	}()

	// Send output and clean up resources when done
	defer cr.cleanBuff()
	defer func() {
		cr.output <- cr.outputBuff.String()
		close(cr.output)
	}()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		select {
		// Wait for command to finish
		case err := <-done:
			return err
		// Handle stop signal
		case <-cr.stopSignal:
			// Send termination signal
			if err := cr.cmd.Process.Signal(syscall.SIGTERM); err != nil {
				/// If termination fails, used kill
				return cr.cmd.Process.Kill()
			}
			return nil
		// Read command output
		default:
			// If buffer is full, send output and reset buffer
			if cr.isBuffFull() {
				cr.output <- cr.outputBuff.String()
				cr.cleanBuff()
			}
			cr.outputBuff.WriteString(scanner.Text() + "\n")
		}
	}
	return nil
}

// stop stops the execution of the command.
func (cr *commandRun) stop() {
	select {
	case <-cr.stopSignal:
		// Channel already closed
	default:
		close(cr.stopSignal)
	}
}

// cleanBuff resets the output buffer.
func (cr *commandRun) cleanBuff() {
	cr.outputBuff.Reset()
}

// isBuffFull checks if the output buffer is full.
func (cr *commandRun) isBuffFull() bool {
	return cr.outputBuff.Len() >= cr.maxBufferSize
}

// getStatus retrieves the exit code of the command process.
func (cr *commandRun) getStatus() int {
	return cr.cmd.ProcessState.ExitCode()
}

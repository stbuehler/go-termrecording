package recording

import (
	"os"
	"syscall"
)

type TemporaryError interface {
	error
	Temporary() bool
}

type Forwarder struct {
	pipeWrite *os.File
}

func Forward(dest *os.File, source *os.File) (f Forwarder, err error) {
	p := make([]int, 2)
	if err = syscall.Pipe2(p, syscall.O_CLOEXEC); err != nil {
		return
	}

	f.pipeWrite = os.NewFile(uintptr(p[1]), "pipe-write")

	pipeReadFd := p[0]

	go func() {
		buffer := make([]byte, 4096)

		pipeRead := os.NewFile(uintptr(pipeReadFd), "pipe-read")
		defer pipeRead.Close()

		sourceFd := int(source.Fd())
		destFd := int(dest.Fd())

		var selIn, selOut syscall.FdSet

		maxFd := pipeReadFd
		if sourceFd > maxFd {
			maxFd = sourceFd
		}
		if destFd > maxFd {
			maxFd = destFd
		}
		maxFd++

		for {
			fd_SET(&selIn, sourceFd)
			fd_SET(&selIn, pipeReadFd)

			syscall.Select(maxFd, &selIn, nil, nil, nil)
			if fd_ISSET(&selIn, pipeReadFd) {
				return
			}
			if !fd_ISSET(&selIn, sourceFd) {
				continue
			}

			n, err := source.Read(buffer)
			if err != nil {
				println("Forward read error: " + err.Error())
				return
			}
			if n == 0 {
				return // EOF
			}

			sendData := buffer[0:n]

			for len(sendData) > 0 {
				fd_UNSET(&selIn, sourceFd)
				fd_SET(&selIn, pipeReadFd)
				fd_SET(&selOut, destFd)
				syscall.Select(maxFd, &selIn, &selOut, nil, nil)
				if fd_ISSET(&selIn, pipeReadFd) {
					return
				}
				if !fd_ISSET(&selOut, destFd) {
					continue
				}

				n, err := dest.Write(sendData)
				sendData = sendData[n:]
				if err != nil {
					if tempErr, ok := err.(TemporaryError); !ok || !tempErr.Temporary() {
						println("Forward write error: " + err.Error())
						return
					}
					// ignore temporary errors
				}
			}
		}
	}()

	return
}

func (this Forwarder) Stop() {
	if this.pipeWrite != nil {
		this.pipeWrite.Write([]byte{0})
		this.pipeWrite.Close()
		this.pipeWrite = nil
	}
}

func fd_SET(p *syscall.FdSet, i int) {
	p.Bits[i/64] |= 1 << uint(i) % 64
}

func fd_UNSET(p *syscall.FdSet, i int) {
	p.Bits[i/64] &^= (1 << uint(i) % 64)
}

func fd_ISSET(p *syscall.FdSet, i int) bool {
	return (p.Bits[i/64] & (1 << uint(i) % 64)) != 0
}

func fd_ZERO(p *syscall.FdSet) {
	for i := range p.Bits {
		p.Bits[i] = 0
	}
}

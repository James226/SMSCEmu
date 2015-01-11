package logger

import (
    "bytes"
    "io"
    "os"
    "testing"
    "strings"
    )

func TestDebugShouldWriteToStream(t *testing.T) {
    r, w, _ := os.Pipe()

    log := InitCustom(w)
    log.Debug().Print("Test")


    outC := make(chan string)
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    w.Close()
    out := <-outC

    if !strings.HasSuffix(out, "Test\n") {
        t.Error("Fail!");
    }
}

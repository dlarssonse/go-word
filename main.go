package goword

import (
	"os"
	"os/exec"

	dlog "gitlab.com/dlarssonse/go-logger"
)

func main() {

	dlog.OutputConsole = true
	dlog.OutputDebug = true
	dlog.OutputInfo = true
	dlog.OutputWarning = true

	dlog.Info("Testing...")

	archive, err := Open("assets/Fuktmätningsmall_V4_RBK.docx")
	if err != nil {
		dlog.Error("Open(): %s", err)
		return
	}

	for _, file := range archive.Files {
		dlog.Info("%s", file.File.Name)
	}

	// Image 2
	if err := archive.ReplaceFile("word/media/image5.jpg", "assets/test.jpg"); err != nil {
		dlog.Error("GetFileFromReader(): %s", err)
		return
	}
	// Image 3
	if err := archive.ReplaceFile("word/media/image3.jpg", "assets/test.jpg"); err != nil {
		dlog.Error("GetFileFromReader(): %s", err)
		return
	}
	// Image 3
	if err := archive.ReplaceFile("word/media/image6.jpeg", "assets/test.jpg"); err != nil {
		dlog.Error("GetFileFromReader(): %s", err)
		return
	}

	if err := archive.SaveAs("assets/Output.docx"); err != nil {
		dlog.Error("Save(): %s", err)
		return
	}

	os.Chdir("assets")
	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", "Output.docx")
	if err := cmd.Run(); err != nil {
		dlog.Error("ConvertToPDF(): %s", err)
		return
	}
	os.Chdir("..")
}

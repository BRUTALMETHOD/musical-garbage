package controllers

import (
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
)

func Draw(w http.ResponseWriter, r *http.Request) {
	nodeName := os.Getenv("NODE_NAME")
	drawFont := os.Getenv("DRAW_FONT")
	myFigure := figure.NewFigure(nodeName, drawFont, true)
	figure.Write(w, myFigure)
}

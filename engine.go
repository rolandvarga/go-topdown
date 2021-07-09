package main

import (
	"errors"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
)

const (
	ASSETS_DIR = "assets/"
)

var (
	ErrLoadAssets   = errors.New("error loading assets")
	ErrFindingAsset = errors.New("error finding asset")
)

type Engine struct {
	Assets map[string]string
}

func newEngine() Engine {
	engine := Engine{}
	engine.Assets = make(map[string]string)
	err := engine.loadAssets()
	if err != nil {
		panic(fmt.Errorf("%v: %q", ErrLoadAssets, err))
	}
	return engine
}

func (e *Engine) loadAssets() error {
	err := filepath.Walk(ASSETS_DIR, collectAssetsIn(e.Assets))
	if err != nil {
		return err
	}
	return nil
}

// collectAssetsIn walks the assets directory and stores each file in Engine.Assets
func collectAssetsIn(assetsMap map[string]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// determine if entry is a file, remove its extension and assign to map as
		// filename -> filepath
		if !info.IsDir() {
			filename := info.Name()
			ext := filepath.Ext(filename)
			trimmed := filename[0 : len(filename)-len(ext)]
			assetsMap[trimmed] = path
		}
		return nil
	}
}

func (e *Engine) loadPictureAt(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func (e *Engine) loadPictureFrom(asset string) (pixel.Picture, error) {
	var picture pixel.Picture
	err := ErrFindingAsset
	if path, ok := e.Assets[asset]; ok {
		picture, err = e.loadPictureAt(path)
	}
	return picture, err
}

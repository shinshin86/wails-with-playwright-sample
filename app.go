package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Screenshot(url string) string {
	runOption := &playwright.RunOptions{
		SkipInstallBrowsers: true,
	}

	err := playwright.Install(runOption)
	if err != nil {
		log.Fatalf("could not install playwright dependencies: %v", err)
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	option := playwright.BrowserTypeLaunchOptions{
		Channel:  playwright.String("chrome"),
		Headless: playwright.Bool(false),
	}

	browser, err := pw.Chromium.Launch(option)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	defer page.Close()

	if _, err = page.Goto(url); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	b, _ := page.Screenshot()

	return base64.StdEncoding.EncodeToString(b)
}

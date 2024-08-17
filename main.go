package main

import (
	hvy "client/hH"
	"context"
	"fmt"
	"net/http"
)

// TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.
const UrlHost = "<TBD>"
const Token = "<TBD>"

func main() {
	ctx := context.Background()
	httpClient := &http.Client{}

	hevyClient := &hvy.HevyClient{
		ApiUrl:     UrlHost,
		ApiKey:     Token,
		ApiVersion: "v1",
		Context:    ctx,
		Client:     httpClient,
	}

	params := hvy.PaginationParams{}
	templates, _ := hevyClient.GetExerciseTemplates(params)
	fmt.Print(templates.ExerciseTemplates)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

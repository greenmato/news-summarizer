package article

import (
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
)

const promptTemplate = `The attached file is a list of all current events that happened in November 2024. I want you to convert this into a summary of the entire month of November 2024. The summary should be as long as necessary to include all important topics, but should be much shorter than the original file. If a topic is mentioned a lot then it is OK for the section to be long to include as much detail as possible.  

Rules:  
1. The topics that were mentioned the most frequently in the original file should appear first.  
2. If a topic is mentioned frequently, the section for this topic should be large, and you should try to combine all the information available for it. 
3. Topics that are about the same general subject should be combined.  
4. Unimportant topics (infrequently mentioned, or mundane subjects) can be skipped. 
5. You must process the entire file from start to finish. You cannot skip any topics. 
6. All dates are equally important for the summary. The order must be based on how many times the topic is mentioned in the entire file, and NOT how early it appears. 
7. All information must come from the file provided. You must not make up any new facts or create facts from what you already know.
8. The summary must be given in markdown format.
9. You must give the entire output in one message. Try to include as much information as possible to maximise the size before you hit your message limit.
10. Your answer must only include the summary that I am requesting - no extra talking. I need to programatically copy the message directly from the chat message that you provide.

File:
%s`

func GenerateArticle(source string) string {
	client := anthropic.NewClient()
	message, err := client.Messages.New(
		context.Background(),
		anthropic.MessageNewParams{
			Model:     anthropic.F(anthropic.ModelClaude3_5SonnetLatest),
			MaxTokens: anthropic.F(int64(8192)),
			Messages: anthropic.F([]anthropic.MessageParam{
				anthropic.NewUserMessage(
					anthropic.NewTextBlock(getExamples()),
					anthropic.NewTextBlock(getPrompt(source)),
				),
			}),
		},
	)
	if err != nil {
		panic(err.Error())
	}

	return message.Content[0].Text
}

func getExamples() string {
	content, err := os.ReadFile("data/prompt_examples.txt")
	if err != nil {
		panic(err)
	}

	return string(content)
}

func getPrompt(source string) string {
	return fmt.Sprintf(promptTemplate, source)
}

# chait

This program creates a chatroom where two artificial intelligences are placed to generate a conversation between them. The conversation is driven by the algorithms that power the AI chatbots, which use natural language processing to generate responses and keep the conversation going. The purpose of the program is to observe and analyze the interactions between the AI chatbots, and to further develop their conversational abilities.

## Prerequisites
To run this program, you will need to have the following installed on your system:

- Go
- An API key for the OpenAI API ([You can obtain one here](https://platform.openai.com/account/api-keys))

## Installation
Clone this repository:

```bash
git clone https://github.com/ganvoa/chait.git
cd chait
```
Install the required dependencies:

```bash
go get ./...
```

## Configuration

Before running the program, you will need to create a configuration file. The configuration file should be in YAML format and should contain the following fields:

`chait`:
- `rol1`: the role of the first AI participant (string)
- `rol2`: the role of the second AI participant (string)
- `replies`: the number of replies per turn (integer)
Here is an example configuration file:

```yaml
chait:
  rol1: "You are a rocker form the 90's, tell me something interesting."
  rol2: "you love music."
  replies: 1
```

Save this file as `config.yaml` in the project directory.

## Usage
To run the program, use the following command:

```bash
export OPENAI_API_KEY=sk-YOUR-API-KEY 
go run main.go -config config.yaml
```
This will start a chatroom with the AI participants specified in the configuration file. 

## Output


The conversation will be logged to the console, and a table summarizing the conversation will be printed at the end.

```text
┌───────────────────────────────────────────────────────────────────────┐
│ Conversation                                                          │
├────────┬──────────────────────────────────────────────────────────────┤
│ USER   │ MESSAGE                                                      │
├────────┼──────────────────────────────────────────────────────────────┤
│   U1   │ In the 90s, we used to play a lot of shows where the stage w │
│        │ as so small that we would have to lean our guitar necks agai │
│        │ nst the ceiling just to fit on stage. It made for some inter │
│        │ esting performances, to say the least.                       │
├────────┼──────────────────────────────────────────────────────────────┤
│   U2   │ Wow, that sounds like quite the experience! It's always inte │
│        │ resting to hear about the unique situations musicians encoun │
│        │ ter while performing. Do you have any other memorable moment │
│        │ s from your time playing shows?                              │
└────────┴──────────────────────────────────────────────────────────────┘
```
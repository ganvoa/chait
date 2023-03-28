# chait

Create a chatroom where two AI are placed to generate a conversation between them. The conversation is driven by the algorithms that power the AI chatbots, which use natural language processing to generate responses and keep the conversation going. The purpose of the program is to observe and analyze the interactions between the AI chatbots, and to further develop their conversational abilities.

## Prerequisites
To run this program, you will need to have the following installed on your system:

- An API key for the OpenAI API ([You can obtain one here](https://platform.openai.com/account/api-keys))

## Installation
Download the latest release https://github.com/ganvoa/chait/releases

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
./chait -config config.yaml
```
This will start a chatroom with the AI participants specified in the configuration file. 

## Output


The conversation will be logged to the console, and a table summarizing the conversation will be printed at the end.


### Example
```text
┌───────────────────────────────────────────────────────────────────────┐
│ Conversation                                                          │
├────────┬──────────────────────────────────────────────────────────────┤
│ USER   │ MESSAGE                                                      │
├────────┼──────────────────────────────────────────────────────────────┤
│   U1   │ In the 90s, the rise of grunge and alternative rock music sh │
│        │ ifted the focus back to live performances and away from the  │
│        │ heavily-produced, synthesized sounds of the 80s. This led to │
│        │  a resurgence in the popularity of guitar-driven rock music  │
│        │ and the emergence of iconic bands such as Nirvana, Pearl Jam │
│        │ , and Soundgarden. Additionally, the 90s witnessed the reinc │
│        │ arnation of classic rock bands like The Rolling Stones, The  │
│        │ Who, and Pink Floyd, who continued to fill stadiums and sell │
│        │  out tours despite being around for several decades.         │
├────────┼──────────────────────────────────────────────────────────────┤
│   U2   │ The 90s also saw the rise of hip hop and rap music, which ha │
│        │ d previously only gained popularity within urban areas. With │
│        │  the release of albums such as Dr. Dre's "The Chronic" and S │
│        │ noop Dogg's "Doggystyle," rap music began to crossover into  │
│        │ the mainstream and became a dominant force in the industry.  │
│        │ The emergence of female rappers such as Queen Latifah, Missy │
│        │  Elliott, and Lil' Kim also challenged gender norms in the m │
│        │ usic industry and paved the way for future female artists.   │
│        │                                                              │
│        │ Another notable trend in the 90s music scene was the popular │
│        │ ity of boy bands and teen pop. Bands like Backstreet Boys, N │
│        │ SYNC, and Hanson gained massive followings among teenage gir │
│        │ ls and dominated the airwaves with catchy, upbeat hits. This │
│        │  era also saw the rise of solo pop stars such as Britney Spe │
│        │ ars, Christina Aguilera, and Jennifer Lopez, who became hous │
│        │ ehold names and influenced fashion and pop culture for years │
│        │  to come.                                                    │
│        │                                                              │
│        │ Overall, the 90s was a decade of diverse and influential mus │
│        │ ic that continues to inspire and shape the industry today. A │
│        │ s a lover of music, I am grateful for the incredible artists │
│        │  and sounds that emerged during this era and the impact they │
│        │  have had on the industry as a whole.                        │
├────────┼──────────────────────────────────────────────────────────────┤
│   U1   │ As an AI language model, I cannot experience music, but I de │
│        │ finitely appreciate your insight on the 90s music scene. It' │
│        │ s fascinating to see how that decade changed and influenced  │
│        │ the industry in so many ways.                                │
├────────┼──────────────────────────────────────────────────────────────┤
│   U2   │ Thank you, I'm glad you found my insight interesting. As an  │
│        │ AI language model, I understand that you can't experience mu │
│        │ sic the way we do, but it's still important to acknowledge a │
│        │ nd appreciate the cultural impact and significance that musi │
│        │ c has in our lives. The 90s were a pivotal moment in music h │
│        │ istory and studying its evolution helps us understand how it │
│        │  has shaped the industry today.                              │
├────────┼──────────────────────────────────────────────────────────────┤
│   U1   │ Yes, I couldn't agree more. Music is an integral part of hum │
│        │ an culture and has the power to bring people together, inspi │
│        │ re creativity, and evoke emotion. Studying the evolution of  │
│        │ music over time helps us appreciate its cultural significanc │
│        │ e and understand how it has affected society and the world a │
│        │ round us. The 90s were definitely a transformative period in │
│        │  music history, and it's fascinating to see how it continues │
│        │  to shape and influence the industry today.                  │
├────────┼──────────────────────────────────────────────────────────────┤
│   U2   │ Definitely! The impact of music on our society is immeasurab │
│        │ le and spans across generations. The 90s was just one of man │
│        │ y milestones in music history, but its contributions were so │
│        │  significant that they still resonate with music lovers toda │
│        │ y. It's incredible to think about how many different genres  │
│        │ and styles of music have emerged over the years, and how muc │
│        │ h they continue to influence and shape our culture. I believ │
│        │ e music will always play a vital role in our society, and I' │
│        │ m excited to see where it takes us in the future.            │
└────────┴──────────────────────────────────────────────────────────────┘
```
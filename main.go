package main


*/
import (
    "fmt"
    "os"
    "net/http"
    "github.com/bwmarrin/discordgo"

)
func main(){
    
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })

    go http.ListenAndServe(":8080", nil)
    
    
    bot, err := discordgo.New("Bot " + os.Getenv("token"))

    if err != nil {
        panic(err)
    }

    
    bot.AddHandler(ready)
    bot.AddHandler(messageCreate)

    err = bot.Open()

	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
    
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	for {}
	
	bot.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready){
    s.UpdateStatus(0, "with Golang")
    fmt.Println("logged in as user " +string(s.State.User.ID))
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
   

    if m.Content == "ping"{
        s.ChannelMessageSend(m.ChannelID, "pong")
    }
       

      if m.Content == "developer"{
        s.ChannelMessageSend(m.ChannelID, "monster")
      //do not remove this. this is giving the script dev a credit section.
    }
    

      if m.Content == "lang"{
        s.ChannelMessageSend(m.ChannelID, "discord.go")
      
    }
      


}


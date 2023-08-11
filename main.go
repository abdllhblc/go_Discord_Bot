package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
   sess ,err := discordgo.New("Bot MTEzOTQ5NTUyMTEyMDE1NzcxNg.GgLVbN.YgK8AAvUFRa6ElUu1O1-qULzQB-3JMcX6WHpwI")

   if err !=nil{
	log.Fatal(err)
   }
    
   sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == s.State.User.ID{
		return
	}
	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID,"world!")
	}})

   sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()

	if err != nil{
		log.Fatal(err)
	}	
   
	defer sess.Close()

	fmt.Println("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}

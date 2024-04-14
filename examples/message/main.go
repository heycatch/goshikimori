package main

import (
  "fmt"
  g "github.com/heycatch/goshikimori"
)

func conf() *g.Configuration {
  return g.Add(
    "APPLICATION_NAME",
    "PRIVATE_KEY",
  )
}

func readMessage() {
  c := conf()
  // method #1. find my id and read inbox.
  fast, status, err := c.FastIdUser("arctica")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  messages, err := fast.UserMessages(&g.Options{Type: "inbox", Page: 1, Limit: 10})
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range messages {
    message, status, err := c.ReadMessage(v.Id)
    if status != 200 || err != nil {
      fmt.Println(status, err)
      return
    }
    fmt.Println(
      message.Body,
      message.From.Id, message.From.Nickname,
      message.To.Id, message.To.Nickname,
    )
  }
  // method #2. read the dialogs, find the kind and id.
  d, status, err := c.Dialogs()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range d {
    if v.Message.Kind == "Private" && v.Message.Body != "" {
      message, status, err := c.ReadMessage(v.Message.Id)
      if status != 200 || err != nil {
        fmt.Println(status, err)
        return
      }
      fmt.Println(
        message.Body,
        message.From.Id, message.From.Nickname,
        message.To.Id, message.To.Nickname,
      )
    }
  }
  // method #3. friend nickname search.
  // In this case all messages from the last to the first are found.
  f, status, err := c.FastIdUser("morr")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  sd, err := f.SearchDialogs()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range sd {
    fmt.Println(
      v.Body,
      v.From.Id, v.From.Nickname,
      v.To.Id, v.To.Nickname,
    )
  }
}

func sendMessage() {
  c := conf()
  me, status, err := c.WhoAmi()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  to, status, err := c.FastIdUser("morr")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  message := "test message from API"
  send, status, err := c.SendMessage(me.Id, to.Id, message)
  if status != 201 || err != nil {
    fmt.Println(status, err)
    return
  }
  fmt.Println(send.Id, send.Body, send.From.Nickname, send.To.Nickname)
}

func changeMesage() {
  c := conf()
  d, status, err := c.Dialogs()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range d {
    if v.Message.Kind == "Private" && v.Message.Body != "" {
      new_message := "changed the message from API :)"
      message, status, err := c.ChangeMessage(v.Message.Id, new_message)
      if status != 200 || err != nil {
        fmt.Println(status, err)
        return
      }
      fmt.Println(
        message.Body,
        message.From.Id, message.From.Nickname,
        message.To.Id, message.To.Nickname,
      )
    }
  }
}

func deleteMessage() {
  c := conf()
  d, status, err := c.Dialogs()
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }
  for _, v := range d {
    if v.Message.Kind == "Private" && v.Message.Body != "" {
      status, err := c.DeleteMessage(v.Message.Id)
      if status != 204 || err != nil {
        fmt.Println(status, err)
        return
      }
      fmt.Println(status) // only status 204 is returned
    }
  }
}

func markReadUnreadMessages() {
  c := conf()
  var count int

  fast, status, err := c.FastIdUser("arctica")
  if status != 200 || err != nil {
    fmt.Println(status, err)
    return
  }

  ids, err := fast.UnreadMessagesIds("messages")
  if err != nil {
    fmt.Println(err)
    return
  }

  messages, err := fast.UserMessages(&g.Options{Type: "inbox", Page: 1, Limit: 10})
  if err != nil {
    fmt.Println(err)
    return
  }
  for _, v := range messages {
    message, status, err := c.ReadMessage(v.Id)
    if status != 200 || err != nil {
      fmt.Println(status, err)
      return
    }
    // You can set all sorts of filters for easy array compilation.
    if message.From.Nickname == "morr" && !message.Read {
      ids[count] = message.Id
      count++
    }
  }

  read, err := c.MarkReadMessages(g.IdsToString(ids), 1)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(read)
}

func readAll() {
  c := conf()
  read, err := c.ReadAllMessages("notifications")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(read)
}

func deleteAll() {
  c := conf()
  del, err := c.DeleteAllMessages("notifications")
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(del)
}

func main() {
  // not to use all requests at the same time!
  readMessage()
  sendMessage()
  changeMesage()
  deleteMessage()

  markReadUnreadMessages()
  readAll()
  deleteAll()
}

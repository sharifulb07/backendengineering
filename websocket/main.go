package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Websocket Upgrader


var upgrader=websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}




// Client

type Client struct{
	conn *websocket.Conn
	send chan[]byte
}


// Hub (manage channels )


type Hub struct{
	clients map[*Client]bool
	broadcast chan []byte
	register chan *Client
	unregister chan *Client
}


func  newHub()*Hub  {
	return &Hub{
		clients: make(map[*Client]bool),
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),

	}
	
}



// run 


func (h *Hub)run(){
	for{
		select{
		case client:= <-h.register:
			h.clients[client]=true
			log.Println("Client Connected")
		case client:= <-h.unregister:

			if _, ok:=h.clients[client]; ok{
				delete(h.clients, client)
				close(client.send)

				log.Println("Client disconnected")

			}

		case message:= <- h.broadcast:
			for client:= range h.clients{
				select{
				case client.send<-message:

				default:
					close(client.send)
					delete(h.clients, client)
				}

			}
		}

	
	}
}




// Read & Write Pumps 


func (c *Client) readPump(h *Hub){
	defer func(){
		h.unregister <- c 
		c.conn.Close()
	}()


	for{
		_, msg, err:=c.conn.ReadMessage()
		if err !=nil{
			break
		}

		h.broadcast <-msg
	}
}


func ( c *Client)writePump(){

	defer c.conn.Close()

	for msg:= range c.send{
		err:=c.conn.WriteMessage(websocket.TextMessage, msg)

		if err!=nil{
			break; 
		}
	}
}





// http Handler 


func serveWS(hub *Hub, w http.ResponseWriter, r *http.Request){

	con, err:=upgrader.Upgrade(w, r, nil )

	if err!=nil{
		log.Println(err)
		return 
	}

	client:=&Client{
		conn: con,
		send: make(chan []byte),
	}

hub.register <- client


go client.writePump()
go client.readPump(hub)
	
}

func main(){


	hub:=newHub()
	go hub.run()


	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(hub,w, r)
	})



	log.Println("Websocket server is running on : 8081")
	log.Println(http.ListenAndServe(":8081", nil ))
	


}
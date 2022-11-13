package main

//var wsu = websocket.Upgrader{
//	// check origin will check the cross region source (note : please not using in production)
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//func main() {
//	r := gin.Default()
//	r.GET("/", func(c *gin.Context) {
//		//upgrade get request to websocket protocol
//		ws, err := src.wsu.Upgrade(c.Writer, c.Request, nil)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer ws.Close()
//		for {
//			//Read Message from client
//			mt, message, err := ws.ReadMessage()
//			if err != nil {
//				fmt.Println(err)
//				break
//			}
//			//If client message is ping will return pong
//			if string(message) == "ping" {
//				message = []byte("pong")
//			}
//			//Response message to client
//			err = ws.WriteMessage(mt, message)
//			if err != nil {
//				fmt.Println(err)
//				break
//			}
//		}
//	})
//	_ = r.Run() // listen and serve on 0.0.0.0:8080
//}

package main

import (
	"github.com/gorilla/websocket" // ใช้ในการสร้าง websocket server
	"log"			// ใช้ในการแสดงข้อความ ดีกว่าใช้ fmt เพราะมีการจัด format ข้อความที่แสดงอยู่แล้ว
	"net/http"		// ใช้ในการสร้าง http server
	"strings"		// ใช้ในการแยกข้อความออกเป็น array โดยใช้ฟังก์ชัน Split
)

// ตัวแปรเก็บข้อมูลที่เป็น websocket 
// Upgrader คือ ตัวที่ใช้ในการอัพเกรด http connection เป็น websocket connection
var upgrader = websocket.Upgrader{}

// todoList คือ ตัวแปรเก็บข้อมูลที่เป็น array ของ string
var todoList []string

// ฟังก์ชัน getCmd รับข้อมูล input แยกคำออกมาแล้วส่งคำสั่งกลับ
func getCmd(input string) string {
	inputArr := strings.Split(input, " ") // แยกข้อความออกด้วยช่องว่าง แล้วเก็บในตัวแปร inputArr
	return inputArr[0]
}

// ฟังก์ชัน getMessage รับข้อมูล input แยกคำออกจากช่องว่าง 
// วนลูปเอาข้อความต่อกันแล้วส่งกลับ
func getMessage(input string) string {
	inputArr := strings.Split(input, " ") // แยกข้อความออกด้วยช่องว่าง แล้วเก็บในตัวแปร inputArr
	var result string
	for i := 1; i < len(inputArr); i++ {
		result += inputArr[i]		// การบวก string จะเป็นการต่อกัน += คือการบวกแล้วเก็บค่าในตัวแปรเดิม
	}
	return result
}

// ฟังก์ชัน updateTodoList รับข้อมูล input 
// วนลูป จาก range ของ todoList ถ้าข้อมูลใน todoList 
// ตรงกับ input ให้ข้ามไป
func updateTodoList(input string) {
	tmpList := todoList			// สร้างตัวแปร tmpList เพื่อเก็บข้อมูลใน todoList
	todoList = []string{}		// ลบข้อมูลใน todoList ให้เป็น array ว่าง

	/*
	   วนลูป จาก range ของ tmpList ถ้าข้อมูลใน tmpList
	   ตัวอย่างเช่น ถ้า tmpList เก็บข้อมูล ["a", "b", "c"] 
	   จะวนลูปเป็น 3 รอบ โดย val จะเก็บค่าเป็น "a" และ "b" และ "c"
	
	*/
	for _, val := range tmpList { // range คือ การวนลูปข้อมูลใน array
		// ถ้าข้อมูลใน todoList ตรงกับ input ให้ข้ามไป ไม่ต้องเพิ่มข้อมูลเข้าไปใน todoList
		if val == input {
			continue	// ข้ามไป ไม่ต้องทำอะไร ให้วนลูปต่อไปอีกรอบ
		}
		// ถ้าไม่ตรงให้เพิ่มข้อมูลเข้าไปใน todoList
		todoList = append(todoList, val)
	}
}

func main() {

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade failed: ", err)
			return
		}
		defer conn.Close()

		// เพิ่มข้อมูลเข้าไปใน todoList
		for {
			mt, message, err := conn.ReadMessage() 	// ReadMessage reads a message from the WebSocket connection.
			if err != nil {							// ถ้ามี error ให้แสดงข้อความและหยุดการทำงาน
				log.Println("read failed:", err)	// แสดงข้อความ error
				break								// หยุดการทำงานของ for loop
			}

			input := string(message)	// แปลงข้อมูลที่รับเข้ามาเป็น string
			cmd := getCmd(input)		// แยกคำสั่งออกมา แล้วเก็บในตัวแปร cmd
			msg := getMessage(input)	// แยกข้อความออกมา แล้วเก็บในตัวแปร msg

			if cmd == "add" {			// ถ้า cmd เป็น add ให้เพิ่มข้อมูลเข้าไปใน todoList
				todoList = append(todoList, msg) // append คือ การเพิ่มข้อมูลเข้าไปใน array
			} else if cmd == "done" {   // ถ้า cmd เป็น done ให้ลบข้อมูลออกจาก todoList
				updateTodoList(msg)		// เรียกใช้ฟังก์ชัน updateTodoList
			}

			output := "Current Todos: \n"		// สร้างตัวแปร output เพื่อเก็บข้อความที่จะส่งกลับไปยัง client (browser)
			for _, todo := range todoList {		// วนลูป จาก range ของ todoList 
				output += "\n - " + todo + "\n"	// ต่อข้อความใน output โดยใช้ += คือการต่อกัน
			}
			output += "\n----------------------------------------"

			message = []byte(output)			// แปลงข้อความใน output เป็น byte เนื่องจาก WriteMessage ต้องรับข้อมูลเป็น byte
			err = conn.WriteMessage(mt, message)	// WriteMessage เขียนข้อมูลไปยัง WebSocket connection ที่เปิดอยู่
			if err != nil {						// ถ้ามี error ให้แสดงข้อความและหยุดการทำงาน
				log.Println("write failed:", err)
				break
			}
		}
	
		// สิ้นสุดการเพิ่มข้อมูลเข้าไปใน todoList
	})

	// สร้าง http server ที่ port 8080
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html") // ส่งไฟล์ websockets.html ไปแสดงผล
	})

	http.ListenAndServe(":8080", nil)			// เริ่มต้น http server ที่ port 8080
}
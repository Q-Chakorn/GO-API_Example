package main //ระบุว่าเป็น package หลัก
import (
	"net/http" // library สำหรับการจัดการ HTTP requests

	"github.com/gin-gonic/gin" // library Gin framework สำหรับการสร้าง RESTful API
) // เรียกใช้ library ที่จำเป็น

// PokeDex struct สำหรับเก็บข้อมูลของ Pokemon
type PokeDex struct {
	ID      string   `json:"id"`
	NAME    string   `json:"name"`
	ELEMENT []string `json:"element"` // slice(ชุดของข้อมูลหลายค่าเหมือน array) ของ string สำหรับเก็บประเภทของ Pokemon
}

// Mock data สำหรับเก็บข้อมูล Pokemon
var pokemons = []PokeDex{
	{ID: "1", NAME: "Bulbasaur", ELEMENT: []string{"grass", "poison"}},
	{ID: "2", NAME: "Ivysaur", ELEMENT: []string{"grass", "poison"}},
	{ID: "3", NAME: "Venusaur", ELEMENT: []string{"grass", "poison"}},
}

// ฟังก์ชัน main สำหรับเริ่มต้นโปรแกรม web server ด้วย Gin
func main() {
	router := gin.Default()                       //สร้าง router(obj กำหนดเส้นทาง HTTP request) หลักสำหรับรับส่ง HTTP request
	router.GET("/pokemons", getPokeDex)           //ดึงข้อมูลโปเกม่อนทั้งหมด
	router.GET("/pokemons/:id", getPokeDexByID)   //ดึงข้อมูลโปเกม่อนด้วย id
	router.POST("/pokemons", addPokemon)          //เพิ่มโปเกม่อนใหม่
	router.PUT("/pokemons/:id", updatePokeDex)    //แก้ไขข้อมูลโปเกม่อนตาม id
	router.DELETE("/pokemons/:id", deletePokemon) //ลบโปเกม่อนด้วย id
	router.Run(":8080")                           //run server ที่ port 8080
}

// ฟังก์ชัน handler(รับข้อมูลจาก request) ที่ใช้ดึงข้อมูลโปเกม่อนทั้งหมด
func getPokeDex(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pokemons) //ส่ง response กลับไปเป็น JSON พร้อมสถานะ 200 และข้อมูลทั้งหมดในตัวแปร pokemons
}

// ฟังก์ชัน handler ที่ใช้ดึงข้อมูลโปเกม่อนตาม ID
func getPokeDexByID(c *gin.Context) {
	id := c.Param("id") // ดึงค่าพารามิเตอร์ id จาก URL ที่ผู้ใช้ส่งมา
	for _, a := range pokemons {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	} // วนลูปดูโปเกม่อนทุกตัวใน slice pokemons ถ้าเจอโปเกม่อนที่มี id ตรงกับที่รับมา ส่งข้อมูลโปเกม่อนตัวนั้นกลับไปในรูปแบบ JSON พร้อมสถานะ 200
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"}) // ถ้าไม่เจอ id ส่งข้อความแจ้งว่าไม่พบข้อมูล (404 Not Found) กลับมา
}
func addPokemon(c *gin.Context) {
	var newPokemon PokeDex
	if err := c.BindJSON(&newPokemon); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID in body does not match ID in path"})
		return
	}
	pokemons = append(pokemons, newPokemon)
	c.IndentedJSON(http.StatusCreated, newPokemon)
}
func updatePokeDex(c *gin.Context) {
	id := c.Param("id")
	var updatePokeDex PokeDex
	if err := c.BindJSON(&updatePokeDex); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, b := range pokemons {
		if b.ID == id {
			if updatePokeDex.ID != "" && updatePokeDex.ID != id {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID in body does not match ID in path"})
				return
			}
			pokemons[i] = updatePokeDex
			pokemons[i].ID = id
			c.IndentedJSON(http.StatusOK, pokemons[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
}
func deletePokemon(c *gin.Context) {
	id := c.Param("id")
	for i, b := range pokemons {
		if b.ID == id {
			pokemons = append(pokemons[:i], pokemons[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Pokemon deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
}

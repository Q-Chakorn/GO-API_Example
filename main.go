package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type PokeDex struct {
	ID	string `json:"id"`
	NAME	string `json:"name"`
	TYPE1	string `json:"type1"`
	TYPE2 string `json:"type2"`
}

var pokedexs = []PokeDex {
	{ID:"1" , NAME:"Bulbasaur" , TYPE1:"grass" , TYPE2:"poison" },
}
func main() {
	router := gin.Default()

	router.GET("/pokedexs", getPokeDex)
	router.GET("/pokedexs/:id", getPokeDexByID)
	router.POST("/pokedexs", addPokemon)
	router.PUT("/pokedexs/:id", updatePokeDex)
	router.DELETE("/pokedexs/:id", deletePokeDex)

	router.Run(":8080")
}
func getPokeDex(c *gin.Context){
	c.IndentedJSON(http.StatusOK,pokedexs)
}

func getPokeDexByID(c *gin.Context){
	id := c.Param("id")
	for _, a := range pokedexs {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}		
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "PokeDex not found"})
}
func addPokemon(c *gin.Context){
	var newPokemon PokeDex
	if err := c.BindJSON(&newPokemon); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID in body does not match ID in path"})
		return
	}
	pokedexs = append(pokedexs, newPokemon)
	c.IndentedJSON(http.StatusCreated, newPokemon)
}
func updatePokeDex(c *gin.Context){
	id := c.Param("id")
	var updatePokeDex PokeDex
	if err := c.BindJSON(&updatePokeDex); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, b := range pokedexs{
		if b.ID == id {
			if updatePokeDex.ID != "" && updatePokeDex.ID != id {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID in body does not match ID in path"})
				return
			}
			pokedexs[i] = updatePokeDex
			pokedexs[i].ID = id
			c.IndentedJSON(http.StatusOK , pokedexs[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "PokeDex not found"})
}
func deletePokeDex(c *gin.Context)  {
	id := c.Param("id")
	for i , b := range pokedexs{
		if b.ID == id {
			pokedexs = append(pokedexs[:i], pokedexs[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "PokeDex deleted"})
			return			
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "PokeDex not found"})
}

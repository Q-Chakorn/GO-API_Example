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

var pokemons = []PokeDex {
	{ID:"1" , NAME:"Bulbasaur" , TYPE1:"grass" , TYPE2:"poison" },
	{ID:"2" , NAME:"Ivysaur" , TYPE1:"grass" , TYPE2:"poison" },
	{ID:"3" , NAME:"Venusaur" , TYPE1:"grass" , TYPE2:"poison" },
}
func main() {
	router := gin.Default()

	router.GET("/pokemons", getPokeDex)
	router.GET("/pokemons/:id", getPokeDexByID)
	router.POST("/pokemons", addPokemon)
	router.PUT("/pokemons/:id", updatePokeDex)
	router.DELETE("/pokemons/:id", deletePokemon)

	router.Run(":8080")
}
func getPokeDex(c *gin.Context){
	c.IndentedJSON(http.StatusOK,pokemons)
}

func getPokeDexByID(c *gin.Context){
	id := c.Param("id")
	for _, a := range pokemons {
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
	pokemons = append(pokemons, newPokemon)
	c.IndentedJSON(http.StatusCreated, newPokemon)
}
func updatePokeDex(c *gin.Context){
	id := c.Param("id")
	var updatePokeDex PokeDex
	if err := c.BindJSON(&updatePokeDex); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, b := range pokemons{
		if b.ID == id {
			if updatePokeDex.ID != "" && updatePokeDex.ID != id {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID in body does not match ID in path"})
				return
			}
			pokemons[i] = updatePokeDex
			pokemons[i].ID = id
			c.IndentedJSON(http.StatusOK , pokemons[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "PokeDex not found"})
}
func deletePokemon(c *gin.Context)  {
	id := c.Param("id")
	for i , b := range pokemons{
		if b.ID == id {
			pokemons = append(pokemons[:i], pokemons[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Pokemon deleted"})
			return			
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
}

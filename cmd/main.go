package main



import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    //"github.com/DjadiHadia/restapiwithgo/database"
    //"database/sql"
    

    //_ "github.com/lib/pq"
   
)

func main() {
    
    //database connexion using gorm
    database.ConnectDb()
    
    myAgency := NewAgency()
    GetInfo(myAgency)
    myAgency.UpdatePhone("00213777777")
    GetInfo(myAgency)
    cars := []car{newcar("00000000000","BMW","black","5","2015"),newcar("11111111111","Mercedes","black","5","2023"),newcar("222222222222","Mini cooper","red","5","2020"),newcar("33333333333","Jeep","black","5","2020")}

    for _, car := range cars{
        if car["color"] == "black" {
            printCar(car)
        }
        
    }
    var client1 client
    client1.first_name="djadi"
    client1.last_name =  "hadia"
    client1.Address= "ALGERIA"
    client1.email="hadianew@hotmail.com"
    client1.Phone="002135555555"

    var client2 client
    client2.first_name="djadi2"
    client2.last_name =  "hadia2"
    client2.Address= "ALGERIA"
    client2.email="hadianew@hotmail.com"
    client2.Phone="002135511115"

    
    
     clients := []client{client1,client2}

     for _, client := range clients{

        //fmt.Println(client)
        fmt.Printf("%+v",client)
    }
   


    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, hadia!")
    })

    app.Listen(":3000")

}



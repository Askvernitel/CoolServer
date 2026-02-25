package server

import "github.com/gin-gonic/gin"



type Server struct {  

}

func (s *Server) Start(){ 
	gin.Default();
}

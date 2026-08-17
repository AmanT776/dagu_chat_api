package logger

import (
	"fmt"
	"log/slog"
	"os"
)


func New(env string) *slog.Logger{
	var handler slog.Handler
	switch env {
	case "production":
		logfile ,err := os.OpenFile("dagu.access.log",os.O_APPEND | os.O_WRONLY | os.O_CREATE, 0o644 )
		if err != nil{
			fmt.Printf("error opening file for loggging: %s",err)
		}
		handler = slog.NewJSONHandler(logfile,&slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
		
	default:
		handler = slog.NewJSONHandler(os.Stderr,&slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}	
	return slog.New(handler)
}
package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	pb "Proyecto_gRPC/proto"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var db *sql.DB

type server struct {
	pb.UnimplementedWarframeServiceServer
}

func (s *server) GetWaframeInfo(ctx context.Context, req *pb.WarframeRequest) (*pb.WarframeResponse, error) {

	var name, wtype string
	var level int

	query := "SELECT * FROM warfraprd.Lotus where Name like  @Name"

	row := db.QueryRowContext(ctx, query, sql.Named("Name", "%"+req.Name+"%"))
	err := row.Scan(&name, &wtype, &level)

	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.WarframeResponse{

				Name:  "Not found",
				Type:  "Not found",
				Level: 0,
			}, nil
		}
		return nil, err
	}

	return &pb.WarframeResponse{

		Name:  name,
		Type:  wtype,
		Level: int32(level),
	}, nil

}

func (s *server) GetWarameList(req *pb.Empty, stream pb.WarframeService_GetWarameListServer) error {
	query := "Select * from warfraprd.Lotus"

	rows, err := db.Query(query)
	if err != nil {
		log.Panic(err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var name, wtype string
		var level int

		if err := rows.Scan(&name, &wtype, &level); err != nil {

			log.Panic(err)
			return err

		}

		if err := stream.Send(&pb.WarframeResponse{
			Name:  name,
			Type:  wtype,
			Level: int32(level),
		}); err != nil {
			log.Panic(err)
			return err
		}
	}

	return nil
}

func (s *server) AddWarframes(stream pb.WarframeService_AddWarframesServer) error {
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AddWarframesResponse{
				Count: count,
			})
		}

		if err != nil {
			log.Panic(err)
			return err
		}

		query := "insert into warfraprd.Lotus (Name, Type, Level) values (@Name, @Type, @Level)"
		_, err = db.Exec(query,
			sql.Named("Name", req.Name),
			sql.Named("Type", req.Type),
			sql.Named("Level", req.Level))

		if err != nil {
			log.Panic(err)
			return err
		}

		count++
		log.Printf("Added %s", req.Name)
	}
}

func (s *server) GetWarframesByType(stream pb.WarframeService_GetWarframesByTypeServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("End of stream")
			return nil
		}

		if err != nil {
			log.Println("Err")
			return err
		}

		query := "Select * from warfraprd.Lotus where Type = @Type"
		rows, err := db.Query(query, sql.Named("Type", req.Type))
		if err != nil {
			log.Panic(err)
			return err
		}
		defer rows.Close()

		for rows.Next() {
			var name, wtype string
			var level int

			if err := rows.Scan(&name, &wtype, &level); err != nil {
				log.Panic(err)
				return err
			}

			if err := stream.Send(&pb.WarframeResponse{
				Name:  name,
				Type:  wtype,
				Level: int32(level),
			}); err != nil {
				log.Panic(err)
				return err
			}
		}
	}

}

func main() {

	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")

	}

	s := os.Getenv("DB_SERVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD_")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, s, port, database)

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Conneted to database")

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterWarframeServiceServer(grpcServer, &server{})

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Ok"))
		})
		log.Println("Starting health check server on port 8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	log.Println("Starting server on 50051")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

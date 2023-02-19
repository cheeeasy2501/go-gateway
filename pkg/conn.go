package pkg

import "google.golang.org/grpc"

type GRPCConnection struct {
	conn *grpc.ClientConn
}

func NewGRPCConnection(target string, opts ...grpc.DialOption) *GRPCConnection {
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		panic(err)
	}

	return &GRPCConnection{
		conn: conn,
	}
}

func (c *GRPCConnection) GetConnection() *grpc.ClientConn {
	return c.conn
}

func (c *GRPCConnection) CloseConnection() {
	c.conn.Close()
}
